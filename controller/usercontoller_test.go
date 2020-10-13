package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"userApi/database"

	"github.com/gin-gonic/gin"
)

func TestFindAll(t *testing.T) {
	handle := func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		database.NewDB()
		FindAll(c)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/users", nil)

	handle(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func TestFindUserbyID(t *testing.T) {
	handle := func(w http.ResponseWriter, r *http.Request) {
		c, _ := gin.CreateTestContext(w)
		database.NewDB()
		FindUserbyID(c)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/users/1", nil)

	handle(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
