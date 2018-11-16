package controllers

import (
	"inotas-back/models"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"inotas-back/enviroment"
	"inotas-back/DAOs"
)

type NFeController struct {}

func (controller NFeController) requestNFe(key string) (NFe* models.NFeRequest, err error){
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
		return NFe, err
	}
}

func (controller NFeController) GetContent(token, key string) (errorR models.Error){
	DAOSeller := DAOs.DAOSeller{}
	DAOProduct := DAOs.DAOProduct{}
	DAOItem := DAOs.DAOItem{}
	DAOShop := DAOs.DAOShop{}

	authControl := AuthController{}
	email, errorR  := authControl.CheckAuth(token)

	if errorR != (models.Error{}) {
		return errorR
	}

	var err error = nil
	NFe, err := controller.requestNFe(key)

	if err != nil {
		return models.ErrorResponse(err, 500)
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





	err = DAOSeller.SaveSeller(&jsonFormat.Seller)
	err = DAOShop.SaveShop(jsonFormat.Shop)
	if err == nil {
		err = DAOProduct.SaveProducts(jsonFormat.Products)
		err = DAOItem.SaveItems(jsonFormat.Items)
		return
	}
	errorR = models.ErrorResponse(err, 500)
	return
}