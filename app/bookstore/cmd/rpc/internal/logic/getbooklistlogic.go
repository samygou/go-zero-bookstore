package logic

import (
	"context"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/app/bookstore/repository"
	"go-zero-bookstore/common/logx"
)

type GetBookListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookListLogic {
	return &GetBookListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取book list
func (l *GetBookListLogic) GetBookList(in *pb.GetBookListReq) (*pb.GetBookListResp, error) {
	page := l.svcCtx.Config.DefaultPageParam.Page
	if in.Page != nil && *in.Page > 1 {
		page = *in.Page
	}

	pageSize := l.svcCtx.Config.DefaultPageParam.PageSize
	if in.PageSize != nil && *in.PageSize > 0 {
		pageSize = *in.PageSize
	}

	results, err := l.svcCtx.BookModel.GetBookList(l.ctx, &repository.GetBookListReq{
		Name:     in.Name,
		Price:    in.Price,
		OrderBy:  in.OrderBy,
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		logx.Error(err)
		return nil, ErrInternalFault
	}

	var books []*pb.Book

	for _, book := range results {
		books = append(books, &pb.Book{
			Id:    book.Id,
			Name:  book.Name,
			Price: book.Price,
			Desc:  book.Desc,
		})
	}

	return &pb.GetBookListResp{
		Books: books,
	}, nil
}
