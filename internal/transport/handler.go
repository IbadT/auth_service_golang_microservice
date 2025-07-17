package transportgrpc

import (
	"context"

	"github.com/IbadT/auth_service_golang_microservice.git/internal/auth"
	authpb "github.com/IbadT/project-protos/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	svc auth.Service
	authpb.UnimplementedAuthServiceServer
}

func NewHandler(s auth.Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	email := req.Email
	password := req.Password
	login, err := h.svc.Login(email, password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &authpb.LoginResponse{
		AccessToken:  login.AccessToken,
		RefreshToken: login.RefreshToken,
	}, nil
}

func (h *Handler) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	email := req.Email
	password := req.Password
	registerResp, err := h.svc.Register(email, password)
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "user is already exist")
	}

	return &authpb.RegisterResponse{
		Id: registerResp.ID.String(),
	}, nil
}

func (h *Handler) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	token := req.GetRefreshToken()
	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "refresh token required")
	}

	newTokens, err := h.svc.RefreshToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &authpb.RefreshTokenResponse{
		AccessToken:  newTokens.AccessToken,
		RefreshToken: newTokens.RefreshToken,
	}, nil
}
