package zConfig

import (
    "encoding/json"
    "errors"
    "path"
    "strings"

    "gopkg.in/yaml.v2"
)

// CommConfig 配置文件
type CommConfig struct {
    config  interface{}
}

// NewCommConfig 新建配置
func NewCommConfig(filepath string) (*CommConfig, error) {
    // 读取配置文件扩展名
    ext := strings.ToLower(path.Ext(filepath))
    // 读取配置文件
    data, err := ReadConfigFile(filepath)
    if nil != err {
        return nil, err
    }
    // 解析配置文件
    c, err := parseCommConfig(data, ext)
    return c, err
}

// parseCommConfig 解析配置文件数据
func parseCommConfig(data []byte, ext string) (*CommConfig, error) {
    var c CommConfig
    var err error
    // 解析配置
    if ".json" == ext {
        err = json.Unmarshal(data, &c)
    } else if ".yaml" == ext {
        err = yaml.Unmarshal(data, &c)
    } else {
        return &c, ErrUnkownConfigFile
    }
    return &c, err
}

// Map 配置文件转换成map[string]interface{}
func (c *CommConfig)Map() (map[string]interface{}, error) {
    if m, ok := (c.config).(map[string]interface{}); ok {
        return m, nil
    }
    return nil, errors.New("类型转换错误")
}