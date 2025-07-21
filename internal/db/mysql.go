package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

type DBSecrets struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"dbname"`
}

func Connect() (*sql.DB, error) {
	secretName := os.Getenv("AWS_SECRET_NAME")
	region := os.Getenv("AWS_REGION")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	client := secretsmanager.NewFromConfig(cfg)
	secretOutput, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	})
	if err != nil {
		return nil, err
	}

	var creds DBSecrets
	err = json.Unmarshal([]byte(*secretOutput.SecretString), &creds)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		creds.Username, creds.Password, creds.Host, creds.Port, creds.Database)

	return sql.Open("mysql", dsn)
}
