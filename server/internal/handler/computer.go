package handler

import(
	"github.com/gin-gonic/gin"
	g "computer_store/internal/global"
)

type QueryComputer struct{
	Id  int  `json:"id" form:"id"`
}

func GetComputerList(c *gin.Context)  {
	var query Pagequery
	if err := c.ShouldBindJSON(&query); err != nil {
		ReturnError(c,g.ErrRequest,err)
		return 
	}

	data,err := model.GetFrontComputerList(GetDB(c),query.Page,query.Size,query.Keyword)
	if err != nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}
	ReturnSuccess(c,data)
}

func GetComputerDetail(c *gin.Context)  {
	var query QueryComputer
	if err := c.ShouldBindJSON(&query); err!= nil {
		ReturnError(c,g.ErrRequest,err)
		return
	}
	data,err := model.GetComputerDetail(GetDB(c),query.Id)
	if err!= nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}

	ReturnSuccess(c,data)
}	

func GetComputerinfo(c *gin.Context){
	var query QueryComputer
	if err := c.ShouldBindJSON(&query); err!= nil {
		ReturnError(c,g.ErrRequest,err)
		return
	}
	data,err := model.GetComputerInfo(GetDB(c),query.Id)
	if err!= nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}

	ReturnSuccess(c,data)
}


func GetComputerIntro(c *gin.Context){
	var query QueryComputer
	if err := c.ShouldBindJSON(&query); err!= nil {
		ReturnError(c,g.ErrRequest,err)
		return
	}
	data,err := model.GetComputerintro(GetDB(c),query.Id)
	if err!= nil {
		ReturnError(c,g.ErrDbOp,err)
		return
	}

	ReturnSuccess(c,data)
}
