package gocto

import (
	"context"
	"net/http"
)

func (g *Gocto) BindCard(ctx context.Context, request BindCardRequest) (*BindCardResponse, error) {
	var resp *BindCardResponse

	request.OctoSecret = g.octoSecret
	request.OctoShopID = g.octoShopID
	request.BindNotifyUrl = g.bindNotifyURL
	request.NotifyUrl = g.notifyURL

	if err := g.do(ctx, http.MethodPost, baseURL+"bind_card", request, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

