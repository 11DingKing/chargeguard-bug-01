package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	ResetTaskHTTPState()
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("POST", "/task", nil))
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("status=%d", rr.Code)
	}
	snapshot := dispatchStore.Snapshot()
	if snapshot.Status != "reported" || snapshot.Owner != "" || snapshot.Audits != 0 {
		t.Fatalf("snapshot=%+v", snapshot)
	}
}
