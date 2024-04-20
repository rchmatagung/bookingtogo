package server

import (
	"bookingtogo/config"
	"bookingtogo/internal/wrapper/handler"
	"bookingtogo/internal/wrapper/repository"
	"bookingtogo/internal/wrapper/usecase"
	"bookingtogo/pkg/infra/db"
	"log"
	"net/http"
	"time"

	coreCustomer "bookingtogo/internal/core/customer"
	coreNationality "bookingtogo/internal/core/nationality"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run(conf *config.Config, dbList *db.DatabaseList, appLoger *logrus.Logger) {

	if dbList.DatabaseApp == nil {
		log.Println("is nil")
	}

	repo := repository.NewRepository(conf, dbList, appLoger)
	usecase := usecase.NewUsecase(repo, conf, dbList, appLoger)
	handler := handler.NewHandler(usecase, conf, appLoger)

	router := mux.NewRouter()
	coreNationality.NewRoutes(router, handler)
	coreCustomer.NewRoutes(router, handler)

	http.Handle("/", router)

	loggedRouter := handlers.LoggingHandler(appLoger.Writer(), router)

	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
