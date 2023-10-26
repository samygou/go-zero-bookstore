package logic

import (
	"context"
	"errors"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/app/bookstore/repository"
	"go-zero-bookstore/common/logx"
)

type GetBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取book
func (l *GetBookLogic) GetBook(in *pb.GetBookReq) (*pb.GetBookResp, error) {
	book, err := l.svcCtx.BookModel.FindOne(l.ctx, in.Id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		logx.Error(err)
		return nil, ErrInternalFault
	}
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrBookNotExist
	}

	return &pb.GetBookResp{
		Book: &pb.Book{
			Id:    book.Id,
			Name:  book.Name,
			Price: book.Price,
			Desc:  book.Desc,
		},
	}, nil
}
