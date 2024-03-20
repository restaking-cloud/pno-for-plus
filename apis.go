package pno

import (
	"net/http"
)

const (
	// Router paths
	pathRoot = "/"
)

func (pno *PNOService) handleRoot(w http.ResponseWriter, _ *http.Request) {
	pno.respondOK(w, "Preferred Node Operator module is running")
}
