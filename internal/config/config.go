package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Config struct {
	projectID string
	DB        *sql.DB
}

func (c Config) ProjectID() string {
	return c.projectID
}

func isLocal() bool {
	return os.Getenv("K_SERVICE") == ""
}
func dsn() string {
	// postgresqlのdnsを環境変数から構成する
	return fmt.Sprintf("host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=disable")
}
func createDB() (*sql.DB, error) {
	return sql.Open("postgres", dsn())
}
func NewConfig() (Config, error) {
	var (
		projectID string
		db        *sql.DB
		err       error
	)
	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if db, err = createDB(); err != nil {
		return Config{}, fmt.Errorf("failed to open database: %w", err)
	}
	return Config{
		projectID: projectID,
		DB:        db,
	}, nil
}
