package logic

import (
	"context"
	"errors"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/app/usercenter/repository"
	"go-zero-bookstore/common/logx"
)

type DeleteBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBookLogic {
	return &DeleteBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 删除book
func (l *DeleteBookLogic) DeleteBook(in *pb.DeleteBookReq) (*pb.DeleteBookResp, error) {
	_, err := l.svcCtx.BookModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		logx.Error(err)
		return nil, ErrInternalFault
	}
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrBookNotExist
	}

	err = l.svcCtx.BookModel.Delete(l.ctx, in.Id)
	if err != nil {
		logx.Error(err)
		return nil, ErrInternalFault
	}

	return &pb.DeleteBookResp{
		Id: in.Id,
	}, nil
}
