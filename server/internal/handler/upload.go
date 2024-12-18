package handler

import (
	"github.com/gin-gonic/gin"
	g "computer_store/internal/global"
	"computer_store/internal/utils/upload"
)

type Upload struct{}

func(*Upload) UploadFile(c *gin.Context){
	_,fileHeader,err := c.Request.FormFile("file")
	if err != nil{
		ReturnError(c,g.ErrFileReceive,err)
		return
	}

	// 文件存储接口
	oss := upload.NewOSS()
	filepath,_,err := oss.UploadFile(fileHeader)
	if err != nil {
		ReturnError(c,g.ErrFileUpload,err)
		return
	}

	ReturnSuccess(c,filepath)
}