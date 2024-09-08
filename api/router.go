package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/xbmlz/baojia/middleware"
	"github.com/xbmlz/baojia/model"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Delims("@{", "}")

	r.StaticFS("/public", http.Dir("public"))
	r.LoadHTMLGlob("templates/*")

	store := gormsessions.NewStore(model.GetDB(), true, []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	apiRouter := r.Group("/api")
	{
		apiRouter.POST("/login", Login)
		apiRouter.POST("/register", Register)

		authRouter := apiRouter.Group("", middleware.LoginRequired())
		{
			authRouter.GET("/products", GetProducts)
			authRouter.GET("/user", GetUserInfo)
			authRouter.POST("/price", UpdatePrice)
		}
	}

	// r.GET("/login.html", AppLoginView)

	// adminRouter := r.Group("admin", middleware.AdminLoginRequired())
	// {
	// 	adminRouter.GET("", AdminView)
	// 	adminRouter.GET("/price", AdminPriceView)

	// 	adminRouter.GET("/api/product", GetProducts)
	// 	adminRouter.POST("/api/price", SavePrice)
	// }

	// appRouter := r.Group("", middleware.AppLoginRequired())
	// {
	// 	appRouter.GET("/", AppIndexView)
	// 	appRouter.GET("/price", AppPriceView)
	// 	appRouter.GET("/sale", AppSaleView)
	// 	appRouter.GET("/my", AppMyView)
	// }

	// apiRouter := r.Group("/api")
	// {
	// 	apiRouter.GET("/product", GetAppProducts)
	// 	apiRouter.GET("/product/picker", GetProductPicker)
	// 	apiRouter.POST("/register", AppRegister)
	// 	apiRouter.POST("/login", AppLogin)
	// 	apiRouter.POST("/logout", AppLogout)
	// }

	return r
}
