// order
package codeswholesale

type Orders struct {
	OrderID string  `json:"orderId"`
	Items   []Order `json:"items"`
}

type Order struct {
	CodeType cwTypeCode `json:"codeType"`
	Code     string     `json:"code"`
	Links    []Link     `json:"links,omitempty"`
	FileName string     `json:"fileName,omitempty"`
}

func (o *Order) IsCodeText() bool {
	if o.CodeType == CODE_TEXT {
		return true
	}
	return false
}

func (o *Order) IsCodePreorder() bool {
	if o.CodeType == CODE_PREORDER {
		return true
	}
	return false
}

func (o *Order) IsCodeImage() bool {
	if o.CodeType == CODE_IMAGE {
		return true
	}
	return false
}
