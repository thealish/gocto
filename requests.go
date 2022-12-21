package gocto

type BindCardRequest struct {
	OctoShopID        int    `json:"octo_shop_id"`
	OctoSecret        string `json:"octo_secret"`
	ShopTransactionId string `json:"shop_transaction_id"`
	Pan               string `json:"pan"`
	Exp               string `json:"exp"`
	Phone             string `json:"phone"`
	CardHolderName    string `json:"cardHolderName"`
	Cvc2              string `json:"cvc2"`
	Method            string `json:"method"`
	ReturnUrl         string `json:"return_url"`
	NotifyUrl         string `json:"notify_url"`
	BindNotifyUrl     string `json:"bind_notify_url"`
}

type BindCardResponse struct {
	Error             int    `json:"error"`
	Status            string `json:"status"`
	ShopTransactionId string `json:"shop_transaction_id"`
	OctoPaymentUUID   string `json:"octo_payment_UUID"`
	OctoPayUrl        string `json:"octo_pay_url"`
	CardToken         string `json:"card_token"`
}

type CardBindingNotification struct {
	ShopTransactionId string `json:"shop_transaction_id"`
	Pan               string `json:"pan"`
	Exp               string `json:"exp"`
	CardHolderName    string `json:"cardHolderName"`
	CardToken         string `json:"card_token"`
	Status            string `json:"status"`
}

type BlockCardRequest struct {
	OctoShopId int    `json:"octo_shop_id"`
	OctoSecret string `json:"octo_secret"`
	CardToken  string `json:"card_token"`
}

type BlockCardResponse struct {
	Error        int         `json:"error"`
	ErrorMessage interface{} `json:"errorMessage"`
	Status       string      `json:"status"`
}
