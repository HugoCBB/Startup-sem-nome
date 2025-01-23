package controllers

import (
	"backend/api/api/models"
	"encoding/json"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Users)
}
