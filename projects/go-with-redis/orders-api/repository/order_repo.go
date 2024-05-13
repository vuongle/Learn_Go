package repository

import (
	"context"
	"encoding/json"
	"errors"
	"orders-api/helpers"
	"orders-api/models"
	"orders-api/utils"

	"github.com/redis/go-redis/v9"
)

type OrderRepo struct {
	Client *redis.Client
}

func (r *OrderRepo) Insert(ctx context.Context, order models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	key := helpers.OrderIDKey(uint(order.OrderID))

	// use a pipeline as a transaction
	txn := r.Client.TxPipeline()

	// Save an order with a separated key
	res := r.Client.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		txn.Discard()
		return err
	}

	// Add the order to a list with another key
	// This is used for FindAll() function
	if err := r.Client.SAdd(ctx, "all_orders", key).Err(); err != nil {
		txn.Discard()
		return err
	}

	// Handler transaction
	if _, err := txn.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) FindByID(ctx context.Context, id uint) (models.Order, error) {
	key := helpers.OrderIDKey(uint(id))

	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return models.Order{}, utils.ErrOrderNotExist
	} else if err != nil {
		return models.Order{}, err
	}

	var order models.Order
	err = json.Unmarshal([]byte(value), &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r *OrderRepo) DeleteByID(ctx context.Context, id uint) error {
	key := helpers.OrderIDKey(uint(id))

	// use a pipeline as a transaction
	txn := r.Client.TxPipeline()

	// Delete an order by a separated key
	err := txn.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return utils.ErrOrderNotExist
	} else if err != nil {
		txn.Discard()
		return err
	}

	if err := txn.SRem(ctx, "all_orders", key).Err(); err != nil {
		txn.Discard()
		return err
	}

	// Handler transaction
	if _, err := txn.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) Update(ctx context.Context, order models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	key := helpers.OrderIDKey(uint(order.OrderID))
	err = r.Client.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		return utils.ErrOrderNotExist
	} else if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepo) FindAll(ctx context.Context, paging models.OrderPaging) (models.OrderResult, error) {
	res := r.Client.SScan(ctx, "all_orders", uint64(paging.Offset), "*", int64(paging.Size))
	keys, cursor, err := res.Result()
	if err != nil {
		return models.OrderResult{}, err
	}

	// Return empty data if there is no keys
	if len(keys) == 0 {
		return models.OrderResult{
			Orders: []models.Order{},
		}, nil
	}

	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return models.OrderResult{}, err
	}

	orders := make([]models.Order, len(xs))

	for i, x := range xs {
		x := x.(string)
		var order models.Order

		err := json.Unmarshal([]byte(x), &order)
		if err != nil {
			return models.OrderResult{}, err
		}

		orders[i] = order
	}

	return models.OrderResult{
		Orders: orders,
		Cursor: cursor,
	}, nil
}
