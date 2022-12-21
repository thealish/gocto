package gocto

import (
	"context"
	"net/http"
)

func (g *Gocto) BlockCard(ctx context.Context, request BlockCardRequest) (*BlockCardResponse, error) {
	var resp *BlockCardResponse

	request.OctoSecret = g.octoSecret
	request.OctoShopId = g.octoShopID

	if err := g.do(ctx, http.MethodPost, baseURL+"block_card_token", request, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
