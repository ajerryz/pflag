# `pflag`的API
方法速记:
```text
定义: String / StringP / StringVar
解析: Parse / Parsed
访问: GetX / Lookup / Changed
帮助: Usage / PrintDefaults
行为: NoOptDefVal / NormalizeFunc
扩展: Var / AddFlagSet
```

# 一、顶层函数
这些方法作用于 默认全局 FlagSet(`pflag.CommandLine`)
1. 定义falag
   - 基础类型(有默认值，无短参数，返回对应类型的指针)
```go
pflag.Bool(name string, value bool, usage string) *bool
pflag.Int(name string, value int, usage string) *int
pflag.Int64(name string, value int64, usage string) *int64
pflag.Uint(name string, value uint, usage string) *uint
pflag.Float64(name string, value float64, usage string) *float64
pflag.String(name string, value string, usage string) *string
pflag.Duration(name string, value time.Duration, usage string) *time.Duration
```
  - 包含Short name(方法带P)
```go
pflag.StringP(name, shorthand, value, usage)
pflag.BoolP(...)
pflag.IntP(...)
```
2. Var绑定(直接绑定变量)
- 直接绑定变量
```go
pflag.StringVar(&cfg, "config", "", "config file")
pflag.IntVarP(&port, "port", "p", 8080, "port")

```
- 直接绑定变量(带short name),也是对应方法后加P
3. Slice/Map 类型
- slice
```go
pflag.StringSlice(name, []string{}, usage)
pflag.IntSlice(...)
pflag.BoolSlice(...)

pflag.StringArray(name, []string{}, usage) // 不 split
```
- map
```go
pflag.StringToString(name, map[string]string{}, usage)
pflag.StringToInt(...)
```
4. 解析与访问
```go
pflag.Parse()
pflag.Args()          // 非 flag 参数
pflag.NArg()          // Args 数量
pflag.Arg(i)          // 第 i 个
```
5. Get 系列(需要 Parsed后都去，若flag不存在 返回error)
```go
pflag.GetString("name")
pflag.GetInt("port")
pflag.GetBool("verbose")

```

# 二、FlagSet方法(核心与进阶)
`FlagSet`是Pflag的灵魂
```go
// 创建新的FlagSet
fs := pflag.NewFlagSet("app",pflag.ExitOnError)
```


6. FlagSet定义方法与顶层一一对应
```go
fs.String(...)
fs.StringP(...)
fs.StringVar(...)
fs.StringVarP(...)
```
- Cobra 框架内部完全基于`FlagSet`

7. FlagSet控制解析
```go
fs.Parse(os.Args[1:])
fs.Parsed() bool
fs.Set("port", "8080")   // 直接赋值（字符串）
```

8. Flag查找 & 元信息
```go
fs.Lookup("port")        // *Flag
fs.ShorthandLookup("p") // *Flag
```

Flag结构体关键字段:
```go
type Flag struct {
    Name        string
    Shorthand   string
    Usage       string
    Value       Value
    DefValue    string
    Changed     bool
    Hidden      bool
}
```

9. Changed判断(非常重要)
```go
fs.Changed("config")
```
🖊️ : CLI > config > default


# 三、Usage/Help相关
10. Help自动生成
```go
fs.PrintDefaults()
fs.PrintDefaults()
```


11. 自定义Usage
```go
fs.Usage = func() {
    fmt.Println("custom usage")
}
```

12. 隐藏/弃用
```go
fs.MarkHidden("debug")
fs.MarkDeprecated("old", "use --new")
fs.MarkShorthandDeprecated("d", "use --debug")
```



# 四、高级 Flag 行为控制
13. NoOptDefVal(可省略值)
```go
fs.Lookup("log").NoOptDefVal = "info"
```
```shell
--log     # 等价 --log=info
```

14. NormalizeFunc(参数名规范化)
```go
fs.SetNormalizeFunc(func(fs *pflag.FlagSet, name string) pflag.NormalizedName {
    return pflag.NormalizedName(strings.ReplaceAll(name, "_", "-"))
})
```
```shell
--log_level == --log-level
```

15. 自定义Value(扩展类型)
```go
type Level string
func (l *Level) Set(s string) error { ... }
func (l *Level) String() string { ... }
func (l *Level) Type() string { return "level" }

fs.Var(&level, "level", "log level")
```


# 五、错误处理模式
16. ErrorHanding
```go
pflag.NewFlagSet("app",
    pflag.ContinueOnError |
    pflag.ExitOnError |
    pflag.PanicOnError,
)
```

# 六、辅助 & 不常用方法
17. 访问所有flags
```go
fs.Visit(func(f *pflag.Flag) {})
fs.VisitAll(func(f *pflag.Flag) {})
```
区别:
- `Visit`: 只访问 Changed
- `VisitAll`: 访问全部

18. 排序控制
```go
fs.SortFlags = false
```

19. 添加外部FlagSet
```go
fs.AddFlagSet(other)
```
- Cobra框架子命令继承核心


20. 初始化 & Reset
```go
fs.Init("name", pflag.ExitOnError)
fs = pflag.NewFlagSet(...)
```


# 七、顶层变量 & 常量
```go
pflag.CommandLine   // 默认 FlagSet
pflag.ErrHelp       // help 错误
```
