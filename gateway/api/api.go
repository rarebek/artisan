package api

import (
	"log"

	"github.com/gin-gonic/gin"
	authhandler "github.com/ruziba3vich/armiya-gateway/api/handlers/auth_handlers"
	producthandler "github.com/ruziba3vich/armiya-gateway/api/handlers/product_handlers"
	"github.com/ruziba3vich/armiya-gateway/config"
	_ "github.com/ruziba3vich/armiya-gateway/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	logger         *log.Logger
	cfg            *config.Config
	authhandler    *authhandler.AuthHandlers
	producthandler *producthandler.ProductHandlers
}

func New(
	cfg *config.Config,
	logger *log.Logger,
	authhandler *authhandler.AuthHandlers,
	producthandler *producthandler.ProductHandlers) *API {
	return &API{
		logger:         logger,
		cfg:            cfg,
		authhandler:    authhandler,
		producthandler: producthandler,
	}
}

// @title API
// @version 1.0
// @description TEST
// @host localhost:9090
// @BasePath /api/v1
func (a *API) RUN() error {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		api.POST("/auth/register", a.authhandler.Register)         //+
		api.POST("/auth/login", a.authhandler.Login)               //------- TODO: jwt token
		api.GET("/auth/profile/:id", a.authhandler.ShowProfile)    //+
		api.PUT("/auth/profile/edit", a.authhandler.EditProfile)   //+
		api.PUT("/auth/usertype/edit", a.authhandler.EditUserType) //+
		api.GET("/auth/users", a.authhandler.GetAllUsers)          //+
		api.DELETE("/auth/delete/:id", a.authhandler.DeleteUser)   //+
		api.POST("/auth/reset", a.authhandler.ResetPassword)       //+

		api.POST("/product/add", a.producthandler.AddProduct)
		api.PUT("/product/edit", a.producthandler.EditProduct)
		api.DELETE("/product/delete/:id", a.producthandler.DeleteProduct)
		api.GET("/products", a.producthandler.GetAllProducts)
		api.GET("/product/:id", a.producthandler.GetProduct)
		api.POST("/product/search", a.producthandler.SearchAndFilterProduct)
		api.POST("/product/rate", a.producthandler.RateProduct)
		api.GET("/product/ratings/:product_id", a.producthandler.GetAllRatings)
		api.POST("/order", a.producthandler.OrderProduct)
		api.PUT("/order/cancel", a.producthandler.CancelOrder)
		api.PUT("/order/status", a.producthandler.ChangeOrderStatus)
		api.GET("/order/all", a.producthandler.GetAllOrders)
		api.GET("/order/:id", a.producthandler.ShowOrderInfo)
		api.POST("/order/pay", a.producthandler.Pay)
		api.GET("/order/payment/status/:order_id", a.producthandler.CheckPaymentStatus)
		api.PUT("/order/shipping", a.producthandler.UpdateShippingDetails)

		api.POST("/category", a.producthandler.AddCategory)
	}

	return router.Run(a.cfg.ServerAddress)
}
