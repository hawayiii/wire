//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

// NewApp
func NewApp() *app1 {

	wire.Build(
		wireSet,
	)

	return &app1{}
}
