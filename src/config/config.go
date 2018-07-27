package config

import (
	"path/filepath"
	"strings"
	"errors"
)

var AppConfig Config

type Config interface {
	GetString(key string) string
	GetStringWithDefVal(key string, defVal string) string

	GetInt(key string) int
	GetIntWithDefVal(key string, defVal int) int

	GetInt64(key string) int64
	GetInt64WithDefVal(key string, defVal int64) int64

	GetFloat32(key string) float32
	GetFloat32WithDefVal(key string, defVal float32) float32

	GetFloat64(key string) float64
	GetFloat64WithDefVal(key string, defVal float64) float64

	GetBool(key string) bool
	GetBoolWithDefVal(key string, defVal bool) bool
}

func init() {
	AppConfig, _ = NewConfig("ini", "config/app.config")
}

func NewConfig(adapter, filename string) (Config, error) {
	path, err := GetCurrentPath(filename)
	if err != nil {
		return nil, err
	}
	switch adapter {
	case "ini":
		return LoadIniConfigFile(path)
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}

func GetCurrentPath(filename string) (path string, err error) {
	path, err = filepath.Abs(filename)
	if err != nil {
		return
	}
	path = strings.Replace(path, "\\", "/", -1)
	path = strings.Replace(path, "\\\\", "/", -1)
	return
}
