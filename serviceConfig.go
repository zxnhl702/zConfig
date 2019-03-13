package zConfig

import (
    "encoding/json"
    "errors"
    "log"
    "os"
    "path"
    "strings"

    "gopkg.in/yaml.v2"
)

// ErrUnkownConfigFile 未知的文件类型
var ErrUnkownConfigFile = errors.New("zconf: unknown file format")

// GetEnvConf 获取OS环境变量
func GetEnvConf(key string) string {
    return os.Getenv(key)
}

// NewServiceConfig 初始化配置
func NewServiceConfig(configFilePath string) (*ServiceConfig, error) {
    // 读取配置文件扩展名
    ext := strings.ToLower(path.Ext(configFilePath))
    // 读取配置文件
    data, err := ReadConfigFile(configFilePath)
    if nil != err {
        return nil, err
    }
    // 解析配置文件
    c, err := parseConfig(data, ext)
    // 生成配置文件中的常量
    c.Configs = getConfigs(c)
    return c, err
}

// parseConfig 解析配置文件数据
func parseConfig(data []byte, ext string) (*ServiceConfig, error) {
    var c ServiceConfig
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

// getConfigs 生成配置文件中的常量
func getConfigs(s *ServiceConfig) map[string]string {
    configs := make(map[string]string)
    // 生成配置文件中consts部门的数据为map[string]string
    configs = getConsts(s.Consts)
    // 查看包含的其他配置文件
    for _, v := range s.Includes {
        // 读取配置文件
        data, err := ReadConfigFile(v)
        if nil != err {
            log.Printf("配置文件%s读取失败", v)
            log.Println(err)
            continue
        }
        // 解析配置文件中的数
        consts, err := parseConsts(data, strings.ToLower(path.Ext(v)))
        if nil != err {
            log.Printf("配置文件%s数据解析失败", v)
            log.Println(err)
            continue
        }
        // 常量数据转换成map[string]string的形式
        constmap := getConsts(consts)
        // 遍历数据 把没有重复的放到一起 重复的直接舍去
        for kk, vv := range constmap {
            // 判断常量中是否已经存在
            if _, ok := configs[kk]; !ok {
                configs[kk] = vv
            } else {
                log.Printf("%s的值已经存在，舍弃%s配置文件中的值", kk, v)
            }
        }
    }
    return configs
}

// parseConsts 解析includes中的配置文件 这些配置文件中数据的格式必须按照Consts结构体来
func parseConsts(data []byte, ext string) ([]*Consts, error) {
    var list []*Consts
    var err error
    // 解析配置
    if ".json" == ext {
        err = json.Unmarshal(data, &list)
    } else if ".yaml" == ext {
        err = yaml.Unmarshal(data, &list)
    } else {
        return list, ErrUnkownConfigFile
    }
    return list, err
}

// getConsts 读取常量配置
func getConsts(cfg []*Consts) map[string]string {
    consts := make(map[string]string)
    for _, v := range cfg {
        consts[v.Key] = v.Value
        // 如果有下一级常量
        for _, vv := range v.Sub {
            // 用上一级的常量拼上下一级的常量来获取完全的常量
            consts[vv.Key] = v.Value + vv.Value
        }
    }
    return consts
}