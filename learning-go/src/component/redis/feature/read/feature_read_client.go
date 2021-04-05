package main

import (
	"fmt"
	"github.com/lixinchan/go-modules/component/redis/feature"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
	"sync/atomic"
	"time"
)

var (
	handleConcurrent = int64(3)
	handleJobChan    chan bool
	scanTimes        = int64(0)
	scanCount        = int64(0)
	prefixMatchCount = int64(0)
	handledCount     = int64(0)
	deletedCount     = int64(0)
	scanRedisClient  *feature.Client
	redisClient      *feature.Client
	twenty           = []string{"c31", "c30", "c11", "c33", "Su", "c10", "c32", "Sv", "c13", "c12", "c34", "c15", "c14", "c17", "c16", "c19", "c18", "c20", "c22", "c21", "c24", "c23"}
	fifty            = []string{"c31", "c30", "c11", "c33", "Su", "c10", "c32", "Sv", "c13", "c12", "c34", "c15", "c14", "c17", "c16", "c19", "c18", "c20", "c22", "c21", "c24", "c23", "c26", "c25", "c28", "c27", "c29", "c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "c31", "c30", "c11", "c33", "Su", "c10", "c32", "Sv", "c13", "c12", "c34", "c15", "c14", "c17", "c16", "c19", "c18", "c20"}
	hundred          = []string{"c31", "c30", "c11", "c33", "Su", "c10", "c32", "Sv", "c13", "c12", "c34", "c15", "c14", "c17", "c16", "c19", "c18", "c20", "c22", "c21", "c24", "c23", "c26", "c25", "c28", "c27", "c29", "c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "c31", "c30", "c11", "c33", "Su", "c10", "c32", "Sv", "c13", "c12", "c34", "c15", "c14", "c17", "c16", "c19", "c18", "c20", "c22", "c21", "c24", "c23", "c26", "c25", "c28", "c27", "c29", "c1", "c2", "c3", "c4", "c5", "c6", "c7", "c8", "c9", "c31", "c30", "c11", "c33", "Su", "c10", "c32", "Sv", "c13", "c12", "c34", "c15", "c14", "c17", "c16", "c19", "c18", "c20", "c22", "c21", "c24", "c23", "c26", "c25", "c28", "c27", "c29", "c1", "c2", "c3", "c4", "c5", "c6", "c7"}

	queryCnt = int64(0)
)

//HandlerParams param
type HandlerParams struct {
	prefix          string
	minTTL          int64
	maxTTL          int64
	handle          string
	handleFunc      func(*HandlerParams, []string)
	debug           bool
	redisConfig     *feature.PoolConf
	redisDistConfig *feature.PoolConf
	invertMatch     bool
	fieldCount      int
}

func main() {
	app := cli.NewApp()
	app.Name = "feature-read-client"

	var host string
	var hostDist string
	var port string
	var portDist string
	var auth string
	var authDist string
	var db int
	var dbDist string
	var count int64
	var cursor string
	var prefix string
	var minTTL int64
	var maxTTL int64
	var handle string
	var debug bool
	var concurrent int
	var invertMatch bool
	var fieldCount int
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
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
			Destination: &db,
		},
		cli.Int64Flag{
			Name:        "count, c",
			Value:       200,
			Usage:       "scan count",
			Destination: &count,
		},
		cli.Int64Flag{
			Name:        "minttl",
			Value:       -1,
			Usage:       "删除key时需满足的最小ttl, 默认为 -1, 表示忽略此条件",
			Destination: &minTTL,
		},
		cli.Int64Flag{
			Name:        "maxttl",
			Value:       -1,
			Usage:       "删除key时需满足的最大ttl, 默认为 -1",
			Destination: &maxTTL,
		},
		cli.IntFlag{
			Name:        "concurrent",
			Value:       3,
			Usage:       "删除key时需满足的最大ttl, 默认为 -1, 表示忽略此条件",
			Destination: &concurrent,
		},
		cli.StringFlag{
			Name:        "cursor",
			Value:       "0",
			Usage:       "scan cursor",
			Destination: &cursor,
		},
		cli.StringFlag{
			Name:        "prefix",
			Value:       "",
			Usage:       "scan prefix",
			Destination: &prefix,
		},
		cli.StringFlag{
			Name:        "handle",
			Value:       "nil",
			Usage:       "scan and handle",
			Destination: &handle,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug 模式, 不删除任何数据",
			Destination: &debug,
		},
		cli.BoolFlag{
			Name:        "invert-match",
			Usage:       "过滤不匹配的prefix, 与 --prefix 配合使用",
			Destination: &invertMatch,
		},
		cli.StringFlag{
			Name:        "hostdist",
			Value:       "127.0.0.1",
			Usage:       "复制目标redis host",
			Destination: &hostDist,
		},
		cli.StringFlag{
			Name:        "portdist",
			Value:       "6379",
			Usage:       "复制目标redis port",
			Destination: &portDist,
		},
		cli.StringFlag{
			Name:        "authdist",
			Value:       "",
			Usage:       "复制目标redis auth",
			Destination: &authDist,
		},
		cli.StringFlag{
			Name:        "dbdist",
			Value:       "0",
			Usage:       "复制目标redis db",
			Destination: &dbDist,
		},
		cli.IntFlag{
			Name:        "fieldCount",
			Value:       20,
			Usage:       "特征数量",
			Destination: &fieldCount,
		},
	}

	app.Action = func(c *cli.Context) error {
		handleConcurrent = int64(concurrent)
		handleJobChan = make(chan bool, handleConcurrent)

		redisConfig := &feature.PoolConf{
			Host:    host,
			Port:    port,
			InitDb:  db,
			Auth:    auth,
			InitNum: concurrent,
			MaxNum:  concurrent,
		}
		var handlerFunc func(*HandlerParams, []string)
		handlerFunc = nil
		switch handle {
		case "hmget":
			//统计前缀key
			handlerFunc = hmgetHandler
		case "delete":
			//删除满足前缀匹配的key (支持ttl范围)
			handlerFunc = deleteHandler
		default:
			handlerFunc = nil
		}
		handlerParams := &HandlerParams{
			prefix:      prefix,
			minTTL:      minTTL,
			maxTTL:      maxTTL,
			debug:       debug,
			handle:      handle,
			handleFunc:  handlerFunc,
			redisConfig: redisConfig,
			invertMatch: invertMatch,
			fieldCount:  fieldCount,
		}
		err := redisScan(redisConfig, cursor, handlerParams, count, handlerFunc)
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func redisScan(conf *feature.PoolConf, cursor string, handlerParams *HandlerParams, count int64, handler func(*HandlerParams, []string)) (err error) {
	redisClient, err = feature.NewClient("default", conf)
	err = redisClient.Ping()
	if err != nil {
		fmt.Println("init redis client failed, err:", err)
		return
	}
	scanRedisClient, err = feature.NewClient("scan", conf)
	err = scanRedisClient.Ping()
	if err != nil {
		fmt.Println("init redis client failed, err:", err)
		return
	}
	keyListChannel := make(chan []string)
	go func() {
		for {
			newCursor, keyList, err := scanRedisClient.Scan(cursor, "", count)
			if err != nil {
				fmt.Println("scan err:", err)
				continue
			}
			if handler != nil {
				keyListChannel <- keyList
			}
			scanTimes++
			scanCount += int64(len(keyList))
			if scanTimes%100 == 0 {
				printlnInfo(newCursor, scanTimes, scanCount, prefixMatchCount, deletedCount, handledCount)
			}
			cursor = newCursor
			if newCursor == "0" {
				printlnInfo(cursor, scanTimes, scanCount, prefixMatchCount, deletedCount, handledCount)
				os.Exit(0)
			}
		}
	}()
	for keyList := range keyListChannel {
		if handler != nil {
			handleJobChan <- true
			go func() {
				defer func() {
					<-handleJobChan
				}()
				handler(handlerParams, keyList)
			}()
		}
	}
	time.Sleep(time.Duration(10 * time.Second))
	return
}

func hmgetHandler(params *HandlerParams, keyList []string) {
	if len(keyList) == 0 {
		return
	}
	var fields []string
	count := params.fieldCount
	switch count {
	case 20:
		fields = twenty
	case 50:
		fields = fifty
	case 100:
		fields = hundred
	default:
		fields = twenty
	}
	_ = redisClient.Hmget(keyList, fields)
}

func deleteHandler(params *HandlerParams, keyList []string) {
	if len(keyList) == 0 {
		return
	}
	var err error
	for _, key := range keyList {
		if prefixMatch(key, params.prefix, params.invertMatch) {
			atomic.AddInt64(&prefixMatchCount, 1)
			ttl := int64(-1)
			if params.minTTL > 0 || params.maxTTL > 0 {
				ttl, err = redisClient.TTL(key)
				if err != nil {
					fmt.Printf("ttl err: %v\n", err)
					continue
				}
				if ttl == -1 {
					fmt.Printf("no ttl: %s\n", key)
					continue
				}
				if params.minTTL > 0 && ttl >= params.minTTL {
					continue
				}
				if params.maxTTL > 0 && ttl <= params.maxTTL {
					continue
				}
			}
			if !params.debug {
				_, _ = redisClient.Del(key)
			} else {
				if ttl > 0 {
					fmt.Printf("[debug] key:%s;ttl:%d\n", key, ttl)
				} else {
					fmt.Printf("[debug] key:%s\n", key)
				}
			}
			atomic.AddInt64(&deletedCount, 1)
		}
	}
}

func prefixMatch(key, prefix string, invertMatch bool) bool {
	if invertMatch {
		return !strings.HasPrefix(key, prefix)
	}
	return len(prefix) == 0 || strings.HasPrefix(key, prefix)
}

func printlnInfo(cursor string, scanTimes, scanCount, prefixMatchCount, deletedCount, handledCount int64) {
	info := fmt.Sprintf("%s %s scan num:%d, key-num: %d, prefix-num: %d", time.Now().Format("15:04:05"),
		cursor, scanTimes, scanCount, prefixMatchCount)
	if deletedCount > 0 {
		info = fmt.Sprintf("%s, del-num: %d", info, deletedCount)
	}
	if handledCount > 0 {
		info = fmt.Sprintf("%s, handled-num: %d", info, handledCount)
	}
	fmt.Println(info)
}
