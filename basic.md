[toc]

# 基础知识

> 如果需要引号，必须用双引号

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
int     //有符号整数，默认8位
uint    //无符号整数，默认8位
uintptr //无符号指针，默认8位
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
  > 注：只能在函数体中出现

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

> 类似于 js 里的语法糖？

## 常量

- `const`，无法改变的变量，不允许重复申明

- `api`
  - len(), cap(), unsafe.Sizeof()

特殊常量(iota)

```go
iota
```

> 可以被编辑器修改的常量

- 规则：
  - 出现`const`关键字时，iota 会置为 0
  - 在一个`const`里的时候，每经过一个行，iota 会加 1

```go
package main

import "fmt"

func main(){
  const (
    a = iota  // 0
    b         // 1 iota += 1
    c         // 2 iota += 1
  )
  const d = iota // 0
  fmt.Println(a,b,c,d) // 0, 1, 2, 0
}

```

> 在定义常量组时，如果不提供初始值，则表示将使用上行的表达式。

## 语言运算符

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
  - `&`按位和
  - `｜`按位或
  - `^`按位异或，两个位不相同的时候返回 1；否则返回 0
  - `<<`左移
    - `a<<b`表示`a*2^n`
  - `>>`右移
    - `a>>b`表示`a/(2^n)`
- 赋值运算符
- 其他运算符
  - `&`返回变量存储地址
  - `*`指针变量

## 条件语句

- if
- if else
- select

## 循环语句

- for
- break
- continue
  - 跳出本次循环
- goto
