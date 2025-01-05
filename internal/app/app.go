package app

import (
	"context"

	"yourproject/config"
	"yourproject/internal/database"
	"yourproject/internal/repository"
	"yourproject/internal/server"
	"yourproject/internal/service"
	"yourproject/pkg/logger"
)

type App struct {
	cfg     *config.Config
	logger  *logger.Logger
	server  *server.Server
	cleanup []func()
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log := logger.New(cfg.LogLevel)

	db, err := database.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	repo := repository.NewPostgresRepository(db)

	svc := service.New(repo, log)

	srv := server.New(cfg.Server, svc, log)

	return &App{
		cfg:    cfg,
		logger: log,
		server: srv,
		cleanup: []func(){
			func() { sqlDB.Close() },
		},
	}, nil
}

func (a *App) Run() error {
	return a.server.Start()
}

func (a *App) Shutdown(ctx context.Context) {
	a.server.Shutdown(ctx)
	for _, cleanup := range a.cleanup {
		cleanup()
	}
}
