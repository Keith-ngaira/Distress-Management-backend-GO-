package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load .env file from the root directory
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Warning: .env file not found. Using environment variables.")
	}

	// Validate required environment variables
	requiredEnvVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_NAME"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Error: %s environment variable is required", envVar)
		}
	}

	log.Println("Connecting to MySQL server...")

	// Connect to MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), // This can be empty for passwordless MySQL
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
	log.Println("Successfully connected to MySQL server")

	// Read schema file
	schemaSQL, err := os.ReadFile("schema_temp.sql")
	if err != nil {
		log.Fatal("Error reading schema file:", err)
	}

	// Execute schema
	log.Println("Executing schema...")
	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		log.Fatal("Error executing schema:", err)
	}

	// Create test cases
	_, err = db.Exec(`
		INSERT INTO cases (reference_number, sender_name, receiving_date, subject, country_of_origin, distressed_person_name, nature_of_case, case_details, status, stage)
		VALUES 
		('REF001', 'John Doe', NOW(), 'Emergency Medical Assistance', 'Kenya', 'Alice Smith', 'Emergency', 'Need immediate medical assistance', 'Pending', 'Front Office Receipt'),
		('REF002', 'Jane Smith', NOW(), 'Lost Passport', 'Uganda', 'Bob Johnson', 'Standard', 'Lost passport during travel', 'Under Review', 'Director Review'),
		('REF003', 'Mike Brown', NOW(), 'Financial Aid', 'Tanzania', 'Carol White', 'Urgent', 'Requires financial assistance', 'In Progress', 'Case Investigation'),
		('REF004', 'Sarah Wilson', NOW(), 'Legal Support', 'Rwanda', 'David Lee', 'Standard', 'Legal consultation needed', 'Assigned', 'Cadet Assignment'),
		('REF005', 'Tom Harris', NOW(), 'Medical Emergency', 'Burundi', 'Eve Taylor', 'Emergency', 'Critical medical condition', 'In Progress', 'Case Investigation')
	`)
	if err != nil {
		log.Fatal("Error creating test cases:", err)
	}
	log.Println("Test cases created successfully!")

	log.Println("Database migration completed successfully!")
}
