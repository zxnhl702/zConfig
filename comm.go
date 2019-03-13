package zConfig

import (
    "io/ioutil"
    "os"
)

// ReadConfigFile 读取配置文件
func ReadConfigFile(configFilePath string) ([]byte, error) {
    var data []byte
    // 打开配置文件
    f, err := os.Open(configFilePath)
    if nil != err {
        return data, err
    }
    // 读取配置文件
    data, err = ioutil.ReadAll(f)
    if nil != err {
        return data, err
    }
    return data, nil
}