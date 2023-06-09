package helloservice

import (
	"github.com/rs/zerolog/log"
	"macinvoice/internal/models"
)

type HelloService interface {
	WriteMessage(message string) models.Hello
}

type service struct {
	Config
}

func NewService(config Config) (HelloService, error) {

	return &service{config}, nil
}

// TODO: ctx is passed for traceability, we need to retrieve the transaction ID
func (s *service) WriteMessage(message string) models.Hello {
	hello := models.Hello{message}

	log.Info().Msg("WriteMessage() executed")
	log.Debug().Msg("message: " + message)

	return hello
}
