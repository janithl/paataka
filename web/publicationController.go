package web

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	id, err := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/publications/"))
	if err != nil {
		fmt.Fprint(w, "Publication ID should be a number!")
		return
	}

	publication := pc.service.RetrievePublication(id)
	fmt.Fprint(w, publication)
}
