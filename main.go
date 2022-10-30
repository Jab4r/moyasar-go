package Moyasar

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Will return a payment by its id
func FindPayemntById(id, Authorization string) (PAYMENT, error) {
	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("GET", "https://api.moyasar.com/v1/payments/"+id, nil)
	if err != nil {
		return PAYMENT{}, err
	}

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
	err = json.Unmarshal(body, &payment)
	if err != nil {
		return PAYMENT{}, err
	}
	return payment, nil
}

// Will return an array of payment structs
func GetAllPayments(Authorization string) ([]PAYMENT, error) {
	client := &http.Client{Timeout: time.Second * 20}
	req, err := http.NewRequest("GET", "https://api.moyasar.com/v1/payments", nil)
	if err != nil {
		return []PAYMENT{}, err
	}

	req.Header.Add("Authorization", "Basic "+Authorization)

	res, err := client.Do(req)
	if err != nil {
		return []PAYMENT{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []PAYMENT{}, err
	}
	var payments PAYMENTS
	err = json.Unmarshal(body, &payments)
	if err != nil {
		return []PAYMENT{}, err
	}
	return payments.Payments, nil
}

// Will capture the payment by its id
func CapturePayemntById(id, Authorization string) (PAYMENT, error) {
	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("POST", "https://api.moyasar.com/v1/payments/"+id+"/capture", nil)
	if err != nil {
		return PAYMENT{}, err
	}

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
	err = json.Unmarshal(body, &payment)
	if err != nil {
		return PAYMENT{}, err
	}
	return payment, nil
} // on error // was paid or incorrect id

// Will void the payment
func VoidPayment(id, Authorization string) (PAYMENT, error) {
	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("POST", "https://api.moyasar.com/v1/payments/"+id+"/void", nil)
	if err != nil {
		return PAYMENT{}, err
	}

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
	err = json.Unmarshal(body, &payment)
	if err != nil {
		return PAYMENT{}, err
	}
	return payment, nil
} // on error // payment voided or void time expired

// Will refund the payment partial or full ## put -1 for full refund
func RefundPayment(id string, amount int, Authorization string) (PAYMENT, error) {
	client := &http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest("POST", "https://api.moyasar.com/v1/payments/"+id+"/refund", nil)
	if err != nil {
		return PAYMENT{}, err
	}

	req.Header.Add("Authorization", "Basic "+Authorization)
	if amount != -1 {
		req.Form.Add("amount", strconv.Itoa(amount))
	}
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
	err = json.Unmarshal(body, &payment)
	if err != nil {
		return PAYMENT{}, err
	}
	return payment, nil
} // on error // payment was refunded or not refundable

// TODO : use the "StatusCode" to verify if id is correct or not
// ELSE : make better error handling to capture errors
// create an error struct to get the error json and handle it
