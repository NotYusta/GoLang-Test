package midtrans

import (
	"fmt"
	"net/url"
)

func InitMidtrans() {
	form := url.Values{}
	form.Add("test", "s")

	fmt.Println(form)
}