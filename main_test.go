package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddData struct {
	a, b   int
	result int
}

func TestSum(t *testing.T) {
	testData := []AddData{
		{1, 3, 4},
		{2, 2, 4},
		{5, 6, 11},
	}

	for _, datas := range testData {
		result := Sum(datas.a, datas.b)

		if result != datas.result {
			t.Errorf("Expected %v got %v", datas.result, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	testData := []AddData{
		{2, 2, 4},
		{6, 6, 36},
		{8, 7, 56},
		{5, 5, 25},
		{6, 2, 12},
	}

	for _, datas := range testData {
		result := Multiply(datas.a, datas.b)

		if result != datas.result {
			t.Errorf("Expected %v got %v", datas.result, result)
		}
	}
}

//https://github.com/gin-gonic/gin#testing
func TestGetping(t *testing.T) {
	router := SetupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	expected := `{"message":"pong"}`

	if w.Body.String() != expected {
		t.Errorf("Expected %v got %v", expected, w.Body.String())
	}
}

func TestGetUsers(t *testing.T) {
	router := SetupServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	// add authorization header to the req
	req.Header.Add("Authorization", "test")
	router.ServeHTTP(w, req)

	expected := `{"message":"{'name': 'test','age': 22}"}`

	if w.Body.String() != expected {
		t.Errorf("Expected %v got %v", expected, w.Body.String())
	}
}
