1.获取结构体字节大小  unsafe.Sizeof(struct{}{})=0
2.如何根据运行环境读取相应环境的配置文件
  在golang的运行环境设置GO_ENV为dev/st/ga读取相应目录下文件
3.linux设置代理地址用户名密码 
  export GO111MODULE=on
  export GOPROXY=https://goproxy.io,direct  //  export GOPROXY=https://goproxy.cn
  export http_proxy="http://root:toor@10.59.79.17:54323"
4.go程序编译没有在go-path的bin目录下发现exe文件
  在当前项目下执行go install即可
5.byte... 数组展开或切片
6.var _ Foo = Dog{}   //var _ Foo = (*Dog)(nil)
  上面用来判断Dog是否实现了Foo, 用作类型断言，如果Dog没有实现Foo，则会报编译错误
7.切片扩容机制：当原切片长度小于1024时，新切片的容量会直接翻倍。而当原切片的容量大于等于1024时，
  会反复地增加25%，直到新容量超过所需要的容量
8.切片扩容时，其会新开一块内存来存储扩容后的切片，相比前后其内存地址已经改变，这在引用中是一个易错点
9.make和new的区别
  new 接受一个参数，这个参数是一种类型，而不是一个值，分配好内存后，返回一个指向该类型内存地址的指针，这个指针指向的内容的值为该类型的零值。
  make 同样用于内存分配，但和new不同，其用于channel，slice和map的分配，而且返回的类型就是这三个类型本身，而不是它们的指针，因为这三种类型本身就是引用类型，
       所以就没必要返回他们的指针了。
10.空指针与野指针 
  空指针: 未被初始化的指针  var p *int 这时如果我们想要对其取值操作 *p, 会报错.
  野指针: 被一片无效的地址空间初始化 var p *int = 0xc00000a0c8
11.栈帧: 用来给函数运行提供内存空间, 取内存于stack 上.当函数调用时, 产生栈帧; 函数调用结束, 释放栈帧.
  数据区保存的是初始化后的数据. 一般 make() 或者 new() 出来的都存储在堆区
  那么栈帧用来存放什么?
  局部变量
  形参
  内存字段描述值
  其中, 形参与局部变量存储地位等同
12.Go中接口的命名约定：接口名以er结尾
13.在Go中，没有隐式类型转换,一般的类型转换可以这么做：int32(i)
14.Golang RPC 之 Thrift
  Thrift是一款高性能、开源的RPC框架，产自Facebook后贡献给了Apache，Thrift囊括了整个RPC的上下游体系，
  自带序列化编译工具，因为Thrift采用的是二进制序列化，并且与gRPC一样使用的都是长连接建立client与server之间的通讯，
  相比于比传统的使用XML，JSON，SOAP等短连接的解决方案性能要快得多。
15.rpcx是Go语言生态圈的Dubbo， 比Dubbo更轻量,不支持跨语言
16.反射机制就是在运行时动态的调用对象的方法和属性
17.判断实例是否实现了某个接口
   	reflect.TypeOf(sStu).Implements(reflect.TypeOf((*IStuService)(nil)).Elem())
18.channel不能close两次,否则会panic。已经关闭channel 不能写入数据。判断channel是否关闭 i, ok := <- ch# hashchain-iris-api
