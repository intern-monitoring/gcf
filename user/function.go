package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/intern-monitoring/backend-intermoni/user"
)

func init() {
	functions.HTTP("InternMonitoring", internMonitoring_User)
}

func internMonitoring_User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://intern-monitoring.github.io")
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, user.Get("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))
}
