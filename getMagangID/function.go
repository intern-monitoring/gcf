package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/intern-monitoring/backend-intermoni/module"
)

func init() {
	functions.HTTP("InternMonitoring", internMonitoringGetMagangID)
}

func internMonitoringGetMagangID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://intern-monitoring.github.io")
	fmt.Fprintf(w, module.GCFHandlerGetMagangFromID("MONGOSTRING", "db_intermoni", r))
}
