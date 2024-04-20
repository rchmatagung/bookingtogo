package nationality

import (
	"bookingtogo/internal/wrapper/handler"

	"github.com/gorilla/mux"
)

func NewRoutes(router *mux.Router, handler handler.Handler) {
	router.HandleFunc("/nationality", handler.Core.Nationality.InsertNationality).Methods("POST")
	router.HandleFunc("/nationality", handler.Core.Nationality.GetAllNationality).Methods("GET")
}
