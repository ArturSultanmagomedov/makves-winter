package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	repo IRepository
}

func NewHandler(repo IRepository) *Handler {
	return &Handler{repo: repo}
}

func (h Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var object struct{ Items []string }
	if err := json.NewDecoder(r.Body).Decode(&object); err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	response := make(map[string]any)
	for _, v := range object.Items {
		item, _ := h.repo.FindItemById(v)
		response[v] = item
	}
	marshal, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(marshal)
}

type IRepository interface {
	FindItemById(id string) (item []string, found bool)
}
