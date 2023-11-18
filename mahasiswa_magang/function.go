package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/intern-monitoring/backend-intermoni/mahasiswa_magang"
)

func init() {
	functions.HTTP("InternMonitoring", internMonitoring_MahasiswaMagang)
}

func internMonitoring_MahasiswaMagang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://intern-monitoring.github.io")
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, mahasiswa_magang.Post("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, mahasiswa_magang.Put("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, mahasiswa_magang.Delete("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, mahasiswa_magang.Get("PASETOPUBLICKEY", "MONGOSTRING", "db_intermoni", r))
}
