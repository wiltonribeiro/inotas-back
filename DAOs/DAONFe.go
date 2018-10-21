package DAOs

import (
	"inotas-back/models"
	"inotas-back/database"
	"strings"
)

type DAONFe struct {}

func (dao DAONFe) SaveProducts(products []models.Product) (err error){
	con, err := Database.OpenConnection()

	for _, product := range products {
		stmt, _ := con.GetDB().Prepare("INSERT INTO product(id, code, description, unity_cost, unity, category_id) values($1,$2,$3,$4,$5,$6)")
		_,err = stmt.Exec(product.Id, product.Code, product.Description, product.UnityCost, product.Unity, product.CategoryId)
	}
	con.Close()
	return
}

func (dao DAONFe) SaveItems(items []models.Item) (err error){
	con, err := Database.OpenConnection()

	for _, item := range items {
		stmt, _ := con.GetDB().Prepare("INSERT INTO item(total_cost, quantity, product_id, shop_key) values($1,$2,$3,$4)")
		_ ,err = stmt.Exec(item.TotalCost, item.Quantity, item.ProductId, item.ShopKey)
	}
	con.Close()
	return
}

func (dao DAONFe) SaveShop(shop models.Shop) (err error){
	con, err := Database.OpenConnection()

	stmt, _ := con.GetDB().Prepare("INSERT INTO shop(nfe_key, total_cost, payment, date, user_email, seller_cnpj, alias) values($1,$2,$3,$4,$5,$6,$7)")
	_ ,err = stmt.Exec(shop.NFeKey, shop.TotalCost, shop.Payment, shop.Date, shop.UserEmail, shop.SellerCnpj, shop.Alias)
	con.Close()
	return
}

func (dao DAONFe) SaveSeller(seller* models.Seller) (err error){

	con, err := Database.OpenConnection()

	stmt, err := con.GetDB().Prepare("INSERT INTO seller(cnpj, name, street, number, postal_code, other_info, district, city_id, state_initials, city)" +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9, $10)" +
		"ON     CONFLICT (cnpj) DO UPDATE " +
		"SET    name = excluded.name," +
		"street = excluded.street," +
		"number = excluded.number," +
		"postal_code = excluded.postal_code," +
		"other_info = excluded.other_info," +
		"district = excluded.district," +
		"city = excluded.city," +
		"city_id = excluded.city_id,state_initials = excluded.state_initials;")

	DAOCity := DAOCity{}
	seller.CityId, _ = DAOCity.GetCityIdByStateAndName(seller.StateInitials,strings.ToUpper(seller.City))
	_ ,err = stmt.Exec(seller.Cnpj, seller.Name, seller.Street, seller.Number, seller.PostalCode, seller.OtherInfo, seller.District, seller.CityId, seller.StateInitials, seller.City)
	return
}