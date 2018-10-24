package DAOs

import (
	"inotas-back/models"
	"inotas-back/database"
)

type DAOItem struct {}

func (dao DAOItem) SaveItems(items []models.Item) (err error){
	con, err := Database.OpenConnection()

	for _, item := range items {
		stmt, _ := con.GetDB().Prepare("INSERT INTO item(total_cost, quantity, product_id, shop_key) values($1,$2,$3,$4)")
		_ ,err = stmt.Exec(item.TotalCost, item.Quantity, item.ProductId, item.ShopKey)
	}
	con.Close()
	return
}

func (dao DAOItem) GetItems(shopKey string) (items []models.ItemRequest, error models.Error){
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "select i.id, i.total_cost, i.quantity, p.* from item i, product p where i.shop_key = $1 and i.product_id = p.id;"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil {
		error = models.ErrorResponse(err,500)
		return
	}

	rows, err := stmt.Query(shopKey)

	if err != nil {
		error = models.ErrorResponse(err,500)
		return
	}
	for rows.Next(){
		var item models.ItemRequest
		rows.Scan(&item.Id, &item.TotalCost, &item.Quantity, &item.Product.Id, &item.Product.Code, &item.Product.Description,
			&item.Product.UnityCost, &item.Product.Unity, &item.Product.CategoryId)

		items = append(items,item)
	}
	return
}
