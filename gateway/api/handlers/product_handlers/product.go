package producthandlers

import (
	"context"
	"log"
	"strconv"

	genprotos "github.com/ruziba3vich/armiya-gateway/genprotos"

	"github.com/gin-gonic/gin"
)

type ProductHandlers struct {
	client genprotos.ProductServiceClient
	logger *log.Logger
}

func NewProductHandlers(client genprotos.ProductServiceClient, logger *log.Logger) *ProductHandlers {
	return &ProductHandlers{
		client: client,
		logger: logger,
	}
}

// AddProduct godoc
// @Summary Add a new product
// @Description Add a new product to the catalog
// @Tags products
// @Accept json
// @Produce json
// @Param product body genprotos.AddProductRequest true "Product"
// @Success 200 {object} genprotos.AddProductResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /product/add [post]
func (h *ProductHandlers) AddProduct(ctx *gin.Context) {
	var req genprotos.AddProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.AddProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// EditProduct godoc
// @Summary Edit an existing product
// @Description Edit the details of an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param product body genprotos.EditProductRequest true "Product"
// @Success 200 {object} genprotos.EditProductResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /product/edit [put]
func (h *ProductHandlers) EditProduct(ctx *gin.Context) {
	var req genprotos.EditProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.EditProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product from the catalog
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} genprotos.Message
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /product/delete/{id} [delete]
func (h *ProductHandlers) DeleteProduct(ctx *gin.Context) {
	var req genprotos.DeleteProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.DeleteProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all products from the catalog
// @Tags products
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page size"
// @Success 200 {object} genprotos.GetAllProductsResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /products [get]
func (h *ProductHandlers) GetAllProducts(ctx *gin.Context) {
	var req genprotos.GetAllProductsRequest

	page, err := strconv.ParseUint(ctx.Query("page"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}
	limit, err := strconv.ParseUint(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}
	req.Page = page
	req.Limit = limit

	resp, err := h.client.GetAllProducts(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Retrieve a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} genprotos.GetProductResponse
// @Failure 500 {object} genprotos.Message
// @Router /products/{id} [get]
func (h *ProductHandlers) GetProduct(ctx *gin.Context) {
	var req genprotos.GetProductRequest
	req.Id = ctx.Param("id")

	resp, err := h.client.GetProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// SearchAndFilterProduct godoc
// @Summary Search and filter products
// @Description Search and filter products based on criteria
// @Tags products
// @Accept json
// @Produce json
// @Param search query string false "Search term"
// @Param filter query string false "Filter criteria"
// @Success 200 {object} genprotos.SearchAndFilterResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /products/search [get]
func (h *ProductHandlers) SearchAndFilterProduct(ctx *gin.Context) {
	var req genprotos.SearchAndFilterRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.SearchAndFilterProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// RateProduct godoc
// @Summary Rate a product
// @Description Rate a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param rating body genprotos.RateProductRequest true "Rating"
// @Success 200 {object} genprotos.RateProductResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /products/rate [post]
func (h *ProductHandlers) RateProduct(ctx *gin.Context) {
	var req genprotos.RateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.RateProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetAllRatings godoc
// @Summary Get all ratings for a product
// @Description Retrieve all ratings for a specific product
// @Tags products
// @Accept json
// @Produce json
// @Param product_id query string true "Product ID"
// @Success 200 {object} genprotos.GetAllRatingsResponse
// @Failure 500 {object} genprotos.Message
// @Router /products/ratings/{product_id} [get]
func (h *ProductHandlers) GetAllRatings(ctx *gin.Context) {
	var req genprotos.GetAllRatingsRequest
	req.ProductId = ctx.Query("product_id")

	resp, err := h.client.GetAllRatings(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// OrderProduct godoc
// @Summary Order a product
// @Description Place an order for a product
// @Tags orders
// @Accept json
// @Produce json
// @Param order body genprotos.OrderRequest true "Order"
// @Success 200 {object} genprotos.OrderResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /order [post]
func (h *ProductHandlers) OrderProduct(ctx *gin.Context) {
	var req genprotos.OrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.OrderProduct(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// CancelOrder godoc
// @Summary Cancel an order
// @Description Cancel a placed order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body genprotos.CancelOrderRequest true "Cancel Order"
// @Success 200 {object} genprotos.CancelOrderResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /order/cancel [put]
func (h *ProductHandlers) CancelOrder(ctx *gin.Context) {
	var req genprotos.CancelOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.CancelOrder(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// ChangeOrderStatus godoc
// @Summary Change order status
// @Description Update the status of an order
// @Tags orders
// @Accept json
// @Produce json
// @Param status body genprotos.ChangeOrderStatusRequest true "Order Status"
// @Success 200 {object} genprotos.ChangeOrderStatusResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /order/status [put]
func (h *ProductHandlers) ChangeOrderStatus(ctx *gin.Context) {
	var req genprotos.ChangeOrderStatusRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.ChangeOrderStatus(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Retrieve all orders
// @Tags orders
// @Accept json
// @Produce json
// @Param page query int true "Page number"
// @Param limit query int true "Page size"
// @Success 200 {object} genprotos.GetAllOrdersResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /order/all [get]
func (h *ProductHandlers) GetAllOrders(ctx *gin.Context) {
	var req genprotos.GetAllOrdersRequest

	page, err := strconv.ParseUint(ctx.Query("page"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid page parameter"})
		return
	}
	limit, err := strconv.ParseUint(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid limit parameter"})
		return
	}
	req.Page = page
	req.Limit = limit

	resp, err := h.client.GetAllOrders(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// ShowOrderInfo godoc
// @Summary Show order information
// @Description Retrieve information for a specific order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} genprotos.ShowOrderInfoResponse
// @Failure 500 {object} genprotos.Message
// @Router /order/{id} [get]
func (h *ProductHandlers) ShowOrderInfo(ctx *gin.Context) {
	var req genprotos.ShowOrderInfoRequest
	req.Id = ctx.Param("id")

	resp, err := h.client.ShowOrderInfo(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// Pay godoc
// @Summary Pay for an order
// @Description Make a payment for a specific order
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body genprotos.PayRequest true "Payment"
// @Success 200 {object} genprotos.PayResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /order/pay [post]
func (h *ProductHandlers) Pay(ctx *gin.Context) {
	var req genprotos.PayRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.Pay(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// CheckPaymentStatus godoc
// @Summary Check payment status
// @Description Check the status of a payment
// @Tags payments
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} genprotos.CheckPaymentStatusResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /order/payment/status/{order_id} [post]
func (h *ProductHandlers) CheckPaymentStatus(ctx *gin.Context) {
	var req genprotos.CheckPaymentStatusRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.CheckPaymentStatus(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// UpdateShippingDetails godoc
// @Summary Update shipping details
// @Description Update the shipping details for an order
// @Tags shipping
// @Accept json
// @Produce json
// @Param shipping body genprotos.UpdateShippingDetailsRequest true "Shipping Details"
// @Success 200 {object} genprotos.UpdateShippingDetailsResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /order/shipping [put]
func (h *ProductHandlers) UpdateShippingDetails(ctx *gin.Context) {
	var req genprotos.UpdateShippingDetailsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.UpdateShippingDetails(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// AddCategory godoc
// @Summary Add a new category
// @Description Add a new category to the product catalog
// @Tags category
// @Accept json
// @Produce json
// @Param category body genprotos.AddCategoryRequest true "Category Details"
// @Success 200 {object} genprotos.AddCategoryResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /category [post]
func (h *ProductHandlers) AddCategory(ctx *gin.Context) {
	var req genprotos.AddCategoryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.AddCategory(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetStatistics godoc
// @Summary Get sales statistics
// @Description Get total sales and revenue within a specified date range
// @Tags statistics
// @Accept json
// @Produce json
// @Param statistics body genprotos.GetStatisticsRequest true "Statistics Request"
// @Success 200 {object} genprotos.GetStatisticsResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /statistics [post]
func (h *ProductHandlers) GetStatistics(ctx *gin.Context) {
	var req genprotos.GetStatisticsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.GetStatistics(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetUserActivity godoc
// @Summary Get user activity
// @Description Get user activity details within a specified date range
// @Tags user
// @Accept json
// @Produce json
// @Param userActivity body genprotos.GetUserActivityRequest true "User Activity Request"
// @Success 200 {object} genprotos.GetUserActivityResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /user/activity [post]
func (h *ProductHandlers) GetUserActivity(ctx *gin.Context) {
	var req genprotos.GetUserActivityRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.GetUserActivity(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetArtisanRankings godoc
// @Summary Get artisan rankings
// @Description Get rankings of artisans based on category
// @Tags artisan
// @Accept json
// @Produce json
// @Param artisanRankings body genprotos.GetArtisanRankingsRequest true "Artisan Rankings Request"
// @Success 200 {object} genprotos.GetArtisanRankingsResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /artisan/rankings [post]
func (h *ProductHandlers) GetArtisanRankings(ctx *gin.Context) {
	var req genprotos.GetArtisanRankingsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.GetArtisanRankings(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}

// GetRecommendations godoc
// @Summary Get product recommendations
// @Description Get personalized product recommendations for a user
// @Tags recommendation
// @Accept json
// @Produce json
// @Param recommendations body genprotos.GetRecommendationsRequest true "Recommendations Request"
// @Success 200 {object} genprotos.GetRecommendationsResponse
// @Failure 400 {object} genprotos.Message
// @Failure 500 {object} genprotos.Message
// @Router /recommendations [post]
func (h *ProductHandlers) GetRecommendations(ctx *gin.Context) {
	var req genprotos.GetRecommendationsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.client.GetRecommendations(context.Background(), &req)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, resp)
}
