package logic

import (
	"tcp/cache"
	logger "tcp/log"
	"tcp/model"

	"google.golang.org/protobuf/proto"
)

type GetHandler struct {
	Name string
}

func (handle *GetHandler) Process(p *model.Request) *model.Response {
	rsp := &model.GetResponse{
		Ret: proto.Int32(model.RetCode_value["SUCCESS"]),
	}
	req := &model.GetRequest{}
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
	if !CheckLimit(req.GetHead().GetName(), handle.Name) {
		rsp.Ret = proto.Int32(model.RetCode_value["ERR_FREQUENCY"])
		return fillResponse(rsp, p.GetCmd().Enum())
	}
	key := cache.MakeKey(prefix, req.GetKey())
	// t1 := time.Now().UnixNano()
	v, err := cache.Client().Get(key)
	// t2 := time.Now().UnixNano()
	if err != nil {
		logger.Errlog.Errorln(err.Error())
	} else {
		//logger.Datalog.Infoln("redis succ cost time:", t2-t1)
		rsp.Val = proto.String(string(v))
	}
	return fillResponse(rsp, p.GetCmd().Enum())
}
