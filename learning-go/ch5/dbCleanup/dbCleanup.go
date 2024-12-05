package exercises

import (
	"context"
	"database/sql"
)

/*
You might be wondering whether there's a way for a deferred
function to examine or modify th ereturn values of its surrounding
function. There is, and it's the best reason to use named return
values. It allows your code to take actions based on an error.
When I talk about errors in Chapter 9, I will discuss a pattern
that uses `defer` to add contextual information to an error
returned from a function. Let's look at a way to handle database
transaction cleanup using named return values and defer:
*/
func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		// detect additional errors during transaction commit after function completes
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) values $1", value1)
	if err != nil {
		return err
	}
	// use tx to do more databse inserts here
	return nil
}
