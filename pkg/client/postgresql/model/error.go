package model

import "log"

func ErrCommit(err error) error {
	log.Printf("failed to commit error: %s", err)
	return err

}

func ErrRollback(err error) error {
	log.Printf("failed to rollback error: %s", err)
	return err
}

func ErrCreateTx(err error) error {
	log.Printf("failed to create tx error: %s", err)
	return err
}

func ErrCreateQuery(err error) error {
	log.Printf("failed to create SQL query error: %s", err)
	return err
}

func ErrScan(err error) error {
	log.Printf("failed to scan error: %s", err)
	return err

}
func ErrDoQuery(err error) error {
	log.Printf("failed to query error: %s", err)
	return err
}
