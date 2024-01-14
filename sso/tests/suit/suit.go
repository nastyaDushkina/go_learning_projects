package suit

// здесь будет общий код для тестов
// например, подключение к базе, подготовка других соединений,
// создание клиента для grpc

import (
	"context"
	"fmt"
	"net"
	"sso/internal/config"
	"strconv"
	"testing"

	ssov1 "github.com/nastyaDushkina/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcHost = "localhost"
)

type Suite struct {
	*testing.T                  //потребуется для вызовов методов *testing.T внутри Suit
	Cfg        *config.Config   // конфигурация приложения
	AuthClient ssov1.AuthClient //клиент для взаимодействия с grpc-сервером
}

func New(t *testing.T) (context.Context, *Suite) {
	// указываем, что функция вспомогательная, чтобы при падении теста
	// был правильный стек вызовов и она не указывалась как финальная
	t.Helper()
	// чтобы тесты могли выполняться параллельно
	t.Parallel()

	cfg := config.MustLoadByPath("../config/local.yaml")

	// создаём контекст с таймаутом
	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	// отменяем контекст
	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	// создаём клиента grpc
	// будем использовать небезопасное соединение
	cc, err := grpc.DialContext(context.Background(),
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	fmt.Println(grpcAddress(cfg))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
