package zConfig

import (

)

// ServiceConfig 配置文件
type ServiceConfig struct {
    HTTP        struct {
        Port        []int           `json:"port" yaml:"port"`
    }                           `json:"http" yaml:"http"`
    HTTPS       struct {
        Port        []int           `json:"port" yaml:"port"`
        CertFile    string          `json:"certFile" yaml:"certFile"`
        KeyFile     string          `json:"keyFile" yaml:"keyFile"`
    }                           `json:"https" yaml:"https"`
    Database    []*ServiceDB    `json:"database" yaml:"database"`
    RabbitMQ    RabbitMQ        `json:"rabbitMQ" yaml:"rabbitMQ"`
    Consts      []*Consts       `json:"consts" yaml:"consts"`
    Includes    []string        `json:"includes" yaml:"includes"`
    Configs     map[string]string
}

// ServiceDB 配置文件的数据库配置
type ServiceDB struct {
    Key         string `json:"key" yaml:"key"`
    Type        string `json:"type" yaml:"type"`
    Driver      string `json:"driver" yaml:"driver"`
    Host        string `json:"host" yaml:"host"`
    Port        int    `json:"port" yaml:"port"`
    User        string `json:"user" yaml:"user"`
    Pwd         string `json:"pwd" yaml:"pwd"`
    Instance    string `json:"instance" yaml:"instance"`
    Mode        string `json:"mode" yaml:"mode"`
}

// RabbitMQ 配置文件的rabbitMQ配置
type RabbitMQ struct {
    On          bool                    `json:"on" yaml:"on"`                   // 是否启用
    Server      []*RabbitMQServer       `json:"server" yaml:"server"`           // 服务配置
}
// RabbitMQServer rabbitMQ的服务器配置
type RabbitMQServer struct {
    User        string                  `json:"user" yaml:"user"`               // 用户名
    Pwd         string                  `json:"pwd" yaml:"pwd"`                 // 密码
    Host        string                  `json:"host" yaml:"host"`               // 主机名 IP地址
    Port        int                     `json:"port" yaml:"port"`               // 端口号
    Exchanges   []*RabbitMQExchange     `json:"exchanges" yaml:"exchanges"`     // rabbitMQ 交换器
}
// RabbitMQExchange rabbitMQ的交换器配置
type RabbitMQExchange struct {
    Name        string                  `json:"name" yaml:"name"`               // 交换器名称
    Kind        string                  `json:"kind" yaml:"kind"`               // 交换器类型 fanout direct topic headers 4类
    QueueName   string                  `json:"queuename" yaml:"queuename"`     // 队列名称
    RoutingKeys []string                `json:"routingKeys" yaml:"routingKeys"` // 路由/绑定的关键词
}

// Consts 配置文件的常量配置
type Consts struct {
    Key     string      `json:"key" yaml:"key"`
    Value   string      `json:"value" yaml:"value"`
    Sub     []*Consts   `json:"sub" yaml:"sub"`
}