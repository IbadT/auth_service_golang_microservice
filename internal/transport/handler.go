package transportgrpc

import (
	"context"

	"github.com/IbadT/auth_service_golang_microservice.git/internal/auth"
	authpb "github.com/IbadT/project-protos/proto/auth"
)

type Handler struct {
	svc auth.Service
	authpb.UnimplementedAuthServiceServer
}

func NewHandler(s auth.Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return &authpb.LoginResponse{}, nil
}

func (h *Handler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	return &authpb.RegisterResponse{}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	return &authpb.RefreshTokenResponse{}, nil
}
