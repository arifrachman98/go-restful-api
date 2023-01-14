//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitDatabaseRepos() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgres,
		NewDatabaseMysql,
		NewDatabaseRepository,
	)

	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitiFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitFooBar() *FooBar {
	wire.Build(
		NewFoo,
		NewBar,
		wire.Struct(
			new(FooBar),
			"Foo",
			"Bar",
		),
	)
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitFooBarUsingValue() *FooBar {
	wire.Build(
		wire.Value(
			fooValue,
		), wire.Value(
			barValue,
		), wire.Struct(
			new(FooBar),
			"*",
		),
	)

	return nil
}

func InitReader() io.Reader {
	wire.Build(
		wire.InterfaceValue(
			new(io.Reader),
			os.Stdin,
		),
	)
	return nil
}

func InitConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(
			new(*Application),
			"Configuration",
		),
	)
	return nil
}

func InitConnection(name string) (*Connection, func()) {
	wire.Build(
		NewConnection,
		NewFile,
	)

	return nil, nil
}
