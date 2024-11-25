package main

import (
	"database/sql"
	"github.com/DiegoJCordeiro/golang-study/chapter11/internal/infra/database"
	"github.com/DiegoJCordeiro/golang-study/chapter11/internal/infra/proto"
	"github.com/DiegoJCordeiro/golang-study/chapter11/internal/infra/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {

	db, err := sql.Open("sqlite3", "./category.db")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(categoryDB)

	grpcServer := grpc.NewServer()

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	proto.RegisterCategoryServiceServer(grpcServer, categoryService)

	serverListen, errNet := net.Listen("tcp", ":50051")

	if errNet != nil {
		panic(errNet)
	}

	if errServer := grpcServer.Serve(serverListen); errServer != nil {
		panic(errServer)
	}
}
