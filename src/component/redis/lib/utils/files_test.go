package utils

import (
	"anygo/lib/utils"
	"math/rand"
	"testing"
	"time"
)

func TestForEachLineConcurrent(t *testing.T) {
	path := "/Users/clx/Documents/code/inc/golang/feature-redis-client/src/data/data.json"
	_ = utils.ForEachLineConcurrent(path, func(line string) error {
		sleppTime := time.Duration(rand.Int63n(int64(2000))) * time.Microsecond
		time.Sleep(sleppTime)
		println(line)
		//fmt.Println(sleppTime)
		return nil
	}, 10)

	//_ = ForEachLine("./files_test.txt", func(line string) error {
	//	println(line)
	//	//sleppTime := time.Duration(rand.Int63n(int64(20000))) * time.Microsecond
	//	time.Sleep(100000000)
	//	//fmt.Println(sleppTime)
	//	return nil
	//})
	println("OK")
}
