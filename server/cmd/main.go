package main

import (
	comstore "computer_store/internal"
	g "computer_store/internal/global"
	"computer_store/internal/middleware"
	"flag"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	configpath := flag.String("c", "../config.yml", "配置文件路径")
	flag.Parse()
	conf := g.ReadConfig(*configpath)
	// 初始化gin引擎
	comstore.InitLogger(conf)
	db := comstore.InitDatabase(conf)
	rdb := comstore.InitRedis(conf)

	gin.SetMode(conf.Server.Mode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"*"})
	r.Use(middleware.CORS()) // 跨域
	r.Use(middleware.WithCookieStore(conf.Session.Name, conf.Session.Salt))
	r.Use(middleware.WithGormDB(db))
	r.Use(middleware.WithRedis(rdb))
	comstore.RegisterAllHandler(r)

	serverAddr := conf.Server.Port
	if serverAddr[0] == ':' || strings.HasPrefix(serverAddr,"0.0.0.0:") { // 没有配域名
			log.Printf("Serving HTTP on (http://localhost:%s/) ... \n",strings.Split(serverAddr, ":")[1])
	}else{ // 配了域名
			log.Printf("Serving HTTP on (http//%s/)\n",serverAddr)
	}
	
	r.Run(serverAddr)
}
