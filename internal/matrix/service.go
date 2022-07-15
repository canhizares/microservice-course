package matrix

import (
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/sumelms/microservice-course/internal/matrix/database"
	"github.com/sumelms/microservice-course/internal/matrix/domain"
	"github.com/sumelms/microservice-course/internal/matrix/transport"
)

func NewHTTPService(router *mux.Router, db *sqlx.DB, logger log.Logger) error {
	service, err := NewService(db, logger)
	if err != nil {
		return err
	}
	transport.NewHTTPHandler(router, service, logger)
	return nil
}

func NewService(db *sqlx.DB, logger log.Logger) (*domain.Service, error) {
	repository, err := database.NewRepository(db)
	if err != nil {
		return nil, err
	}
	return domain.NewService(repository, logger), nil
}
