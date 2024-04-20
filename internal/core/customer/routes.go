package customer

import (
	"bookingtogo/internal/wrapper/handler"

	"github.com/gorilla/mux"
)

func NewRoutes(router *mux.Router, handler handler.Handler) {
	router.HandleFunc("/customer", handler.Core.Customer.InsertCustomer).Methods("POST")
	router.HandleFunc("/customer", handler.Core.Customer.GetAllCustomer).Methods("GET")
	router.HandleFunc("/customer/{customerId}", handler.Core.Customer.GetCustomerById).Methods("GET")
	router.HandleFunc("/customer", handler.Core.Customer.UpdateCustomer).Methods("PUT")
}
