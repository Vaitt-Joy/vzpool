package config

import (
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)

type IniConfig struct {
	ConfigMap map[string]string
	tempKey   string
}

func LoadIniConfigFile(fileName string) (*IniConfig, error) {
	middle := "."
	config := new(IniConfig)
	config.ConfigMap = make(map[string]string)
	// load file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
		return nil, err
	}

	defer file.Close()

	read := bufio.NewReader(file)
	isAnnotation := false
	for {
		b, _, err := read.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		str := strings.TrimSpace(string(b))

		if strings.Index(str, "#") == 0 || strings.Index(str, "//") == 0 {
			continue
		}

		if isAnnotation {
			if strings.Index(str, "*/") == 0 {
				isAnnotation = false
			}
			continue
		}

		if strings.Index(str, "/*") == 0 {
			isAnnotation = true
			continue
		}

		// 读取小节配置

		s1 := strings.Index(str, "[")
		s2 := strings.Index(str, "]")

		if s1 > -1 && s2 > -1 && s2 > s1+1 {
			config.tempKey = strings.TrimSpace(str[s1+1:s2])
			continue
		}

		if len(config.tempKey) < 1 {
			continue
		}

		// 等号 =
		eqIndex := strings.Index(str, "=")
		if eqIndex < 0 {
			continue
		}
		eqLeft := strings.TrimSpace(str[0:eqIndex])

		if len(eqLeft) < 1 {
			continue
		}

		eqRight := strings.TrimSpace(str[eqIndex+1: ])

		pos := strings.Index(eqRight, "\t#")

		val := eqRight

		if pos > -1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}

		pos = strings.Index(eqRight, " #")
		if pos > -1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}

		pos = strings.Index(eqRight, "\t//")
		if pos > -1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}

		pos = strings.Index(eqRight, " //")
		if pos > -1 {
			val = strings.TrimSpace(eqRight[0:pos])
		}
		if len(val) < 1 {
			continue
		}
		key := config.tempKey + middle + eqLeft
		config.ConfigMap[key] = strings.TrimSpace(val)
	}

	return config, nil
}

func (self *IniConfig) Get(key string) string {
	return self.GetWithDefVal(key, "")
}

func (self *IniConfig) GetWithDefVal(key string, defVal string) string {
	v, ok := self.ConfigMap[key]
	if ok {
		return v
	}
	return defVal
}

func (self *IniConfig) GetString(key string) string {
	return self.GetWithDefVal(key, "")
}

func (self *IniConfig) GetStringWithDefVal(key string, defVal string) string {
	return self.GetWithDefVal(key, defVal)
}

func (self *IniConfig) GetInt(key string) int {
	return self.GetIntWithDefVal(key, 0)
}

func (self *IniConfig) GetIntWithDefVal(key string, defVal int) int {
	val, err := strconv.Atoi(self.Get(key))
	if err != nil {
		panic(err)
		return defVal
	}
	return val
}

func (self *IniConfig) GetInt64(key string) int64 {
	return self.GetInt64WithDefVal(key, 0)
}

func (self *IniConfig) GetInt64WithDefVal(key string, defVal int64) int64 {
	val, err := strconv.ParseInt(self.Get(key), 10, 64)
	if err != nil {
		panic(err)
		return defVal
	}
	return val
}

func (self *IniConfig) GetFloat32(key string) float32 {
	return self.GetFloat32WithDefVal(key, 0)
}

func (self *IniConfig) GetFloat32WithDefVal(key string, defVal float32) float32 {
	val, err := strconv.ParseFloat(self.Get(key), 32)
	if err != nil {
		panic(err)
		return float32(defVal)
	}
	return float32(val)
}

func (self *IniConfig) GetFloat64(key string) float64 {
	return self.GetFloat64WithDefVal(key, 0)
}

func (self *IniConfig) GetFloat64WithDefVal(key string, defVal float64) float64 {
	val, err := strconv.ParseFloat(self.Get(key), 64)
	if err != nil {
		panic(err)
		return float64(defVal)
	}
	return val
}

func (self *IniConfig) GetBool(key string) bool {
	return self.GetBoolWithDefVal(key, false)
}

func (self *IniConfig) GetBoolWithDefVal(key string, defVal bool) bool {
	val, err := strconv.ParseBool(self.Get(key))
	if err != nil {
		panic(err)
		return defVal
	}
	return val
}
