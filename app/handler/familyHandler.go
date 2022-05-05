package handler

import "net/http"

type FamilyHandler interface {
	GetAllFamilies(w http.ResponseWriter, r *http.Request)
}
