package basic

import (
	"fmt"
	"net/url"
	"testing"
)

func Test(t *testing.T) {
	//`license` like '%贵V22458%
	str := "%60license%60%20like%20'%25%E8%B4%B5V22458%25"
	fmt.Printf("str=%s\n", str)

	//url解码
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		return
	}
	fmt.Printf("unescape=%s\n", unescape)

	//url编码
	escape := url.QueryEscape(unescape)
	fmt.Printf("escape=%s\n", escape)

	//输出
	//str=%60license%60%20like%20'%25%E8%B4%B5V22458%25
	//unescape=`license` like '%贵V22458%
	//escape=%60license%60+like+%27%25%E8%B4%B5V22458%25  这里like的前后多了加号，但是在下面解码出来还是没问题的

	unescape, err = url.QueryUnescape(escape)

	fmt.Printf("unescape=%s\n", unescape)

}
