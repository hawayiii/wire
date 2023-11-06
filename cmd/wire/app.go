package main

import "github.com/google/wire"

var wireSet = wire.NewSet(
	wire.Struct(new(app1), "*"),
	wire.Bind(new(interface1), new(*app1)),

	wire.Struct(new(app2), "*"),
	wire.Bind(new(interface2), new(*app2)),

	wire.Struct(new(app3), "*"),
	newApp4,
	wire.Struct(new(app5), "*"),
)

func newApp4() app4 {
	return app4{}
}

type interface1 interface {
	Do1() string
}

type app1 struct {
	interface2
	//app3
	app5
}

type app3 struct {
	app5
	app4
}

type app4 struct {
}

type app5 struct {
}

func (a *app1) Do1() string {
	return "app1"
}

type interface2 interface {
	Do2() string
}

type app2 struct {
	interface1
}

func (a *app2) Do2() string {
	return "app2"
}

//type app3 struct {
//	//Simida string
//	App1 *app1
//}
//
//type app4 struct {
//	Number int
//}
//
//func newApp4() *app4 {
//	return &app4{}
//}
