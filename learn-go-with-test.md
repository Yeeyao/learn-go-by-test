# content

- 注意，因为GFW的问题，设置了GOPROXY

- 函数名称以小写字母开头。在 Go 中，公共函数以大写字母开始，私有函数以小写字母开头。

- 关注重点内容以及自己容易出错的部分

## hello world

- 将你「领域」内的代码和外部世界（会引起副作用）分离开会更好。

- 注意，你不必在多个测试框架之间进行选择，然后理解如何安装它们。你需要的一切都内建在语言中，语法与你将要编写的其余代码相同。

## 编写测试

- 测试文件程序需要在一个名为 xxx_test.go 中编写

- 测试函数命名必须以 Test 开始

- 测试函数只接受一个参数 t *testing.T

- 类型为 *testing.T 的变量 t 是你在测试框架中的 "hook"（钩子），所以当你想让测试失败时可以执行 t.Fail() 之类的操作。

- TDD 测试驱动开发

- 测试库中的工具 -- 子测试，针对不同场景进行子测试

    - 有多个测试函数，其中调用同一个子测试函数进行测试，多个测试函数负责处理不同的测试

    - t.Helper()

- 示例 存在于包xxx_test.go的一个函数

- 先使用最少的代码来让失败的测试先跑起来

## 规律

- 编写一个测试

- 让编译通过

- 运行测试，查看失败原因并检查错误消息

- 编写足够的代码来通过测试

- 重构，支持测试的安全性以确保我们拥有易于使用的精心制作的代码

- 如果不写测试，你提交的时候通过运行软件来手动检查你的代码，这会打破你的流畅状态，而且你任何时候都无法将自己从这种状态中拯救出来，尤其是从长远来看。

## 基准测试

- 基准测试默认是顺序运行的

- ns/op 表示执行一次该基准函数需要多少时间
