package handlers

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapGRPCError(w http.ResponseWriter, err error) {
	code := status.Code(err)

	var httpStatus int

	switch code {
	case codes.Unauthenticated:
		httpStatus = http.StatusUnauthorized
	case codes.AlreadyExists:
		httpStatus = http.StatusConflict
	case codes.InvalidArgument:
		httpStatus = http.StatusBadRequest
	case codes.NotFound:
		httpStatus = http.StatusNotFound
	default:
		httpStatus = http.StatusInternalServerError
	}

	http.Error(w, status.Convert(err).Message(), httpStatus)
}
