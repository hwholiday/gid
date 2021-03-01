package grpc

import (
	"context"
	"errors"
	gidSrv "gid/api"
)

func (s *Server) GetId(ctx context.Context, in *gidSrv.ReqId) (*gidSrv.ResId, error) {
	if in.GetBizTag() == "" {
		return nil, errors.New("biz_tag is empty")
	}
	id, err := s.srv.GetId(in.GetBizTag())
	if err != nil {
		return nil, err
	}
	return &gidSrv.ResId{
		Id: id,
	}, nil
}
