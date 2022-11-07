package adaptor

//适配器模式是作为两个不兼容的接口之间的桥梁
//适配器模式一般包含三种角色：
//目标角色（Target）：也就是我们期望的接口；
//源角色（Adaptee）：存在于系统中，内容满足客户需求（需转换），但接口不匹配的接口实例
//适配器（Adapter）：将源角色（Adaptee）转化为目标角色（Target）的类实例

// Adaptee 被适配的接口
type Adaptee interface {
	SpecificRequest() string
}

// adapteeImpl 被适配的目标类
type adapteeImpl struct{}

// SpecificRequest 被适配的目标类方法
func (a *adapteeImpl) SpecificRequest() string {
	return "adaptee method impl content"
}

// NewAdaptee 构建被适配目标类
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// adapter 将adaptee -> target的适配器
type adapter struct {
	adaptee Adaptee
}

func (a *adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

// Target 适配的目标接口
type Target interface {
	Request() string
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}
