package pno

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (pno *PNOService) startServer() error {
	pno.server = &http.Server{
		Addr:    pno.cfg.ListenAddress.Host,
		Handler: pno.getRouter(),
	}

	pno.log.WithField("listenAddr", pno.cfg.ListenAddress.String()).Info("Started PNO server")

	return pno.server.ListenAndServe()
}

func (pno *PNOService) stopServer() error {
	if pno.server == nil {
		return nil
	}

	err := pno.server.Close()
	if err != nil {
		return err
	}

	pno.server = nil

	return nil
}

func (pno *PNOService) ListenAddress() string {
	return pno.cfg.ListenAddress.String()
}

func (pno *PNOService) getRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc(pathRoot, pno.handleRoot).Methods(http.MethodGet)

	r.Use(mux.CORSMethodMiddleware(r))
	loggedRouter := LoggingMiddleware(pno.log, r)
	return loggedRouter
}
