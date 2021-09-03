package SendtoWx

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strings"
)

func Send(text string)  {
	url := viper.GetString("qiyeweixin.url")
	method := "POST"

	payload := strings.NewReader(text)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}


