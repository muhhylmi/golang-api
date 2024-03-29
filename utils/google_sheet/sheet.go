package googlesheet

import (
	"context"
	"encoding/json"
	"golang-api/utils/config"
	"golang-api/utils/logger"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewGsheetService(log *logger.Logger, config *config.Configurations) (GoogleSheetServiceInterface, error) {
	if log == nil {
		log = logger.Newlogger()
	}
	l := log.LogWithContext("dbConnection", "NewDatabase")
	var (
		sheetKey Key = Key{
			Type:         config.SHEET_KEY_TYPE,
			ProjectID:    config.SHEET_KEY_PROJECT_ID,
			PrivateKeyID: config.SHEET_KEY_PRIVATE_KEY_ID,
			PrivateKey:   config.SHEET_KEY_PRIVATE_KEY,
			ClientEmail:  config.SHEET_KEY_CLIENT_EMAIL,
			ClientID:     config.SHEET_KEY_CLIENT_ID,
			AuthURI:      config.SHEET_KEY_AUTH_URI,
			TokenURI:     config.SHEET_KEY_TOKEN_URI,
			AuthProvider: config.SHEET_KEY_AUTH_PROVIDER,
			Client:       config.SHEET_KEY_CLIENT_CERT_URI,
		}
	)
	credential, err := json.Marshal(sheetKey)
	if err != nil {
		l.Error(err)
		return nil, err
	}

	srv, err := sheets.NewService(context.Background(), option.WithCredentialsJSON(credential))
	if err != nil {
		l.Error(err)
		return nil, err
	}

	return &GSheetsServices{
		client: srv,
		Logger: log,
		ctx:    context.Background(),
	}, nil
}
