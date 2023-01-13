package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMySQL Database

func NewDatabasePostgres() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{
		Name: "PostgreSQL",
	})
}

func NewDatabaseMysql() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{
		Name: "MySQL",
	})
}

type DatabaseRepository struct {
	DBpostgres *DatabasePostgreSQL
	DBmysql    *DatabaseMySQL
}

func NewDatabaseRepository(pstg *DatabasePostgreSQL, mysql *DatabaseMySQL) *DatabaseRepository {
	return &DatabaseRepository{
		DBpostgres: pstg,
		DBmysql:    mysql,
	}
}
