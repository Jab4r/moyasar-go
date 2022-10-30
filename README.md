# moyasar-go

## installation
```
go get -u github.com/Jab4r/moyasar-go
```

## required
start your code with this 

```
var (
	SecretKey     = "sk_YOUR_SECRET_KEY_HERE"
	PublicKey     = "pk_YOUR_PUBLIC_KEY_HERE"
	Authorization = base64.StdEncoding.EncodeToString([]byte(SecretKey + ":" + PublicKey))
)
```

## example

```
	payment, err := Moyasar.FindPayemntById("123ac-1234a-1234567b-12345bc123", Authorization)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(payment.Currency) // view the currency
	fmt.Println(payment.Amount) // view amount paid
```
