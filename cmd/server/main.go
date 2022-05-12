package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	logger "github.com/safinst/go_tls_server/log"

	"net"

	"github.com/safinst/go_tls_server/logic"
	"github.com/safinst/go_tls_server/model"
	"github.com/safinst/go_tls_server/util"

	"google.golang.org/protobuf/proto"

	"net/http"
	_ "net/http/pprof"
)

var (
	getreqcounter       uint32 = 0
	setreqcounter       uint32 = 0
	loginreqcounter     uint32 = 0
	getrspcounter       uint32 = 0
	setrspcounter       uint32 = 0
	loginrspcounter     uint32 = 0
	lastgetreqcounter   uint32 = 0
	lastsetreqcounter   uint32 = 0
	lastloginreqcounter uint32 = 0
	lastgetrspcounter   uint32 = 0
	lastsetrspcounter   uint32 = 0
	lastloginrspcounter uint32 = 0
)

var (
	starttime int64 = 0
	endtime   int64 = 0
)

func handleConn(c net.Conn) {
	defer c.Close()
	rbuf := bufio.NewReader(c)
	//wbuf := bufio.NewWriter(c)
	for {
		data, err := util.Decode(rbuf)
		if err == nil && len(data) > 0 {
			p := &model.Request{}
			err := proto.Unmarshal(data, p)
			if err != nil {
				logger.Errlog.Errorln("Unmarshal Err:", err.Error())
				//logger.Errlog.Errorf("%v", data)
				continue
			}
			var rsp *model.Response
			switch p.GetCmd() {
			case *model.Cmd_GET.Enum():
				l := &logic.GetHandler{
					Name: "get",
				}
				atomic.AddUint32(&getreqcounter, 1)
				rsp = l.Process(p)
			case *model.Cmd_SET.Enum():
				l := &logic.SetHandler{
					Name: "set",
				}
				atomic.AddUint32(&setreqcounter, 1)
				rsp = l.Process(p)
			case *model.Cmd_LOGIN.Enum():
				l := &logic.LoginHandler{
					Name: "login",
				}
				atomic.AddUint32(&loginreqcounter, 1)
				rsp = l.Process(p)
			}
			buf, err := proto.Marshal(rsp)
			if err != nil {
				logger.Errlog.Errorln(" Marshal err :", err.Error())
				continue
			}
			err = util.Encode(c, buf)
			switch p.GetCmd() {
			case *model.Cmd_GET.Enum():
				atomic.AddUint32(&getrspcounter, 1)
			case *model.Cmd_SET.Enum():
				atomic.AddUint32(&setrspcounter, 1)
			case *model.Cmd_LOGIN.Enum():
				atomic.AddUint32(&loginrspcounter, 1)
			}
			if err != nil {
				logger.Errlog.Errorln(" Encode err :", err.Error())
				return
			}
		}
		if err != nil && err != io.EOF {
			logger.Errlog.Errorln(" read error", err.Error())
		}
	}
}

func fillConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		logger.Errlog.Errorln("LoadX509KeyPair Error ", err.Error())
		return nil, err
	}
	certBytes, err := ioutil.ReadFile("client.pem")
	if err != nil {
		logger.Errlog.Errorln("Unable to read cert.pem ", err.Error())
		return nil, err
	}
	clientCertPool := x509.NewCertPool()
	if ok := clientCertPool.AppendCertsFromPEM(certBytes); !ok {
		logger.Errlog.Errorln("failed to parse root certificate")
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}, nil
}

func statistic() {
	endtime = time.Now().Unix()
	totalreq := getreqcounter + setreqcounter + loginreqcounter - lastgetreqcounter - lastsetreqcounter - lastloginreqcounter
	totalrsp := getrspcounter + setrspcounter + loginrspcounter - lastgetrspcounter - lastsetrspcounter - lastloginrspcounter
	costtime := endtime - starttime
	logger.Statisticlog.Infoln("recv req total:", totalreq)
	logger.Statisticlog.Infoln("send rsp total:", totalrsp)
	logger.Statisticlog.Infoln("req q/s:", float32(totalreq)/float32(costtime))
	logger.Statisticlog.Infoln("rsp q/s:", float32(totalrsp)/float32(costtime))
	starttime = endtime
	atomic.SwapUint32(&lastgetreqcounter, getreqcounter)
	atomic.SwapUint32(&lastsetreqcounter, setreqcounter)
	atomic.SwapUint32(&lastloginreqcounter, loginreqcounter)
	atomic.SwapUint32(&lastgetrspcounter, getrspcounter)
	atomic.SwapUint32(&lastsetrspcounter, setrspcounter)
	atomic.SwapUint32(&lastloginrspcounter, loginrspcounter)
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			logger.Datalog.Infoln("signal:", sig)
			statistic()
			os.Exit(0)
		}
	}()

}
func main() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			statistic()
		}
	}()
	config, err := fillConfig()
	if err != nil {
		logger.Errlog.Errorln("fill config failed ", err.Error())
		return
	}
	ln, err := tls.Listen("tcp", ":10000", config)
	if err != nil {
		logger.Errlog.Errorln(err.Error())
		return
	}
	logger.Datalog.Println("server start ok(on *:443)")
	defer ln.Close()
	starttime = time.Now().Unix()
	go http.ListenAndServe("0.0.0.0:6060", nil)
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Errlog.Errorln("accept error ", err.Error())
			continue
		}
		go handleConn(conn)
	}
}
