package handler

import (
	g "computer_store/internal/global"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

//查询结构体 
type Pagequery struct{
	Page   int  `form:"page_num"`
	Size   int  `form:"page_size"`
	Keyword string `form:"keyword"`
}

// HTTP code + code + msg + data
func ReturnHttpResponse(c *gin.Context, httpcode int, code int, msg string, data any) {
	c.JSON(httpcode, Response[any]{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// Result + data
func ReturnResponse(c *gin.Context, r g.Result, data any) {
	ReturnHttpResponse(c, http.StatusOK, r.Code, r.Msg, data)
}

// data
func ReturnSuccess(c *gin.Context, data any) {
	ReturnResponse(c, g.OkReresult, data)
}

// 预料中的错误  = 业务错误 + 系统错误  在业务层处理，返回200 http状态码
// 意外的错误  = 触发panic， 在中间件中被捕获，返回500 http状态码
// data是错误数据（可以是error和string）， error是业务错误
func ReturnError(c *gin.Context, r g.Result, data any) {
	slog.Info("[FUNC-RETURN-ERROR]] :" + r.Msg)

	var val string = r.Msg

	if data != nil {
		switch v := data.(type) {
		case error:
			val = v.Error()

		case string:
			val = v
		}
		slog.Error(val)
	}

	c.AbortWithStatusJSON(
		http.StatusOK,
		Response[any]{
			Code:    r.Code,
			Message: r.Msg,
			Data:    val,
		},
	)
}

func GetDB(c *gin.Context) *gorm.DB{
	return c.MustGet(g.CTX_DB).(*gorm.DB)
}

func GetRDB(c *gin.Context) *redis.Client{
	return c.MustGet(g.CTX_RDB).(*redis.Client)
}

// 1. 从Context中获取UserInfo，如果获取到说明context中已经保存了
// 2. 从Session中获取UserInfo Uid(通过解析token获取)
// 3. 根据id从数据库中获取  然后在设置在gin context中
func CurrentUserAuth(c *gin.Context)(*model.UserInfo,error){
	var key string = g.CTX_USER_AUTH
	if cache,exit := c.Get(key);exit && cache != nil {
		slog.Debug("[FUNC-GET-USER-AUTH]: 从gin context中获取")
		return cache.(*model.UserInfo),nil
	}

	//2.
	s := sessions.Default(c)
	id := s.Get(key)
	
	// 3.
	if id != nil{
		db := GetDB(c)
		userinfo,err := model.GetUserInfo(db,id.(int))

		if err!= nil {
			return nil,err
		}

		c.Set(key,userinfo)
		return userinfo,nil
	}

	return nil,errors.New("session中没有 user_id")
}

