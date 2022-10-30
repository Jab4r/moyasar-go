package Moyasar

import "time"

type PAYMENT struct {
	Id              string      `json:"id"`
	Status          string      `json:"status"`
	Amount          int         `json:"amount"`
	Fee             int         `json:"fee"`
	Currency        string      `json:"currency"`
	Refunded        int         `json:"refunded"`
	Refunded_at     *time.Time  `json:"refunded_at"` // nullable
	Captured        int         `json:"captured"`
	Captured_at     *time.Time  `json:"captured_at"` // nullable
	Voided_at       *time.Time  `json:"voided_at"`   // nullable
	Description     string      `json:"description"`
	Amount_format   string      `json:"amount_format"`
	Fee_format      string      `json:"fee_format"`
	Refunded_format string      `json:"refunded_format"`
	Captured_format string      `json:"captured_format"`
	Invoice_id      *string     `json:"invoice_id"` // nullable
	Ip              string      `json:"ip"`
	Callback_url    string      `json:"callback_url"`
	Created_at      time.Time   `json:"created_at"`
	Updated_at      time.Time   `json:"updated_at"`
	Metadata        interface{} `json:"metadata"` // work it out
	Source          interface{} `json:"source"`   // this too
}

type PAYMENTS struct {
	Payments []PAYMENT `json:"payments"`
}
