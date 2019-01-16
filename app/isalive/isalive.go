package isalive

import (
	"net/http"

	raven "github.com/getsentry/raven-go"
	"github.com/gorilla/mux"

	"github.com/iandkenzt/kafka-golang/utils"
)

// Alive Struct
type Alive struct {
	IsAlive bool `json:"is_alive"`
}

// IsAlive Check machine status
func IsAlive(res http.ResponseWriter, req *http.Request) {

	var status Alive
	status.IsAlive = true

	utils.Response(res, 0, "Success", status, http.StatusOK, nil)
}

// Raven ...
var Raven = raven.RecoveryHandler

// BuildIsaliveRoutes ...
func BuildIsaliveRoutes(r *mux.Router) {
	r.Methods("GET").Path("/is_alive").HandlerFunc(Raven(IsAlive))
}
