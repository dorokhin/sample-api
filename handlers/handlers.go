package handlers

import (
	"encoding/json"
	"github.com/dorokhin/sample-api/models"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	users = append(users, models.User{ID: 1, FirstName: "Andrew", LastName: "Dorokhin", Email: "andrew@dorokhin.moscow"})
	users = append(users, models.User{ID: 2, FirstName: "Vasily", LastName: "Gorev", Email: "vasya_gor@email.com"})
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}
