package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestUrlShortener(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	var jsonStr = []byte(`{"url":"google.com"}`)
	req, err := http.NewRequest("POST", "/submit", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(submitHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	r, _ := regexp.Compile("is ([a-zA-Z0-9]{8}) and")
	res := r.FindStringSubmatch(rr.Body.String())
	if r.MatchString(rr.Body.String()) != true {
		t.Errorf("handler returned unexpected body")
	} else {
		t.Log(res[1])
	}
}
