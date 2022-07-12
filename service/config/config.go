package config

import "net/http"

const (
	Host     = "172.17.0.2"
	Port     = 5432
	User     = "postgres"
	Password = "test"
	Dbname   = "data"
)

const (
	OK                    = http.StatusOK
	CREATED               = http.StatusCreated
	NO_CONTENT            = http.StatusNoContent
	BAD_REQUEST           = http.StatusBadRequest
	UNAUTHORIZED          = http.StatusUnauthorized
	FORBIDDEN             = http.StatusForbidden
	NOT_FOUND             = http.StatusNotFound
	INTERNAL_SERVER_ERROR = http.StatusInternalServerError
	SERVICE_UNAVAILABLE   = http.StatusServiceUnavailable
)
