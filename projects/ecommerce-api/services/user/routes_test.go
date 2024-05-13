package user

import (
	"bytes"
	"ecommerce-api/types"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) GetUserById(id string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
func TestUser(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if user payload invalid",
		func(t *testing.T) {
			payload := types.RegisterUserPayload{
				FirstName: "user",
				LastName:  "1",
				Email:     "",
				Password:  "123",
			}
			marshalled, _ := json.Marshal(payload)
			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
			if err != nil {
				log.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/register", handler.handleRegister)
			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusBadRequest {
				t.Errorf("expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
			}
		})
}
