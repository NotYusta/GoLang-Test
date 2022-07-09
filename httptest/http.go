package httptest

import (
	"net/http"

	"me.yusta/config"
)

func Init() {
	req, _ := http.NewRequest("POST", "nothing", nil)

	req.SetBasicAuth(config.Yaml.Midtrans.ServerKey, "")

}