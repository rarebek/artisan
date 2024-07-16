package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"armiya/equipment-service/genprotos"
	"armiya/equipment-service/internal/config"
)

type Product struct {
	db           *sql.DB
	queryBuilder squirrel.StatementBuilderType
}

func New(config *config.Config) (*Product, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	return &Product{
		db:           db,
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}, nil
}

func (p *Product) AddProduct(ctx context.Context, req *genprotos.AddProductRequest) (*genprotos.AddProductResponse, error) {
	// Begin a transaction
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Check if category ID exists
	var categoryId string
	queryCategory := "SELECT id FROM product_categories WHERE id = $1"
	err = tx.QueryRowContext(ctx, queryCategory, req.CategoryId).Scan(&categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category ID %s not found", req.CategoryId)
		}
		return nil, fmt.Errorf("failed to verify category ID: %v", err)
	}

	// Prepare product data
	data := map[string]interface{}{
		"id":          uuid.NewString(),
		"name":        req.Name,
		"description": req.Description,
		"artisan_id":  req.ArtisanId,
		"price":       req.Price,
		"category_id": req.CategoryId,
		"quantity":    req.Quantity,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	// Build SQL query to insert product
	query, args, err := p.queryBuilder.Insert("products").
		SetMap(data).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	// Execute SQL query
	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	// Return response
	return &genprotos.AddProductResponse{
		Id:          data["id"].(string),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		CategoryId:  req.CategoryId,
		Quantity:    req.Quantity,
		CreatedAt:   data["created_at"].(time.Time).String(),
	}, nil
}

func (p *Product) EditProduct(ctx context.Context, req *genprotos.EditProductRequest) (*genprotos.EditProductResponse, error) {
	data := map[string]interface{}{
		"id":    req.Id,
		"name":  req.Name,
		"price": req.Price,
	}

	query, args, err := p.queryBuilder.Update("products").
		SetMap(data).
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	var updatedProduct genprotos.EditProductResponse
	err = p.db.QueryRowContext(ctx, "SELECT id, name, description, price, category_id, quantity, updated_at FROM products WHERE id = $1", req.Id).
		Scan(&updatedProduct.Id, &updatedProduct.Name, &updatedProduct.Description, &updatedProduct.Price, &updatedProduct.CategoryId, &updatedProduct.Quantity, &updatedProduct.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated product: %v", err)
	}

	return &updatedProduct, nil
}

func (p *Product) DeleteProduct(ctx context.Context, req *genprotos.DeleteProductRequest) (*genprotos.Message, error) {
	query, args, err := p.queryBuilder.Delete("products").
		Where(squirrel.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	result, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("product with ID %s not found", req.Id)
	}

	return &genprotos.Message{Message: fmt.Sprintf("Product with ID %s deleted successfully", req.Id)}, nil
}

func (p *Product) GetProduct(ctx context.Context, req *genprotos.GetProductRequest) (*genprotos.GetProductResponse, error) {
	var product genprotos.GetProductResponse
	err := p.db.QueryRowContext(ctx, "SELECT id, name, description, price, category_id, quantity, created_at, updated_at FROM products WHERE id = $1", req.Id).
		Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Quantity, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product: %v", err)
	}

	return &product, nil
}

func (p *Product) GetAllProducts(ctx context.Context, req *genprotos.GetAllProductsRequest) (*genprotos.GetAllProductsResponse, error) {
	var products []*genprotos.Product
	var total uint64

	rows, err := p.db.QueryContext(ctx, "SELECT id, name, description, price, category_id, quantity FROM products LIMIT $1 OFFSET $2", req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product genprotos.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		products = append(products, &product)
	}

	err = p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM products").Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total number of products: %v", err)
	}

	return &genprotos.GetAllProductsResponse{
		Products: products,
		Total:    total,
		Page:     req.Page,
		Limit:    req.Limit,
	}, nil
}

