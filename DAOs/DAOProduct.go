package DAOs

import (
	"inotas-back/models"
	"inotas-back/database"
)

type DAOProduct struct {}

func (dao DAOProduct) SaveProducts(products []models.Product) (err error){
	con, err := Database.OpenConnection()

	for _, product := range products {
		stmt, _ := con.GetDB().Prepare("INSERT INTO product(id, code, description, unity_cost, unity, category_id) values($1,$2,$3,$4,$5,$6)")
		_,err = stmt.Exec(product.Id, product.Code, product.Description, product.UnityCost, product.Unity, product.CategoryId)
	}
	con.Close()
	return
}

