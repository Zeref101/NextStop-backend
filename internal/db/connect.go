package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error){

	dsn := "user=postgres.mianfanfnhpopqdbbdzi password=thedarkslayer@59 dbname=postgres sslmode=require host=aws-0-ap-south-1.pooler.supabase.com port=6543"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        return nil, fmt.Errorf("failed to connect to db: %w", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, fmt.Errorf("failed to get database connection: %w", err)
    }

    err = sqlDB.Ping()
    if err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    log.Println("Database connection successfully established")

	return db, nil

}

func CloseDB(dbInstance *gorm.DB){


	sqlDB, err := dbInstance.DB()

	if err != nil{
		log.Fatal("Failed to close db connection")
	}

	
	if err := sqlDB.Close(); err != nil{
		log.Fatal("Failed to close db connection")
	}
}