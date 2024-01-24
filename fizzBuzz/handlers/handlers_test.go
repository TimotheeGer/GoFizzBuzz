package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestCheckQueryParamIsValid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?int1=3&int2=5&limit=10&str1=fizz&str2=buzz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CheckQueryParam(c)

	if err != nil {
		t.Errorf("CheckQueryParam a retourné une erreur pour des paramètres valides: %v", err)
	}
}

func TestCheckQueryParamMissing(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?int1=3&int2=5&limit=10&str1=fizz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CheckQueryParam(c)

	if err == nil {
		t.Errorf("CheckQueryParam n'a pas retourner d'erreur pour le Nombre de paramètres manquan: %v", err)
	}
}

func TestCheckQueryParamTooMuch(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?int1=3&int2=5&limit=10&str1=fizz&str2=buzz&str3=lizz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CheckQueryParam(c)

	if err == nil {
		t.Errorf("CheckQueryParam n'a pas retourner d'erreur pour le Nombre de paramètres en trop: %v", err)
	}
}

func TestCheckQueryParamWhitIntString(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?int1=bonjour&int2=5&limit=10&str1=fizz&str2=buzz&str3=lizz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CheckQueryParam(c)

	if err == nil {
		t.Errorf("CheckQueryParam n'a pas retourner d'erreur pour int1=bonjour: %v", err)
	}
}

func TestCheckQueryParamWhitIntNegative(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?int1=-78&int2=5&limit=10&str1=fizz&str2=buzz&str3=lizz", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := CheckQueryParam(c)

	if err == nil {
		t.Errorf("CheckQueryParam n'a pas retourner d'erreur pour int negative: %v", err)
	}
}