func (p *Product) SearchAndFilterProducts(ctx context.Context, req *genprotos.SearchAndFilterRequest) (*genprotos.SearchAndFilterResponse, error) {
	var products []*genprotos.Product
	var total uint64

	baseQuery := "SELECT id, name, description, price, category_id, quantity FROM products WHERE 1=1"
	countQuery := "SELECT COUNT(*) FROM products WHERE 1=1"

	if req.Name != "" {
		baseQuery += fmt.Sprintf(" AND name ILIKE '%%%s%%'", req.Name)
		countQuery += fmt.Sprintf(" AND name ILIKE '%%%s%%'", req.Name)
	}
	if req.Category != "" {
		baseQuery += fmt.Sprintf(" AND category_id = '%s'", req.Category)
		countQuery += fmt.Sprintf(" AND category_id = '%s'", req.Category)
	}
	if req.MinPrice > 0 {
		baseQuery += fmt.Sprintf(" AND price >= %f", req.MinPrice)
		countQuery += fmt.Sprintf(" AND price >= %f", req.MinPrice)
	}
	if req.MaxPrice > 0 {
		baseQuery += fmt.Sprintf(" AND price <= %f", req.MaxPrice)
		countQuery += fmt.Sprintf(" AND price <= %f", req.MaxPrice)
	}

	baseQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", req.Limit, (req.Page-1)*req.Limit)

	rows, err := p.db.Query(baseQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product genprotos.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.CategoryId, &product.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		products = append(products, &product)
	}

	err = p.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total number of filtered products: %v", err)
	}

	return &genprotos.SearchAndFilterResponse{
		Products: products,
		Total:    total,
		Page:     req.Page,
		Limit:    req.Limit,
	}, nil
}

func (p *Product) RateProduct(ctx context.Context, req *genprotos.RateProductRequest) (*genprotos.RateProductResponse, error) {
	data := map[string]interface{}{
		"id":         uuid.NewString(),
		"user_id":    req.UserId,
		"product_id": req.ProductId,
		"rating":     req.Rating,
		"comment":    req.Comment,
		"created_at": time.Now(),
	}

	query, args, err := p.queryBuilder.Insert("ratings").
		SetMap(data).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return &genprotos.RateProductResponse{
		Id:        data["id"].(string),
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Rating:    req.Rating,
		Comment:   req.Comment,
		CreatedAt: data["created_at"].(time.Time).String(),
	}, nil
}

func (p *Product) GetAllRatings(ctx context.Context, req *genprotos.GetAllRatingsRequest) (*genprotos.GetAllRatingsResponse, error) {
	var ratings []*genprotos.Rating

	rows, err := p.db.QueryContext(ctx, "SELECT id, user_id, rating, comment, created_at FROM ratings WHERE product_id = $1", req.ProductId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ratings: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var rating genprotos.Rating
		err := rows.Scan(&rating.Id, &rating.UserId, &rating.Rating, &rating.Comment, &rating.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rating row: %v", err)
		}
		ratings = append(ratings, &rating)
	}

	var averageRating float32
	var totalRatings uint64

	err = p.db.QueryRowContext(ctx, "SELECT AVG(rating), COUNT(*) FROM ratings WHERE product_id = $1", req.ProductId).
		Scan(&averageRating, &totalRatings)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate average rating: %v", err)
	}

	return &genprotos.GetAllRatingsResponse{
		Ratings:       ratings,
		AverageRating: averageRating,
		TotalRatings:  totalRatings,
	}, nil
}

func (p *Product) GetAllOrders(ctx context.Context, req *genprotos.GetAllOrdersRequest) (*genprotos.GetAllOrdersResponse, error) {
	var orders []*genprotos.Order
	var total uint64

	rows, err := p.db.QueryContext(ctx, "SELECT id, user_id, total_amount, status, created_at FROM orders LIMIT $1 OFFSET $2", req.Limit, (req.Page-1)*req.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order genprotos.Order
		err := rows.Scan(&order.Id, &order.UserId, &order.TotalAmount, &order.Status, &order.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order row: %v", err)
		}
		orders = append(orders, &order)
	}

	err = p.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to get total number of orders: %v", err)
	}

	return &genprotos.GetAllOrdersResponse{
		Orders: orders,
		Total:  total,
		Page:   req.Page,
		Limit:  req.Limit,
	}, nil
}

