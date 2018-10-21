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

func (dao *DAOShop)  GetShop(email string) (shops []models.ShopComplete ,error models.Error){
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "SELECT shop.*, seller.* FROM shop INNER JOIN seller ON shop.seller_cnpj = seller.cnpj WHERE shop.user_email = $1;"
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
		var shop models.ShopComplete
		rows.Scan(&shop.NFeKey, &shop.TotalCost, &shop.Payment, &shop.Date, &shop.UserEmail, &shop.SellerCnpj, &shop.Alias,
			&shop.Seller.Cnpj, &shop.Seller.Name, &shop.Seller.Street, &shop.Seller.Number, &shop.Seller.PostalCode, &shop.Seller.OtherInfo, &shop.Seller.District, &shop.Seller.CityId, &shop.Seller.StateInitials)
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

