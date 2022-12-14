package px

type OrderList1 struct {
	Data []struct {
		ConsignmentInfo struct {
			FourPxInboundDate     int64  `json:"4px_inbound_date"`
			FourPxOutboundDate    int64  `json:"4px_outbound_date"`
			FourPxTrackingNo      string `json:"4px_tracking_no"`
			ConsignmentCreateDate int64  `json:"consignment_create_date"`
			ConsignmentStatus     string `json:"consignment_status"`
			DsConsignmentNo       string `json:"ds_consignment_no"`
			HasCheckOda           string `json:"has_check_oda"`
			InsureStatus          string `json:"insure_status"`
			InsureType            string `json:"insure_type"`
			IsHoldSign            string `json:"is_hold_sign"`
			LogisticsChannelNo    string `json:"logistics_channel_no"`
			LogisticsProductCode  string `json:"logistics_product_code"`
			LogisticsProductName  string `json:"logistics_product_name"`
			OdaResultSign         string `json:"oda_result_sign"`
			RefNo                 string `json:"ref_no"`
		} `json:"consignment_info"`
		ParcelConfirmInfo struct {
			ConfirmParcelQty      string `json:"confirm_parcel_qty"`
			ConfirmParcelWeight   int    `json:"confirm_parcel_weight"`
			ParcelListConfirmInfo []struct {
				ConfirmBatteryType      string `json:"confirm_battery_type"`
				ConfirmChargeWeight     int    `json:"confirm_charge_weight"`
				ConfirmHigh             string `json:"confirm_high"`
				ConfirmIncludeBattery   string `json:"confirm_include_battery"`
				ConfirmLength           string `json:"confirm_length"`
				ConfirmVolumeWeight     int    `json:"confirm_volume_weight"`
				ConfirmWeight           string `json:"confirm_weight"`
				ConfirmWidth            string `json:"confirm_width"`
				CurrencyCode            string `json:"currency_code"`
				ParcelTotalValueConfirm string `json:"parcel_total_value_confirm"`
			} `json:"parcel_list_confirm_info"`
		} `json:"parcel_confirm_info"`
	} `json:"data"`
	Msg    string `json:"msg"`
	Result string `json:"result"`
}

type OrderList struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Size     int         `json:"size"`
	OrderBy  interface{} `json:"orderBy"`
	StartRow int         `json:"startRow"`
	EndRow   int         `json:"endRow"`
	Total    int         `json:"total"`
	Pages    int         `json:"pages"`
	List     []struct {
		ID              int64       `json:"id"`
		CustomerOrderNo string      `json:"customerOrderNo"`
		ServerOrderNo   string      `json:"serverOrderNo"`
		CoFpxTrackNo    string      `json:"coFpxTrackNo"`
		ProductName     string      `json:"productName"`
		ProductEname    string      `json:"productEname"`
		CountryName     string      `json:"countryName"`
		CountryEname    string      `json:"countryEname"`
		ConsigneeName   string      `json:"consigneeName"`
		OsCode          string      `json:"osCode"`
		DsCode          interface{} `json:"dsCode"`
		GroupSign       string      `json:"groupSign"`
		StatusName      interface{} `json:"statusName"`
		ExtraStatus     struct {
			IsSendBack     string `json:"isSendBack"`
			IsNeedSendBack string `json:"isNeedSendBack"`
			HasInsurance   string `json:"hasInsurance"`
			HasFba         string `json:"hasFba"`
			IsPrint        string `json:"isPrint"`
			IsOda          string `json:"isOda"`
			IsCod          string `json:"isCod"`
		} `json:"extraStatus"`
		ProblemSign       interface{} `json:"problemSign"`
		CreatedTime       int64       `json:"createdTime"`
		ReceivedTime      int64       `json:"receivedTime"`
		Status            string      `json:"status"`
		TrialAmount       float64     `json:"trialAmount"`
		DocumentCodeSign  string      `json:"documentCodeSign"`
		CustomerWeight    float64     `json:"customerWeight"`
		IsHold            string      `json:"isHold"`
		NumberExcpContent interface{} `json:"numberExcpContent"`
		PkCode            string      `json:"pkCode"`
		ChnCode           string      `json:"chnCode"`
		CtCode            string      `json:"ctCode"`
		CmID              int         `json:"cmId"`
		CheckInDate       interface{} `json:"checkInDate"`
		CheckOutDate      interface{} `json:"checkOutDate"`
		FinishDate        interface{} `json:"finishDate"`
		ResendCode        interface{} `json:"resendCode"`
		SortCode          string      `json:"sortCode"`
		PrintDate         interface{} `json:"printDate"`
	} `json:"list"`
	FirstPage        int   `json:"firstPage"`
	PrePage          int   `json:"prePage"`
	NextPage         int   `json:"nextPage"`
	LastPage         int   `json:"lastPage"`
	IsFirstPage      bool  `json:"isFirstPage"`
	IsLastPage       bool  `json:"isLastPage"`
	HasPreviousPage  bool  `json:"hasPreviousPage"`
	HasNextPage      bool  `json:"hasNextPage"`
	NavigatePages    int   `json:"navigatePages"`
	NavigatepageNums []int `json:"navigatepageNums"`
}

