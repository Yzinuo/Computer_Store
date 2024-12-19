package handler

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"computer_store/internal/model"
	g "computer_store/internal/global"

	"github.com/redis/go-redis/v9"
)

var rdbctx = context.Background()

func AddCarouselCache(rdb *redis.Client,carousel []model.CarouselVO) error {
	data,err := json.Marshal(carousel)
	if err !=nil {return errors.New("序列化失败")}

	return rdb.Set(rdbctx,g.CAROUSEL,string(data),0).Err()
}



func GetCarouselCache(rdb *redis.Client)(cache []model.CarouselVO,err error) {
	val,err := rdb.Get(rdbctx,g.CAROUSEL).Result()
	if err != nil {return nil,err}
	
	if err = json.Unmarshal([]byte(val),&cache);err!= nil{return nil,err}

	return cache,nil
}


// email
func SetMailInfo (rdb *redis.Client,info string,expire time.Duration) error{
	return rdb.Set(rdbctx,info,true,expire).Err()
}
func GetMailInfo (rdb *redis.Client,info string) (bool,error){
	return rdb.Get(rdbctx,info).Bool()
}
func DeleteMailInfo(rdb *redis.Client,info string) error{
	return rdb.Del(rdbctx,info).Err()
}