package logic

import "go-zero-bookstore/common/errorx"

var (
	ErrAccountAlreadyExists = errorx.NewError(errorx.CodeAlreadyExists, "Mobile already exists")
	ErrAccountInternalFault = errorx.NewError(errorx.CodeInternal, "internal failed")
	ErrUsernameIsEmpty      = errorx.NewError(errorx.CodeInvalidArgument, "username can not been null")
	ErrPasswordIncorrect    = errorx.NewError(errorx.CodeInvalidArgument, "The password is incorrect")
	ErrPasswordUnEqual      = errorx.NewError(errorx.CodeInvalidArgument, "The tow password are different")
	ErrMobileIncorrect      = errorx.NewError(errorx.CodeInvalidArgument, "The format of the mobile number is incorrect")
	ErrAccountNotExists     = errorx.NewError(errorx.CodeNotFound, "The account not exist")
	ErrAccountLoginFailed   = errorx.NewError(errorx.CodeUnauthenticated, "The mobile or password is incorrect")
)
