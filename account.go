// account, model
package codeswholesale

type Account struct {
	FullName       string  `json:"fullName"`
	Email          string  `json:"email"`
	CurrentBalance float32 `json:"currentBalance"`
	CurrentCredit  float32 `json:"currentCredit"`
	TotalToUse     float32 `json:"totalToUse"`
	Links          []Link  `json:"links"`
}
