package CRP

type PostgreSQLConnection struct {
}

func NewPostgreSQLConnection() *PostgreSQLConnection {
	return &PostgreSQLConnection{}
}

func (PostgreSQLConnection) GetConnection() string {
	return "postgresql conn"
}
