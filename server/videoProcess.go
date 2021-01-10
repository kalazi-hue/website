package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

var redisdb *redis.Client
func xmain() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "104.225.150.234:6379", // use default Addr
		Password: "yfxe8tixjHi4PBsY",               // no password set
		DB:       0,                // use default DB
	})
	//for {
		res, err := redisdb.RPop("video-process").Result()
		if err != nil {
			fmt.Println("暂无任务")
			time.Sleep(time.Second * 10)
		}
		info := strings.Split(res, ",")//https://xx.com/finish/movie/upload/2021-01-09/c20ad4d76fe97759aa27a0c99bff6710/origin/ori.mp4
		fmt.Println(info[1])
		re := compressVideo(info[1])
		fmt.Println(re)
	//}

}

/**
	path : https://xx.com/finish/movie/upload/2021-01-09/c20ad4d76fe97759aa27a0c99bff6710/origin/ori.mp4
 */
func compressVideo(path string) string {
	basePath := "/www/wwwroot/www.8dafu.com/server"
	//basePath := "/Users/madman/Downloads/gin-vue-admin-master-new/server"
	domain, orgPath := getDownPath(path)// https://xx.com , /finish/movie/upload/2021-01-09/c20ad4d76fe97759aa27a0c99bff6710/origin/ori.mp4
	tmp := strings.Split(orgPath, "origin")//
	compressPath := basePath + tmp[0] + "compress/"
	fmt.Println(compressPath)
	compressWithFile := compressPath + "compress.mp4"
	fmt.Println(compressWithFile)

	err := os.Mkdir(compressPath, 0777)
	if err != nil {
		fmt.Println("mkdir fail:" + err.Error())
	}

	cmd := fmt.Sprintf(" -i %s -vcodec libx264 -crf 20 %s", basePath + orgPath, compressPath + "compress.mp4")
	c := exec.Command("/usr/local/bin/ffmpeg", cmd)
	err = c.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	return domain + compressWithFile
}


func getDownPath(downUrl string) (string, string) {
	res,_ := url.Parse(downUrl)
	return res.Host,res.Path
}

