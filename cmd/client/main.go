package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/safinst/go_tls_server/model"
	"github.com/safinst/go_tls_server/util"

	logger "github.com/safinst/go_tls_server/log"

	"github.com/lucasepe/codename"
	"google.golang.org/protobuf/proto"
)

var keys = make([]string, 10000)

func init() {
	file, err := os.OpenFile("data.txt", os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		logger.Errlog.Errorln("open file error")
		panic(err)
	}
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		keys = append(keys, string(a))
	}
	logger.Datalog.Infoln("keys size:", len(keys))
}

func fillConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		logger.Errlog.Errorln("LoadX509KeyPair err:", err.Error())
		return nil, err
	}
	certBytes, err := ioutil.ReadFile("client.pem")
	if err != nil {
		logger.Errlog.Errorln("ReadFile err:", err.Error())
		return nil, err
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		logger.Errlog.Errorln("failed to parse root certificate :", err.Error())
		return nil, err
	}
	return &tls.Config{
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}, nil
}

func startNewConn() {
	defer func() {
		if err := recover(); err != nil {
			logger.Errlog.Panicln("Recovered. Error:", err)
		}
	}()
	conf, err := fillConfig()
	if err != nil {
		logger.Errlog.Errorln("get config err:", err.Error())
		return
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		logger.Errlog.Errorln("dial err:", err.Error())
		return
	}
	defer conn.Close()

	logger.Datalog.Infof("dial %s ok", conn.RemoteAddr())

	go func() {
		rbuf := bufio.NewReader(conn)
		for {
			data, err := util.Decode(rbuf)
			if err == nil && len(data) > 0 {
				data, err := util.Decode(rbuf)
				if err == nil && len(data) > 0 {
					p := &model.Response{}
					err := proto.Unmarshal(data, p)
					if err != nil {
						logger.Errlog.Errorln(err.Error())
						continue
					}
					var rsp proto.Message
					switch p.GetCmd() {
					case *model.Cmd_GET.Enum():
						rsp = &model.GetResponse{}
					case *model.Cmd_SET.Enum():
						rsp = &model.SetResponse{}
					case *model.Cmd_LOGIN.Enum():
						rsp = &model.LoginResponse{}
					}
					err = proto.Unmarshal(p.Buf, rsp)
					if err != nil {
						panic(err)
					}
					logger.Datalog.Infoln("rsp ", p)
				}
				if err != nil {
					logger.Errlog.Errorln("err=", err.Error())
				}
			}
		}
	}()

	for {
		req := fillRequest()
		logger.Datalog.Infoln("req ", req)
		buf, err := proto.Marshal(req)
		err = util.Encode(conn, buf)
		if err != nil {
			logger.Errlog.Errorln("send req error:", err.Error())
			return
		}
	}
}

func fillRequest() *model.Request {
	var buf []byte
	var err error
	var cmd *model.Cmd
	rand.Seed(int64(time.Now().UnixNano()))
	i := rand.Int() % 1000000
	var reqtype string
	t := i % 100
	if t <= 80 {
		reqtype = "get"
	} else if i > 80 && i <= 95 {
		reqtype = "set"
	}
	name := strconv.Itoa(i)
	key := name
	token := name + key
	rng, err := codename.DefaultRNG()
	if err != nil {
		logger.Errlog.Error("DefaultRNG:", err.Error())
		return nil
	}
	switch reqtype {
	case "get":
		if i > len(keys) {
			key := codename.Generate(rng, 1000)
			getReq := &model.GetRequest{
				Key: proto.String(key),
				Head: &model.CommonHead{
					Name:  proto.String(name),
					Token: proto.String(token),
				},
			}
			buf, err = proto.Marshal(getReq)
		} else {
			getReq := &model.GetRequest{
				Key: proto.String(keys[i]),
				Head: &model.CommonHead{
					Name:  proto.String(name),
					Token: proto.String(token),
				},
			}
			buf, err = proto.Marshal(getReq)
		}
		cmd = model.Cmd_GET.Enum()
	case "set":
		if i > len(keys) {
			key := codename.Generate(rng, 1000)
			val := codename.Generate(rng, 1000)
			setReq := &model.SetRequest{
				Key: proto.String(key),
				Val: proto.String(val),
				Head: &model.CommonHead{
					Name:  proto.String(name),
					Token: proto.String(token),
				},
			}
			buf, err = proto.Marshal(setReq)
		} else {
			val := codename.Generate(rng, 1000)
			setReq := &model.SetRequest{
				Key: proto.String(keys[i]),
				Val: proto.String(val),
				Head: &model.CommonHead{
					Name:  proto.String(name),
					Token: proto.String(token),
				},
			}
			buf, err = proto.Marshal(setReq)
		}
		cmd = model.Cmd_SET.Enum()
	case "login":
		loginReq := &model.LoginRequest{
			Name:       proto.String(name),
			Key:        proto.String(key),
			ReadTimes:  proto.Uint32(util.Min(100000, uint32(i*3+3))),
			WriteTimes: proto.Uint32(util.Min(100000, uint32(i+3))),
		}
		buf, err = proto.Marshal(loginReq)
		cmd = model.Cmd_LOGIN.Enum()
	}
	if err != nil {
		panic(err)
	}
	return &model.Request{
		Cmd: cmd,
		Buf: buf,
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			startNewConn()
			wg.Done()
		}()
	}
	wg.Wait()
}
