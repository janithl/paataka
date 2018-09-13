package main

import (
	"fmt"
	"net/http"

	"github.com/janithl/paataka/database"
	"github.com/janithl/paataka/domain"
	"github.com/janithl/paataka/web"
)

func main() {
	fmt.Println("Paataka v1.0")

	pd := database.NewSQLPublicationsRepository("sql")
	ps := domain.NewPublicationServiceImpl(pd)
	pc := web.NewPublicationController(ps)

	http.Handle("/publications", pc)
	http.ListenAndServe(":8080", nil)
}
