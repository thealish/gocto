package gocto

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (g *Gocto) do(ctx context.Context, method, path string, request, response interface{}) error {
	requestBytes, err := json.Marshal(request)

	if err != nil {
		return err
	}

	rdr := bytes.NewReader(requestBytes)

	r, err := http.NewRequestWithContext(
		ctx,
		method,
		path,
		rdr,
	)

	r.Header.Set("Content-Type", contentTypeHeader)

	if err != nil {
		return err
	}

	res, err := g.hcl.Do(r)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	// response should always be pointer to struct

	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
