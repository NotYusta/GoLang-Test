package httptest

import (
	"fmt"
	"net/url"
)

func Init() {
	form := url.Values{}
	form.Add("test", "s")
	form.Add("test", "s2")

	fmt.Println(form)
}