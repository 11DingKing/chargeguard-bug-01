package httpapi

import (
	"chargeguard/internal/charging"
	"errors"
	"net/http"
)

var dispatchStore = charging.NewDispatchStore()

func ResetTaskHTTPState() { dispatchStore = charging.NewDispatchStore() }
func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	dispatchStore.FailAudit = true
	err := dispatchStore.Dispatch("operator-a")
	if errors.Is(err, charging.ErrDispatchAudit) {
		snapshot := dispatchStore.Snapshot()
		if snapshot.Status != "reported" || snapshot.Owner != "" || snapshot.Audits != 0 {
			http.Error(w, "dispatch rollback incomplete", http.StatusInternalServerError)
			return
		}
		http.Error(w, "audit failed and dispatch rolled back", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
