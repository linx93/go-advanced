```

// 空结构体的特点
// 1. 零内存占用
// 2. 地址相同
// 3. 无状态

// 空结构体的使用场景
// 1. 实现Set集合 go本身没有自带的Set,我们可以利用 map 类型来实现一个 Set 集合。由于 map 的 key 具有唯一性，我们可以将元素存储为 key，而 value 没有实际作用，为了节省内存，我们可以使用空结构体作为 value 的值
// 2. 用于通道信号
// 3. 作为方法接收器

```
```
func main() {
	//空结构体的特点
	//1. 零内存占用
	var a int
	var b string
	var c struct{}
	fmt.Println(unsafe.Sizeof(a)) // 4 int默认是4
	fmt.Println(unsafe.Sizeof(b)) // 8 string默认是8
	fmt.Println(unsafe.Sizeof(c)) // 0 struct{}默认就是0

	//2. 地址相同
	var s1 struct{}
	var s2 struct{}
	fmt.Printf("%p\n", &s1) //0xe29540
	fmt.Printf("%p\n", &s2) //0xe29540
	fmt.Println(&s1 == &s2) //true

	//3. 无状态
	//空结构体本身不包含任何字段，不能记录任何的值，因此他不能有状态。
	//所以这使得空结构体在表示无状态得对象或情况时非常有用

	//--------------------------------------------------------------------------------------------------------------------------------

	// 空结构体的使用场景
	// 1. 实现Set集合 go本身没有自带的Set,我们可以利用 map 类型来实现一个 Set 集合。由于 map 的 key 具有唯一性，我们可以将元素存储为 key，而 value 没有实际作用，为了节省内存，我们可以使用空结构体作为 value 的值
	mySet := MySet[string]{}
	element := "i am linx93"
	mySet.Add(element)
	mySet.Add("利用map+struct{}实现Set集合")
	mySet.Add("哈哈")
	contains := mySet.Contains(element)
	fmt.Printf("Set中存在元素 i am linx93:%v\n", contains) //Set中存在元素 i am linx93:true
	mySet.Remove(element)
	fmt.Printf("mySet:%#v", mySet) //mySet:main.MySet[string]{"利用map+struct{}实现Set集合":struct {}{}, "哈哈":struct {}{}}

	// 2. 用于通道信号
	// 3. 作为方法接收器
}

// MySet 1. 实现Set集合 利用map的key唯一，MySet的元素作为map的key,value无意义用空结构体填充
type MySet[E comparable] map[E]struct{}

// Add 添加元素
func (set MySet[E]) Add(element E) {
	//元素放到map的key，value用空结构体填充
	set[element] = struct{}{}
}
```