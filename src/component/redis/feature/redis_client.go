package feature

import (
	"github.com/garyburd/redigo/redis"
	"github.com/lixinchan/go-modules/component/redis/lib/log"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const (
	// keyDelimiter ,
	keyDelimiter string = ","
	// valDelimiter value and expire time delimiter |
	valDelimiter string = "|"
)

// Client redis client
type Client struct {
	Tag  string
	Pool *redis.Pool
}

var (
	clients = map[string]*Client{}
)

// PoolConf redis pool config
type PoolConf struct {
	Host           string
	Port           string
	Auth           string
	InitDb         int
	Wait           bool
	InitNum        int
	MaxNum         int
	ConnectTimeout int //ms
	ReadTimeout    int //ms
	WriteTimeout   int //ms
	IdleTimeout    int //ms
}

// Feature feature
type Feature struct {
	Field  string `redis:"field"`
	Val    string `redis:"val"`
	Expire int64  `redis:"expire"`
}

// init client
func initClient(tag string, pool *redis.Pool) *Client {
	return &Client{Tag: tag, Pool: pool}
}

// NewClient new redis client
func NewClient(tag string, conf *PoolConf) (client *Client, err error) {
	if redisPool, err := InitRedis(conf); err != nil {
		client = initClient(tag, redisPool)
		clients[client.Tag] = client
		return client, nil
	}
	return
}

// InitRedis init redis
func InitRedis(conf *PoolConf) (pool *redis.Pool, err error) {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: time.Second * 60,
		Wait:        true,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				log.Warn(err)
			}
			return err
		},
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp", conf.Host+":"+conf.Port,
				redis.DialConnectTimeout(time.Duration(5000)*time.Millisecond),
				redis.DialReadTimeout(time.Duration(5000)*time.Millisecond),
				redis.DialWriteTimeout(time.Duration(5000)*time.Millisecond))
			if err != nil {
				log.Warn(err)
				return nil, err
			}
			if len(conf.Auth) > 0 {
				if _, err = c.Do("AUTH", conf.Auth); err != nil {
					_ = c.Close()
					log.Warn(err)
					return nil, err
				}
			}
			if _, err = c.Do("SELECT", conf.InitDb); err != nil {
				log.Warn(err)
			}
			return c, err
		},
	}
	return
}

func (c *Client) getConn() redis.Conn {
	return c.Pool.Get()
}

// Hmset hmset
func (c *Client) Hmset(key string, features []Feature, expireTime int) (result string, err error) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close()
	}()
	filedVal := make(map[string]string)
	var f *Feature
	for _, fe := range features {
		f = &fe
		filedVal[f.Field] = f.Val + valDelimiter + strconv.FormatInt(f.Expire, 10)
	}
	ret, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(filedVal)...)
	_, _ = conn.Do("expire", key, expireTime)
	if err != nil {
		return redis.String(ret, err)
	}
	return
}

// Hmget hmget
func (c *Client) Hmget(keys, fields []string) (result map[string][]string) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close()
	}()
	queryKey := strings.Join(keys, keyDelimiter)
	ret, err := conn.Do("HMGET", redis.Args{}.Add(queryKey).AddFlat(fields)...)
	if err == nil {
		if res, err := redis.Strings(ret, err); err == nil {
			result = make(map[string][]string)
			var idx = int32(0)
			for _, val := range keys {
				result[val] = res[idx:len(fields)]
				atomic.AddInt32(&idx, int32(len(fields)))
			}
			return result
		}
	}
	return
}

// Hget redis hget
func (c *Client) Hget(key string, fields []string) (result map[string][]string) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close()
	}()
	key = strings.TrimSpace(key)
	ret, err := conn.Do("HGET", redis.Args{}.Add(key).AddFlat(fields)...)
	if res, err := redis.Strings(ret, err); err == nil {
		result = make(map[string][]string)
		result[key] = res
	}
	return
}

// Hdel redis hdel
func (c *Client) Hdel(key string, fields []string) (result int64) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close()
	}()
	key = strings.TrimSpace(key)
	ret, err := conn.Do("HDEL", redis.Args{}.Add(key).AddFlat(fields)...)
	if res, err := redis.Int64(ret, err); err == nil {
		result = res
	}
	return
}

// Hclear redis hclear
func (c *Client) Hclear(key string) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close()
	}()
	key = strings.TrimSpace(key)
	_, _ = conn.Do("HCLEAR", key)
}

// Del del key
func (c *Client) Del(key string) (result bool, err error) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close()
	}()
	if ret, err := conn.Do("del", key); err == nil {
		result, err = redis.Bool(ret, err)
	}
	return
}

// Expire set expire
func (c *Client) Expire(key string, expireTime int64) (result int64, err error) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close()
	}()
	if ret, err := conn.Do("expire", key, expireTime); err == nil {
		result, err = redis.Int64(ret, err)
	}
	return
}

// Ping ping
func (c *Client) Ping() error {
	conn := c.getConn()
	defer func() {
		_ = conn.Close
	}()
	_, err := conn.Do("ping")
	return err
}

// Scan scan
func (c *Client) Scan(cursor, match string, count int64) (newCursor string, keys []string, err error) {
	var (
		res interface{}
	)
	args := []interface{}{cursor}
	if len(match) > 0 {
		args = append(args, "match", match)
	}
	args = append(args, "count", count)
	conn := c.getConn()
	defer func() {
		_ = conn.Close
	}()
	if res, err = conn.Do("scan", args...); err == nil {
		if v, ok := res.([]interface{}); ok {
			if len(v) > 0 {
				newCursor, err = redis.String(v[0], err)
			}
			if err == nil && len(v) > 1 {
				keys, err = redis.Strings(v[1], err)
			}
		}
	}
	return
}

// TTL ttl
func (c *Client) TTL(key string) (res int64, err error) {
	conn := c.getConn()
	defer func() {
		_ = conn.Close
	}()
	r, err := conn.Do("ttl", key)
	if err == nil {
		res, err = redis.Int64(r, err)
	}
	return
}
