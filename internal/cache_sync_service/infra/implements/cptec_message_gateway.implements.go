package implements

import "context"

type CptecMessageGateway struct{}

func NewCptecMessageGateway() *CptecMessageGateway {
	return &CptecMessageGateway{}
}

func (CptecMessageGateway) MessageByLocation(ctx context.Context, city string, state string) (string, error) {
	return city + ":" + state, nil
}
