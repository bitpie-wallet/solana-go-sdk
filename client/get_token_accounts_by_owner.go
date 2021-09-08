package client

import (
	"context"

	"github.com/portto/solana-go-sdk/common"
)

type TokenAmount struct {
	Amount         string  `json:"amount"`
	Decimals       int32   `json:"decimals"`
	UIAmount       float64 `json:"uiAmount"`
	UIAmountString string  `json:"uiAmountString"`
}
type Info struct {
	Delegate        string      `json:"delegate"`
	DelegatedAmount TokenAmount `json:"delegatedAmount"`
	IsInitialized   bool        `json:"isInitialized"`
	IsNative        bool        `json:"isNative"`
	Mint            string      `json:"mint"`
	Owner           string      `json:"owner"`
	TokenAmount     TokenAmount `json:"tokenAmount"`
}
type Parsed struct {
	AccountType string `json:"accountType"`
	Info        Info   `json:"info"`
}
type Data struct {
	Parsed  Parsed `json:"parsed"`
	Program string `json:"program"`
}
type Account struct {
	Data       Data   `json:"data"`
	Executable bool   `json:"executable"`
	Lamports   int64  `json:"lamports"`
	Owner      string `json:"owner"`
	RentEpoch  int64  `json:"rentEpoch"`
}
type Accounts struct {
	Account Account `json:"account"`
	Pubkey  string  `json:"pubkey,omitempty"`
}

func (s *Client) GetTokenAccountsByOwner(ctx context.Context, account string) ([]Accounts, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context    `json:"context"`
			Value   []Accounts `json:"value"`
		} `json:"result"`
	}{}
	params := []interface{}{account,
		map[string]interface{}{"programId": common.TokenProgramID.ToBase58()},
		map[string]interface{}{
			"encoding": "jsonParsed",
		}}

	err := s.request(ctx, "getTokenAccountsByOwner", params, &res)
	if err != nil {
		return []Accounts{}, err
	}
	return res.Result.Value, nil
}
func (s *Client) GetTokenAccountByMint(ctx context.Context, account string, mint string) ([]Accounts, error) {
	res := struct {
		GeneralResponse
		Result struct {
			Context Context    `json:"context"`
			Value   []Accounts `json:"value"`
		} `json:"result"`
	}{}
	params := []interface{}{account,
		map[string]interface{}{"mint": mint},
		map[string]interface{}{
			"encoding": "jsonParsed",
		}}

	err := s.request(ctx, "getTokenAccountsByOwner", params, &res)
	if err != nil {
		return []Accounts{}, err
	}
	return res.Result.Value, nil
}
