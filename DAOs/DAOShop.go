package DAOs

import (
	"inotas-back/database"
	"inotas-back/models"
	"strings"
)

type DAOShop struct {}

func (dao *DAOShop) UpdateProductsCategories(products []models.Product) (error models.Error) {
	con, err := Database.OpenConnection()
	defer con.Close()

	if err != nil {
		return models.ErrorResponse(err, 500)
	}

	for _, product := range products {
		if err := dao.updateProductCategory(product, con); err != nil {
			return models.ErrorResponse(err, 500)
		}
	}


	return
}

func (dao *DAOShop) updateProductCategory(product models.Product, con Database.Connection) (err error){

	query := "UPDATE product SET category_id = $1 WHERE id = $2;"
	stmt, err := con.GetDB().Prepare(query)

	if err != nil {
		return err
	}

	_,err = stmt.Exec(product.CategoryId, product.Id)

	return err
}

func (dao *DAOShop)  GetShop(email string) (shops []models.ShopRequest,error models.Error){
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "SELECT sp.nfe_key, sp.alias, sp.date, sp.payment, sp.total_cost, sl.cnpj, sl.name, sl.state_initials, city.name FROM shop sp, seller sl, city where sp.seller_cnpj = sl.cnpj and sp.user_email = $1 and city.id = sl.city_id;"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil {
		error = models.ErrorResponse(err,500)
		return
	}

	rows, err := stmt.Query(email)

	if err != nil {
		error = models.ErrorResponse(err,500)
		return
	}
	for rows.Next(){
		var shop models.ShopRequest
		rows.Scan(&shop.NFeKey, &shop.Alias, &shop.Date, &shop.Payment, &shop.TotalCost, &shop.Seller.Cnpj,
			&shop.Seller.Name, &shop.Seller.StateInitials, &shop.Seller.City)

		shops = append(shops,shop)
	}


	return
}


func (dao *DAOShop) UpdateShopAlias(shop models.Shop) (error models.Error){
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "UPDATE shop SET alias = $1 WHERE nfe_key = $2;"
	stmt, err :=con.GetDB().Prepare(query)
	if err != nil {
		return models.ErrorResponse(err, 500)
	}

	upperAlias := strings.ToUpper(shop.Alias)
	_, err = stmt.Exec(strings.TrimSpace(upperAlias), shop.NFeKey)
	return
}

func (dao DAOShop) SaveShop(shop models.Shop) (err error){
	con, err := Database.OpenConnection()

	stmt, _ := con.GetDB().Prepare("INSERT INTO shop(nfe_key, total_cost, payment, date, user_email, seller_cnpj, alias) values($1,$2,$3,$4,$5,$6,$7)")
	_ ,err = stmt.Exec(shop.NFeKey, shop.TotalCost, shop.Payment, shop.Date, shop.UserEmail, shop.SellerCnpj, shop.Alias)
	con.Close()
	return
}