func (p *Product) ShowOrderInfo(ctx context.Context, req *genprotos.ShowOrderInfoRequest) (*genprotos.ShowOrderInfoResponse, error) {
	order := &genprotos.ShowOrderInfoResponse{
		Items:           make([]*genprotos.Item, 0),
		ShippingAddress: &genprotos.ShippingAddress{},
	}

	var shippingAddressJSON []byte

	err := p.db.QueryRowContext(ctx, `
        SELECT id, user_id, total_amount, status, shipping_address, created_at, updated_at 
        FROM orders 
        WHERE id = $1
    `, req.Id).Scan(
		&order.OrderId,
		&order.UserId,
		&order.TotalAmount,
		&order.Status,
		&shippingAddressJSON,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order: %v", err)
	}

	// Unmarshal the shipping address JSON
	err = json.Unmarshal(shippingAddressJSON, order.ShippingAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal shipping address: %v", err)
	}

	rows, err := p.db.QueryContext(ctx, "SELECT product_id, quantity FROM order_items WHERE order_id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := &genprotos.Item{}
		err := rows.Scan(&item.ProductId, &item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order item row: %v", err)
		}
		order.Items = append(order.Items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over order items rows: %v", err)
	}

	return order, nil
}

func (p *Product) CancelOrder(ctx context.Context, req *genprotos.CancelOrderRequest) (*genprotos.CancelOrderResponse, error) {
	data := map[string]interface{}{
		"status":     "cancelled",
		"updated_at": time.Now(),
	}

	query, args, err := p.queryBuilder.Update("orders").
		SetMap(data).
		Where(squirrel.Eq{"id": req.OrderId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return &genprotos.CancelOrderResponse{
		Id:        req.OrderId,
		Status:    "cancelled",
		UpdatedAt: data["updated_at"].(time.Time).String(),
	}, nil
}

func (p *Product) ChangeOrderStatus(ctx context.Context, req *genprotos.ChangeOrderStatusRequest) (*genprotos.ChangeOrderStatusResponse, error) {
	data := map[string]interface{}{
		"status":     req.Status,
		"updated_at": time.Now(),
	}

	query, args, err := p.queryBuilder.Update("orders").
		SetMap(data).
		Where(squirrel.Eq{"id": req.OrderId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return &genprotos.ChangeOrderStatusResponse{
		Id:        req.OrderId,
		Status:    req.Status,
		UpdatedAt: data["updated_at"].(time.Time).String(),
	}, nil
}

func (p *Product) OrderProduct(ctx context.Context, req *genprotos.OrderRequest) (*genprotos.OrderResponse, error) {
	orderID := uuid.New()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	totalAmount := calculateTotalAmount(ctx, p.db, req.Items)

	// Marshal ShippingAddress to JSON
	shippingAddressJSON, err := json.Marshal(req.ShippingAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal shipping address: %v", err)
	}

	_, err = tx.ExecContext(ctx, `
        INSERT INTO orders (id, user_id, total_amount, status, shipping_address, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, orderID, req.UserId, totalAmount, "pending", shippingAddressJSON, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to insert order into database: %v", err)
	}

	for _, item := range req.Items {
		price, err := getProductPrice(ctx, p.db, item.ProductId)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch product price: %v", err)
		}
		id := uuid.NewString()
		_, err = tx.ExecContext(ctx, `
            INSERT INTO order_items (id, order_id, product_id, quantity, price)
            VALUES ($1, $2, $3, $4, $5)
        `, id, orderID, item.ProductId, item.Quantity, price)
		if err != nil {
			return nil, fmt.Errorf("failed to insert order item into database: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &genprotos.OrderResponse{
		Id:              orderID.String(),
		UserId:          req.UserId,
		TotalAmount:     float32(totalAmount),
		Status:          "pending",
		ShippingAddress: req.ShippingAddress,
		CreatedAt:       time.Now().Format(time.RFC3339),
		Items:           req.Items,
	}, nil
}

func calculateTotalAmount(ctx context.Context, db *sql.DB, items []*genprotos.Item) float64 {
	var total float64
	for _, item := range items {
		price, err := getProductPrice(ctx, db, item.ProductId)
		if err != nil {
			log.Printf("failed to fetch product price for product %s: %v", item.ProductId, err)
			continue
		}
		total += float64(item.Quantity) * price
	}
	return total
}

func getProductPrice(ctx context.Context, db *sql.DB, productID string) (float64, error) {
	var price float64
	err := db.QueryRowContext(ctx, "SELECT price FROM products WHERE id = $1", productID).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch product price: %v", err)
	}
	return price, nil
}

func (p *Product) Pay(ctx context.Context, req *genprotos.PayRequest) (*genprotos.PayResponse, error) {
	paymentID := uuid.New()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	totalAmount, err := calculateTotalAmountForPayment(ctx, p.db, req.OrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total amount for payment: %v", err)
	}

	_, err = tx.ExecContext(ctx, `
        INSERT INTO payments (id, order_id, amount, status, transaction_id, payment_method, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `, paymentID, req.OrderId, totalAmount, "paid", "dummy_transaction_id", req.PaymentMethod, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to insert payment into database: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &genprotos.PayResponse{
		OrderId:       req.OrderId,
		PaymentId:     paymentID.String(),
		Amount:        float32(totalAmount),
		Status:        "paid",
		TransactionId: "dummy_transaction_id",
		CreatedAt:     time.Now().Format(time.RFC3339),
	}, nil
}

func calculateTotalAmountForPayment(ctx context.Context, db *sql.DB, orderID string) (float64, error) {
	var totalAmount float64
	rows, err := db.QueryContext(ctx, `
        SELECT quantity, price
        FROM order_items
        WHERE order_id = $1
    `, orderID)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch order items: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var quantity uint64
		var price float64
		if err := rows.Scan(&quantity, &price); err != nil {
			return 0, fmt.Errorf("failed to scan order item: %v", err)
		}
		totalAmount += float64(quantity) * price
	}
	return totalAmount, nil
}

func (p *Product) UpdateShippingDetails(ctx context.Context, req *genprotos.UpdateShippingDetailsRequest) (*genprotos.UpdateShippingDetailsResponse, error) {
	// Create the shipping details map
	shippingDetails := map[string]interface{}{
		"tracking_number":         req.TrackingNumber,
		"carrier":                 req.Carrier,
		"estimated_delivery_date": req.EstimatedDeliveryDate,
	}

	// Convert the shipping details to JSON
	shippingDetailsJSON, err := json.Marshal(shippingDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal shipping details: %v", err)
	}

	// Update the order in the database
	query, args, err := p.queryBuilder.Update("orders").
		Set("shipping_address", shippingDetailsJSON).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": req.OrderId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	if _, err := p.db.ExecContext(ctx, query, args...); err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}

	// Fetch the updated order details
	var updatedOrder genprotos.UpdateShippingDetailsResponse
	var shippingAddressJSON []byte

	err = p.db.QueryRowContext(ctx, "SELECT id, shipping_address, updated_at FROM orders WHERE id = $1", req.OrderId).
		Scan(&updatedOrder.OrderId, &shippingAddressJSON, &updatedOrder.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated order: %v", err)
	}

	// Unmarshal the JSON into the response struct fields
	var shippingDetailsMap map[string]interface{}
	if err := json.Unmarshal(shippingAddressJSON, &shippingDetailsMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal shipping address: %v", err)
	}

	updatedOrder.TrackingNumber = shippingDetailsMap["tracking_number"].(string)
	updatedOrder.Carrier = shippingDetailsMap["carrier"].(string)
	updatedOrder.EstimatedDeliveryDate = shippingDetailsMap["estimated_delivery_date"].(string)

	return &updatedOrder, nil
}

func (p *Product) AddCategory(ctx context.Context, req *genprotos.AddCategoryRequest) (*genprotos.AddCategoryResponse, error) {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()
	id := uuid.NewString()
	var categoryID string
	err = tx.QueryRowContext(ctx, `
		INSERT INTO product_categories (id, name, description, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, id, req.Name, req.Description, time.Now()).Scan(&categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to add category: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &genprotos.AddCategoryResponse{
		Id:          categoryID,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}, nil
}

func (p *Product) GetStatistics(ctx context.Context, req *genprotos.GetStatisticsRequest) (*genprotos.GetStatisticsResponse, error) {
	var (
		totalSales   uint64
		totalRevenue float64
	)

	err := p.db.QueryRowContext(ctx, `
		SELECT COUNT(*) AS total_sales, SUM(total_amount) AS total_revenue
		FROM orders
		WHERE created_at >= $1 AND created_at <= $2
	`, req.StartDate, req.EndDate).Scan(&totalSales, &totalRevenue)
	if err != nil {
		return nil, fmt.Errorf("failed to get statistics: %v", err)
	}

	return &genprotos.GetStatisticsResponse{
		TotalSales:    totalSales,
		TotalRevenue:  float32(totalRevenue),
		TopProducts:   nil,
		TopCategories: nil,
	}, nil
}

func (p *Product) GetUserActivity(ctx context.Context, req *genprotos.GetUserActivityRequest) (*genprotos.GetUserActivityResponse, error) {
	var (
		ordersPlaced   uint64
		totalSpent     float64
		reviewsWritten uint64
	)

	err := p.db.QueryRowContext(ctx, `
		SELECT COUNT(*) AS orders_placed, SUM(total_amount) AS total_spent, COUNT(*) AS reviews_written
		FROM orders
		WHERE user_id = $1 AND created_at >= $2 AND created_at <= $3
	`, req.UserId, req.StartDate, req.EndDate).Scan(&ordersPlaced, &totalSpent, &reviewsWritten)
	if err != nil {
		return nil, fmt.Errorf("failed to get user activity: %v", err)
	}

	return &genprotos.GetUserActivityResponse{
		UserId:         req.UserId,
		OrdersPlaced:   ordersPlaced,
		TotalSpent:     float32(totalSpent),
		ReviewsWritten: reviewsWritten,
	}, nil
}

func (p *Product) GetArtisanRankings(ctx context.Context, req *genprotos.GetArtisanRankingsRequest) (*genprotos.GetArtisanRankingsResponse, error) {
	var rankings []*genprotos.ArtisanRanking

	query := `
        SELECT p.artisan_id, a.full_name, AVG(r.rating) as average_rating, COUNT(o.id) as total_sales, SUM(o.total_price) as total_revenue
        FROM products p
        JOIN ratings r ON p.id = r.product_id
        JOIN orders o ON p.id = o.product_id
        JOIN artisans a ON p.artisan_id = a.id
        WHERE p.category = $1
        GROUP BY p.artisan_id, a.full_name
        ORDER BY average_rating DESC
        LIMIT $2
    `

	rows, err := p.db.QueryContext(ctx, query, req.Category, req.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch artisan rankings: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ranking genprotos.ArtisanRanking
		err := rows.Scan(&ranking.ArtisanId, &ranking.FullName, &ranking.AverageRating, &ranking.TotalSales, &ranking.TotalRevenue)
		if err != nil {
			return nil, fmt.Errorf("failed to scan artisan ranking row: %v", err)
		}
		rankings = append(rankings, &ranking)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to process artisan ranking rows: %v", err)
	}

	return &genprotos.GetArtisanRankingsResponse{
		Rankings: rankings,
	}, nil
}

func (p *Product) GetRecommendations(ctx context.Context, req *genprotos.GetRecommendationsRequest) (*genprotos.GetRecommendationsResponse, error) {
	var recommendations []*genprotos.Recommendation

	query := `
        SELECT p.id, p.name, p.price, p.category_id
        FROM products p
        JOIN recommendations r ON p.id = r.product_id
        WHERE r.user_id = $1
        ORDER BY r.score DESC
        LIMIT $2
    `

	rows, err := p.db.QueryContext(ctx, query, req.UserId, req.Limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recommendations: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var recommendation genprotos.Recommendation
		err := rows.Scan(&recommendation.Id, &recommendation.Name, &recommendation.Price, &recommendation.CategoryId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan recommendation row: %v", err)
		}
		recommendations = append(recommendations, &recommendation)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to process recommendation rows: %v", err)
	}

	return &genprotos.GetRecommendationsResponse{
		Recommendations: recommendations,
	}, nil
}

func (p *Product) CheckPaymentStatus(ctx context.Context, req *genprotos.CheckPaymentStatusRequest) (*genprotos.CheckPaymentStatusResponse, error) {
	query, args, err := p.queryBuilder.
		Select("order_id, payment_id, amount, status, transaction_id, created_at").
		From("payments").
		Where(squirrel.Eq{"order_id": req.OrderId}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %v", err)
	}

	var payment genprotos.CheckPaymentStatusResponse
	err = p.db.QueryRowContext(ctx, query, args...).Scan(
		&payment.OrderId,
		&payment.PaymentId,
		&payment.Amount,
		&payment.Status,
		&payment.TransactionId,
		&payment.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found for order_id: %s", req.OrderId)
		}
		return nil, fmt.Errorf("failed to fetch payment status: %v", err)
	}

	return &payment, nil
}
