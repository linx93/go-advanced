package init_code

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// Instance 定义一个简单的实列 借鉴grpc中的创建方式
type Instance struct {
	Code string `json:"code"` //必传字段

	HttpProxyConfig HttpProxyConfig `json:"httpProxyConfig"`

	Message   string `json:"message"`
	Name      string `json:"name"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt int64  `json:"createdAt"`
	Version   string `json:"version" linx:"column:sid;type:like"`
}

// Option 定义一个option
type Option func(*Instance)

//func New(code string, opts ...Option) *Instance {
//	ins := &Instance{}
//	for _, o := range opts {
//		o(ins)
//	}
//	return ins
//}

func New(opts ...Option) *Instance {
	ins := &Instance{}
	for _, o := range opts {
		o(ins)
	}
	return ins
}

// HttpProxyConfig 代理配置
type HttpProxyConfig struct {
	HttpProxy *Proxy `json:"httpProxy" yaml:"httpProxy"`
}

type Proxy struct {
	Ip     string `json:"ip" yaml:"ip"`
	Port   string `json:"port" yaml:"port"`
	Enable bool   `json:"enable" yaml:"enable"`
}

func WithHttpProxyConfig(yamlFile string) Option {

	file, err := os.ReadFile(yamlFile)
	if err != nil {
		return func(i *Instance) {}
	}

	c := HttpProxyConfig{}

	yaml.Unmarshal(file, &c)

	return func(i *Instance) {
		i.HttpProxyConfig = c
	}
}

func WithName(name string) Option {
	return func(i *Instance) {
		i.Name = name
	}
}

func WithVersion(version string) Option {
	return func(i *Instance) {
		i.Version = version
	}
}

func (i *Instance) String() string {
	return fmt.Sprintf("Ins[name=%s,version=%s,createdAt=%d,updatedAt=%d,message=%s,HttpProxyConfig=%v]", i.Name, i.Version, i.CreatedAt, i.UpdatedAt, i.Message, i.HttpProxyConfig)
}

func (i *HttpProxyConfig) String() string {
	return fmt.Sprintf("[httpProxy=%v]", i.HttpProxy)
}

func (i *Proxy) String() string {
	return fmt.Sprintf("[ip=%s,port=%s,enabled=%v]", i.Ip, i.Port, i.Enable)
}
