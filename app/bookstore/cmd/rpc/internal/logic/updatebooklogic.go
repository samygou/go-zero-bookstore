package logic

import (
	"context"
	"errors"
	"time"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/app/bookstore/repository"
	"go-zero-bookstore/common/logx"
)

type UpdateBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookLogic {
	return &UpdateBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 更新book
func (l *UpdateBookLogic) UpdateBook(in *pb.UpdateBookReq) (*pb.UpdateBookResp, error) {
	if len(in.Name) == 0 {
		return nil, ErrBookNameIsNull
	}

	if in.Price <= 0 {
		return nil, ErrBookPriceIsIncorrect
	}

	_, err := l.svcCtx.BookModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		logx.Error(err)
		return nil, ErrInternalFault
	}
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrBookNotExist
	}

	err = l.svcCtx.BookModel.Update(l.ctx, &repository.Books{
		Id: in.Id,
		//CreateAt: book.CreateAt,
		UpdateAt: time.Now().Unix(),
		Name:     in.Name,
		Price:    in.Price,
		Desc:     in.Desc,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &pb.UpdateBookResp{
		Id: in.Id,
	}, nil
}
