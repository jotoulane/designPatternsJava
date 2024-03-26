package CRP

import "testing"

func TestDb(t *testing.T) {
	dao := NewProductDao(NewMySQLConnection())
	dao.addProduct()
}
