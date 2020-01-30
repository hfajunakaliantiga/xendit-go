package balance

import (
	"net/url"

	"github.com/xendit/xendit-go"
)

// GetParams contains parameters for Get
type GetParams struct {
	AccountType xendit.BalanceAccountTypeEnum `json:"account_type"`
}

// QueryString creates query string from GetAllParams, ignores nil values
func (p *GetParams) QueryString() string {
	urlValues := &url.Values{}

	if p.AccountType != "" {
		urlValues.Add("account_type", string(p.AccountType))
	}

	return urlValues.Encode()
}