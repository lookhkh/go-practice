package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lookhkh/data-downlodaer/pkgquery"
)

func setartTestHttpServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello World")
			}))
}

func TestFetchPackageData(t *testing.T) {
	ts := setartTestHttpServer()
	defer ts.Close()
	packages, err := pkgquery.FetchPackageData(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if len(packages) != 2 {
		t.Fatalf("Expected 2 but %v", len(packages))
	}
}

func TestFetchRemoteData(t *testing.T) {
	ts := setartTestHttpServer()
	defer ts.Close()
	expected := "Hello World"
	data, err := fetchRemoteResource(ts.URL)

	if err != nil {
		t.Fatal(err)
	}

	if expected != string(data) {
		t.Errorf("expected response to be %v, got : %v", expected, string(data))
	}

}
