package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/intern-monitoring/backend-intermoni/mentor"
)

func init() {
	functions.HTTP("InternMonitoring", internMonitoring_Mentor)
}

func internMonitoring_Mentor(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	w.Header().Set("Access-Control-Allow-Origin", "https://intern-monitoring.github.io")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method == http.MethodPost{
		fmt.Fprintf(w, mentor.Post("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))
		return
	}
	if r.Method == http.MethodPut{
		fmt.Fprintf(w, mentor.Put("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, mentor.Get("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))

}
