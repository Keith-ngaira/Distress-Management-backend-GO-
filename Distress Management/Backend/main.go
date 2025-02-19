package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"distress-management/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found. Using environment variables.")
	}

	// Validate required environment variables
	requiredEnvVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_NAME"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Error: %s environment variable is required", envVar)
		}
	}

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll("./uploads", 0755); err != nil {
		log.Fatal("Error creating uploads directory:", err)
	}

	// Initialize database with retry mechanism
	var db *sql.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"), // This can be empty for passwordless MySQL
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME")))
		if err != nil {
			log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
			if i < maxRetries-1 {
				time.Sleep(time.Second * 5)
				continue
			}
			log.Fatal("Error connecting to database:", err)
		}
		break
	}

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	// Initialize router and handlers
	router := mux.NewRouter()
	app := &handlers.App{
		DB: db,
	}

	// API routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Cases routes
	apiRouter.HandleFunc("/cases", app.GetCases).Methods("GET")
	apiRouter.HandleFunc("/cases", app.CreateCase).Methods("POST")
	apiRouter.HandleFunc("/cases/{id}", app.GetCase).Methods("GET")
	apiRouter.HandleFunc("/cases/{id}", app.UpdateCase).Methods("PUT")
	apiRouter.HandleFunc("/cases/{id}/status", app.UpdateCaseStatus).Methods("PATCH")

	// Documents routes
	apiRouter.HandleFunc("/cases/{id}/documents", app.UploadDocument).Methods("POST")
	apiRouter.HandleFunc("/cases/{id}/documents", app.GetDocuments).Methods("GET")
	apiRouter.HandleFunc("/cases/{id}/documents/{docId}", app.DeleteDocument).Methods("DELETE")

	// Progress notes routes
	apiRouter.HandleFunc("/cases/{id}/notes", app.AddProgressNote).Methods("POST")
	apiRouter.HandleFunc("/cases/{id}/notes", app.GetProgressNotes).Methods("GET")

	// Dashboard routes
	apiRouter.HandleFunc("/dashboard/stats", app.GetDashboardStats).Methods("GET")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		Debug:           true,
	})

	// Wrap router with CORS and logging middleware
	handler := c.Handler(router)
	handler = logRequestFunc(handler)

	// Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// Middleware to log requests
func logRequestFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
