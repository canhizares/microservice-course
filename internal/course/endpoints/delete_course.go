package endpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/sumelms/microservice-course/internal/course/domain"
)

type deleteCourseRequest struct {
	UUID string `json:"uuid" validate:"required"`
}

func NewDeleteCourseHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeDeleteCourseEndpoint(s),
		decodeDeleteCourseRequest,
		encodeDeleteCourseResponse,
		opts...,
	)
}

func makeDeleteCourseEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteCourseRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		id := uuid.MustParse(req.UUID)
		if err := s.DeleteCourse(ctx, id); err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func decodeDeleteCourseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["uuid"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return deleteCourseRequest{UUID: id}, nil
}

func encodeDeleteCourseResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
