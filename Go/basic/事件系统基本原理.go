package main

// 实例化一个通过字符串映射函数切片的map
var eventByName = make(map[string][]func(interface{}))

// 注册事件， 提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})){

	// 通过名字查找事件列表
	list := eventByName[name]

	// 在列表切片中添加函数
	list = append(list,callback)

	// 将修改的事件列表切片保存回去
	eventByName[name] = list
}

// 调用事件
func CallEvent(name string, param interface{}){

	// 通过名字找到事件列表
	list := eventByName[name]

	// 遍历这个事件的所有回调
	for _, callback := range list{

		// 传入参数调用回调
		callback(param)
	}
}

func main() {


	//fmt.Println(reflect.TypeOf(l))
}


/*

事件系统基本原理

事件系统可以将事件派发者与事件处理者解耦。
例如，网络底层可以生成各种事件，在网络连接上后，网络底层只需将事件派发出去，
而不需要关心到底哪些代码来响应连接上的逻辑。或者再比如，你注册、关注或者订阅某“大V”的社交消息后，
“大V”发生的任何事件都会通知你，但他并不用了解粉丝们是如何为她喝彩或者疯狂的。如下图所示为事件系统基本原理图


第 4 行，创建一个 map 实例，这个 map 通过事件名（string）关联回调列表（[]func(interface{}），同一个事件名称可能存在多个事件回调，因此使用回调列表保存。回调的函数声明为 func(interface{})。

第 7 行，提供给外部的通过事件名注册响应函数的入口。

第 10 行，eventByName 通过事件名（name）进行查询，返回回调列表（[]func(interface{}）。

第 13 行，为同一个事件名称在已经注册的事件回调的列表中再添加一个回调函数。

第 16 行，将修改后的函数列表设置到 map 的对应事件名中。


事件调用

事件调用方和注册方是事件处理中完全不同的两个角色。事件调用方是事发现场，负责将事件和事件发生的参数通过事件系统派发出去，而不关心事件到底由谁处理；事件注册方通过事件系统注册应该响应哪些事件及如何使用回调函数处理这些事件。事件调用的详细实现请参考上面代码的 CallEvent() 函数。

代码说明如下：
第 20 行，调用事件的入口，提供事件名称 name 和参数 param。事件的参数表示描述事件具体的细节，例如门打开的事件触发时，参数可以传入谁进来了。
第 23 行，通过注册事件回调的 eventByName 和事件名字查询处理函数列表 list。
第 26 行，遍历这个事件列表，如果没有找到对应的事件，list 将是一个空切片。
第 29 行，将每个函数回调传入事件参数并调用，就会触发事件实现方的逻辑处理。
*/
