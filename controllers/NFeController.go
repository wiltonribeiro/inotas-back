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
		Products []models.Produto `json:"products"`
		Items []models.Item `json:"items"`
		Shop models.Compra `json:"shop"`
		Seller models.Vendedor `json:"seller"`
	}{}

	for i:=0;i<cap(c);i++ {
		data := <-c
		switch data.(type) {
		case []models.Produto:
			jsonFormat.Products = data.([]models.Produto)
		case []models.Item:
			jsonFormat.Items = data.([]models.Item)
		case models.Compra:
			jsonFormat.Shop = data.(models.Compra)
		case models.Vendedor:
			jsonFormat.Seller = data.(models.Vendedor)
		}
	}


	err = controller.saveSeller(&jsonFormat.Seller)
	err = controller.saveShop(jsonFormat.Shop)
	err = controller.saveProducts(jsonFormat.Products)
	err = controller.saveItems(jsonFormat.Items)

	return jsonFormat, err
}

func (controller NFeController) saveProducts(products []models.Produto) (err error){
	for _, product := range products {
		stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO produto(id, code, descricao, valor_unidade, unidade, id_categoria) values($1,$2,$3,$4,$5,$6)")
		_,err = stmt.Exec(product.Id, product.Code, product.Descricao, product.ValorUnidade, product.Unidade, product.CetegoriaId)
	}
	return
}

func (controller NFeController) saveItems(items []models.Item) (err error){
	for _, item := range items {
		stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO item(valor_total, quantidade, id_produto, compra_key) values($1,$2,$3,$4)")
		_ ,err = stmt.Exec(item.ValorTotal, item.Qntd, item.ProdutoId, item.CompraNFeKey)
	}
	return
}

func (controller NFeController) saveShop(shop models.Compra) (err error){
	stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO compra(nfe_key, valor_total, pagamento, data_emissao, usuario_email, vendedor_cnpj) values($1,$2,$3,$4,$5,$6)")
	_ ,err = stmt.Exec(shop.NFeKey, shop.Total, shop.FormaPagamento, shop.DataCompra, shop.UsuarioEmail, shop.VendedorCnpj)
	return
}

func (controller NFeController) saveSeller(seller* models.Vendedor) (err error){
	stmt, _ := controller.DataBase.GetDB().Prepare("INSERT INTO vendedor(cnpj,nome,rua,numero,cep,info_adicional,distrito,id_cidade,sigla_estado)" +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)" +
		"ON     CONFLICT (cnpj) DO UPDATE " +
		"SET    nome = excluded.nome," +
		"rua = excluded.rua," +
		"numero = excluded.numero," +
		"cep = excluded.cep," +
		"info_adicional = excluded.info_adicional," +
		"distrito = excluded.distrito," +
		"id_cidade = excluded.id_cidade,sigla_estado = excluded.sigla_estado;")

	locationController := LocationController{controller.DataBase}
	seller.IdCidade = locationController.GetIdCityByStateAndName(seller.SiglaEstado,strings.ToUpper(seller.Cidade))
	_ ,err = stmt.Exec(seller.Cnpj, seller.Nome, seller.Rua, seller.Numero, seller.Cep, seller.InfoAdicional, seller.Distrito, seller.IdCidade, seller.SiglaEstado)
	return
}