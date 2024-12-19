package handler

import (
	g "computer_store/internal/global"
	"computer_store/internal/model"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type QueryComputer struct{
	ID  int  `form:"id"`
}

type Computer struct{}

func (*Computer)GetComputerList(c *gin.Context)  {
	var query Pagequery
	if err := c.ShouldBindQuery(&query); err != nil {
		ReturnError(c,g.ErrRequest,err)
		return 
	}

	data,err := model.GetComputerList(GetDB(c),query.Page,query.Size,query.Keyword)
	if err != nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}
	ReturnSuccess(c,data)
}

func (*Computer)GetComputerDetail(c *gin.Context)  {
	var query QueryComputer
	if err := c.ShouldBindQuery(&query); err!= nil {
		slog.Info(err.Error())
		ReturnError(c,g.ErrRequest,err)
		return
	}
	data,err := model.GetComputerDetail(GetDB(c),query.ID)
	if err!= nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}

	ReturnSuccess(c,*data)
}	

func (*Computer)GetComputerinfo(c *gin.Context){
	var query QueryComputer
	if err := c.ShouldBindQuery(&query); err!= nil {
		ReturnError(c,g.ErrRequest,err)
		return
	}
	data,err := model.GetComputerInfo(GetDB(c),query.ID)
	if err!= nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}

	ReturnSuccess(c,*data)
}


func (*Computer)GetComputerIntro(c *gin.Context){
	var query QueryComputer
	if err := c.ShouldBindQuery(&query); err!= nil {
		ReturnError(c,g.ErrRequest,err)
		return
	}
	data,err := model.GetComputerImage(GetDB(c),query.ID)
	if err!= nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}

	ReturnSuccess(c,data)
}
