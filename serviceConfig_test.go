package zConfig

import (
    "log"
    "testing"

    "gopkg.in/yaml.v2"
)

func TestConfMapGet(t *testing.T) {
    c, err := NewServiceConfig("config.yaml")
    if nil != err {
        log.Fatalln(err)
    }
    t.Log(c.Includes)
    t.Log(c.Configs["filevideoRootUrl"])
}

func TestConfig2(t *testing.T) {
    // 读取配置文件
    data, err := ReadConfigFile("config.yaml")
    if nil != err {
        log.Println(err)
        return
    }
    var cfg map[string]interface{}
    err = yaml.Unmarshal(data, &cfg)
    if nil != err {
        log.Println(err)
        return
    }
    log.Println(cfg)
}

func TestConfig3(t *testing.T) {
    // 读取配置文件
    cfg, err := NewServiceConfig("config.yaml")
    if nil != err {
        log.Println(err)
        return
    }
    log.Println(cfg.RabbitMQ.On)
    log.Println(cfg.RabbitMQ.Server[0].Exchanges[0])
    log.Println(cfg.RabbitMQ.Server[0].Exchanges[0].RoutingKeys)
}

func TestConfig4(t *testing.T) {
    // 读取配置文件
    cfg, err := NewServiceConfig("config.json")
    if nil != err {
        log.Println(err)
        return
    }
    log.Println(cfg.RabbitMQ.On)
    log.Println(cfg.RabbitMQ.Server[0].Exchanges[0])
    log.Println(cfg.RabbitMQ.Server[0].Exchanges[0].RoutingKeys)
}