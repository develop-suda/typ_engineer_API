package router

import (
	"net/http"

	"github.com/develop-suda/typ_engineer_API/api/handler"
)

func Router() {
	http.HandleFunc("/api/typWord", handler.TypWordSelectHandler)
	http.HandleFunc("/api/types", handler.GetTypeHandler)
	http.HandleFunc("/api/partsofspeech", handler.GetPartsOfSpeechHandler)
	// http.HandleFunc("/api/userRegist", handler.UserRegistHandler)
	http.ListenAndServe(":8888", nil)
}
