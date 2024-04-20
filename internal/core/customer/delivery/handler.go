package delivery

import (
	"bookingtogo/config"
	"bookingtogo/internal/core/customer/models"
	"bookingtogo/internal/wrapper/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CustomerHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewCustomerHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) CustomerHandler {
	return CustomerHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (c CustomerHandler) InsertCustomer(w http.ResponseWriter, r *http.Request) {
	dataReq := new(models.CustomerRequest)
	if err := json.NewDecoder(r.Body).Decode(dataReq); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	insertCustomer, err := c.Usecase.Core.Customer.InsertCustomer(r.Context(), *dataReq)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(insertCustomer)
}

func (c CustomerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request) {

	getAllCustomer, err := c.Usecase.Core.Customer.GetAllCustomer(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllCustomer)

}

func (c CustomerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {

	customerIdStr := mux.Vars(r)["customerId"]
	customerId, err := strconv.ParseInt(customerIdStr, 10, 64)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	getCustomerById, err := c.Usecase.Core.Customer.GetCustomerById(r.Context(), customerId)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getCustomerById)

}
