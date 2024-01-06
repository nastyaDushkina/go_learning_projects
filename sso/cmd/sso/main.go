package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"syscall"
)

const (
	envDev   = "dev"
	envLocal = "local"
	envProd  = "prods"
)

func main() {
	// TODO инициализация объекта конфига
	cfg := config.MustLoad()

	// TODO инициалзаация логгера

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("cfg", cfg))
	//log.Debug("debug message")
	//log.Error("error message")
	//log.Warn("warn message")

	// TODO инициализация приложения (app)
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	// TODO  запустит gRPC сервер
	// запуск в отдельной горутине
	go application.GRPCSrv.MustRun()

	// TODO Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// зависаем на моменте, пока в каанал что-то не придёт, чтобы прочитать оттуда
	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()

	log.Info("application stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
