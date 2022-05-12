package logic

import (
	"time"

	"github.com/safinst/go_tls_server/cache"

	"github.com/safinst/go_tls_server/model"

	"google.golang.org/protobuf/proto"

	logger "github.com/safinst/go_tls_server/log"
)

type SetHandler struct {
	Name string
}

func (h *SetHandler) Process(p *model.Request) *model.Response {
	rsp := &model.SetResponse{
		Ret: proto.Int32(model.RetCode_value["SUCCESS"]),
	}
	req := &model.SetRequest{}
	err := proto.Unmarshal(p.Buf, req)
	logger.Accesslog.Infof("%v", req)
	if err != nil {
		rsp.Ret = proto.Int32(model.RetCode_value["ERR_REQ_PARA"])
		return fillResponse(rsp, p.GetCmd().Enum())
	}
	if !CheckToken(req.GetHead().GetName(), req.GetHead().GetToken()) {
		rsp.Ret = proto.Int32(model.RetCode_value["ERR_TOKEN"])
		return fillResponse(rsp, p.GetCmd().Enum())
	}
	if !CheckLimit(req.GetHead().GetName(), h.Name) {
		rsp.Ret = proto.Int32(model.RetCode_value["ERR_FREQUENCY"])
		return fillResponse(rsp, p.GetCmd().Enum())
	}
	key := cache.MakeKey(prefix, req.GetKey())
	t1 := time.Now().UnixNano()
	err = cache.Client().SetExpire(key, 10, req.GetVal())
	t2 := time.Now().UnixNano()
	if err != nil {
		logger.Errlog.Errorln(" get redis error:", err.Error())
	} else {
		logger.Datalog.Infoln("redis succ cost time ", t2-t1)
	}
	return fillResponse(rsp, p.GetCmd().Enum())
}
