package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not set in the environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	DB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Failed to create connection pool: %v", err)
	}

	if err := DB.Ping(ctx); err != nil {
		log.Fatalf("❌ Unable to ping DB: %v", err)
	}

	fmt.Println("✅ Connected to Postgres DB")
}
