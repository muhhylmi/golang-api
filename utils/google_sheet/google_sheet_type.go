package googlesheet

import (
	"context"
	"golang-api/utils/logger"

	"google.golang.org/api/sheets/v4"
)

type Key struct {
	Type         string `json:"type"`
	ProjectID    string `json:"project_id"`
	PrivateKeyID string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	ClientEmail  string `json:"client_email"`
	ClientID     string `json:"client_id"`
	AuthURI      string `json:"auth_uri"`
	TokenURI     string `json:"token_uri"`
	AuthProvider string `json:"auth_provider_x509_cert_url"`
	Client       string `json:"client_x509_cert_url"`
}

type GSheetsServices struct {
	Logger *logger.Logger
	ctx    context.Context

	client *sheets.Service
}

type GoogleSheetServiceInterface interface {
	GetSheetData(sheetId string, sheetRange string) (*sheets.ValueRange, error)
	GetSheetDataWithFilter(sheetId string, sheetRange string) (*sheets.ValueRange, []int, error)
}
