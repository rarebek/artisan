package storage

import (
	"armiya/equipment-service/genprotos"
	"armiya/equipment-service/internal/config"
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
)

type (
	Auth struct {
		db           *sql.DB
		queryBuilder sq.StatementBuilderType
	}
)

func New(config *config.Config) (*Auth, error) {
	db, err := ConnectDB(*config)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	return &Auth{
		db:           db,
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, nil
}

func (e *Auth) Register(ctx context.Context, req *genprotos.RegisterRequest) (*genprotos.RegisterResponse, error) {
	data := map[string]interface{}{
		"id":         uuid.NewString(),
		"username":   req.Username,
		"email":      req.Email,
		"password":   req.Password,
		"full_name":  req.FullName,
		"user_type":  req.UserType,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	query, args, err := e.queryBuilder.Insert("users").
		SetMap(data).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	if _, err = e.db.ExecContext(ctx, query, args...); err != nil {
		pp.Println(err)
		return nil, err
	}

	return &genprotos.RegisterResponse{
		Id:        data["id"].(string),
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
		FullName:  req.FullName,
		UserType:  req.UserType,
		CreatedAt: data["created_at"].(time.Time).String(),
	}, nil
}

func (e *Auth) Login(ctx context.Context, req *genprotos.LoginRequest) (*genprotos.LoginResponse, error) {
	query, args, err := e.queryBuilder.Select("password").
		From("users").
		Where(sq.Eq{"email": req.Email}).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	var storedPassword string
	err = e.db.QueryRowContext(ctx, query, args...).Scan(&storedPassword)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	if storedPassword != req.Password {
		return &genprotos.LoginResponse{True: false}, nil
	}

	return &genprotos.LoginResponse{True: true}, nil
}

func (e *Auth) ShowProfile(ctx context.Context, req *genprotos.ShowProfileRequest) (*genprotos.ShowProfileResponse, error) {
	query, args, err := e.queryBuilder.Select("id", "username", "email", "password", "full_name", "user_type", "created_at", "updated_at").
		From("users").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	var user genprotos.ShowProfileResponse
	err = e.db.QueryRowContext(ctx, query, args...).Scan(
		&user.Id, &user.Username, &user.Email, &user.Password, &user.FullName,
		&user.UserType, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	return &user, nil
}

func (e *Auth) EditProfile(ctx context.Context, req *genprotos.EditProfileRequest) (*genprotos.EditProfileResponse, error) {
	data := map[string]interface{}{
		"full_name":  req.FullName,
		"bio":        req.Bio,
		"updated_at": time.Now(),
	}

	query, args, err := e.queryBuilder.Update("users").
		SetMap(data).
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	_, err = e.db.ExecContext(ctx, query, args...)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	selectQuery, selectArgs, err := e.queryBuilder.Select("username", "email", "password", "full_name", "user_type", "created_at").
		From("users").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	var response genprotos.EditProfileResponse
	err = e.db.QueryRowContext(ctx, selectQuery, selectArgs...).Scan(
		&response.Username, &response.Email, &response.Password, &response.FullName,
		&response.UserType, &response.CreatedAt,
	)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	response.Id = req.Id

	return &response, nil
}

func (e *Auth) EditUserType(ctx context.Context, req *genprotos.EditUserTypeRequest) (*genprotos.EditUserTypeResponse, error) {
	var username string
	data := map[string]interface{}{
		"user_type":  req.UserType,
		"updated_at": time.Now(),
	}

	selectQuery, args, err := e.queryBuilder.Select("username").
		From("users").
		Where(sq.Eq{
			"id": req.Id,
		}).ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	if err = e.db.QueryRow(selectQuery, args...).Scan(&username); err != nil {
		pp.Println(err)
		return nil, err
	}

	query, args, err := e.queryBuilder.Update("users").
		SetMap(data).
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	_, err = e.db.ExecContext(ctx, query, args...)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	return &genprotos.EditUserTypeResponse{
		Id:        req.Id,
		Username:  username,
		UserType:  req.UserType,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (e *Auth) GetAllUsers(ctx context.Context, req *genprotos.GetAllUsersRequest) (*genprotos.GetAllUsersResponse, error) {
	offset := (req.Page - 1) * req.Limit
	query, args, err := e.queryBuilder.Select("id", "username", "full_name", "user_type").
		From("users").
		Limit(req.Limit).
		Offset(offset).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	rows, err := e.db.QueryContext(ctx, query, args...)
	if err != nil {
		pp.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []*genprotos.User
	for rows.Next() {
		var user genprotos.User
		if err := rows.Scan(&user.Id, &user.Username, &user.FullName, &user.UserType); err != nil {
			pp.Println(err)
			return nil, err
		}
		users = append(users, &user)
	}

	query, args, err = e.queryBuilder.Select("COUNT(*)").From("users").ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	var total uint64
	err = e.db.QueryRowContext(ctx, query, args...).Scan(&total)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	return &genprotos.GetAllUsersResponse{
		Users: users,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, nil
}

func (e *Auth) DeleteUser(ctx context.Context, req *genprotos.DeleteUserRequest) (*genprotos.AuthMessage, error) {
	query, args, err := e.queryBuilder.Delete("users").
		Where(sq.Eq{"id": req.Id}).
		ToSql()
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	_, err = e.db.ExecContext(ctx, query, args...)
	if err != nil {
		pp.Println(err)
		return nil, err
	}

	return &genprotos.AuthMessage{Message: "User deleted successfully"}, nil
}

func (e *Auth) ResetPassword(ctx context.Context, req *genprotos.ResetPasswordRequest) (*genprotos.AuthMessage, error) {
	return &genprotos.AuthMessage{Message: "Password reset send to your email."}, nil
}
