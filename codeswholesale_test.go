// codeswholesale_test
package codeswholesale

import (
	"testing"
)

const (
	CLIENT_ID        = "ff72ce315d1259e822f47d87d02d261e"
	CLIENT_SECRET    = "$2a$10$E2jVWDADFA5gh6zlRVcrlOOX01Q/HJoT6hXuDMJxek.YEo.lkO2T6"
	TEST_PRODUCT     = "6313677f-5219-47e4-a067-7401f55c5a3a"
	TEST_BUY_PRODUCT = "ffe2274d-5469-4b0f-b57b-f8d21b09c24c"
)

func getAPI() *CodesWholesale {
	return New(CLIENT_ID, CLIENT_SECRET, MODE_SANDBOX)
}

func TestAccount(t *testing.T) {
	t.Log("Fetching account")

	acc, err := getAPI().GetAccount()
	if err != nil {
		t.Log(err)
		return
	}

	t.Logf("Account: %+v\n", *acc)
}

func TestProducts(t *testing.T) {

	t.Log("Fetching products")
	prods, err := getAPI().GetProducts()
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("Products: %+v\n", prods)

	t.Log("Testing unvalid product by ID")
	prod, err := getAPI().GetProductByID("blabla-this-must-fail")
	if prod != nil {
		t.Log(err)
		return
	}
	t.Logf("This fail OK: %s\n", err)

	t.Log("Testing valid product by ID")
	prod, err = getAPI().GetProductByID(TEST_PRODUCT)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("Product: %+v\n", prod)
}

func TestSingleCodeOrder(t *testing.T) {
	t.Log("Single code order")

	order, err := getAPI().SingleCodeOrder(TEST_BUY_PRODUCT)
	if err != nil {
		t.Log(err)
		return
	}

	t.Logf("Order: %+v\n", order)
}

func TestMultipleCodeOrder(t *testing.T) {
	t.Log("Multiple code order")

	order, err := getAPI().MultipleCodeOrder(TEST_BUY_PRODUCT, 5)
	if err != nil {
		t.Log(err)
		return
	}

	t.Logf("Multiple orders: %+v\n", order)
}
