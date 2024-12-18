package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CSV_data represents the table structure
type CSV_data struct {
	ID                    uint `gorm:"primaryKey"`
	SiteID                string
	FixletID              string
	Name                  string
	Criticality           string
	RelevantComputerCount int
}

func main() {
	// Database connection string
	dbcon := "host=postgres user=postgres password=password dbname=CSV_db port=5432 sslmode=disable TimeZone=UTC"

	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dbcon), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&CSV_data{})

	// Open the CSV file
	file, err := os.Open("file.csv")
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file content: %v", err)
	}

	// Insert each row into the database
	for _, row := range records[1:] { // Skip header row
		data := CSV_data{
			SiteID:                row[0],
			FixletID:              row[1],
			Name:                  row[2],
			Criticality:           row[3],
			RelevantComputerCount: atoi(row[4]),
		}
		db.Create(&data)
	}

	fmt.Println("CSV data imported successfully.")
}

// Helper function to convert string to integer
func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
