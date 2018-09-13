package web

import (
	"fmt"
	"net/http"

	"github.com/janithl/paataka/domain"
)

// PublicationController is a struct for the publications controller
type PublicationController struct {
	service domain.PublicationService
}

// NewPublicationController returns a new instance of PublicationController
func NewPublicationController(ps domain.PublicationService) *PublicationController {
	return &PublicationController{
		service: ps,
	}
}

func (pc *PublicationController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World")
}
