package comstore

import (
	"computer_store/internal/handler"
	"computer_store/internal/middleware"

	"github.com/gin-gonic/gin"
)

var (
	frontAPI    handler.Front
	userAPI     handler.User
	ComputerAPI handler.Computer
	UploadAPI   handler.Upload
)

func RegisterAllHandler(r *gin.Engine){
	RegisterBaseHandler(r)
	RegisterFrontHandler(r)
}

func RegisterBaseHandler(r *gin.Engine){
	base := r.Group("/api")
	base.POST("/login",userAPI.Login)
	base.POST("/register",userAPI.Register)
	base.GET("/email/verify",userAPI.VerifyCode)
	base.GET("/logout",userAPI.Logout)
	base.GET("/product",ComputerAPI.GetComputerinfo)
	base.GET("/product-details",ComputerAPI.GetComputerDetail)
	base.GET("/product-images",ComputerAPI.GetComputerIntro)
}

func RegisterFrontHandler(r *gin.Engine){
	front := r.Group("/api/front")
	front.GET("/carousel",frontAPI.GetCarouselList)
	front.GET("/products",ComputerAPI.GetComputerList)

	front.Use(middleware.JWTAuth())
	{
		front.GET("/user",userAPI.GetUserInfo)
		front.POST("/upload",UploadAPI.UploadFile)
		front.PUT("/userinfo",userAPI.UpdateUserInfo)
	}
}