type OrderDetail struct {
	OrderID int64 `json:"orderId"`
	Order   struct {
		ID                       int64       `json:"id"`
		Ids                      interface{} `json:"ids"`
		DataSourceOrigin         string      `json:"dataSourceOrigin"`
		CustomerID               int         `json:"customerId"`
		NewCustomerID            interface{} `json:"newCustomerId"`
		ShipmentChannelID        interface{} `json:"shipmentChannelId"`
		CustomerOrderNo          string      `json:"customerOrderNo"`
		ServerOrderNo            string      `json:"serverOrderNo"`
		ChangedOrderNo           interface{} `json:"changedOrderNo"`
		CoFpxTrackNo             string      `json:"coFpxTrackNo"`
		ProductCode              string      `json:"productCode"`
		TransformProductCode     interface{} `json:"transformProductCode"`
		ShipmentCategory         string      `json:"shipmentCategory"`
		PaymentMode              string      `json:"paymentMode"`
		OriginCountry            string      `json:"originCountry"`
		DestinationCountry       string      `json:"destinationCountry"`
		Pieces                   int         `json:"pieces"`
		InsuranceType            string      `json:"insuranceType"`
		InsuranceStatus          string      `json:"insuranceStatus"`
		IsCheckedOda             string      `json:"isCheckedOda"`
		IsOda                    string      `json:"isOda"`
		IsHold                   string      `json:"isHold"`
		InsuranceValue           float32     `json:"insuranceValue"`
		BuyerID                  interface{} `json:"buyerId"`
		ShipperWangWangID        interface{} `json:"shipperWangWangId"`
		CreatorID                int         `json:"creatorId"`
		CreatedTime              int64       `json:"createdTime"`
		ModifiedTime             int64       `json:"modifiedTime"`
		PrintedTime              interface{} `json:"printedTime"`
		ConfirmTime              interface{} `json:"confirmTime"`
		PostTime                 int64       `json:"postTime"`
		CheckinTime              interface{} `json:"checkinTime"`
		CheckoutTime             interface{} `json:"checkoutTime"`
		Status                   string      `json:"status"`
		CustomerOrderSerialNo    interface{} `json:"customerOrderSerialNo"`
		RefuseInsuranceReason    interface{} `json:"refuseInsuranceReason"`
		OrderNoChangedSign       string      `json:"orderNoChangedSign"`
		DataJoinUpType           string      `json:"dataJoinUpType"`
		IsNeedSendBack           string      `json:"isNeedSendBack"`
		IsSendBack               string      `json:"isSendBack"`
		FbaCode                  interface{} `json:"fbaCode"`
		PostalShipmentCategory   string      `json:"postalShipmentCategory"`
		EbusinessPlatformID      int         `json:"ebusinessPlatformId"`
		CustomerCode             string      `json:"customerCode"`
		HasBattery               string      `json:"hasBattery"`
		HoldProduct              string      `json:"holdProduct"`
		HoldEproduct             string      `json:"holdEproduct"`
		HoldCountry              string      `json:"holdCountry"`
		HoldEcountry             string      `json:"holdEcountry"`
		IsSyncOrderNo            interface{} `json:"isSyncOrderNo"`
		HasSignature             string      `json:"hasSignature"`
		IsDeclare                string      `json:"isDeclare"`
		RouteCount               int         `json:"routeCount"`
		CommentType              interface{} `json:"commentType"`
		CommentDesc              interface{} `json:"commentDesc"`
		ArrivalMode              int         `json:"arrivalMode"`
		BizCode                  interface{} `json:"bizCode"`
		PortCode                 interface{} `json:"portCode"`
		ZoneCode                 interface{} `json:"zoneCode"`
		ParcelSource             string      `json:"parcelSource"`
		ParcelMark               interface{} `json:"parcelMark"`
		GrossWeight              interface{} `json:"grossWeight"`
		ChargeWeight             interface{} `json:"chargeWeight"`
		VolumeWeight             interface{} `json:"volumeWeight"`
		SystemOrderNo            interface{} `json:"systemOrderNo"`
		ParcelCount              int         `json:"parcelCount"`
		LstChildNumber           interface{} `json:"lstChildNumber"`
		LabelURL                 interface{} `json:"labelUrl"`
		InvoiceURL               interface{} `json:"invoiceUrl"`
		ChnCode                  string      `json:"chnCode"`
		IsVip                    interface{} `json:"isVip"`
		CommentContent           interface{} `json:"commentContent"`
		CustomerWeightg          interface{} `json:"customerWeightg"`
		SubAccountID             interface{} `json:"subAccountId"`
		ServiceName              interface{} `json:"serviceName"`
		ServiceCode              interface{} `json:"serviceCode"`
		SalName                  interface{} `json:"salName"`
		SalCode                  interface{} `json:"salCode"`
		ProductName              interface{} `json:"productName"`
		CrmCmCode                interface{} `json:"crmCmCode"`
		CodServerNo              interface{} `json:"codServerNo"`
		CodSign                  string      `json:"codSign"`
		CurrencyDeclareInsurance interface{} `json:"currencyDeclareInsurance"`
		IsArrears                interface{} `json:"isArrears"`
		CheckIndate              interface{} `json:"checkIndate"`
		CountryName              interface{} `json:"countryName"`
		RefunderNo               interface{} `json:"refunderNo"`
		IsCodAgainDelivery       interface{} `json:"isCodAgainDelivery"`
		FinishTime               interface{} `json:"finishTime"`
		RejectTime               interface{} `json:"rejectTime"`
		GeneratedOrderNo         bool        `json:"generatedOrderNo"`
	} `json:"order"`
	OrderAttach struct {
		OrderID                  interface{} `json:"orderId"`
		CustomerWeight           float64     `json:"customerWeight"`
		GrossWeight              interface{} `json:"grossWeight"`
		VolumeWeight             interface{} `json:"volumeWeight"`
		ChargeWeight             interface{} `json:"chargeWeight"`
		TransactionID            interface{} `json:"transactionId"`
		Warning                  interface{} `json:"warning"`
		Error                    interface{} `json:"error"`
		OrderNote                interface{} `json:"orderNote"`
		DeclinationReason        interface{} `json:"declinationReason"`
		ChannelID                interface{} `json:"channelId"`
		InvoiceReferenceSign     interface{} `json:"invoiceReferenceSign"`
		SenderAccount            interface{} `json:"senderAccount"`
		PayerAccount             interface{} `json:"payerAccount"`
		BttCode                  interface{} `json:"bttCode"`
		InvoicePrintDate         interface{} `json:"invoicePrintDate"`
		DutyPayMode              string      `json:"dutyPayMode"`
		PrintFormat              interface{} `json:"printFormat"`
		TariffNo                 interface{} `json:"tariffNo"`
		MasterHawbCode           interface{} `json:"masterHawbCode"`
		LabelCode                interface{} `json:"labelCode"`
		DeclareValue             interface{} `json:"declareValue"`
		CustomerWeightLimit      interface{} `json:"customerWeightLimit"`
		WarehouseCode            interface{} `json:"warehouseCode"`
		ChannelChangeFlag        interface{} `json:"channelChangeFlag"`
		Width                    interface{} `json:"width"`
		Length                   interface{} `json:"length"`
		Height                   interface{} `json:"height"`
		BusinessType             string      `json:"businessType"`
		SalesPlatform            interface{} `json:"salesPlatform"`
		SellerID                 interface{} `json:"sellerId"`
		IsReturnOnOversea        string      `json:"isReturnOnOversea"`
		IsReturnOnDomestic       string      `json:"isReturnOnDomestic"`
		ValueAddedServices       interface{} `json:"valueAddedServices"`
		IsCommercialInvoice      string      `json:"isCommercialInvoice"`
		CodPrice                 float32     `json:"codPrice"`
		CodCurrency              string      `json:"codCurrency"`
		TotalCBM                 interface{} `json:"totalCBM"`
		ShipmentID               interface{} `json:"shipmentId"`
		ReferenceID              interface{} `json:"referenceId"`
		Mark                     interface{} `json:"mark"`
		InvoiceNo                interface{} `json:"invoiceNo"`
		HblNo                    interface{} `json:"hblNo"`
		DateOfExportation        interface{} `json:"dateOfExportation"`
		MixNo                    interface{} `json:"mixNo"`
		TotalCartons             interface{} `json:"totalCartons"`
		ParcelVolumeList         interface{} `json:"parcelVolumeList"`
		ChnCode                  string      `json:"chnCode"`
		ChnLock                  string      `json:"chnLock"`
		FreightCharges           interface{} `json:"freightCharges"`
		CoParcelinspect          interface{} `json:"coParcelinspect"`
		SortCode                 string      `json:"sortCode"`
		SubAccountID             interface{} `json:"subAccountId"`
		BatteryType              string      `json:"batteryType"`
		DistributionType         int         `json:"distributionType"`
		DistributorCode          interface{} `json:"distributorCode"`
		BtName                   interface{} `json:"btName"`
		CodStringPrice           interface{} `json:"codStringPrice"`
		IossNo                   string      `json:"iossNo"`
		LucidNo                  string      `json:"lucidNo"`
		DeclareInsurance         interface{} `json:"declareInsurance"`
		CurrencyFreight          interface{} `json:"currencyFreight"`
		CurrencyDeclareInsurance interface{} `json:"currencyDeclareInsurance"`
		CurrencyCustoms          string      `json:"currencyCustoms"`
		ExchangeRateCustoms      string      `json:"exchangeRateCustoms"`
		FreightChargesExchange   int         `json:"freightChargesExchange"`
		DeclareInsuranceExchange int         `json:"declareInsuranceExchange"`
		CurrentCpRescode         interface{} `json:"currentCpRescode"`
		LogisticsOrderCreatetime interface{} `json:"logisticsOrderCreatetime"`
		Ext                      interface{} `json:"ext"`
		DhlecStatus              interface{} `json:"dhlecStatus"`
		DeliveryArea             string      `json:"deliveryArea"`
		EcCode                   string      `json:"ecCode"`
		DeclareService           string      `json:"declareService"`
		PickUpStationName        interface{} `json:"pickUpStationName"`
		PickUpStationCode        interface{} `json:"pickUpStationCode"`
	} `json:"orderAttach"`
	ShipperConsignee struct {
		OrderID                interface{} `json:"orderId"`
		ShipperName            string      `json:"shipperName"`
		ShipperAddress1        string      `json:"shipperAddress1"`
		ShipperAddress2        string      `json:"shipperAddress2"`
		ShipperAddress3        string      `json:"shipperAddress3"`
		ShipperPostcode        string      `json:"shipperPostcode"`
		ShipperTelephone       string      `json:"shipperTelephone"`
		ShipperFax             interface{} `json:"shipperFax"`
		ConsigneeName          string      `json:"consigneeName"`
		ConsigneeFirstName     string      `json:"consigneeFirstName"`
		ConsigneeLastName      string      `json:"consigneeLastName"`
		ConsigneeAddress1      string      `json:"consigneeAddress1"`
		ConsigneeAddress2      string      `json:"consigneeAddress2"`
		ConsigneeAddress3      interface{} `json:"consigneeAddress3"`
		ConsigneePostcode      string      `json:"consigneePostcode"`
		ConsigneeTelephone     string      `json:"consigneeTelephone"`
		ConsigneeFax           interface{} `json:"consigneeFax"`
		ConsigneeMail          string      `json:"consigneeMail"`
		ShipperCompany         string      `json:"shipperCompany"`
		ConsigneeCompany       string      `json:"consigneeCompany"`
		ConsigneeCity          string      `json:"consigneeCity"`
		ShipperAreaCode        interface{} `json:"shipperAreaCode"`
		ConsigneeAreaCode      interface{} `json:"consigneeAreaCode"`
		ShipperProvince        string      `json:"shipperProvince"`
		ShipperCity            string      `json:"shipperCity"`
		ShipperMobile          string      `json:"shipperMobile"`
		ShipperEmail           string      `json:"shipperEmail"`
		ShipperCountry         string      `json:"shipperCountry"`
		ShipperStreet          string      `json:"shipperStreet"`
		ShipperIDCode          interface{} `json:"shipperIdCode"`
		ConsigneeMobile        string      `json:"consigneeMobile"`
		ConsigneeCountry       string      `json:"consigneeCountry"`
		ConsigneeStreet        string      `json:"consigneeStreet"`
		ConsigneeDoorNo        string      `json:"consigneeDoorNo"`
		ConsigneeIDCode        interface{} `json:"consigneeIdCode"`
		ShipperWangWangID      interface{} `json:"shipperWangWangId"`
		ConsigneeWangWangID    interface{} `json:"consigneeWangWangId"`
		ConsigneeProvince      string      `json:"consigneeProvince"`
		ShipperCredentialsType interface{} `json:"shipperCredentialsType"`
		CredentialsType        interface{} `json:"credentialsType"`
		CredentialsPeriod      interface{} `json:"credentialsPeriod"`
		ConsigneeCityID        interface{} `json:"consigneeCityId"`
		ShipperVatNo           interface{} `json:"shipperVatNo"`
		ConsigneeVatNo         interface{} `json:"consigneeVatNo"`
		ShipperLastName        interface{} `json:"shipperLastName"`
		ShipperidFronturl      interface{} `json:"shipperidFronturl"`
		ShipperidBackurl       interface{} `json:"shipperidBackurl"`
		ConsigneeidFronturl    interface{} `json:"consigneeidFronturl"`
		ConsigneeidBackturl    interface{} `json:"consigneeidBackturl"`
		ShipperDistrict        interface{} `json:"shipperDistrict"`
		ConsigneeDistrict      string      `json:"consigneeDistrict"`
		DataSourceOriginCode   interface{} `json:"dataSourceOriginCode"`
	} `json:"shipperConsignee"`
	DeclareList []struct {
		OrderID                interface{} `json:"orderId"`
		OpID                   interface{} `json:"opId"`
		InvoiceID              interface{} `json:"invoiceId"`
		Name                   string      `json:"name"`
		Ename                  string      `json:"ename"`
		Cname                  string      `json:"cname"`
		UnitCode               string      `json:"unitCode"`
		Pcs                    int         `json:"pcs"`
		UnitPrice              float64     `json:"unitPrice"`
		Note                   interface{} `json:"note"`
		NamePrefix             interface{} `json:"namePrefix"`
		Hscode                 string      `json:"hscode"`
		Weight                 float64     `json:"weight"`
		URL                    interface{} `json:"url"`
		TotalPrice             interface{} `json:"totalPrice"`
		PkCode                 interface{} `json:"pkCode"`
		DiHscodeExport         interface{} `json:"diHscodeExport"`
		DiHscodeImport         interface{} `json:"diHscodeImport"`
		DeclareUnitPriceExport string      `json:"declareUnitPriceExport"`
		CurrencyExport         interface{} `json:"currencyExport"`
		DeclareUnitPriceImport string      `json:"declareUnitPriceImport"`
		CurrencyImport         interface{} `json:"currencyImport"`
		BrandExport            interface{} `json:"brandExport"`
		BrandImport            interface{} `json:"brandImport"`
		CountryExport          interface{} `json:"countryExport"`
		CountryImport          interface{} `json:"countryImport"`
		CategoryID             interface{} `json:"categoryId"`
		CategoryName           interface{} `json:"categoryName"`
		TaxRate                interface{} `json:"taxRate"`
		Cartons                interface{} `json:"cartons"`
		DeclareWeight          interface{} `json:"declareWeight"`
		Cbm                    interface{} `json:"cbm"`
		ProductModel           interface{} `json:"productModel"`
		ProductSize            interface{} `json:"productSize"`
		Material               string      `json:"material"`
		MeterialRatio          interface{} `json:"meterialRatio"`
		Uses                   string      `json:"uses"`
		Brand                  interface{} `json:"brand"`
		Currency               string      `json:"currency"`
		DeclareSpecification   string      `json:"declareSpecification"`
		SkuCode                string      `json:"skuCode"`
	} `json:"declareList"`
	OrderParcelList  interface{} `json:"orderParcelList"`
	Refunder         interface{} `json:"refunder"`
	RefunderOversea  interface{} `json:"refunderOversea"`
	SelfSendArrival  interface{} `json:"selfSendArrival"`
	DoorPickUp       interface{} `json:"doorPickUp"`
	ExpressArrival   interface{} `json:"expressArrival"`
	OrderCommentList []struct {
		ID               int64       `json:"id"`
		OrderID          int64       `json:"orderId"`
		Type             string      `json:"type"`
		Content          string      `json:"content"`
		CreatorID        interface{} `json:"creatorId"`
		CreatorCode      interface{} `json:"creatorCode"`
		CreatorShortName interface{} `json:"creatorShortName"`
		CreateTime       int64       `json:"createTime"`
		CreateTimeString string      `json:"createTimeString"`
	} `json:"orderCommentList"`
	InsuranceAttachItem   interface{} `json:"insuranceAttachItem"`
	RowNum                interface{} `json:"rowNum"`
	MqEventType           interface{} `json:"mqEventType"`
	GsdpAttachDTO         interface{} `json:"gsdpAttachDTO"`
	DeliverToRecipientDTO interface{} `json:"deliverToRecipientDTO"`
	ParcelVolumeList      []struct {
		Length  interface{} `json:"length"`
		Width   interface{} `json:"width"`
		Heigh   interface{} `json:"heigh"`
		Weight  string      `json:"weight"`
		OrderID interface{} `json:"orderId"`
	} `json:"parcelVolumeList"`
	OrderLabelURLDTO struct {
		BsID              int64       `json:"bsId"`
		LabelURL          interface{} `json:"labelUrl"`
		CustomLableURL    interface{} `json:"customLableUrl"`
		PackageLabelURL   interface{} `json:"packageLabelUrl"`
		InvoiceLabelURL   interface{} `json:"invoiceLabelUrl"`
		BatteryLetterURL  interface{} `json:"batteryLetterUrl"`
		RequestNo         interface{} `json:"requestNo"`
		InvoiceBatteryURL interface{} `json:"invoiceBatteryUrl"`
	} `json:"orderLabelUrlDTO"`
	OrderDeclareInvoicePictureDTO interface{}   `json:"orderDeclareInvoicePictureDTO"`
	AllInvoicetotalPrice          interface{}   `json:"allInvoicetotalPrice"`
	ExtraServices                 []interface{} `json:"extraServices"`
	CargoVolumes                  []interface{} `json:"cargoVolumes"`
	OrderAccessoryList            []interface{} `json:"orderAccessoryList"`
	SourceFlag                    interface{}   `json:"sourceFlag"`
	AllInvoiceCurrency            interface{}   `json:"allInvoiceCurrency"`
	Deliver                       interface{}   `json:"deliver"`
	JudgeSvip                     interface{}   `json:"judgeSvip"`
	ServiceRefunder               interface{}   `json:"serviceRefunder"`
	Locale                        interface{}   `json:"locale"`
}
