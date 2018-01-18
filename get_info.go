package yobit

import (
	"context"
	"net/url"
)

type GetInfo struct {
	Success int `json:"success"`
	Return  struct {
		Funds           interface{} `json:"funds,omitempty"`
		FundsInclOrders interface{} `json:"funds_incl_orders,omitempty"`
		Rights          struct {
			Info     int `json:"info"`
			Trade    int `json:"trade"`
			Withdraw int `json:"withdraw"`
		} `json:"rights"`
		TransactionCount int `json:"transaction_count"`
		OpenOrders       int `json:"open_orders"`
		ServerTime       int `json:"server_time"`
	} `json:"return"`
}

func (c *Client) GetInfo(ctx context.Context, params url.Values) (GetInfo, error) {
	req, err := c.newAuthenticatedRequest(ctx, "POST", "tapi", params)
	if err != nil {
		return GetInfo{}, err
	}

	var ret = &GetInfo{}

	_, err = c.do(req, ret)
	if err != nil {
		return *ret, err
	}

	return *ret, nil
}