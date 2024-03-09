package gameservice

import (
	"context"

	"gitlab.com/morbackend/mor_services/mpb"
)

func (svc *GameService) AdminRecoverEnergy(ctx context.Context, req *mpb.ReqUserId) (*mpb.Empty, error) {
	err := svc.dao.recoverEnergy(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &mpb.Empty{}, nil
}
