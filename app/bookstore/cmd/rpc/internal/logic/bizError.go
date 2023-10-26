package logic

import "go-zero-bookstore/common/errorx"

var (
	ErrInternalFault        = errorx.NewError(errorx.CodeInternal, "internal err")
	ErrBookNameIsNull       = errorx.NewError(errorx.CodeInvalidArgument, "The book name must not been null")
	ErrBookPriceIsIncorrect = errorx.NewError(errorx.CodeInvalidArgument, "The price of the book is incorrect")
	ErrBookNotExist         = errorx.NewError(errorx.CodeNotFound, "The book record not exist")
)
