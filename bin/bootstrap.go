package main

import (
	"fmt"
	app_type "golang-api/utils/app_opts"
	"golang-api/utils/config"
	"golang-api/utils/database"
	googlesheet "golang-api/utils/google_sheet"
	utils_grpc "golang-api/utils/grpc"
	"golang-api/utils/logger"
	utils "golang-api/utils/validator"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func panicRecoveryHandler(err any) error {
	message := fmt.Sprintf("Unexpected error: %v", err)
	fmt.Println(message)
	return status.Errorf(codes.Internal, "%s", err)
}

func Init() *app_type.App {
	config := config.GetConfig()
	logger := logger.Newlogger()
	e := echo.New()
	validator := utils.NewValidationUtil()

	PostgresUsername := config.DB_POSTGRES_USERNAME
	PostgresPassword := config.DB_POSTGRES_PASSWORD
	PostgresHost := config.DB_POSTGRES_HOST
	PostgresPort := config.DB_POSTGRES_PORT
	PostgresDBName := config.DB_POSTGRES_NAME
	PostgresSchema := config.DB_POSTGRES_SCHEMA
	dsn := "host=" + PostgresHost + " user=" + PostgresUsername + " password=" + PostgresPassword + " dbname=" + PostgresDBName + " port=" + PostgresPort + " sslmode=disable TimeZone=Asia/Shanghai" + " search_path=" + PostgresSchema

	db, err := database.NewDatabase(&database.DBServiceVar{
		Logger:      logger,
		PostgresUri: &dsn,
	})
	if err != nil {
		panic(err)
	}

	gsheetService, err := googlesheet.NewGsheetService(logger, config)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(panicRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(panicRecoveryHandler)),
		),
	)
	grpcServices, err := utils_grpc.NewGrpcServices(&utils_grpc.GrpcServiceHosts{
		BookHost: &config.BOOK_GRPC_HOST,
	})
	if err != nil {
		panic(err)
	}

	return &app_type.App{
		DBService:     db,
		Apps:          e,
		Validator:     validator,
		Logger:        logger,
		GlobalConfig:  config,
		GsheetService: gsheetService,
		GrpcServices:  grpcServices,

		GRPC: grpcServer,
	}

}
