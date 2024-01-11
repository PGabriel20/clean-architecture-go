package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/PGabriel20/expenses-go/config"
	"github.com/PGabriel20/expenses-go/internal/infra/database/repository"
	"github.com/PGabriel20/expenses-go/internal/infra/handler"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func main() {
	//Load env file
	config, err := config.LoadConfig("../")
	if err != nil {
		panic(fmt.Errorf("Failed loading config file %v", err))
	}

	//Open DB connection
	db, err := sql.Open(config.DBDriver, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))
	if err != nil {
		panic(fmt.Errorf("Failed to connect to database: %v", err))
	}

	defer db.Close()

	//Repositories
	userRepo := repository.NewUserRepositoryPostgres(db)

	//CHI router instance
	r := chi.NewRouter()

	//HTTP handlers
	handler.NewUserHandler(userRepo, r)
	
	fmt.Println("Server running on port 3000")

	http.ListenAndServe(":3000", r)
}

func init(){
	

}