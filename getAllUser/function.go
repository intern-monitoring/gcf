package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/intern-monitoring/backend-intermoni/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	functions.HTTP("InternMonitoring", internMonitoringGetUser)
}

func internMonitoringGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://intern-monitoring.github.io")
	fmt.Fprintf(w, module.GCFHandlerGetAll("MONGOSTRING", "db_intermoni", "user", []User{}))
}

type User struct {
	ID           	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email  			string             `bson:"email,omitempty" json:"email,omitempty"`
	Role     		string			   `bson:"role,omitempty" json:"role,omitempty"`
}
