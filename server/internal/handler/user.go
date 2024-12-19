package handler

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	g "computer_store/internal/global"
	"computer_store/internal/utils"
	"computer_store/internal/utils/jwt"

	"computer_store/internal/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct{}

type Loginreq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginVO struct {
	model.User
	Token string `json:"token"`
}
type Registereq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateUserReq struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
}

func (*User) Login(c *gin.Context) {
	var loginreq Loginreq
	if err := c.ShouldBindJSON(&loginreq); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	auth, err := model.GetUserInfoByEmail(GetDB(c), loginreq.Email) // 获取用户所有信息
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ReturnError(c, g.ErrUserNotExist, err)
			return
		}
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	if !utils.BcryptCheck(loginreq.Password, auth.Password) {
		ReturnError(c, g.ErrPassword, err)
		return
	}

	conf := g.Conf.JWT
	token, err := jwt.GenToken(conf.Secret, conf.Issuer, int(conf.Expire), auth.ID)
	if err != nil {
		ReturnError(c, g.ErrTokenCreate, err)
		return
	}

	slog.Info("用户登录成功") //保存登录状态10分钟
	session := sessions.Default(c)
	session.Set(g.CTX_USER_AUTH, auth.ID)
	session.Save()

	ReturnSuccess(c, LoginVO{
		User:  *auth,
		Token: token,
	})
}

func (*User) Logout(c *gin.Context) {
	c.Set(g.CTX_USER_AUTH, nil)

	session := sessions.Default(c)
	session.Delete(g.CTX_USER_AUTH)
	session.Save()

	ReturnSuccess(c, nil)
}

func (*User) Register(c *gin.Context) {
	var regreq Registereq
	if err := c.ShouldBindJSON(&regreq); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}
	regreq.Email = utils.Format(regreq.Email)

	// 检查用户名是否存在，避免重复注册
	auth, err := model.GetUserInfoByEmail(GetDB(c), regreq.Email)
	if err != nil {
		var flag bool = false
		if errors.Is(err, gorm.ErrRecordNotFound) {
			flag = true
		}
		if !flag {
			ReturnError(c, g.ErrDbOp, err)
			return
		}
	}

	if auth != nil {
		ReturnError(c, g.ErrUserExist, err)
		return
	}

	// 通过邮箱验证
	info := utils.GenEmailVerificationInfo(regreq.Email, regreq.Password)
	SetMailInfo(GetRDB(c), info, 15*time.Minute) // 15分钟过期
	EmailData := utils.GetEmailData(regreq.Email, info)
	err = utils.SendEmail(regreq.Email, EmailData)
	if err != nil {
		ReturnError(c, g.ErrSendEmail, err)
		return
	}

	ReturnSuccess(c, nil)
}

func (*User) VerifyCode(c *gin.Context) {
	var code string
	if code = c.Query("info"); code == "" {
		returnErrorPage(c)
		return
	}

	// 验证是否有code在数据库中
	ifExist, err := GetMailInfo(GetRDB(c), code)
	if err != nil {
		returnErrorPage(c)
		return
	}
	if !ifExist {
		returnErrorPage(c)
		return
	}

	DeleteMailInfo(GetRDB(c), code)

	username, password, err := utils.ParseEmailVerificationInfo(code)
	if err != nil {
		returnErrorPage(c)
		return
	}

	// 注册用户
	err = model.CreateNewUser(GetDB(c), username, password)
	if err != nil {
		returnErrorPage(c)
		return
	}

	// 注册成功，返回成功页面
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
        <!DOCTYPE html>
        <html lang="zh-CN">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>注册成功</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f4f4f4;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    margin: 0;
                }
                .container {
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 8px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                    text-align: center;
                }
                h1 {
                    color: #5cb85c;
                }
                p {
                    color: #333;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>注册成功</h1>
                <p>恭喜您，注册成功！</p>
            </div>
        </body>
        </html>
    `))
}

func returnErrorPage(c *gin.Context) {
	c.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte(`
        <!DOCTYPE html>
        <html lang="zh-CN">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>注册失败</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f4f4f4;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    margin: 0;
                }
                .container {
                    background-color: #fff;
                    padding: 20px;
                    border-radius: 8px;
                    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
                    text-align: center;
                }
                h1 {
                    color: #d9534f;
                }
                p {
                    color: #333;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>注册失败</h1>
                <p>请重试。</p>
            </div>
        </body>
        </html>
    `))
}

func (*User) GetUserInfo(c *gin.Context) {
	userinfo, err := CurrentUserAuth(c)
	if err != nil {
		ReturnError(c, g.ErrTokenNotExist, err)
		return
	}

	ReturnSuccess(c, userinfo)
}

func (*User) UpdateUserInfo(c *gin.Context) {
	var req UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	auth, err := CurrentUserAuth(c)
	if err != nil {
		ReturnError(c, g.ErrTokenNotExist, err)
		return
	}

	userinfo,err := model.UpdateUserInfo(GetDB(c), auth.ID, req.Nickname, req.Email, req.Avatar, req.Intro)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	ReturnSuccess(c, userinfo)
}
