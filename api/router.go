package api

import (
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/middleware"
	"github.com/xbmlz/baojia/model"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	store := gormsessions.NewStore(model.GetDB(), true, []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/login", Login)
		apiRouter.POST("/register", Register)
		apiRouter.POST("/upload", UploadFile)

		authRouter := apiRouter.Group("", middleware.JwtAuthRequired())
		{
			authRouter.GET("/products", GetProducts)
			authRouter.GET("/user", GetUserInfo)
			authRouter.POST("/price", UpdatePrice)
			authRouter.POST("/sale", CreateSale)
			authRouter.GET("/sales", GetSales)
		}
	}
	return r
}
