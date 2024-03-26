package CRP

type MySQLConnection struct {
}

func NewMySQLConnection() *MySQLConnection {
	return &MySQLConnection{}
}

func (MySQLConnection) GetConnection() string {
	return "mysql conn"
}
