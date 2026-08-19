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
		http.Error(w, "audit failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
