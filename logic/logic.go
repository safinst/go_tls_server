package logic

import "github.com/safinst/go_tls_server/model"

type Handler interface {
	Process(*model.Request) *model.Response
}
