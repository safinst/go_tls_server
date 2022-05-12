package log

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	Datalog      = logrus.New()
	Accesslog    = logrus.New()
	Statisticlog = logrus.New()
	Errlog       = logrus.New()
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)

	file, err := os.OpenFile("data.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err == nil {
		Datalog.SetOutput(file)
	} else {
		log.Errorln("open file error")
		Datalog.SetOutput(os.Stdout)
	}
	Datalog.SetLevel(log.DebugLevel)
	Datalog.SetFormatter(&log.JSONFormatter{})
	Datalog.SetReportCaller(true)

	file, err = os.OpenFile("statistic.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err == nil {
		Statisticlog.SetOutput(file)
	} else {
		log.Errorln("open file error")
		Statisticlog.SetOutput(os.Stdout)
	}
	Statisticlog.SetFormatter(&log.JSONFormatter{})
	Statisticlog.SetReportCaller(true)
	Statisticlog.SetLevel(log.DebugLevel)

	file, err = os.OpenFile("access.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err == nil {
		Accesslog.SetOutput(file)
	} else {
		log.Errorln("open file error")
		Accesslog.SetOutput(os.Stdout)
	}
	Accesslog.SetFormatter(&log.JSONFormatter{})
	Accesslog.SetReportCaller(true)
	Accesslog.SetLevel(log.DebugLevel)

	file, err = os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err == nil {
		Errlog.SetOutput(file)
	} else {
		log.Errorln("open file error")
		Errlog.SetOutput(os.Stdout)
	}
	Errlog.SetFormatter(&log.JSONFormatter{})
	Errlog.SetReportCaller(true)
	Errlog.SetLevel(log.DebugLevel)
}
