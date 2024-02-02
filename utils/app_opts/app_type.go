package app_type

import (
	"golang-api/utils/config"
	"golang-api/utils/database"
	googlesheet "golang-api/utils/google_sheet"
	utils_grpc "golang-api/utils/grpc"
	"golang-api/utils/logger"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type App struct {
	DBService     *database.DBService
	Apps          *echo.Echo
	Validator     echo.Validator
	Logger        *logger.Logger
	GlobalConfig  *config.Configurations
	GsheetService googlesheet.GoogleSheetServiceInterface
	GrpcServices  utils_grpc.GrpcServicesInterface
	GRPC          *grpc.Server
}
