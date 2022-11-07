package chain

import (
	"log"
	"strings"
)

//职责链模式主要包含以下角色
//抽象处理者（Handler）角色：定义一个处理请求的接口，包含抽象处理方法和一个后继连接。
//具体处理者（Concrete Handler）角色：实现抽象处理者的处理方法，判断能否处理本次请求，如果可以处理请求则处理，否则将该请求转给它的后继者。
//客户类（Client）角色：创建处理链，并向链头的具体处理者对象提交请求，它不关心处理细节和请求的传递过程。

type Handler interface {
	Handle(content string)
	Next(handler Handler)
}

type AdHandler struct {
	handler Handler
}

func (ad *AdHandler) Handle(content string) {
	newContent := strings.ReplaceAll(content, "广告", "**")
	log.Printf("newContent:%v\n", newContent)

	if ad.handler != nil {
		ad.handler.Handle(newContent)
	}
}

func (ad *AdHandler) Next(handler Handler) {
	if handler != nil {
		ad.handler = handler
	}
}

// YellowHandler 涉黄过滤
type YellowHandler struct {
	handler Handler
}

func (yellow *YellowHandler) Handle(content string) {
	newContent := strings.ReplaceAll(content, "涉黄", "**")

	log.Printf("newContent:%v\n", newContent)

	if yellow.handler != nil {
		yellow.handler.Handle(newContent)
	}

}

func (yellow *YellowHandler) Next(handler Handler) {
	if handler != nil {
		yellow.handler = handler
	}
}

// SensitiveHandler 敏感词过滤
type SensitiveHandler struct {
	handler Handler
}

func (sensitive *SensitiveHandler) Handle(content string) {
	newContent := strings.ReplaceAll(content, "敏感词", "**")

	log.Printf("newContent:%v\n", newContent)

	if sensitive.handler != nil {
		sensitive.handler.Handle(newContent)
	}

}

func (sensitive *SensitiveHandler) Next(handler Handler) {
	if handler != nil {
		sensitive.handler = handler
	}
}

type Client struct {
}
