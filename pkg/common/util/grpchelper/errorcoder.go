package grpchelper

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"google.golang.org/grpc/codes"
)

type (
	errCode struct {
		matcher preseterrors.Matcher
		code    codes.Code
	}
)

var errCodes []errCode = make([]errCode, 0)

func GetCode(err error) codes.Code {
	for _, esc := range errCodes {
		if esc.matcher(err) {
			return esc.code
		}
	}
	return codes.Internal
}

func RegisterError(matcher preseterrors.Matcher, code codes.Code) {
	errCodes = append(errCodes, errCode{matcher: matcher, code: code})
}
