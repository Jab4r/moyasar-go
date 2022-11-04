# Moyasar-go 
This is a Golang wrapper library for the payment processing website "**[Moyasar](https://moyasar.com)**"
# How to start
First install the package
```
go get -u github.com/Jab4r/moyasar-go
``` 
Now import the package like so
```
Moyasar "github.com/Jab4r/moyasar-go"
```
Now this is **IMPORTANT !!**
```
var (
	SecretKey     = "sk_YOUR_SECRET_KEY_HERE"
	PublicKey     = "pk_YOUR_PUBLIC_KEY_HERE"
	Authorization = base64.StdEncoding.EncodeToString([]byte(SecretKey + ":" + PublicKey))
)
```

You are all done !

# functions & examples

## GetAllPayments()
```
// lets get all the payments
payments, err := Moyasar.GetAllPayments(Authorization)
if err != nil {
fmt.Println(err)
return
}

// we go through every payment and print the amount paid
for  _, payment := range payments {
fmt.Println(payment.Amount)
}
```
## FindPayemntById()
```
// lets get one payment by its id
payment, err := Moyasar.FindPayemntById("123ac-1234a-1234567b-12345bc123", Authorization)

if err != nil {

fmt.Println(err)

return

}

fmt.Println(payment.Currency) // view the currency

fmt.Println(payment.Amount) // view amount paid
```
## CapturePayemntById()
```
// lets capture a payment
captured, err := Moyasar.CapturePayemntById("1234-1234-1ac34-ca123",Authorization)

if err != nil {
// capturing payment faild
fmt.Println(err)
return
}

fmt.Println(captured.Status)
```

## VoidPayment()
```
// lets void a payment
// you cant void old payments that are over 1 or 2 hours
voided, err := Moyasar.VoidPayment("1234-1234-1ac34-ca123", Authorization)

if err != nil {

// capturing payment faild

fmt.Println(err)

return

}

fmt.Println(voided.Status) // view its status after void
```

## RefundPayment()
> to refund the whole payment set the second param to "Moyasar.RefundAll" like so RefundPayment("id", Moyasar.RefundAll ,  Authorization)
```
// lets refund part or whole payment
refundedPayment, err := Moyasar.RefundPayment("1234-1234-1ac34-ca123", 1000, Authorization)
if err != nil {

// capturing payment faild

fmt.Println(err)

return
}
fmt.Println(refundedPayment.Refunded_at) // view the time of refund
```

