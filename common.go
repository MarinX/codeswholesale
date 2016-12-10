// common, model for common structs for other models
package codeswholesale

// Order response codes
const (
	CODE_TEXT     cwTypeCode = "CODE_TEXT"
	CODE_PREORDER cwTypeCode = "CODE_PREORDER"
	CODE_IMAGE    cwTypeCode = "CODE_IMAGE"
)

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Image struct {
	URL    string `json:"image"`
	Format string `json:"format"`
}

type Price struct {
	Value float32 `json:"value"`
	From  float32 `json:"from"`
	To    float32 `json:"to"`
}

type CWError struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
