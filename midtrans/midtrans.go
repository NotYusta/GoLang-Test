package midtrans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"me.yusta/config"
)

type TransactionDetails struct {
	OrderID string `json:"order_id"`
	GrossAmount int `json:"gross_amount"`
}

type TransactionRequest struct {
	PaymentType string `json:"payment_type"`
	TransactionDetails TransactionDetails  `json:"transaction_details"`
}
func InitMidtrans() {

	c := &http.Client{}


	value := &TransactionRequest{
		PaymentType: "qris",
		TransactionDetails: TransactionDetails{
			OrderID: "nothin-01",
			GrossAmount: 1000,
		},
	}

	val, _ := json.Marshal(value)

	fmt.Println(string(val))
	req, _ := http.NewRequest(http.MethodPost, "https://api.midtrans.com/v2/charge", bytes.NewBuffer(val))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(config.Yaml.Midtrans.ServerKey, "")


	res, err := c.Do(req)
	if err != nil {
		fmt.Println("an error has occured ", err.Error())
		return
	}

	var thing interface{}
	json.NewDecoder(res.Body).Decode(&thing)

	fmt.Println(thing)

}