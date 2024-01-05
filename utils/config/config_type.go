package config

type Configurations struct {
	DB_USERNAME          string
	DB_PASSWORD          string
	DB_HOST              string
	DB_PORT              string
	DB_NAME              string
	DB_POSTGRES_USERNAME string
	DB_POSTGRES_PASSWORD string
	DB_POSTGRES_HOST     string
	DB_POSTGRES_PORT     string
	DB_POSTGRES_NAME     string
	DB_POSTGRES_SCHEMA   string
	HOST                 string
	BASIC_AUTH_USERNAME  string
	BASIC_AUTH_PASSWORD  string
	JWT_SECRET_KEY       string
	GRPC_PORT            string

	SPREAD_SHEET_ID           string
	SHEET_KEY_TYPE            string
	SHEET_KEY_PROJECT_ID      string
	SHEET_KEY_PRIVATE_KEY_ID  string
	SHEET_KEY_PRIVATE_KEY     string
	SHEET_KEY_CLIENT_ID       string
	SHEET_KEY_CLIENT_EMAIL    string
	SHEET_KEY_AUTH_URI        string
	SHEET_KEY_TOKEN_URI       string
	SHEET_KEY_AUTH_PROVIDER   string
	SHEET_KEY_CLIENT_CERT_URI string
	SHEET_KEY_UNIV_DOMAIN     string
}
