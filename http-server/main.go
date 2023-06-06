package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type requestContextKey struct{}
type requestContextValue struct {
	requestID string
}
type logLine struct {
	UserIP string `json : "user_ip"`
	Event  string `json : "event"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)

	for {
		var l logLine
		err := dec.Decode(&l)
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(l.UserIP, l.Event)
	}
	fmt.Fprint(w, "ok")

}

func addRequestID(r *http.Request, requestID string) *http.Request {
	c := requestContextValue{
		requestID: requestID,
	}
	currentCtx := r.Context()
	newContext := context.WithValue(currentCtx, requestContextKey{}, c)
	return r.WithContext(newContext)
}

func logRequest(r *http.Request) {
	ctx := r.Context()
	v := ctx.Value(requestContextKey{})

	if m, ok := v.(requestContextKey); ok {
		log.Printf("processing req ID : %s", m)
	}
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprintf(w, "request process")
}

func main() {

	listenAddr := "localhost:8080"
	mux := http.NewServeMux()
	setUpHanlder(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}

func setUpHanlder(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/healthz", healthCheckHanlder)
	mux.HandleFunc("/decode", decodeHandler)
	mux.HandleFunc("/log", longRunningProcessHandler)

}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	requestID := "api-123"
	r = addRequestID(r, requestID)
	processRequest(w, r)
	fmt.Fprint(w, "Hello, world")

}

func healthCheckHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func longRunningProcess(logwriter *io.PipeWriter) {
	for i := 0; i <= 20; i++ {
		fmt.Fprintf(logwriter, `{"id" : %d, "user_ip":"123.123.123.1", "event":"click"}`, i)
		time.Sleep(1 * time.Second)
	}
	logwriter.Close()

}

func longRunningProcessHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan struct{})
	logReader, logWriter := io.Pipe()
	go longRunningProcess(logWriter)
	go processStreamer(logReader, w, done)
	<-done
}

func processStreamer(logReader *io.PipeReader, w http.ResponseWriter, done chan struct{}) {
	buf := make([]byte, 500)
	f, flushSupported := w.(http.Flusher)
	defer logReader.Close()

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	for {
		n, err := logReader.Read(buf)
		if err == io.EOF {
			break
		}

		w.Write(buf[:n])
		if flushSupported {
			f.Flush()
		}
	}

	done <- struct{}{}

}
