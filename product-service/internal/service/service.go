package service

import (
	"armiya/equipment-service/genprotos"
	"armiya/equipment-service/internal/storage"
	"context"
	"log"
	"os"
)

type (
	Productservice struct {
		genprotos.UnimplementedProductServiceServer
		productService storage.Product
		logger         *log.Logger
	}
)

func New(service storage.Product) *Productservice {
	return &Productservice{
		productService: service,
		logger:         log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *Productservice) AddProduct(ctx context.Context, req *genprotos.AddProductRequest) (*genprotos.AddProductResponse, error) {
	s.logger.Println("Add Product Request")
	return s.productService.AddProduct(ctx, req)
}

func (s *Productservice) EditProduct(ctx context.Context, req *genprotos.EditProductRequest) (*genprotos.EditProductResponse, error) {
	s.logger.Println("Edit Product Request")
	return s.productService.EditProduct(ctx, req)
}

func (s *Productservice) DeleteProduct(ctx context.Context, req *genprotos.DeleteProductRequest) (*genprotos.Message, error) {
	s.logger.Println("Delete Product Request")
	return s.productService.DeleteProduct(ctx, req)
}

func (s *Productservice) GetProduct(ctx context.Context, req *genprotos.GetProductRequest) (*genprotos.GetProductResponse, error) {
	s.logger.Println("Get Product Request")
	return s.productService.GetProduct(ctx, req)
}

func (s *Productservice) GetAllProducts(ctx context.Context, req *genprotos.GetAllProductsRequest) (*genprotos.GetAllProductsResponse, error) {
	s.logger.Println("Get All Products Request")
	return s.productService.GetAllProducts(ctx, req)
}

func (s *Productservice) SearchAndFilterProduct(ctx context.Context, req *genprotos.SearchAndFilterRequest) (*genprotos.SearchAndFilterResponse, error) {
	s.logger.Println("Search and Filter Product Request")
	return s.productService.SearchAndFilterProducts(ctx, req)
}

func (s *Productservice) RateProduct(ctx context.Context, req *genprotos.RateProductRequest) (*genprotos.RateProductResponse, error) {
	s.logger.Println("Rate Product Request")
	return s.productService.RateProduct(ctx, req)
}

func (s *Productservice) GetAllRatings(ctx context.Context, req *genprotos.GetAllRatingsRequest) (*genprotos.GetAllRatingsResponse, error) {
	s.logger.Println("Get All Ratings Request")
	return s.productService.GetAllRatings(ctx, req)
}

func (s *Productservice) GetAllOrders(ctx context.Context, req *genprotos.GetAllOrdersRequest) (*genprotos.GetAllOrdersResponse, error) {
	s.logger.Println("Get All Orders Request")
	return s.productService.GetAllOrders(ctx, req)
}

func (s *Productservice) ShowOrderInfo(ctx context.Context, req *genprotos.ShowOrderInfoRequest) (*genprotos.ShowOrderInfoResponse, error) {
	s.logger.Println("Get Order Info Request")
	return s.productService.ShowOrderInfo(ctx, req)
}

func (s *Productservice) CancelOrder(ctx context.Context, req *genprotos.CancelOrderRequest) (*genprotos.CancelOrderResponse, error) {
	s.logger.Println("Cancel order Request")
	return s.productService.CancelOrder(ctx, req)
}

func (s *Productservice) ChangeOrderStatus(ctx context.Context, req *genprotos.ChangeOrderStatusRequest) (*genprotos.ChangeOrderStatusResponse, error) {
	s.logger.Println("Get Order Info Request")
	return s.productService.ChangeOrderStatus(ctx, req)
}

func (s *Productservice) OrderProduct(ctx context.Context, req *genprotos.OrderRequest) (*genprotos.OrderResponse, error) {
	return s.productService.OrderProduct(ctx, req)
}

func (s *Productservice) Pay(ctx context.Context, req *genprotos.PayRequest) (*genprotos.PayResponse, error) {
	return s.productService.Pay(ctx, req)
}

func (s *Productservice) UpdateShippingDetails(ctx context.Context, req *genprotos.UpdateShippingDetailsRequest) (*genprotos.UpdateShippingDetailsResponse, error) {
	return s.productService.UpdateShippingDetails(ctx, req)
}

func (s *Productservice) AddCategory(ctx context.Context, req *genprotos.AddCategoryRequest) (*genprotos.AddCategoryResponse, error) {
	return s.productService.AddCategory(ctx, req)
}

func (s *Productservice) GetStatistics(ctx context.Context, req *genprotos.GetStatisticsRequest) (*genprotos.GetStatisticsResponse, error) {
	return s.productService.GetStatistics(ctx, req)
}

func (s *Productservice) GetUserActivity(ctx context.Context, req *genprotos.GetUserActivityRequest) (*genprotos.GetUserActivityResponse, error) {
	return s.productService.GetUserActivity(ctx, req)
}

func (s *Productservice) GetArtisanRankings(ctx context.Context, req *genprotos.GetArtisanRankingsRequest) (*genprotos.GetArtisanRankingsResponse, error) {
	return s.productService.GetArtisanRankings(ctx, req)
}

func (s *Productservice) GetRecommendations(ctx context.Context, req *genprotos.GetRecommendationsRequest) (*genprotos.GetRecommendationsResponse, error) {
	return s.productService.GetRecommendations(ctx, req)
}

func (s *Productservice) CheckPaymentStatus(ctx context.Context, req *genprotos.CheckPaymentStatusRequest) (*genprotos.CheckPaymentStatusResponse, error) {
	return s.productService.CheckPaymentStatus(ctx, req)
}
