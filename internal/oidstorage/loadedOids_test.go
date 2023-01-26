package oidstorage

import "testing"

func TestSomething(t *testing.T) {
	db := InitializeDb()
	// really just here to catch lack of testing coverage
	// TODO : write unit tests for this struct
	NewLoadedOids(db)
}
