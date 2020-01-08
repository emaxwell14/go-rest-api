package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetErrorResponse_statusCode(t *testing.T) {
	w := httptest.NewRecorder()
	SetErrorResponse(w, http.StatusBadRequest, "This is a custom message")
	res := w.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code was incorrect, got: %d, wanted: %d.", res.StatusCode, http.StatusBadRequest)
	}
}
func TestSetErrorResponse_defaultResponseMessage(t *testing.T) {
	w := httptest.NewRecorder()
	SetErrorResponse(w, http.StatusBadRequest, "")
	res := w.Result()

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	expected := http.StatusText(http.StatusBadRequest)
	actual := string(bodyBytes)

	if expected != actual {
		t.Errorf("Default response string is incorrect, got: %s, wanted: %s.", actual, expected)
	}
}

func TestSetErrorResponse_customResponseMessage(t *testing.T) {
	customMessage := "This is a custom message"
	w := httptest.NewRecorder()
	SetErrorResponse(w, http.StatusBadRequest, customMessage)
	res := w.Result()

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	actual := string(bodyBytes)

	if customMessage != actual {
		t.Errorf("Custom response string is incorrect, got: %s, wanted: %s.", actual, customMessage)
	}
}
