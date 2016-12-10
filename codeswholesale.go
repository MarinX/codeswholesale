// codeswholesale
package codeswholesale

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type cwMode int
type cwTypeCode string

// API Constants
const (
	API_STATUS_SUCCESS   = 0
	API_VERSION          = "v1"
	API_LIVE_ENDPOINT    = "https://api.codeswholesale.com"
	API_SANDBOX_ENDPOINT = "https://sandbox.codeswholesale.com"
)

// API modes
const (
	MODE_LIVE    cwMode = 1
	MODE_SANDBOX cwMode = 0
)

type CodesWholesale struct {
	mode   cwMode
	client *http.Client
}

// Creates new REST API client, to get api key:
// 1. Login into your account on app.codeswholesale.com
// 2. Locate Web API tab on the left bar
// 3. Generate new credentials (Keep in mind that a client secret will be visible only once)
func New(clientID string, clientSecret string, mode cwMode) *CodesWholesale {

	cw := &CodesWholesale{
		mode: mode,
	}

	cfg := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     cw.GetEndpoint() + "/oauth/token",
	}

	cw.client = cfg.Client(oauth2.NoContext)

	return cw
}

// Get endpoint based by mode: MODE_LIVE or MODE_SANDBOX
// It’s highly recommended to start working on your custom implementatnion on SANDBOX environment
func (c *CodesWholesale) GetEndpoint() string {
	if c.mode == MODE_LIVE {
		return API_LIVE_ENDPOINT
	}

	if c.mode == MODE_SANDBOX {
		return API_SANDBOX_ENDPOINT
	}
	// default to sandbox
	return API_SANDBOX_ENDPOINT
}

// Account Details provide you whole information about your account from name to actual balance with credit included.
func (c *CodesWholesale) GetAccount() (*Account, error) {

	acc := new(Account)

	err := c.callGet("accounts/current", acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

// Here you can easily ask for all products from CodesWholesale Price List.
// In response you will receive important informations about each product.
func (c *CodesWholesale) GetProducts() ([]Product, error) {

	var itm Item

	err := c.callGet("products", &itm)
	if err != nil {
		return itm.Items, err
	}

	return itm.Items, nil
}

// You can ask our API for only one product using ID.
// Tip: It’s good way to keep your stock and prices always updated.
func (c *CodesWholesale) GetProductByID(cwID string) (*Product, error) {

	prod := new(Product)

	err := c.callGet(fmt.Sprintf("products/%s", cwID), prod)
	if err != nil {
		return nil, err
	}

	return prod, nil
}

// Using your before recieved productId you can order this game in our API via this CURL.
// Remember it is only for one code.
func (c *CodesWholesale) SingleCodeOrder(cwID string) (*Order, error) {

	mm := make(map[string]string)
	mm["productId"] = cwID

	ord := new(Order)
	if err := c.callPost("orders", mm, ord); err != nil {
		return nil, err
	}

	return ord, nil
}

// Using your productID you can order games. Everything depends on quantity which you have to provide in your request.
// Create multiple codes order
func (c *CodesWholesale) MultipleCodeOrder(cwID string, quantity uint) (*Orders, error) {

	mm := make(map[string]string)
	mm["productId"] = cwID
	mm["quantity"] = fmt.Sprintf("%d", quantity)

	ord := new(Orders)
	if err := c.callPost("orders", mm, ord); err != nil {
		return ord, err
	}

	return ord, nil
}

// Constructs the version API
func (c *CodesWholesale) constructApi(resource string) string {
	return fmt.Sprintf("%s/%s/%s", c.GetEndpoint(), API_VERSION, resource)
}

// Calls GET HTTP method with error handling
func (c *CodesWholesale) callGet(resource string, v interface{}) error {

	resp, err := c.client.Get(c.constructApi(resource))
	if err != nil {
		return fmt.Errorf("[codeswholesale] %s ", err)
	}

	if err := responseToJSON(resp, v); err != nil {
		return fmt.Errorf("[codeswholesale] %s", err)
	}

	return nil
}

// Calls POST HTTP method with error handling
func (c *CodesWholesale) callPost(resource string, values map[string]string, v interface{}) error {

	urlValues := make(url.Values)
	for val, key := range values {
		urlValues[key] = []string{val}
	}

	resp, err := c.client.PostForm(c.constructApi(resource), urlValues)
	if err != nil {
		return fmt.Errorf("[codeswholesale] %s ", err)
	}

	if err := responseToJSON(resp, v); err != nil {
		return fmt.Errorf("[codeswholesale] %s", err)
	}

	return nil
}
