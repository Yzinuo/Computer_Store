package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	g "computer_store/internal/global"
)

type Front struct{}

func (*Front) GetCarouselList(c *gin.Context) {
	data, err := GetCarouselCache(GetRDB(c))
	if data != nil && err == nil {
		ReturnSuccess(c,data)
		return
	}

	switch err{
	case redis.Nil:
		break
	default:
		 ReturnError(c,g.ErrRedisOp,err)
		 return
	}

	data,err  = model.GetCarousel()
	if err != nil{
		ReturnError(c,g.ErrDbOp,err)
		return
	}
	
	if err = AddCarouselCache(GetRDB(c),data);err!= nil{
		ReturnError(c,g.ErrRedisOp,err)
		return
	}

	ReturnSuccess(c,data)
}

