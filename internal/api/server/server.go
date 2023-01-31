package server

import (
	"BMinder/internal/config"
	"BMinder/internal/personstore"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type Server struct {
	store *personstore.PersonStore
}

func NewServer(store *personstore.PersonStore) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Start(cfg *config.Config, ctx context.Context, logger zerolog.Logger) error {

	//mux.HandleFunc("/person/", func(w http.ResponseWriter, request *http.Request) {
	//
	//	var res Response
	//
	//	if request.Method == http.MethodGet {
	//		persons := s.store.GetPersonAll()
	//		res = NewResponse(persons, 200)
	//	}
	//
	//	w.Header().Set("Content-Type", "application/json")
	//	json, err := res.GetJson()
	//	if err != nil {
	//		logger.Error().Err(err).Msg("Не удалось вставить данные в ответ_")
	//	}
	//
	//	_, err = w.Write(json)
	//	if err != nil {
	//		logger.Error().Err(err).Msg("Не удалось вставить данные в ответ")
	//	}
	//})

	r := mux.NewRouter()

	r.HandleFunc("/person", func(w http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			persons := s.store.GetPersonAll()

			data, err := json.Marshal(persons)
			if err != nil {
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(data)
			if err != nil {
				logger.Error().Err(err)
				return
			}
		}

	})

	logger.Info().Msg("Инициализация http сервера")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3333",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Info().Msg("test")
	var err error
	go func() {
		logger.Info().Msg("test2")
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatal().Err(err).Msg("Ошибка запуска http сервера")
		}

		logger.Info().Msg("http сервер готов принимать запросы")
	}()

	return err
}
