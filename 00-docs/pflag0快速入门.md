# pflag笔记

# 一、`pflag` 是什么？ 为什么不用 `flag`
`pflag` 是 Go 官方`flag`包的增强版，完全兼容 GNU/POSIX 风格参数(`--long`,`-s`,`--key=value`),也是 `Cobra`框架的底层参数解析库。


`pflag`与`flag`的对比

| 特性            | flag（标准库） | pflag |
| ------------- | --------- | ----- |
| `--long` 参数   | ❌         | ✅     |
| `-s` 短参数      | ⚠️（有限）    | ✅     |
| `--key=value` | ❌         | ✅     |
| POSIX/GNU 风格  | ❌         | ✅     |
| Cobra 支持      | ❌         | ✅     |

结论：
- 写 CLI 工具 -> `pflag`
- 写服务器简单参数-> `flag`也行


# 二、安装导入`pflag`
```shell
go get github.com/spf13/pflag
```
```go
import "github.com/spf13/pflag"
```


# 三、基本使用
[快速入门](./samples/pflag_quickstart/pflag_quickstart.go)
