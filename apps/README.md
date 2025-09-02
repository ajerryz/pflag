# spf13/pflag了解

# pflag 是什么?
pflag 是 Go 的 flag 包的替代品，实现了 POSIX/GNU 风格的 --flags。

```shell
# 安装
go get github.com/spf13/pflag

# 测试
go test github.com/spf13/pflag
```

# 主要特性
1. 支持更多类型：除了标准库的基本类型，还支持 time.Duration、ip、ip mask、ip net、count 等
2. 更灵活的参数定义：可以定义短选项（-v）和长选项（--verbose）
3. 支持参数别名：可以为同一个参数定义多个名称
4. 支持嵌套子命令：适合复杂命令行工具（如 kubectl）
5. 兼容标准库 flag：可以平滑迁移

**用法**:
pflag 是 Go 原生 flag 包的直接替代品。如果您以“flag”名称导入 pflag，则所有代码均可继续正常运行。
```go
import flag "github.com/spf13/pflag"
```


# 基本使用流程
1. 导入 pflag 包
2. 定义命令行参数
3. 解析命令行参数
4. 使用解析后的参数值

[快入开始](./quickstart/main.go)



