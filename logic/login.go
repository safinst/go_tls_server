package logic

import (
	"time"

	logger "github.com/safinst/go_tls_server/log"

	"github.com/safinst/go_tls_server/model"

	"github.com/safinst/go_tls_server/db"

	"google.golang.org/protobuf/proto"
)

type LoginHandler struct {
	Name string
}

/*
	process login request
	store data to mysql
*/
func (h *LoginHandler) Process(p *model.Request) *model.Response {
	rsp := &model.LoginResponse{
		Ret: proto.Int32(model.RetCode_value["SUCCESS"]),
	}
	req := &model.LoginRequest{}
	err := proto.Unmarshal(p.Buf, req)
	logger.Accesslog.Infof("%v", req)
	if err != nil {
		rsp.Ret = proto.Int32(model.RetCode_value["ERR_REQ_PARA"])
		return fillResponse(rsp, p.GetCmd().Enum())
	}

	// todo token apply jwt
	token := req.GetName() + req.GetKey()
	var loginModel = &model.LoginTab{}
	loginModel.Name = req.GetName()
	loginModel.Key = req.GetKey()
	loginModel.ReadTimes = req.GetReadTimes()
	loginModel.WrittenTimes = req.GetWriteTimes()
	loginModel.Token = token
	db := db.Client().Table(loginModel.Table())
	err = db.Create(loginModel).Error
	if err != nil {
		rsp.Ret = proto.Int32(model.RetCode_value["ERR_TOKEN"])
		logger.Errlog.Errorln(err.Error())
	} else {
		SetCacheData(req.GetName(), cacheData{
			Name:         req.GetName(),
			Token:        token,
			ReadTimes:    req.GetReadTimes(),
			WriteTimes:   req.GetWriteTimes(),
			CurrentRead:  0,
			CurrentWrite: 0,
			UpateTime:    time.Now().Unix(),
		})
	}
	return fillResponse(rsp, p.GetCmd().Enum())
}
