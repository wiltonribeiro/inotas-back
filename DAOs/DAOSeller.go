package DAOs

import (
	"inotas-back/models"
	"inotas-back/database"
	"strings"
)

type DAOSeller struct {}

func (dao DAOSeller) SaveSeller(seller* models.Seller) (err error){

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
