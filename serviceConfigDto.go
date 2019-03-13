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

// Consts 配置文件的常量配置
type Consts struct {
    Key     string      `json:"key" yaml:"key"`
    Value   string      `json:"value" yaml:"value"`
    Sub     []*Consts   `json:"sub" yaml:"sub"`
}