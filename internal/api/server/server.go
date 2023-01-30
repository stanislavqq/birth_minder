package server

import (
	"BMinder/internal/config"
	"BMinder/internal/personstore"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

type Server struct {
	store *personstore.PersonStore
}

func NewServer(store *personstore.PersonStore) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) Start(cfg *config.Config, logger zerolog.Logger) error {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		persons := s.store.GetPersonAll()

		_, err := io.WriteString(w, "Root Bitch, HTTP!\n")
		if err != nil {
			panic(err)
		}
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, request *http.Request) {

		_, err := io.WriteString(w, "Hello MZF, HTTP!\n")
		if err != nil {
			panic(err)
		}
	})

	var err error
	go func() {
		logger.Info().Msg("Инициализация http сервера")
		err = http.ListenAndServe(":3333", mux)
		if err != nil {
			logger.Error().Err(err)
			return
		}
	}()

	if err != nil {
		return err
	}
	return nil
}
