package metric

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/api/heartbeat", h.Heartbeat)

}

// Heartbeat
// @Summary Heartbeat metric
// @Tags Metric
// @Success 204
// @Failure 400
// @Router /api/heartbeat [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	log.Print("heartbeat")
	w.WriteHeader(204)
}
