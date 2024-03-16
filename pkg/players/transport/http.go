package transport

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ibad69/go.fields/pkg/players"
	repo "github.com/ibad69/go.fields/pkg/players/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	playerService players.Service
}

func Activate(ro chi.Router, db *mongo.Client) {

	// get the player service but doing new on it and also give the repo / database initialization to it
	ps := players.New(repo.New(db))
	NewHandler(ps, ro)

}

func NewHandler(ps players.Service, ro chi.Router) {
	h := &handler{playerService: ps}
	ro.Route("/players", func(r chi.Router) {
		r.Get("/getPlayers", h.GetPlayers)
		r.Post("/signUp", h.SignUp)
		// r.Get("/{id}", h.GetPlayer)
		r.Put("/{id}", h.UpdatePlayer)
		r.Delete("/{id}", h.DeletePlayer)
	})
	// ro.Mount('/players', router)
}

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	// get all the players
	results, err := h.playerService.Get(r.Context(), "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (h *handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {

}

func (h *handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {

}
func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var p players.Player
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	results, err := h.playerService.SignUp(r.Context(), p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
