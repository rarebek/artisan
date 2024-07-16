// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: product.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProductService_AddProduct_FullMethodName             = "/ProductService/AddProduct"
	ProductService_EditProduct_FullMethodName            = "/ProductService/EditProduct"
	ProductService_DeleteProduct_FullMethodName          = "/ProductService/DeleteProduct"
	ProductService_GetAllProducts_FullMethodName         = "/ProductService/GetAllProducts"
	ProductService_GetProduct_FullMethodName             = "/ProductService/GetProduct"
	ProductService_SearchAndFilterProduct_FullMethodName = "/ProductService/SearchAndFilterProduct"
	ProductService_RateProduct_FullMethodName            = "/ProductService/RateProduct"
	ProductService_GetAllRatings_FullMethodName          = "/ProductService/GetAllRatings"
	ProductService_OrderProduct_FullMethodName           = "/ProductService/OrderProduct"
	ProductService_CancelOrder_FullMethodName            = "/ProductService/CancelOrder"
	ProductService_ChangeOrderStatus_FullMethodName      = "/ProductService/ChangeOrderStatus"
	ProductService_GetAllOrders_FullMethodName           = "/ProductService/GetAllOrders"
	ProductService_ShowOrderInfo_FullMethodName          = "/ProductService/ShowOrderInfo"
	ProductService_Pay_FullMethodName                    = "/ProductService/Pay"
	ProductService_CheckPaymentStatus_FullMethodName     = "/ProductService/CheckPaymentStatus"
	ProductService_UpdateShippingDetails_FullMethodName  = "/ProductService/UpdateShippingDetails"
	ProductService_AddCategory_FullMethodName            = "/ProductService/AddCategory"
	ProductService_GetStatistics_FullMethodName          = "/ProductService/GetStatistics"
	ProductService_GetUserActivity_FullMethodName        = "/ProductService/GetUserActivity"
	ProductService_GetRecommendations_FullMethodName     = "/ProductService/GetRecommendations"
	ProductService_GetArtisanRankings_FullMethodName     = "/ProductService/GetArtisanRankings"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error)
	EditProduct(ctx context.Context, in *EditProductRequest, opts ...grpc.CallOption) (*EditProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Message, error)
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	SearchAndFilterProduct(ctx context.Context, in *SearchAndFilterRequest, opts ...grpc.CallOption) (*SearchAndFilterResponse, error)
	RateProduct(ctx context.Context, in *RateProductRequest, opts ...grpc.CallOption) (*RateProductResponse, error)
	GetAllRatings(ctx context.Context, in *GetAllRatingsRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error)
	OrderProduct(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	ChangeOrderStatus(ctx context.Context, in *ChangeOrderStatusRequest, opts ...grpc.CallOption) (*ChangeOrderStatusResponse, error)
	GetAllOrders(ctx context.Context, in *GetAllOrdersRequest, opts ...grpc.CallOption) (*GetAllOrdersResponse, error)
	ShowOrderInfo(ctx context.Context, in *ShowOrderInfoRequest, opts ...grpc.CallOption) (*ShowOrderInfoResponse, error)
	Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
	CheckPaymentStatus(ctx context.Context, in *CheckPaymentStatusRequest, opts ...grpc.CallOption) (*CheckPaymentStatusResponse, error)
	UpdateShippingDetails(ctx context.Context, in *UpdateShippingDetailsRequest, opts ...grpc.CallOption) (*UpdateShippingDetailsResponse, error)
	AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*AddCategoryResponse, error)
	GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsResponse, error)
	GetUserActivity(ctx context.Context, in *GetUserActivityRequest, opts ...grpc.CallOption) (*GetUserActivityResponse, error)
	GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error)
	GetArtisanRankings(ctx context.Context, in *GetArtisanRankingsRequest, opts ...grpc.CallOption) (*GetArtisanRankingsResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddProductResponse)
	err := c.cc.Invoke(ctx, ProductService_AddProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) EditProduct(ctx context.Context, in *EditProductRequest, opts ...grpc.CallOption) (*EditProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditProductResponse)
	err := c.cc.Invoke(ctx, ProductService_EditProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, ProductService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllProductsResponse)
	err := c.cc.Invoke(ctx, ProductService_GetAllProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, ProductService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SearchAndFilterProduct(ctx context.Context, in *SearchAndFilterRequest, opts ...grpc.CallOption) (*SearchAndFilterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchAndFilterResponse)
	err := c.cc.Invoke(ctx, ProductService_SearchAndFilterProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) RateProduct(ctx context.Context, in *RateProductRequest, opts ...grpc.CallOption) (*RateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RateProductResponse)
	err := c.cc.Invoke(ctx, ProductService_RateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllRatings(ctx context.Context, in *GetAllRatingsRequest, opts ...grpc.CallOption) (*GetAllRatingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllRatingsResponse)
	err := c.cc.Invoke(ctx, ProductService_GetAllRatings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) OrderProduct(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, ProductService_OrderProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, ProductService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ChangeOrderStatus(ctx context.Context, in *ChangeOrderStatusRequest, opts ...grpc.CallOption) (*ChangeOrderStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeOrderStatusResponse)
	err := c.cc.Invoke(ctx, ProductService_ChangeOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAllOrders(ctx context.Context, in *GetAllOrdersRequest, opts ...grpc.CallOption) (*GetAllOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllOrdersResponse)
	err := c.cc.Invoke(ctx, ProductService_GetAllOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ShowOrderInfo(ctx context.Context, in *ShowOrderInfoRequest, opts ...grpc.CallOption) (*ShowOrderInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShowOrderInfoResponse)
	err := c.cc.Invoke(ctx, ProductService_ShowOrderInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, ProductService_Pay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CheckPaymentStatus(ctx context.Context, in *CheckPaymentStatusRequest, opts ...grpc.CallOption) (*CheckPaymentStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckPaymentStatusResponse)
	err := c.cc.Invoke(ctx, ProductService_CheckPaymentStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateShippingDetails(ctx context.Context, in *UpdateShippingDetailsRequest, opts ...grpc.CallOption) (*UpdateShippingDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateShippingDetailsResponse)
	err := c.cc.Invoke(ctx, ProductService_UpdateShippingDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*AddCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCategoryResponse)
	err := c.cc.Invoke(ctx, ProductService_AddCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetStatistics(ctx context.Context, in *GetStatisticsRequest, opts ...grpc.CallOption) (*GetStatisticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatisticsResponse)
	err := c.cc.Invoke(ctx, ProductService_GetStatistics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetUserActivity(ctx context.Context, in *GetUserActivityRequest, opts ...grpc.CallOption) (*GetUserActivityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserActivityResponse)
	err := c.cc.Invoke(ctx, ProductService_GetUserActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRecommendationsResponse)
	err := c.cc.Invoke(ctx, ProductService_GetRecommendations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetArtisanRankings(ctx context.Context, in *GetArtisanRankingsRequest, opts ...grpc.CallOption) (*GetArtisanRankingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArtisanRankingsResponse)
	err := c.cc.Invoke(ctx, ProductService_GetArtisanRankings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error)
	EditProduct(context.Context, *EditProductRequest) (*EditProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*Message, error)
	GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	SearchAndFilterProduct(context.Context, *SearchAndFilterRequest) (*SearchAndFilterResponse, error)
	RateProduct(context.Context, *RateProductRequest) (*RateProductResponse, error)
	GetAllRatings(context.Context, *GetAllRatingsRequest) (*GetAllRatingsResponse, error)
	OrderProduct(context.Context, *OrderRequest) (*OrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	ChangeOrderStatus(context.Context, *ChangeOrderStatusRequest) (*ChangeOrderStatusResponse, error)
	GetAllOrders(context.Context, *GetAllOrdersRequest) (*GetAllOrdersResponse, error)
	ShowOrderInfo(context.Context, *ShowOrderInfoRequest) (*ShowOrderInfoResponse, error)
	Pay(context.Context, *PayRequest) (*PayResponse, error)
	CheckPaymentStatus(context.Context, *CheckPaymentStatusRequest) (*CheckPaymentStatusResponse, error)
	UpdateShippingDetails(context.Context, *UpdateShippingDetailsRequest) (*UpdateShippingDetailsResponse, error)
	AddCategory(context.Context, *AddCategoryRequest) (*AddCategoryResponse, error)
	GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsResponse, error)
	GetUserActivity(context.Context, *GetUserActivityRequest) (*GetUserActivityResponse, error)
	GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error)
	GetArtisanRankings(context.Context, *GetArtisanRankingsRequest) (*GetArtisanRankingsResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) AddProduct(context.Context, *AddProductRequest) (*AddProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedProductServiceServer) EditProduct(context.Context, *EditProductRequest) (*EditProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (UnimplementedProductServiceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) SearchAndFilterProduct(context.Context, *SearchAndFilterRequest) (*SearchAndFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAndFilterProduct not implemented")
}
func (UnimplementedProductServiceServer) RateProduct(context.Context, *RateProductRequest) (*RateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateProduct not implemented")
}
func (UnimplementedProductServiceServer) GetAllRatings(context.Context, *GetAllRatingsRequest) (*GetAllRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRatings not implemented")
}
func (UnimplementedProductServiceServer) OrderProduct(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderProduct not implemented")
}
func (UnimplementedProductServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedProductServiceServer) ChangeOrderStatus(context.Context, *ChangeOrderStatusRequest) (*ChangeOrderStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeOrderStatus not implemented")
}
func (UnimplementedProductServiceServer) GetAllOrders(context.Context, *GetAllOrdersRequest) (*GetAllOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
func (UnimplementedProductServiceServer) ShowOrderInfo(context.Context, *ShowOrderInfoRequest) (*ShowOrderInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowOrderInfo not implemented")
}
func (UnimplementedProductServiceServer) Pay(context.Context, *PayRequest) (*PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}
func (UnimplementedProductServiceServer) CheckPaymentStatus(context.Context, *CheckPaymentStatusRequest) (*CheckPaymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPaymentStatus not implemented")
}
func (UnimplementedProductServiceServer) UpdateShippingDetails(context.Context, *UpdateShippingDetailsRequest) (*UpdateShippingDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShippingDetails not implemented")
}
func (UnimplementedProductServiceServer) AddCategory(context.Context, *AddCategoryRequest) (*AddCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedProductServiceServer) GetStatistics(context.Context, *GetStatisticsRequest) (*GetStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatistics not implemented")
}
func (UnimplementedProductServiceServer) GetUserActivity(context.Context, *GetUserActivityRequest) (*GetUserActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserActivity not implemented")
}
func (UnimplementedProductServiceServer) GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendations not implemented")
}
func (UnimplementedProductServiceServer) GetArtisanRankings(context.Context, *GetArtisanRankingsRequest) (*GetArtisanRankingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtisanRankings not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_EditProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).EditProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_EditProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).EditProduct(ctx, req.(*EditProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetAllProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAllProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SearchAndFilterProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAndFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SearchAndFilterProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SearchAndFilterProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SearchAndFilterProduct(ctx, req.(*SearchAndFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_RateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).RateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_RateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).RateProduct(ctx, req.(*RateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAllRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAllRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetAllRatings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAllRatings(ctx, req.(*GetAllRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_OrderProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).OrderProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_OrderProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).OrderProduct(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ChangeOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ChangeOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ChangeOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ChangeOrderStatus(ctx, req.(*ChangeOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetAllOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAllOrders(ctx, req.(*GetAllOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ShowOrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowOrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ShowOrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ShowOrderInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ShowOrderInfo(ctx, req.(*ShowOrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_Pay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Pay(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CheckPaymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPaymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CheckPaymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CheckPaymentStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CheckPaymentStatus(ctx, req.(*CheckPaymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateShippingDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShippingDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateShippingDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateShippingDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateShippingDetails(ctx, req.(*UpdateShippingDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_AddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).AddCategory(ctx, req.(*AddCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetStatistics(ctx, req.(*GetStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetUserActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetUserActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetUserActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetUserActivity(ctx, req.(*GetUserActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetRecommendations(ctx, req.(*GetRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetArtisanRankings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtisanRankingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetArtisanRankings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetArtisanRankings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetArtisanRankings(ctx, req.(*GetArtisanRankingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _ProductService_AddProduct_Handler,
		},
		{
			MethodName: "EditProduct",
			Handler:    _ProductService_EditProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _ProductService_GetAllProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "SearchAndFilterProduct",
			Handler:    _ProductService_SearchAndFilterProduct_Handler,
		},
		{
			MethodName: "RateProduct",
			Handler:    _ProductService_RateProduct_Handler,
		},
		{
			MethodName: "GetAllRatings",
			Handler:    _ProductService_GetAllRatings_Handler,
		},
		{
			MethodName: "OrderProduct",
			Handler:    _ProductService_OrderProduct_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _ProductService_CancelOrder_Handler,
		},
		{
			MethodName: "ChangeOrderStatus",
			Handler:    _ProductService_ChangeOrderStatus_Handler,
		},
		{
			MethodName: "GetAllOrders",
			Handler:    _ProductService_GetAllOrders_Handler,
		},
		{
			MethodName: "ShowOrderInfo",
			Handler:    _ProductService_ShowOrderInfo_Handler,
		},
		{
			MethodName: "Pay",
			Handler:    _ProductService_Pay_Handler,
		},
		{
			MethodName: "CheckPaymentStatus",
			Handler:    _ProductService_CheckPaymentStatus_Handler,
		},
		{
			MethodName: "UpdateShippingDetails",
			Handler:    _ProductService_UpdateShippingDetails_Handler,
		},
		{
			MethodName: "AddCategory",
			Handler:    _ProductService_AddCategory_Handler,
		},
		{
			MethodName: "GetStatistics",
			Handler:    _ProductService_GetStatistics_Handler,
		},
		{
			MethodName: "GetUserActivity",
			Handler:    _ProductService_GetUserActivity_Handler,
		},
		{
			MethodName: "GetRecommendations",
			Handler:    _ProductService_GetRecommendations_Handler,
		},
		{
			MethodName: "GetArtisanRankings",
			Handler:    _ProductService_GetArtisanRankings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
