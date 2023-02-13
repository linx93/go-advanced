package init_code

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

const (
	DafaultVsersion = "1.0.0"
)

type Ins struct {
	Name            string          `json:"type"`
	Age             int8            `json:"age,omitempty"`
	HttpProxyConfig HttpProxyConfig `json:"httpProxyConfig"`
	Version         string          `json:"version"`
	UpdatedAt       int64           `json:"updatedAt"`
	CreatedAt       int64           `json:"createdAt"`
}

func NewIns() *Ins {
	return &Ins{
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Version:   DafaultVsersion,
	}
}

func (i *Ins) SetName(name string) *Ins {
	i.Name = name
	return i
}

func (i *Ins) SetAge(age int8) *Ins {
	i.Age = age
	return i
}

func (i *Ins) SetConfig(yamlFile string) *Ins {
	file, err := os.ReadFile(yamlFile)
	if err != nil {
		log.Fatalf("read yaml file error: %v", err)
		return nil
	}
	c := HttpProxyConfig{}
	yaml.Unmarshal(file, &c)

	i.HttpProxyConfig = c
	return i
}

func (i *Ins) String() string {
	return fmt.Sprintf("Ins[name=%s,age=%d,Config=%v,Version=%s,CreatedAt=%d,UpdatedAt=%d]", i.Name, i.Age, i.HttpProxyConfig, i.Version, i.CreatedAt, i.UpdatedAt)
}
