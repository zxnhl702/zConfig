package zConfig

import (
    "log"
    "testing"
)

func TestCommConfigGetYAML(t *testing.T) {
    cfg, err := NewCommConfig("./cfg.yaml")
    log.Println(err)
    log.Println(cfg)
    n, err := Get(cfg.config, "database.0.host")
    log.Println(err)
    log.Println(n)
}

func TestCommConfigGetJSON(t *testing.T) {
    cfg, err := NewCommConfig("./cfg.json")
    log.Println(err)
    log.Println(cfg)
    n, err := Get(cfg.config, "database.0.host")
    log.Println(err)
    log.Println(n)
}