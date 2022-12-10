package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckFirstLogin(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	result := check_login("artyom.leskco@gmail.com", "123", w, r)
	if result != true {
		t.Errorf("Incorrect result. Except %t", result)
	}
}

func TestCheckSecondLogin(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	result := check_login("root@root", "root", w, r)
	if result != true {
		t.Errorf("Incorrect result. Except %t", result)
	}
}

func TestCheckFirstFalseLogin(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	result := check_login("artyom.leskco@mail", "123", w, r)
	if result == true {
		t.Errorf("Incorrect result. Except %t", result)
	}
}

func TestCheckSecondFalseLogin(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	result := check_login("root@roots", "root", w, r)
	if result == true {
		t.Errorf("Incorrect result. Except %t", result)
	}
}

func TestCheckFirstFalsePassword(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	result := check_login("artyom.leskco@gmail", "1234", w, r)
	if result == true {
		t.Errorf("Incorrect result. Except %t", result)
	}
}

func TestCheckSecondFalsePassword(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	result := check_login("root@root", "roo", w, r)
	if result == true {
		t.Errorf("Incorrect result. Except %t", result)
	}
}