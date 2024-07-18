package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDB(){

	dsn := "user=postgres.mianfanfnhpopqdbbdzi password=thedarkslayer@59 dbname=postgres sslmode=require host=aws-0-ap-south-1.pooler.supabase.com port=6543"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("Failed to get database connection:", err)
    }

    err = sqlDB.Ping()
    if err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    log.Println("Database connection successfully established")

}

func CloseDB(){

	sqlDB, err := db.DB()

	if err != nil{
		log.Fatal("Failed to close db connection")
	}
	if err := sqlDB.Close(); err != nil{
		log.Fatal("Failed to close db connection")
	}
}