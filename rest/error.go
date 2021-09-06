package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bradleyshawkins/rent"
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Errors interface {
	Message() string
	ErrorCode() string
}

func ErrorHandler(h func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Printf("Unexpected Error: %v", err)
			statusCode := http.StatusInternalServerError
			msg := "An unknown error occurred"
			code := rent.CodeUnknown
			re, ok := err.(*rent.Error)
			if ok {
				statusCode = re.HttpStatusCode()
				msg = re.UserMessage()
				code = re.Code()
			}
			w.WriteHeader(statusCode)
			_ = json.NewEncoder(w).Encode(Error{
				Message: msg,
				Code:    int(code),
			})
		}
	}
}
