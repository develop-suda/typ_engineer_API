package router

import (
	"net/http"
    "github.com/develop-suda/typ_engineer_API/api/handler"
)

func Router() {
	http.HandleFunc("/api", handler.SelectWordsHandler)
    http.HandleFunc("/api/levels", handler.GetLevelsHandler)
	http.ListenAndServe(":8888", nil)
}