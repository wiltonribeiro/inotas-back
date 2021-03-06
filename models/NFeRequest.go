package models

type NFeRequest struct {
	IssuedOn       string `json:"issuedOn"`
	CheckCodeDigit float64    `json:"checkCodeDigit"`
	CityCode       float64    `json:"cityCode"`
	Serie          float64    `json:"serie"`
	CodeModel      float64    `json:"codeModel"`
	StateCode      float64    `json:"stateCode"`
	Number         float64    `json:"number"`
	CheckCode      float64    `json:"checkCode"`
	Issuer         struct {
		Cnae          float64    `json:"cnae"`
		CodeTaxRegime string `json:"codeTaxRegime"`
		Address       struct {
			City struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"city"`
			State                 string `json:"state"`
			District              string `json:"district"`
			AdditionalInformation string `json:"additionalInformation"`
			Street                string `json:"street"`
			Number                string `json:"number"`
			PostalCode            string `json:"postalCode"`
			Country               string `json:"country"`
			Phone                 string `json:"phone"`
		} `json:"address"`
		Im               string `json:"im"`
		StateTaxNumber   string `json:"stateTaxNumber"`
		TradeName        string `json:"tradeName"`
		Name             string `json:"name"`
		FederalTaxNumber string `json:"federalTaxNumber"`
	} `json:"issuer"`
	Buyer struct {
		StateTaxNumberIndicator float64 `json:"stateTaxNumberIndicator"`
		Address                 struct {
			City struct {
				Name string `json:"name"`
				Code string `json:"code"`
			} `json:"city"`
			State      string `json:"state"`
			District   string `json:"district"`
			Street     string `json:"street"`
			Number     string `json:"number"`
			PostalCode string `json:"postalCode"`
			Country    string `json:"country"`
			Phone      string `json:"phone"`
		} `json:"address"`
		Email            string `json:"email"`
		Name             string `json:"name"`
		FederalTaxNumber string `json:"federalTaxNumber"`
	} `json:"buyer"`
	Totals struct {
		Icms struct {
			FederalTaxesAmount       float64 `json:"federalTaxesAmount"`
			InvoiceAmount            float64 `json:"invoiceAmount"`
			ProductAmount            float64 `json:"productAmount"`
			IpiDevolAmount           float64     `json:"ipiDevolAmount"`
			FcpstAmount              float64     `json:"fcpstAmount"`
			FcpAmount                float64     `json:"fcpAmount"`
			IcmsufSenderAmount       float64     `json:"icmsufSenderAmount"`
			IcmsufDestinationAmount  float64     `json:"icmsufDestinationAmount"`
			FcpufDestinationAmount   float64     `json:"fcpufDestinationAmount"`
			OthersAmount             float64     `json:"othersAmount"`
			CofinsAmount             float64     `json:"cofinsAmount"`
			PisAmount                float64     `json:"pisAmount"`
			IpiAmount                float64     `json:"ipiAmount"`
			IiAmount                 float64     `json:"iiAmount"`
			DiscountAmount           float64     `json:"discountAmount"`
			InsuranceAmount          float64     `json:"insuranceAmount"`
			FreightAmount            float64     `json:"freightAmount"`
			StAmount                 float64     `json:"stAmount"`
			StCalculationBasisAmount float64     `json:"stCalculationBasisAmount"`
			IcmsExemptAmount         float64     `json:"icmsExemptAmount"`
			IcmsAmount               float64     `json:"icmsAmount"`
			BaseTax                  float64     `json:"baseTax"`
		} `json:"icms"`
	} `json:"totals"`
	Transport struct {
		FreightModality float64 `json:"freightModality"`
		TransportGroup  struct {
			State          string `json:"state"`
			FullAddress    string `json:"fullAddress"`
			StateTaxNumber string `json:"stateTaxNumber"`
			Name           string `json:"name"`
			CityName       string `json:"cityName"`
		} `json:"transportGroup"`
		Volume struct {
			GrossWeight      float64    `json:"grossWeight"`
			NetWeight        float64    `json:"netWeight"`
			VolumeQuantity   float64    `json:"volumeQuantity"`
			VolumeNumeration string `json:"volumeNumeration"`
			Species          string `json:"species"`
		} `json:"volume"`
	} `json:"transport"`
	AdditionalInformation struct {
		XMLAuthorized []int  `json:"xmlAuthorized"`
		Contract      string `json:"contract"`
		Order         string `json:"order"`
		Effort        string `json:"effort"`
		Taxpayer      string `json:"taxpayer"`
	} `json:"additionalInformation"`
	Protocol struct {
		ReceiptOn       string `json:"receiptOn"`
		StatusCode      float64    `json:"statusCode"`
		EnvironmentType string `json:"environmentType"`
		ValidatorDigit  string `json:"validatorDigit"`
		ProtocolNumber  string `json:"protocolNumber"`
		AccessKey       string `json:"accessKey"`
	} `json:"protocol"`
	Items []struct {
		TaxUnitAmount  float64 `json:"taxUnitAmount"`
		QuantityTax    float64     `json:"quantityTax"`
		TotalAmount    float64 `json:"totalAmount"`
		UnitAmount     float64 `json:"unitAmount"`
		Quantity       float64     `json:"quantity"`
		TotalIndicator bool    `json:"totalIndicator"`
		Cfop           float64     `json:"cfop"`
		Tax            struct {
			TotalTax float64 `json:"totalTax"`
			Icms     struct {
				Csosn  string `json:"csosn"`
				Origin string `json:"origin"`
			} `json:"icms"`
			Ipi struct {
				Amount             float64    `json:"amount"`
				Rate               float64    `json:"rate"`
				Base               string `json:"base"`
				Cst                string `json:"cst"`
				ClassificationCode string `json:"classificationCode"`
				StampCode          string `json:"stampCode"`
				ProducerCNPJ       string `json:"producerCNPJ"`
				Classification     string `json:"classification"`
			} `json:"ipi"`
			Pis struct {
				Cst string `json:"cst"`
			} `json:"pis"`
			Cofins struct {
				Cst string `json:"cst"`
			} `json:"cofins"`
		} `json:"tax"`
		ItemNumberOrderBuy float64    `json:"itemNumberOrderBuy"`
		NumberOrderBuy     string `json:"numberOrderBuy"`
		Cest               string `json:"cest"`
		UnitTax            string `json:"unitTax"`
		EanTaxableCode     string `json:"eanTaxableCode"`
		Unit               string `json:"unit"`
		Extipi             string `json:"extipi"`
		Ncm                string `json:"ncm"`
		Description        string `json:"description"`
		CodeGTIN           string `json:"codeGTIN"`
		Code               string `json:"code"`
	} `json:"items"`
	Payment []struct {
		PaymentDetail []struct {
			Amount float64 `json:"amount"`
			Method string  `json:"method"`
		} `json:"paymentDetail"`
	} `json:"payment"`
	CurrentStatus   string `json:"currentStatus"`
	PaymentType     string `json:"paymentType"`
	OperationType   string `json:"operationType"`
	Destination     string `json:"destination"`
	PrintType       string `json:"printType"`
	IssueType       string `json:"issueType"`
	EnvironmentType string `json:"environmentType"`
	PurposeType     string `json:"purposeType"`
	ConsumerType    string `json:"consumerType"`
	PresenceType    string `json:"presenceType"`
	ProcessType     string `json:"processType"`
	OperationOn     string `json:"operationOn"`
	XMLVersion      string `json:"xmlVersion"`
	InvoiceVersion  string `json:"invoiceVersion"`
	OperationNature string `json:"operationNature"`
}
