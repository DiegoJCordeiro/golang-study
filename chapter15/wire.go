//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import "database/sql"

func NewRepository(db sql.DB) {

}
