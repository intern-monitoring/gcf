package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/intern-monitoring/backend-intermoni/module"
)

func init() {
	functions.HTTP("InternMonitoring", internMonitoringGetMagang)
}

func internMonitoringGetMagang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://intern-monitoring.github.io")
	fmt.Fprintf(w, module.GCFHandlerGetAllMagang("MONGOSTRING", "db_intermoni"))
}
