package helper

import (
	"database/sql"
)

func RollbackOrCommit(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollaback := tx.Rollback()
		PanicIFError(errorRollaback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIFError(errorCommit)
	}
}
