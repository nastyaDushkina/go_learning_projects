package main

import (
	"errors"
	"flag"
	"fmt"

	//библиотека для миграций
	"github.com/golang-migrate/migrate/v4"
	// драйвер для выполнения миграций sqlite3
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// драйвер для получения миграций из файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// go run ./cmd/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./migrations
// go run ./cmd/migrator/main.go --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_tests

// migrationsTable - задаём для возможности хранения версии тестовой базы (для тестов)
// для тестов будут тестовые миграции
func main() {
	var storagePath, migrationsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to database")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of igrations table")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	//создаём экземпляр мигратора
	//x-migrations-table - отвечает за путь до таблицы версий миграций
	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied")
}
