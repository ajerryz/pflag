# `pflag` 功能特性

# 支持的类型

- 基本类型
    - `bool`
    - `int` / `int8` / `int16` / `int32` / `int64`
    - `uint*`
    - `float` / `flat64`
    - `string`
    - `duration` 例:`pflag.Duration("timeout",time.Second,"timeout")`
- 切片类型(标准flag库不支持)
    - `StringSlice`
        - 例：`pflag.StringSlice("hosts",[]string{},"host list")`
        - flag参数: `--hosts=a,b,c` 或 `--hosts a -- hosts b`
    - `IntSlice`
    - `BoolSlice`
    - `StringArray`(不split)
- Map类型
    - 例:`pflag.StringToString("labels",nil, "key=value pairs")`
    - flag参数:`--labels a= 1,b=2`

# 类型的默认值与缺省行为

bool

```go
pflag.Bool("verbose", false, "verbose mode")

// --verbose  # true
```

可省略值(`NoOptDefVal`)

```go
pflag.Lookup("log").NoOptDefVal = "info"
```

```shell
--log # 等价 --log=info
```

# 参数作用域与分组

1. `FlagSet`核心能力

```go
fs := pflag.NewFlasgSet("sub", pflag.ExitOnError)
```

用途:

- 子命令
- 模块化参数
- CObra command 内部实现


2. 全局/本地 flag(结合 Cobra)

- `PersistentFlags()`: 父命令可继承
- `Flags()`: 仅当前命令

# 参数绑定能力

1. 直接绑定变量

```go
var port int
pflag.IntVar(&port,"port", 8080, "listen port")
```

2. 与 Viper 集成(强项)
```go
viper.BindPflag("server.port", flag.Lookup("port"))
```
支持:
- flag
- env
- config file
- 默认值


# 参数解析与访问
1. 解析
```go
pflag.Parse()
```
2. 获取值
```go
pflag.String("name","","")
pflag.CommandLine.GetString("name")
```
3. 判断是否显示设置值
```go
pflag.CommandLine.Changed("config")
```


# 参数校验与自定义类型
1. 自定义Value
```go
type IPValue net.IP
func(i *IPValue) Set(s string) error {...}
```
2. 参数范围校验(结合 Parse后逻辑)
```go
if port < 1 || port > 65535 {
    return errors.New("invalid port")
}
```

# 帮助信息与使用信息(Usage)
1. 自动生成 help
```shell
--help
-h
```
2. 自定义Usage
```go
pflag.Usage = func() {
    fmt.Println("custom helo")
}
```
3. 隐藏参数(高级)
```go
pflag.CommandLine.MarkHidden("name")
```
4. 标记过期
```go
pflag.CommandLine.MarkDeprecated("old","use --new")
```

# 错误处理模式
```go
pflag.NewFlagSet("app",
    pflag.ContineOnError | pflag.ExitOnError | pflag.PanicOnError
    )
```


# 与标准`flag`的差异
| 特性          | flag | pflag |
| ----------- | ---- | ----- |
| `--long`    | ❌    | ✅     |
| `-abc`      | ❌    | ✅     |
| slice/map   | ❌    | ✅     |
| Changed     | ❌    | ✅     |
| NoOptDefVal | ❌    | ✅     |
| Cobra 集成    | ❌    | ✅     |
| POSIX/GNU   | ❌    | ✅     |
