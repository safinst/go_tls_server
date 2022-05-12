package logic

import (
	"sync"
	"tcp/db"
	logger "tcp/log"
	"tcp/model"
	"time"

	"google.golang.org/protobuf/proto"
)

const prefix = "testtesttesttest_"

var memCache sync.Map

type Result struct {
	Name       string `gorm:"column:name"`
	Token      string `gorm:"column:token"`
	ReadTimes  uint32 `gorm:"column:readtimes"`
	WriteTimes uint32 `gorm:"column:writetimes"`
}
type cacheData struct {
	Name         string
	Token        string
	ReadTimes    uint32
	WriteTimes   uint32
	CurrentRead  uint32
	CurrentWrite uint32
	UpateTime    int64
}

var list = make([]Result, 0)

func init() {
	var loginModel = &model.LoginTab{}
	db.Client().Table(loginModel.Table()).Raw("SELECT name, token, readtimes, writetimes FROM test_login_tab").Scan(&list)
	t := time.Now().Unix()
	for _, item := range list {
		d := cacheData{
			Name:         item.Name,
			Token:        item.Token,
			ReadTimes:    item.ReadTimes,
			WriteTimes:   item.WriteTimes,
			CurrentRead:  0,
			CurrentWrite: 0,
			UpateTime:    t,
		}
		memCache.Store(item.Name, d)
	}
	logger.Datalog.Infoln(" load data size:", len(list))
}
func GetCacheData(key string) (any interface{}, ok bool) {
	return memCache.Load(key)
}

func SetCacheData(key string, val interface{}) {
	memCache.Store(key, val)
}

func CheckToken(name, token string) bool {
	cd, ok := memCache.Load(name)
	if !ok {
		logger.Errlog.Errorln("CheckToken GetCacheData Fail ")
		return false
	}
	val, ok := cd.(cacheData)
	if !ok {
		logger.Errlog.Errorln("CheckToken GetCacheData Fail ")
		return false
	}
	if token != val.Token {
		logger.Errlog.Errorln("Error token")
		return false
	}
	return true
}

func CheckLimit(key, t string) bool {
	cd, ok := memCache.Load(key)
	if !ok {
		logger.Errlog.Errorln("CheckLimit GetCacheData Fail ")
		return false
	}
	//logger.Datalog.Debugf("%v", cd)
	val, ok := cd.(cacheData)
	if !ok {
		logger.Errlog.Errorln("CheckLimit GetCacheData Fail ")
		return false
	}
	if time.Now().Unix()-val.UpateTime > 1 {
		val.CurrentRead = 0
		val.CurrentWrite = 0
		val.UpateTime = time.Now().Unix()
	}
	switch t {
	case "get":
		val.CurrentRead++
	case "set":
		val.CurrentWrite++
	}
	memCache.Store(key, val)
	if val.CurrentRead > val.ReadTimes {
		return false
	}
	if val.CurrentWrite > val.WriteTimes {
		return false
	}
	return true
}

func fillResponse(m proto.Message, cmd *model.Cmd) *model.Response {
	buf, err := proto.Marshal(m)
	if err != nil {
		logger.Errlog.Errorln("Marshal Error:", err.Error())
	}
	logger.Accesslog.Infof("%+v", m)
	return &model.Response{
		Cmd: cmd,
		Buf: buf,
	}
}
