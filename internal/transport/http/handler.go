package http

import (
	"encoding/json"
	"net/http"
	"tictactoe/internal/domain"
)

type Handler struct {
	svc domain.Service
	mux *http.ServeMux
}

func NewHandler(svc domain.Service) *Handler {
	h := &Handler{svc: svc, mux: http.NewServeMux()}
	h.mux.HandleFunc("POST /game/{uuid}", h.handleMove)
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) handleMove(w http.ResponseWriter, r *http.Request) {
	uuid := r.PathValue("uuid")
	if uuid == "" {
		http.Error(w, "missing uuid", http.StatusBadRequest)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	g, err := h.svc.Load(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if done, _ := h.svc.Done(g.Board); done {
		http.Error(w, domain.ErrGameAlreadyOver.Error(), http.StatusBadRequest)
		return
	}

	b := fromRequest(&req)

	if err := h.svc.Validate(g, b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	done, winner := h.svc.Done(b)
	if !done {
		b = h.svc.Move(b)
		done, winner = h.svc.Done(b)
	}

	if err := h.svc.Save(&domain.Game{UUID: uuid, Board: b}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toResponse(&domain.Game{UUID: uuid, Board: b}, winner, done))
}
