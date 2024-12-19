package middleware

import (
	g "computer_store/internal/global"
	"computer_store/internal/handler"
	"log/slog"
	"strings"
	"time"
	"computer_store/internal/model"
	"computer_store/internal/utils/jwt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context){
		slog.Debug("[middleware JWTAuth]]  JUST Begin to use Jwt ")
		db := c.MustGet(g.CTX_DB).(*gorm.DB)

		authorization := c.Request.Header.Get("Authorization")
		if authorization == ""{
			handler.ReturnError(c,g.ErrTokenNotExist,nil)
			return
		}

		splitauth := strings.Split(authorization, " ")
		//正常来说长度是2 不是2说明被攻击
		if len(splitauth) != 2 || splitauth[0] != "Bearer"{
			handler.ReturnError(c,g.ErrTokenType,nil)
			return
		}
		token := splitauth[1]

		claims,err := jwt.ParseToken(g.Conf.JWT.Secret,token)
		if err != nil{
			handler.ReturnError(c,g.ErrTokenWrong,err)
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix(){
			handler.ReturnError(c,g.ErrTokenRuntime,nil)
			return
		}
		
		user,err := model.GetUserInfoById(db,claims.UserId)
		if err !=nil{
			handler.ReturnError(c,g.ErrDbOp,err)
			return
		}

		sess := sessions.Default(c)
		sess.Set(g.CTX_USER_AUTH,claims.UserId)
		sess.Save()

		c.Set(g.CTX_USER_AUTH,user)
	}
}