package app

import (
	"log/slog"
	grpc_app "sso-grpc/internal/app/grpc"
	"sso-grpc/internal/config"
)

type App struct {
	GRPCSrv *grpc_app.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	dbConfig config.DBConfig,
	tokensConfig config.TokenConfig,
) *App {
	grpcApp := grpc_app.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
