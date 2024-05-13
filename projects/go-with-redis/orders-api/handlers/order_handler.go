package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"orders-api/models"
	"orders-api/repository"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type OrderHandler struct {
	Repo *repository.OrderRepo
}

func (o *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CustomerID uuid.UUID         `json:"customer_id"`
		LineItems  []models.LineItem `json:"line_items"`
	}

	// Decode request body to struct
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()
	order := models.Order{
		OrderID:    rand.Uint64(),
		CustomerID: body.CustomerID,
		LineItems:  body.LineItems,
		CreatedAt:  &now,
	}

	err := o.Repo.Insert(r.Context(), order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Return data to client
	res, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)
}

func (o *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	cursorStr := r.URL.Query().Get("cursor")
	if cursorStr == "" {
		cursorStr = "0"
	}

	const decinal = 10
	const bitSize = 10
	cursor, err := strconv.ParseUint(cursorStr, decinal, bitSize)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const size = 50
	res, err := o.Repo.FindAll(r.Context(), models.OrderPaging{
		Size:   size,
		Offset: uint(cursor),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response struct {
		Items []models.Order `json:"items"`
		Next  uint64         `json:"next,omitempty"`
	}
	response.Items = res.Orders
	response.Next = res.Cursor

	data, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (o *OrderHandler) GetById(w http.ResponseWriter, r *http.Request) {

}

func (o *OrderHandler) UpdateById(w http.ResponseWriter, r *http.Request) {

}

func (o *OrderHandler) DeleteById(w http.ResponseWriter, r *http.Request) {

}
