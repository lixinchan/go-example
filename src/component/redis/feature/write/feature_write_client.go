package main

import (
	"component/redis/feature/redis"
	"component/redis/lib/utils"
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var (
	expireTime       = 1000
	insertedCnt      = int64(0)
	writeCnt         = int64(0)
	handleConcurrent = int(10)
	path             = ""
	prefix           = ""
	redisClient      *redis.Client
)

func main() {
	app := cli.NewApp()
	app.Name = "feature_write_client"

	var host string
	var port string
	var auth string
	var initDb int
	var path string
	var prefix string
	var concurrent int
	var writeCnt int64
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host, h",
			Value:       "127.0.0.1",
			Usage:       "redis host",
			Destination: &host,
		},
		cli.StringFlag{
			Name:        "port, p",
			Value:       "6379",
			Usage:       "redis port",
			Destination: &port,
		},
		cli.StringFlag{
			Name:        "auth, a",
			Value:       "",
			Usage:       "redis auth",
			Destination: &auth,
		},
		cli.IntFlag{
			Name:        "db, d",
			Value:       0,
			Usage:       "redis db",
			Destination: &initDb,
		},
		cli.StringFlag{
			Name:        "path",
			Value:       "",
			Usage:       "data path",
			Destination: &path,
		},
		cli.StringFlag{
			Name:        "prefix",
			Value:       "",
			Usage:       "prefix",
			Destination: &prefix,
		},
		cli.IntFlag{
			Name:        "concurrent, c",
			Value:       20,
			Usage:       "concurrent",
			Destination: &concurrent,
		}, cli.Int64Flag{
			Name:        "wirteCnt, cnt",
			Value:       0,
			Usage:       "redis write cnt",
			Destination: &writeCnt,
		},
	}

	app.Action = func(c *cli.Context) (err error) {
		writeCnt = int64(writeCnt)
		handleConcurrent = int(concurrent)
		path = string(path)
		prefix = string(prefix)
		redisConfig := &redis.PoolConf{
			Host:    host,
			Port:    port,
			InitDb:  initDb,
			Auth:    auth,
			InitNum: concurrent,
			MaxNum:  concurrent,
		}
		insertData(redisConfig, path, prefix, handleConcurrent)
		return
	}
}

func insertData(redisConf *redis.PoolConf, path, prefix string, concurrent int) {
	if utils.IsValidDir(path) {
		fileArray := utils.ListAllFile(path)
		if len(fileArray) == 0 {
			return
		}
		redisClient, err := redis.NewClient("write", redisConf)
		if err != nil {
			fmt.Println("redis init failed!", err)
			os.Exit(1)
		}
		for _, val := range fileArray {
			_ = utils.ForEachLineConcurrent(val, func(line string) error {
				sleepTime := time.Duration(rand.Int63n(int64(2000))) * time.Microsecond
				time.Sleep(sleepTime)
				ln := strings.Fields(line)
				var featureData map[string]string
				if len(ln) > 2 {
					featureData = make(map[string]string)
					if err := json.Unmarshal([]byte(ln[1]), &featureData); err == nil {
						features := make([]redis.Feature, 0)
						for k, v := range featureData {
							f := redis.Feature{Field: k, Val: v, Expire: randExpireTime(1000, 3600)}
							features = append(features, f)
						}
						_, err = redisClient.Hmset(prefix+":"+ln[0], features, 86400)
						if err == nil {
							atomic.AddInt64(&insertedCnt, int64(1))
							if insertedCnt > writeCnt {
								os.Exit(1)
							}
						}
						if insertedCnt%1000 == 0 {
							fmt.Printf("%s inserted-num: %d\n", time.Now().Format("15:04:05"), insertedCnt)
						}
					}
				}
				return nil
			}, concurrent)
		}
	}
}

func randExpireTime(min, max int64) int64 {
	if (min > max) || (min == 0) || (max == 0) {
		return max
	}
	r, _ := crand.Int(crand.Reader, big.NewInt(max))
	return r.Int64() + min
}
