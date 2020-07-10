package grpc

import (
	"context"
	gid "gid/api/grpc"
	"gid/entity"
	"net/http"
)

func (s *Server) Ping(ctx context.Context, in *gid.ReqPing) (out *gid.ResPong, err error) {
	out = &gid.ResPong{
		Status: &gid.Status{
			Code: http.StatusOK,
		},
		Data: "pong",
	}
	return
}

func (s *Server) GetRandId(ctx context.Context, in *gid.ReqRandId) (out *gid.ResRandId, err error) {
	out = &gid.ResRandId{
		Status: &gid.Status{
			Code: http.StatusOK,
		},
		Id: s.srv.SnowFlakeGetId(),
	}
	return
}

func (s *Server) GetId(ctx context.Context, in *gid.ReqId) (out *gid.ResId, err error) {
	var id int64
	out = &gid.ResId{
		Status: &gid.Status{
			Code: http.StatusOK,
		},
	}
	if in.GetTag() == "" {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = "parameter error"
		return
	}
	if id, err = s.srv.GetId(in.GetTag()); err != nil {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = err.Error()
		return
	}
	out.Id = id
	return
}
func (s *Server) CreateTag(ctx context.Context, in *gid.ReqTagCreate) (out *gid.ResTagCreate, err error) {
	out = &gid.ResTagCreate{
		Status: &gid.Status{
			Code: http.StatusOK,
		},
	}
	if in.GetTag() == "" || in.GetStep() == 0 {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = "parameter error"
		return
	}
	if err = s.srv.CreateTag(&entity.Segments{
		BizTag: in.GetTag(),
		MaxId:  in.GetMaxId(),
		Step:   in.GetStep(),
		Remark: in.GetRemark(),
	}); err != nil {
		out.Status.Code = http.StatusInternalServerError
		out.Status.Msg = err.Error()
		return
	}
	return
}
