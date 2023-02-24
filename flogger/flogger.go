package flogger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type commonMinimumDenominator interface {
	Errorf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Debugf(format string, a ...interface{})
	Warnf(format string, a ...interface{})
}

func getLogrus() commonMinimumDenominator {
	l := logrus.New()
	l.SetLevel(logrus.DebugLevel)
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "15:04:05"
	l.SetFormatter(customFormatter)
	return l
}

func get() commonMinimumDenominator {
	return getLogrus()
}

func Pretty(a interface{}) {
	get().Debugf("%+v", a)
}

func Errorf(format string, a ...interface{}) {
	if len(a) == 0 {
		get().Errorf(format)
		return
	}
	if a[0] == nil {
		get().Errorf(format)
		return
	}
	get().Errorf(format, a...)
}

func Infof(format string, a ...interface{}) {
	if len(a) == 0 {
		get().Infof(format)
		return
	}
	if a[0] == nil {
		get().Infof(format)
		return
	}
	get().Infof(format, a...)
}

func Debugf(format string, a ...interface{}) {
	if len(a) == 0 {
		get().Debugf(format)
		return
	}
	if a[0] == nil {
		get().Debugf(format)
		return
	}
	get().Debugf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	if len(a) == 0 {
		get().Warnf(format)
		return
	}
	if a[0] == nil {
		get().Warnf(format)
		return
	}
	get().Warnf(format, a...)
}

func Localf(format string, a ...interface{}) {
	_, local := os.LookupEnv("LOCAL_ENVIRONMENT")
	if !local {
		return
	}
	Debugf(format, a...)
}
