package account

import (
	"encoding/json"

	"github.com/vultr/govultr/v3"
	"gopkg.in/yaml.v3"
)

// AccountPrinter ...
type AccountPrinter struct {
	Account *govultr.Account `json:"account"`
}

// JSON ...
func (s *AccountPrinter) JSON() []byte {
	prettyJSON, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		panic("move this into byte")
	}

	return prettyJSON
}

// YAML ...
func (s *AccountPrinter) YAML() []byte {
	yam, err := yaml.Marshal(s)
	if err != nil {
		panic("move this into byte")
	}
	return yam
}

// Columns ...
func (a *AccountPrinter) Columns() map[int][]interface{} {
	return map[int][]interface{}{0: {"BALANCE", "PENDING CHARGES", "LAST PAYMENT DATE", "LAST PAYMENT AMOUNT", "NAME", "EMAIL", "ACLS"}}
}

// Data ...
func (a *AccountPrinter) Data() map[int][]interface{} {
	return map[int][]interface{}{0: {a.Account.Balance, a.Account.PendingCharges, a.Account.LastPaymentDate, a.Account.LastPaymentAmount, a.Account.Name, a.Account.Email, a.Account.ACL}}
}

// Paging ...
func (a *AccountPrinter) Paging() map[int][]interface{} {
	return nil
}
