package pno

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/sirupsen/logrus"
)

func httpClientDisallowRedirects(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}

var (
	errHTTPErrorResponse = errors.New("HTTP error response")
)

type httpErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (pno *PNOService) respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := httpErrorResp{code, message}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		pno.log.WithField("response", resp).WithError(err).Error("Couldn't write error response")
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (pno *PNOService) respondOK(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		pno.log.WithField("response", response).WithError(err).Error("Couldn't write OK response")
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func LoggingMiddleware(logger *logrus.Entry, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				method := ""
				url := ""
				if r != nil {
					method = r.Method
					url = r.URL.EscapedPath()
				}

				logger.WithFields(logrus.Fields{
					"err":    err,
					"trace":  string(debug.Stack()),
					"method": r.Method,
				}).Error(fmt.Sprintf("http request panic: %s %s", method, url))
			}
		}()
		start := time.Now()
		fmt.Println("\n*******************************")
		fmt.Println("*******************************")
		logger.Info(r.RequestURI)
		next.ServeHTTP(w, r)
		logger.WithFields(logrus.Fields{
			"method":   r.Method,
			"path":     r.URL.EscapedPath(),
			"duration": fmt.Sprintf("%f", time.Since(start).Seconds()),
		}).Info(fmt.Sprintf("http: %s %s", r.Method, r.URL.EscapedPath()))
		fmt.Println("*******************************")
		fmt.Println("*******************************\n")
	})
}
