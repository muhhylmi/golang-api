package utils

import (
	"context"
	"encoding/json"
	"golang-api/config"

	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
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

var (
	sheetKey Key = Key{
		Type:         config.GetConfig().SHEET_KEY_TYPE,
		ProjectID:    config.GetConfig().SHEET_KEY_PROJECT_ID,
		PrivateKeyID: config.GetConfig().SHEET_KEY_PRIVATE_KEY_ID,
		PrivateKey:   config.GetConfig().SHEET_KEY_PRIVATE_KEY,
		ClientEmail:  config.GetConfig().SHEET_KEY_CLIENT_EMAIL,
		ClientID:     config.GetConfig().SHEET_KEY_CLIENT_ID,
		AuthURI:      config.GetConfig().SHEET_KEY_AUTH_URI,
		TokenURI:     config.GetConfig().SHEET_KEY_TOKEN_URI,
		AuthProvider: config.GetConfig().SHEET_KEY_AUTH_PROVIDER,
		Client:       config.GetConfig().SHEET_KEY_CLIENT_CERT_URI,
	}
)

func GetSheetConfig(log *logrus.Entry) (*sheets.Service, error) {
	credential, err := json.Marshal(sheetKey)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	srv, err := sheets.NewService(context.Background(), option.WithCredentialsJSON(credential))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return srv, nil
}

func GetSheetData(service *sheets.Service, sheetId string, sheetRange string) (*sheets.ValueRange, error) {
	values, err := service.Spreadsheets.Values.Get(sheetId, sheetRange).Do()
	if err != nil {
		return nil, err
	}
	return values, nil
}
