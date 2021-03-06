package SkuPi

import (
	"github.com/warnstar/tao"
	"context"
)

type Pi struct {
	ConnId     int64
	Ctx     	context.Context
	ConnWriter tao.WriteCloser
	Info       *Info
	IsTimeSync bool
	IsTsiPreStart bool
	IsTsiPreStop bool
	IsTsiTestStart bool
	IsTsiTestStop bool
	IsWriteKbFinish bool
	IsUploadDac bool
	IsSendResult bool
}

type Info struct {
	Name       string
	ConnectNow int
	Member [] int
}
