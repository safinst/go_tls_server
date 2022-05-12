package logic

import "tcp/model"

type Handler interface {
	Process(*model.Request) *model.Response
}
