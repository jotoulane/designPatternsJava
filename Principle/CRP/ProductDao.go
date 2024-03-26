package CRP

type ProductDao struct {
	dbConn DBConnection
}

func NewProductDao(dbConn DBConnection) *ProductDao {
	return &ProductDao{dbConn: dbConn}
}

func (p ProductDao) addProduct() {
	connection := p.dbConn.GetConnection()
	println(connection)
}
