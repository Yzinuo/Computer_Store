package middleware

import (
	g "computer_store/internal/global"
	"time"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

//把数据库的连接放到context中，方便后续使用
func WithGormDB(db *gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		c.Set(g.CTX_DB, db)
		c.Next()
	}
}

func WithRedis(rdb *redis.Client) gin.HandlerFunc{
	return func(c *gin.Context){
		c.Set(g.CTX_RDB, rdb)
		c.Next()
	}
}

//session 放入cookie中
func WithCookieStore(name,secret string) gin.HandlerFunc{
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{Path: "/", MaxAge:600})// 指定session的作用域
	return sessions.Sessions(name, store)
}

// 完成跨域
func CORS() gin.HandlerFunc{
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET","POST","PUT","DELETE","OPTION","PATCH"},
		AllowHeaders: []string{"Origin","Authorization","Content-Type","X-Requested-With"},
		ExposeHeaders: []string{"Content-Type"},
		AllowCredentials: true, // 允许发送cookie
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           24 * time.Hour, // 缓存预检请求，提高服务性能
	})
	
}