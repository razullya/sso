package app

import (
	"log/slog"
	appgrpc "sso/internal/app/grpc"
	"time"
)

type App struct {
	GPRCSrv *appgrpc.App
}

func New(

	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTl time.Duration,
) *App {
	grpcApp := appgrpc.New(log, grpcPort)
	return &App{
		GPRCSrv: grpcApp,
	}
}
