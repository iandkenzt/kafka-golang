package producers

import (
	"github.com/getsentry/raven-go"
	"github.com/gorilla/mux"
)

// Raven ...
var Raven = raven.RecoveryHandler

// BuildProducersRoutes ...
func BuildProducersRoutes(r *mux.Router) {
	r.Methods("POST").Path("/user_activity").HandlerFunc(Raven(UserActivity))
}
