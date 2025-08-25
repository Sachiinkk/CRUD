package student

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Sachiinkk/student-api/internal/response"
	"github.com/Sachiinkk/student-api/internal/types"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return

		}

		if r.Method != http.MethodPost {
			http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}
		//validator

		if err := validator.New().Struct(student); err != nil{
			validatesError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest , response.Validation(validatesError))
			return 
		}
		response.WriteJson(w, http.StatusCreated, map[string]string{"Success": "OK"})

	}

}
