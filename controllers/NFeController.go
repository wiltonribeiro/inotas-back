package controllers

import (
	"inotas-back/models"
	"inotas-back/database"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"strings"
	"inotas-back/enviroment"
)

type NFeController struct {
	DataBase* database.Connection
}

func (controller NFeController) requestNFe(key string) (err error ,NFe* models.NFeRequest){
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://nfe.api.nfe.io/v2/productinvoices/"+key,nil)
	req.Header.Add("Authorization",enviroment.AuthAPI)
	resp, _ := client.Do(req)

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal([]byte(body), &NFe)
		return err, NFe
	}
}

func (controller NFeController) GetContent(email, key string) (interface{}, error){
	var err error = nil
	err , NFe := controller.requestNFe(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	filter := NFeFilter{}
	c := filter.FilterData(email, *NFe)

	jsonFormat := struct {
		Products []models.Product `json:"products"`
		Items []models.Item `json:"items"`
		Shop models.Shop `json:"shop"`
		Seller models.Seller `json:"seller"`
	}{}

	for i:=0;i<cap(c);i++ {
		data := <-c
		switch data.(type) {
		case []models.Product:
			jsonFormat.Products = data.([]models.Product)
		case []models.Item:
			jsonFormat.Items = data.([]models.Item)
		case models.Shop:
			jsonFormat.Shop = data.(models.Shop)
		case models.Seller:
			jsonFormat.Seller = data.(models.Seller)
		}
	}


	err = controller.saveSeller(&jsonFormat.Seller)
	err = controller.saveShop(jsonFormat.Shop)
	err = controller.saveProducts(jsonFormat.Products)
	err = controller.saveItems(jsonFormat.Items)

	return jsonFormat, err
}

func (controller NFeController) saveProducts(products []models.Product) (err error){
	for _, product := range products {
		stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO product(id, code, description, unity_cost, unity, category_id) values($1,$2,$3,$4,$5,$6)")
		_,err = stmt.Exec(product.Id, product.Code, product.Description, product.UnityCost, product.Unity, product.CategoryId)
	}
	return
}

func (controller NFeController) saveItems(items []models.Item) (err error){
	for _, item := range items {
		stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO item(total_cost, quantity, product_id, shop_key) values($1,$2,$3,$4)")
		_ ,err = stmt.Exec(item.TotalCost, item.Quantity, item.ProductId, item.ShopKey)
	}
	return
}

func (controller NFeController) saveShop(shop models.Shop) (err error){
	stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO shop(nfe_key, total_cost, payment, date, user_email, seller_cnpj) values($1,$2,$3,$4,$5,$6)")
	_ ,err = stmt.Exec(shop.NFeKey, shop.TotalCost, shop.Payment, shop.Date, shop.UserEmail, shop.SellerCnpj)
	return
}

func (controller NFeController) saveSeller(seller* models.Seller) (err error){
	stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO seller(cnpj, name, street, number, postal_code, other_info, district, city_id, state_initials)" +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)" +
		"ON     CONFLICT (cnpj) DO UPDATE " +
		"SET    name = excluded.name," +
		"street = excluded.street," +
		"number = excluded.number," +
		"postal_code = excluded.postal_code," +
		"other_info = excluded.other_info," +
		"district = excluded.district," +
		"city_id = excluded.city_id,state_initials = excluded.state_initials;")

	locationController := LocationController{controller.DataBase}
	seller.CityId = locationController.GetIdCityByStateAndName(seller.StateInitials,strings.ToUpper(seller.City))
	_ ,err = stmt.Exec(seller.Cnpj, seller.Name, seller.Street, seller.Number, seller.PostalCode, seller.OtherInfo, seller.District, seller.CityId, seller.StateInitials)
	return
}