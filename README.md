# Go CodesWholesale API

[![Build Status](https://travis-ci.org/MarinX/codeswholesale.svg)](https://travis-ci.org/MarinX/codeswholesale)
[![GoDoc](https://godoc.org/github.com/MarinX/codeswholesale?status.svg)](https://godoc.org/github.com/MarinX/codeswholesale)

codeswholesale is a Go client library for the [CodesWholesale API](https://docs.codeswholesale.com/api-documentation/).


## Start using it

1. Download and install it:

    ```sh
    $ go get github.com/MarinX/codeswholesale
    ```

2. Import it in your code:

    ```go
    import "github.com/MarinX/codeswholesale"
    ```

3. (Optional) Run test 

    ```sh 
    $ go test -v
    ```

## API Examples

#### Get Account

```go
	cw := codeswholesale.New("ClientID", "ClientSecret", codeswholesale.MODE_SANDBOX)

	account, err := cw.GetAccount()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", account)
```

#### Get all products

```go
	cw := codeswholesale.New("ClientID", "ClientSecret", codeswholesale.MODE_SANDBOX)

	products, err := cw.GetProducts()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range products {
		fmt.Printf("%+v\n", val)
	}
```

#### Get Single Product

```go
	cw := codeswholesale.New("ClientID", "ClientSecret", codeswholesale.MODE_SANDBOX)

	product, err := cw.GetProductByID("product-cwid")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("%+v\n", product)
```

#### Order single code

```go
	cw := codeswholesale.New("ClientID", "ClientSecret", codeswholesale.MODE_SANDBOX)

	order, err := cw.SingleCodeOrder("product-cwid")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", order)
```

#### Multiple code order

```go
	cw := codeswholesale.New("ClientID", "ClientSecret", codeswholesale.MODE_SANDBOX)

	orders, err := cw.MultipleCodeOrder("product-cwid",5)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Order ID %s", orders.OrderID)

	for _, val := range orders.Items {
		fmt.Printf("%+v\n", val)
	}
```

## License

This library is under the MIT License
