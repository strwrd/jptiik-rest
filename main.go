package main

import (
	"log"

	"github.com/strwrd/jptiik-rest/delivery/http"
	"github.com/strwrd/jptiik-rest/repository/mysql"
	"github.com/strwrd/jptiik-rest/usecase"
)

func main() {
	// Initial mysql repository object
	mysqlRepo, err := mysql.NewRepository()
	if err != nil {
		log.Fatal(err)
	}

	// Create usecase object
	ucase := usecase.NewUsecase(mysqlRepo)

	// Create server object
	server := http.NewHandler(ucase)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
