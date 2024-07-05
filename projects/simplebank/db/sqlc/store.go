package db

import (
	"context"
	"database/sql"
	"fmt"
)

// The Store struct is likely used as a wrapper around the Queries struct to provide a convenient way to perform database
// operations while also keeping a reference to the database connection.
// The store includes all functions in the Queries struct. This is a kind of inheritance (Store inherits from Queries).
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx: executes a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Init a new Queries struct that supports the transaction
	q := New(tx)

	// Call the function within the transaction
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Start a new money transfer between two accounts. It includes the following business logic:
//
// 1.Create a new transfer record
//
// 2.Add account entries
//
//   - One entry describes the transfer amount moves out of the from account (native amount)
//   - One entry describes the transfer amount go in the to account (native amount)
//
// 3.Update account balances
func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		// 1.Create a new transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// 2.Add account entries
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount, //negative because money moves out of arg.FromAccountID
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount, //positive because money goes in of arg.ToAccountID
		})
		if err != nil {
			return err
		}

		// 3.Update accounts' balances
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount, //negative because money moves out of arg.FromAccountID
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount, //positive because money goes in of arg.ToAccountID
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
