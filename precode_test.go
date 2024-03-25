package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// тест для сценария: запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
}

// тест для сценария: город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа
func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=podolsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

// тест для сценария: если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=6&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	}

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, len(list), totalCount)
}
