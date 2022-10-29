package Moyasar

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func FindPayemntById(id, Authorization string) (PAYMENT, error) {
	client := &http.Client{Timeout: time.Second * 10}
	req, _ := http.NewRequest("GET", "https://api.moyasar.com/v1/payments/"+id, nil)

	req.Header.Add("Authorization", "Basic "+Authorization)

	res, err := client.Do(req)
	if err != nil {
		return PAYMENT{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PAYMENT{}, err
	}
	var payment PAYMENT
	json.Unmarshal(body, &payment)
	return payment, nil
}
