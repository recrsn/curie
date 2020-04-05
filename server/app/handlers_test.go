package app

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func router() *gin.Engine {
	return gin.Default()
}

func Test_ping(t *testing.T) {
	t.Run("Should return pong", func(t *testing.T) {
		router := router()
		router.GET("/", ping)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "pong", w.Body.String())
	})
}

func Test_versionInfo(t *testing.T) {
	t.Run("Should return version info", func(t *testing.T) {
		router := router()
		router.GET("/", versionInfo)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.JSONEq(t, "{ \"version\": \"" + Version +"\"}", w.Body.String())
	})
}