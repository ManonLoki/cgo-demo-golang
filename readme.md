# CGO 交互演示项目

这个项目主要是演示将Go写的代码导出给 C/C++使用，
涉及知识点

* 如何使用及导出C Api
* Go基础类型与C / C++ 基础类型转换
* 复杂数据的交互  Struct
* 整合其他Go包
* 导出为 静态库(.a /.lib)
* 导出为 动态库(.so / .dylib /.dll)

## 如何使用及导出C Api
如果想使用CGO，需要导入 C 包

同时package必须是main，并实现main函数

对于要导出的函数 使用注释形式进行导出 "//export [MethodName]"

举个例子

```go
import "C"

//export Todo
func Todo(){
	// Todo
}

func main(){}
```


## Go类型与C/C++基础类型转换
* Go -> C / C++
* int -> C.int
* float -> C.float
* double -> C.double
* char -> C.char
* string -> *C.char
* boolean -> C.boolean

## Go类型与C/C++ 复杂类型转换 Struct
对于C中存在的数据结构，可以通过对 对C包增加注释的方式进行定义，举个例子
```go
/*
typedef struct User{
  char* name; // 这里是CString
  age  int;
} User;
 */
import "C"

 
func DisplayUser(user C.User){
	// Todo
} 

```

## 整合其他Go包
直接用就完了

## 导出为静态库
```shell
go build -ldflags "-s -w" -buildmode=c-archive -o [LibName] ./{Go File Name} 
# 例子
go build -ldflags "-s -w" -buildmode=c-archive -o cgo_demo.a ./
```

## 导出为动态链接库

```shell
# 命令
go build  -ldflags "-s -w" -buildmode=c-shared -o [LibName] ./{Go File Name}

# 例子
go build  -ldflags "-s -w" -buildmode=c-shared  -o cgo_demo.dylib ./
```

## Mac下对其他sdk的支持
```shell
# 前面添加
CGO_CFLAGS=-mmacosx-version-min=[Sdk Version] CGO_LDFLAGS=-mmacosx-version-min=[Sdk Version] 

#例子
CGO_CFLAGS=-mmacosx-version-min=10.10 CGO_LDFLAGS=-mmacosx-version-min=10.10 go build  -ldflags "-s -w" -buildmode=c-shared  -o cgo_demo.dylib ./
```

## 注释方式指定连接参数
```shell
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#cgo LDFLAGS: -framework Security
```

 