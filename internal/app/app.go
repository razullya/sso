package app

import (
	"log/slog"
	appgrpc "sso/internal/app/grpc"
	"sso/internal/services/auth"
	"sso/internal/storage/sqlite"
	"time"
)

type App struct {
	GPRCSrv *appgrpc.App
}

func New(

	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {

	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := appgrpc.New(log, grpcPort, authService)
	return &App{
		GPRCSrv: grpcApp,
	}
}
