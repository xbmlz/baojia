package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/middleware"
	"github.com/xbmlz/baojia/model"
	"github.com/xbmlz/baojia/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest = RegisterRequest

func Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// check username is exist
	_, exist, err := model.GetUserByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if exist {
		ctx.JSON(http.StatusConflict, gin.H{
			"code": 5001,
			"msg":  "用户名已存在",
		})
		return
	}

	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := model.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code": 0,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, exist, err := model.GetUserByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 4011,
			"msg":  "用户不存在",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 4012,
			"msg":  "密码错误",
		})
		return
	}

	// session := sessions.Default(ctx)
	// session.Set(middleware.SessionUserKey, user.ID)
	// session.Save()

	token, err := token.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": token,
	})
}

func Logout(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// session.Delete(middleware.SessionUserKey)
	// session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "退出成功",
	})
}

func GetUserInfo(ctx *gin.Context) {
	// session := sessions.Default(ctx)
	// userID := session.Get(middleware.SessionUserKey)
	user, err := getLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 4013,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": user,
	})
}

func getLoginUser(ctx *gin.Context) (*model.User, error) {
	userID := ctx.GetInt(middleware.CurrentUserKey)
	if userID == 0 {
		return nil, errors.New("用户未登录")
	}
	user, err := model.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
