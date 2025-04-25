package reg

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/niiilov/e-commerce/internal/db"
	"github.com/niiilov/e-commerce/internal/models"
)

func RegHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var response models.Response

	response.Success = false

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "Invalid JSON format"
		sendResponse(w, http.StatusBadRequest, response) //400
		return
	}

	if len(user.Username) < 3 {
		response.Message = "Invalid Username format"
		sendResponse(w, http.StatusBadRequest, response) //400
		return
	}

	if len(user.Password) < 8 {
		response.Message = "Invalid Password format"
		sendResponse(w, http.StatusBadRequest, response) //400
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		response.Message = "Invalid Email format"
		sendResponse(w, http.StatusBadRequest, response) //400
		return
	}

	user.HashPassword()

	_, err = db.DB.Exec("INSERT INTO user (username, email, pass) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		response.Message = "Invalid DB"
		sendResponse(w, http.StatusInternalServerError, response) //500
		return
	}

	response.Success = true
	response.Message = "User registred"
	sendResponse(w, http.StatusCreated, response) //201
}

func sendResponse(w http.ResponseWriter, status int, response models.Response) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
