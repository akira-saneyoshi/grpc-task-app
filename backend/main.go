package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	"github.com/akira-saneyoshi/task-app/infrastructure/persistence/model/db"
	"github.com/akira-saneyoshi/task-app/interfaces/di"
	"github.com/akira-saneyoshi/task-app/interfaces/interceptor"
	"github.com/akira-saneyoshi/task-app/interfaces/proto/auth/v1/authv1connect"
	"github.com/akira-saneyoshi/task-app/interfaces/proto/task/v1/taskv1connect"
	"github.com/akira-saneyoshi/task-app/interfaces/proto/user/v1/userv1connect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	grpcreflect "connectrpc.com/grpcreflect"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("[ERROR] Error loading .env file, using environment variables")
	}
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var issuer, keyPath, url string
	var ok bool
	if issuer, ok = os.LookupEnv("REPO_NAME"); !ok {
		return fmt.Errorf("[ERROR] issuer not set: %s", issuer)
	}
	if keyPath, ok = os.LookupEnv("PRIVATE_KEY_PATH"); !ok {
		return fmt.Errorf("[ERROR] private-key-path not set: %s", keyPath)
	}
	if url, ok = os.LookupEnv("DATABASE_URL"); !ok {
		return fmt.Errorf("[ERROR] database-url not set: %s", url)
	}

	dbConn, err := sql.Open("mysql", url)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to connect to MySQL: %w", err)
	}
	defer dbConn.Close()

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(1)
	dbConn.SetConnMaxLifetime(time.Hour)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = dbConn.PingContext(ctx); err != nil {
		return fmt.Errorf("[ERROR] failed to ping MySQL: %w", err)
	}

	qry := db.New(dbConn)

	timeout := 1 * time.Hour

	authServer, err := di.InitAuth(issuer, keyPath, qry, timeout)
	if err != nil {
		return err
	}
	userServer := di.InitUser(qry)
	taskServer := di.InitTask(qry)

	authInterceptor := connect.WithInterceptors(interceptor.NewAuthInterceptor(issuer, keyPath))

	reflector := grpcreflect.NewStaticReflector(
		// 登録するサービス名を指定 (proto ファイルの service 名)
		authv1connect.AuthServiceName,
		userv1connect.UserServiceName,
		taskv1connect.TaskServiceName,
	)

	mux := http.NewServeMux()
	mux.Handle(authv1connect.NewAuthServiceHandler(authServer))
	mux.Handle(userv1connect.NewUserServiceHandler(userServer))
	mux.Handle(taskv1connect.NewTaskServiceHandler(taskServer, authInterceptor))
	// リフレクション用のパスを登録
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return http.ListenAndServe(
		"localhost:8080",
		cors.AllowAll().Handler(
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
}
