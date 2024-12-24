package api

import (
	"net/http"
)

func InitializeRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/test", func(w http.ResponseWriter, r *http.Request) {
		TestAPi(w, r)
	})

	nextMux := Logging(CORS(mux))

	return nextMux
}