package zConfig

import (
    "encoding/json"
    "errors"
    "fmt"
    "path"
    "strings"
    "strconv"

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
    var cfg interface{}
    var err error
    // 解析配置
    if ".json" == ext {
        err = json.Unmarshal(data, &cfg)
    } else if ".yaml" == ext {
        if err = yaml.Unmarshal(data, &cfg); err != nil {
            return nil, err
        }
        cfg, err = normalizeValue(cfg)
    } else {
        return &c, ErrUnkownConfigFile
    }
    c.config = cfg
    return &c, err
}

// Map 配置文件转换成map[string]interface{}
func (c *CommConfig)Map() (map[string]interface{}, error) {
    if m, ok := (c.config).(map[string]interface{}); ok {
        return m, nil
    }
    return nil, errors.New("类型转换错误")
}

// Fetching -------------------------------------------------------------------

// Get returns a child of the given value according to a dotted path.
func Get(cfg interface{}, path string) (interface{}, error) {
    parts := strings.Split(path, ".")
    // Normalize path.
    for k, v := range parts {
        if v == "" {
            if k == 0 {
                parts = parts[1:]
            } else {
                return nil, fmt.Errorf("Invalid path %q", path)
            }
        }
    }
    // Get the value.
    for pos, part := range parts {
        switch c := cfg.(type) {
        case []interface{}:
            if i, error := strconv.ParseInt(part, 10, 0); error == nil {
                if int(i) < len(c) {
                    cfg = c[i]
                } else {
                    return nil, fmt.Errorf(
                        "Index out of range at %q: list has only %v items",
                        strings.Join(parts[:pos+1], "."), len(c))
                }
            } else {
                return nil, fmt.Errorf("Invalid list index at %q",
                    strings.Join(parts[:pos+1], "."))
            }
        case map[string]interface{}:
            if value, ok := c[part]; ok {
                cfg = value
            } else {
                return nil, fmt.Errorf("Nonexistent map key at %q",
                    strings.Join(parts[:pos+1], "."))
            }
        default:
            return nil, fmt.Errorf(
                "Invalid type at %q: expected []interface{} or map[string]interface{}; got %T",
                strings.Join(parts[:pos+1], "."), cfg)
        }
    }

    return cfg, nil
}

// normalizeValue normalizes a unmarshalled value. This is needed because
// encoding/json doesn't support marshalling map[interface{}]interface{}.
func normalizeValue(value interface{}) (interface{}, error) {
    switch value := value.(type) {
    case map[interface{}]interface{}:
        node := make(map[string]interface{}, len(value))
        for k, v := range value {
            key, ok := k.(string)
            if !ok {
                return nil, fmt.Errorf("Unsupported map key: %#v", k)
            }
            item, err := normalizeValue(v)
            if err != nil {
                return nil, fmt.Errorf("Unsupported map value: %#v", v)
            }
            node[key] = item
        }
        return node, nil
    case map[string]interface{}:
        node := make(map[string]interface{}, len(value))
        for key, v := range value {
            item, err := normalizeValue(v)
            if err != nil {
                return nil, fmt.Errorf("Unsupported map value: %#v", v)
            }
            node[key] = item
        }
        return node, nil
    case []interface{}:
        node := make([]interface{}, len(value))
        for key, v := range value {
            item, err := normalizeValue(v)
            if err != nil {
                return nil, fmt.Errorf("Unsupported list item: %#v", v)
            }
            node[key] = item
        }
        return node, nil
    case bool, float64, int, string, nil:
        return value, nil
    }
    return nil, fmt.Errorf("Unsupported type: %T", value)
}