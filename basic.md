[toc]
# 基础知识

## go 语言基础组成

### 包声明
```go
package main
```
### 引入包
```go
import "fmt"
```

### 函数
```go
func main(){

}
```

### 变量
```go
a
```
### 语句 表达式
```go
// 第一种写法
var a string = "hello world"

// 第二种写法
a := "hello world"
```
### 注释
```go
// 
/**/
两种形式
```


> 标识符大写开头时，能被外部包使用；标识符小写时，只能被内部使用

## go 语言类型

### 布尔型
```go
bool
```
### 数字类型
```go
int
```
### 字符串类型
```go 
string
```
### 派生类型

- 指针
- 数组
- 结构化
- Channel
- 函数
- 切片
- 接口
- map

## 变量

```go
var a int
var a,b int = 1,2
```

> 用着好别扭

### 变量声明

- 指定变量类型，如果没有初始化，默认为 0 值

```go
/*
  int 0
  boolean false
  string ""
*/

```

- 根据值自动判断类型
- `:=`表示声明并初始化变量，重复声明会报错

```go
/*
  会报错
*/
var a string
string := 1


/*
  正常
*/
var a string = 1
/*
  或者
*/
string := 1
```
>类似于js里的语法糖？

## 常量
```go
var => const
```

特殊常量(iota)
```go
iota
```
>可以被编辑器修改的常量