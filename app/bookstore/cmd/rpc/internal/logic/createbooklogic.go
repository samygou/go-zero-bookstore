package logic

import (
	"context"
	"time"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/app/bookstore/repository"
	"go-zero-bookstore/common/logx"
)

type CreateBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBookLogic {
	return &CreateBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 添加book
func (l *CreateBookLogic) CreateBook(in *pb.CreateBookReq) (*pb.CreateBookResp, error) {
	if len(in.Name) == 0 {
		return nil, ErrBookNameIsNull
	}

	if in.Price <= 0 {
		return nil, ErrBookPriceIsIncorrect
	}

	now := time.Now().Unix()

	result, err := l.svcCtx.BookModel.Insert(l.ctx, &repository.Books{
		CreateAt: now,
		UpdateAt: now,
		Name:     in.Name,
		Price:    in.Price,
		Desc:     in.Desc,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	bookId, err := result.LastInsertId()
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &pb.CreateBookResp{
		Id: bookId,
	}, nil
}
