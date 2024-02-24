package db

import (
	"context"
	"database/sql"
	"fmt"
)

/*
	This store will provide all functions to run database queries individually,
	as well as their combination within a transaction.
*/

/*
However, each query only do 1 operation on 1 specific table.
So Queries struct doesn’t support transaction, that’s why we have to extend its functionality
by embedding it inside the Store struct like this. By embedding Queries inside Store, all individual query functions provided by Queries will be available to Store.
*/
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

/*
A context and a callback function as input, then it will start a new db transaction,
create a new Queries object with that transaction call the callback function with the created Queries,
and finally commit or rollback the transaction based on the error returned by that function.
*/
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Call New() function with the created transaction tx, and get back a new Queries object.
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
