# 代码规范

- [代码规范](#代码规范)
  - [代码风格](#代码风格)
    - [代码格式](#代码格式)
    - [声明、初始化和定义](#声明初始化和定义)
    - [错误处理](#错误处理)
    - [panic 处理](#panic-处理)
    - [单元测试](#单元测试)
    - [类型断言](#类型断言)
  - [命名规范](#命名规范)
    - [包命名](#包命名)
    - [函数命名](#函数命名)
    - [文件命名](#文件命名)
    - [结构体命名](#结构体命名)
    - [接口命名](#接口命名)
    - [变量命名](#变量命名)
    - [常量命名](#常量命名)
    - [Error 命名](#error-命名)
  - [注释规范](#注释规范)
  - [数据类型使用规范](#数据类型使用规范)
  - [控制结构](#控制结构)
  - [函数规范](#函数规范)

## 代码风格

### 代码格式

- 所有的代码都必须用 [gofmt](https://pkg.go.dev/cmd/gofmt) 进行格式化。
  ```shell
  gofmt -s -w .
  ```
- 运算符和操作数之间要留空格。
  ```go
  a := b + c
  ```
- 避免过长的行，一行代码不超过 120 个字符，超过部分，请采用合适的换行方式换行。但也有些例外场景，例如 import 行、工具自动生成的代码、带 tag 的 struct 字段。
- 文件长度不能超过 800 行。
- 函数长度不能超过 90 行。
- 代码都必须用 goimports 进行格式化（建议将代码 Go 代码编辑器设置为：保存时运行 goimports）。
- 不要使用相对路径引入包，例如 import …/util/net 。
- 包名称与导入路径的最后一个目录名不匹配时，或者多个相同包名冲突时，则必须使用导入别名。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go
  import (
    "github.com/dgrijalva/jwt-go/v4"
  )
  ```

    </td><td>

  ```go
  import (
    jwt "github.com/dgrijalva/jwt-go/v4"
  )
  ```

    </td></tr>
    </tbody>
  </table>

- 对导入的包进行分组，匿名包的引用使用一个新的分组，并对匿名包引用进行说明。

  ```go
  import (
    // go 标准包
    "fmt"

    // 第三方包
      "github.com/jinzhu/gorm"
      "github.com/spf13/cobra"
      "github.com/spf13/viper"

    // 匿名包单独分组，并对匿名包引用进行说明
      // import mysql driver
      _ "github.com/jinzhu/gorm/dialects/mysql"

    // 内部包
      v1 "github.com/marmotedu/api/apiserver/v1"
      metav1 "github.com/marmotedu/apimachinery/pkg/meta/v1"
      "github.com/marmotedu/iam/pkg/cli/genericclioptions"
  )
  ```

- 使用 Go Modules 作为依赖管理的项目时，必须提交 go.sum 文件。

### 声明、初始化和定义

- 当函数中需要使用到多个变量时，可以在函数开始处使用 var 声明。在函数外部声明必须使用 var ，不要采用 := ，容易踩到变量的作用域的问题。
  ```go
  var (
    Width  int
    Height int
  )
  ```
- 在初始化结构引用时，请使用 &T{}代替 new(T)，以使其与结构体初始化一致。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go
    a := new(T)
    a.Name = "Tom"
  ```

    </td><td>

  ```go
    a := &T{Name:"tom"}

  ```

    </td></tr>
    </tbody>
  </table>

- struct 声明和初始化格式采用多行，定义如下。

  ```go
  type User struct{
    Username  string
    Email     string
  }

  user := User{
    Username: "colin",
    Email: "colin404@foxmail.com",
  }
  ```

- 相似的声明放在一组，同样适用于常量、变量和类型声明。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  const a = "strategy_a"
  const b = "strategy_b"

  ```

    </td><td>

  ```go
  const (
    a = "strategy_a"
    b = "strategy_b"
  )
  ```

    </td></tr>
    </tbody>
  </table>

- 尽可能指定容器容量，以便为容器预先分配内存，例如：
  ```go
  intStrMap := make(map[int]string, 4)
  strList := make([]string, 0, 4)
  ```
- 在顶层，使用标准 var 关键字时，切勿指定类型，除非它与表达式的类型不同。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  var _s string = F()
  func F() string { return "A" }


  ```

    </td><td>

  ```go

  // 由于 F 已经明确了返回一个字符串类型，因此我们没有必要显式指定_s 的类型
  var _s = F()
  func F() string { return "A" }

  ```

    </td></tr>
    </tbody>
  </table>

- 对于未导出的顶层常量和变量，使用 \_ 作为前缀。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  const (
    defaultHost = "127.0.0.1"
    defaultPort = 8080
  )

  ```

    </td><td>

  ```go

  const (
    _defaultHost = "127.0.0.1"
    _defaultPort = 8080
  )

  ```

    </td></tr>
    </tbody>
  </table>

- 嵌入式类型（例如 mutex）应位于结构体内的字段列表的顶部，并且必须有一个空行将嵌入式字段与常规字段分隔开。
   <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  type Client struct {
    version int
    http.Client
  }


  ```

    </td><td>

  ```go

  type Client struct {
    http.Client

    version int
  }

  ```

    </td></tr>
    </tbody>
  </table>

### 错误处理

- error 作为函数的值返回，必须对 error 进行处理，或将返回值赋值显式忽略。对于 defer xx.Close() 可以不用显式处理。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  func process() error {
    //...
  }

  func main() {
    process()
  }

  ```

    </td><td>

  ```go

  func process() error {
    //...
  }

  func main() {
    _ = process()
  }

  ```

    </td></tr>
    </tbody>
  </table>

- 函数有多个返回且包含 error 时，error 必须是最后一个参数。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  func process() (error, string, int) {
    //...
  }

  ```

    </td><td>

  ```go

  func process() (string, int, error) {
    //...
  }

  ```

    </td></tr>
    </tbody>
  </table>

- 降低错误处理的[圈复杂度](https://baike.baidu.com/item/%E5%9C%88%E5%A4%8D%E6%9D%82%E5%BA%A6/828737)。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  func main()  {
    //...
    if err != nil {
      // error code
    } else {
      // normal code
    }
  }

  ```

    </td><td>

  ```go

  func main() {
    //...
    if err != nil {
      // error code
    }
    // normal code
  }


  ```

    </td></tr>
    </tbody>
  </table>

- 如果需要在 if 之外使用函数调用的结果时，应单独判断。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  func main()  {
    //...
    if result, err := check(); err != nil {
      // error handling
      // result handling
    }
  }

  ```

    </td><td>

  ```go

  func main() {
    //...
    result, err := check()
    if err != nil {
      // error handling
    }
    // result handling
  }

  ```

    </td></tr>
    </tbody>
  </table>

- 错误处理应单独判断，不与其他逻辑组合判断。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  func main()  {
    //...
    if result, err := check();
    if result != nil || err != nil {
      // error handling
    }


  }



  ```

    </td><td>

  ```go

  func main() {
    //...
    result, err := check()
    if err != nil {
      // error handing
    }
    if result != nil {
      // result handling
      return errors.New("invalid value")
    }
  }

  ```

    </td></tr>
    </tbody>
  </table>

- 错误描述建议

  - 告诉用户他们可以做什么，而不是告诉他们不能做什么。
  - 当声明一个需求时，用 must 而不是 should。例如，must be greater than 0、must match regex ‘[a-z]+’。
  - 当声明一个格式不对时，用 must not。例如，must not contain。
  - 当声明一个动作时用 may not。例如，may not be specified when otherField is empty、only name may be specified。
  - 引用文字字符串值时，请在单引号中指示文字。例如，ust not contain ‘…’。
  - 当引用另一个字段名称时，请在反引号中指定该名称。例如，must be greater than request。
  - 指定不等时，请使用单词而不是符号。例如，must be less than 256、must be greater than or equal to 0 (不要用 larger than、bigger than、more than、higher than)。
  - 指定数字范围时，请尽可能使用包含范围。
  - Go 1.13 以上，error 生成方式为 fmt.Errorf("module xxx: %w", err)。
  - 错误描述用小写字母开头，结尾不要加标点符号。
      <table>
      <thead>
        <tr><th>Bad</th><th>Good</th></tr>
      </thead>
      <tbody>
      <tr><td>

    ```go

    errors.New("Redis connection failed")
    errors.New("redis connection failed.")

    ```

      </td><td>

    ```go

    errors.New("redis connection failed")


    ```

      </td></tr>
      </tbody>
    </table>

### panic 处理

- 在业务逻辑处理中禁止使用 panic。
- 在 main 包中，只有当程序完全不可运行时使用 panic，例如无法打开文件、无法连接数据库导致程序无法正常运行。
- 在 main 包中，使用 log.Fatal 来记录错误，这样就可以由 log 来结束程序，或者将 panic 抛出的异常记录到日志文件中，方便排查问题。
- 可导出的接口一定不能有 panic。
- 包内应采用 error 而不是 panic 来传递错误。

### 单元测试

- 单元测试文件名命名规范为 example_test.go。
- 每个重要的可导出函数都要编写测试用例。
- 因为单元测试文件内的函数都是不对外的，所以可导出的结构体、函数等可以不带注释。
- 如果存在 func (b \*Bar) Foo ，单测函数可以为 func TestBar_Foo。

### 类型断言

- type assertion 的单个返回值针对不正确的类型将产生 panic。应当使用 “comma ok”的惯用法。
  <table>
    <thead>
      <tr><th>Bad</th><th>Good</th></tr>
    </thead>
    <tbody>
    <tr><td>

  ```go

  t := n.(int)




  ```

    </td><td>

  ```go

  t, ok := n.(int)
  if !ok {
    // error handling
  }

  ```

    </td></tr>
    </tbody>
  </table>

## 命名规范

### 包命名

- 包名必须和目录名一致，尽量采取有意义、简短的包名，不要和标准库冲突。
- 包名全部小写，没有大写或下划线，使用多级目录来划分层级。
- 项目名可以通过中划线来连接多个单词。
- 包名以及包所在的目录名，不要使用复数，例如，是 net/url，而不是 net/urls。
- 不要用 common、util、shared 或者 lib 这类宽泛的、无意义的包名。
- 包名要简单明了，例如 net、time、log。

### 函数命名

- 函数名采用驼峰式，首字母根据访问控制决定使用大写或小写，例如：MixedCaps 或者 mixedCaps。
- 代码生成工具自动生成的代码 (如 xxxx.pb.go) 和为了对相关测试用例进行分组，而采用的下划线 (如 TestMyFunction_WhatIsBeingTested) 排除此规则。

### 文件命名

- 文件名要简短有意义。
- 文件名应小写，并使用下划线分割单词。

### 结构体命名

- 采用驼峰命名方式，首字母根据访问控制决定使用大写或小写，例如 MixedCaps 或者 mixedCaps。
- 结构体名不应该是动词，应该是名词，比如 Node、NodeSpec。
- 避免使用 Data、Info 这类无意义的结构体名。
- 结构体的声明和初始化应采用多行，例如：

  ```go

    type User struct {
      Name string
      Email string
    }
    // 多行初始化
    u := User{
      UserName: "colin",
      Email: "colin404@foxmail.com",
    }

  ```

### 接口命名

- 接口命名的规则，基本和结构体命名规则保持一致：

  - 单个函数的接口名以 “er"”作为后缀（例如 Reader，Writer），有时候可能导致蹩脚的英文，但是没关系。
  - 两个函数的接口名以两个函数名命名，例如 ReadWriter。
  - 三个以上函数的接口名，类似于结构体名。

  ```go

    type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
    }

    type ReadWriter interface {
      Reader
      Writer
    }

  ```

### 变量命名

- 变量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。
- 在相对简单（对象数量少、针对性强）的环境中，可以将一些名称由完整单词简写为单个字母，例如：
  - user 可以简写为 u；
  - userID 可以简写 uid。
- 特有名词时，需要遵循以下规则：

  - 如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient。
  - 其他情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID。

  ```go

  // 一些常见的特有名词
  var LintGonicMapper = GonicMapper{
      "API":   true,
      "ASCII": true,
      "CPU":   true,
      "CSS":   true,
      "DNS":   true,
      "EOF":   true,
      "GUID":  true,
      "HTML":  true,
      "HTTP":  true,
      "HTTPS": true,
      "ID":    true,
      "IP":    true,
      "JSON":  true,
      "LHS":   true,
      "QPS":   true,
      "RAM":   true,
      "RHS":   true,
      "RPC":   true,
      "SLA":   true,
      "SMTP":  true,
      "SSH":   true,
      "TLS":   true,
      "TTL":   true,
      "UI":    true,
      "UID":   true,
      "UUID":  true,
      "URI":   true,
      "URL":   true,
      "UTF8":  true,
      "VM":    true,
      "XML":   true,
      "XSRF":  true,
      "XSS":   true,
  }
  ```

- 若变量类型为 bool 类型，则名称应以 Has，Is，Can 或 Allow 开头，例如：
  ```go
  var hasConflict bool
  var isExist bool
  var canManage bool
  var allowGitHook bool
  ```
- 局部变量应当尽可能短小，比如使用 buf 指代 buffer，使用 i 指代 index。
- 代码生成工具自动生成的代码可排除此规则 (如 xxx.pb.go 里面的 Id)。

### 常量命名

- 常量名必须遵循驼峰式，首字母根据访问控制决定使用大写或小写。
- 如果是枚举类型的常量，需要先创建相应类型：

  ```go

    type Code int


    const (
        ErrUnknown Code = iota
        ErrFatal
    )

  ```

### Error 命名

- Error 类型应该写成 FooError 的形式。

  ```go

  type ExitError struct {
    // ....
  }

  ```

- Error 变量写成 ErrFoo 的形式。

  ```go

  var ErrFormat = errors.New("unknown format")

  ```

## 注释规范

## 数据类型使用规范

## 控制结构

## 函数规范
