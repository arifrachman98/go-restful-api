//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
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
