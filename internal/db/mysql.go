package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

type DBSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"dbname"`
	Port     int    `json:"port"`
}

func LoadDBSecret(secretName string) (DBSecret, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return DBSecret{}, err
	}

	client := sm.NewFromConfig(cfg)
	result, err := client.GetSecretValue(context.TODO(), &sm.GetSecretValueInput{
		SecretId: &secretName,
	})

	if err != nil {
		return DBSecret{}, err
	}

	var secret DBSecret
	json.Unmarshal([]byte(*result.SecretString), &secret)
	return secret, nil
}

func Connect(secretName string) (*sql.DB, error) {
	secret, err := LoadDBSecret(secretName)
	if err != nil {
		log.Fatal("Failed to load secret:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		secret.Username,
		secret.Password,
		secret.Host,
		secret.Port,
		secret.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
