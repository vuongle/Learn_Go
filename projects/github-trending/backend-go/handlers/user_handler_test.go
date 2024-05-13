package handlers

import (
	"encoding/json"
	"fmt"
	"github-trending-api/helper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Create a test case called "TestHandleSignUp"
func TestHandleSignUp(t *testing.T) {

	t.Cleanup(func() {
		t.Log("clean up the test case")
	})

	t.Run("should return 400 status ok", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/user/sign-up", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := UserHandler{}
		handler.HandleSignUp(c)

		res := rec.Body.String()
		fmt.Println(res)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should validate email fail", func(t *testing.T) {

		var userJSON = `{"full_name":"Jhon Doe","email":"doe@gmail.com", "password":"123"}`

		e := echo.New()
		// Init a custom validator and assign it to echo validator
		structValidator := helper.NewStructValidator()
		structValidator.RegisterValidate()
		e.Validator = structValidator

		req := httptest.NewRequest(http.MethodPost, "/user/sign-up", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		handler := UserHandler{}
		handler.HandleSignUp(c)

		// convert json string to map
		var decodedData map[string]interface{}
		json.NewDecoder(rec.Body).Decode(&decodedData)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, "Mật khẩu tối thiểu 4 kí tự", decodedData["message"])
	})
}
