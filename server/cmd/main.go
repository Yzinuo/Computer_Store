package main

import (
	comstore "computer_store/internal"
	g "computer_store/internal/global"
	"computer_store/internal/middleware"
	"flag"
	"github.com/gin-gonic/gin"
)

func main() {
	configpath := flag.String("c", "../cinfig.yml", "配置文件路径")
	flag.Parse()
	conf := g.ReadConfig(*configpath)
	g.Test()
	// 初始化gin引擎
	comstore.InitLogger(conf)
	db := comstore.InitDatabase(conf)
	rdb := comstore.InitRedis(conf)

	gin.SetMode(conf.Server.Mode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"*"})
	r.Use(middleware.CORS()) // 跨域
	r.Use(middleware.WithCokkieStore(conf.Session.Name, conf.Session.Salt))
	r.Use(middleware.WithGormDB(db))
	r.Use(middleware.WithRedis(rdb))
}
