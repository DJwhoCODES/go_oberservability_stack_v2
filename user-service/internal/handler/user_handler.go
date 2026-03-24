package handler

import (
	"context"
	pb "instagram/user-service/internal/gen/user/proto"
	"instagram/user-service/internal/model"
	"instagram/user-service/internal/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.svc.Signup(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.SignupResponse{UserId: user.ID}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: user.ID}, nil
}
