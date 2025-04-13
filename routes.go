package main

import (
	"cv-api/internal/handler"
	"net/http"
)

func SetupRoutes(cvHandler *handler.CVHandler) {
	http.HandleFunc("/cv", cvHandler.GetCV)
}
