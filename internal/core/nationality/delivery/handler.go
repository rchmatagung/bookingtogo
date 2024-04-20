package delivery

import (
	"bookingtogo/config"
	"bookingtogo/internal/core/nationality/models"
	"bookingtogo/internal/wrapper/usecase"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type NationalityHandler struct {
	Usecase usecase.Usecase
	Conf    *config.Config
	Log     *logrus.Logger
}

func NewNationalityHandler(uc usecase.Usecase, conf *config.Config, logger *logrus.Logger) NationalityHandler {
	return NationalityHandler{
		Usecase: uc,
		Conf:    conf,
		Log:     logger,
	}
}

func (n NationalityHandler) InsertNationality(w http.ResponseWriter, r *http.Request) {
	dataReq := new(models.NationalityRequest)
	if err := json.NewDecoder(r.Body).Decode(dataReq); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	insertNationality, err := n.Usecase.Core.Nationality.InsertNationality(r.Context(), *dataReq)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insertNationality)
}

func (n NationalityHandler) GetAllNationality(w http.ResponseWriter, r *http.Request) {
	getAllNationality, err := n.Usecase.Core.Nationality.GetAllNationality(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllNationality)
}
