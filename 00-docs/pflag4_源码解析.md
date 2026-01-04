# pflag.Flag 源码
```go
type Flag struct {
	Name                string              // flag名称
	Shorthand           string              // flag短名称
	Usage               string              // Usage/help信息
	Value               Value               // Value,真正保存值的地方
	DefValue            string              // 字符串形式的默认值
	Changed             bool                // 是否被用户显式改变过
	NoOptDefVal         string              // --flag 不带值时使用
	Deprecated          string              // 过期的，描述过期
	Hidden              bool                // 是否在Help中隐藏该flag
	ShorthandDeprecated string              // 短名称是否过期的描述
	Annotations         map[string][]string // Cobra.command bash autoComplete代码使用
}

```

# pflag.FlagSet 源码
- FlagSet 是一次命令行解析的上下文对象
```go
type FlagSet struct {
	Usage func()   // Usage 是在解析标志时发生错误时调用的函数。该字段是一个可以更改为自定义错误处理程序的函数（不是方法）。
	SortFlags bool // helo输出是否排序

	// ParseErrorsAllowlist is used to configure an allowlist of errors
	ParseErrorsAllowlist ParseErrorsAllowlist

	// ParseErrorsAllowlist is used to configure an allowlist of errors.
	//
	// Deprecated: use [FlagSet.ParseErrorsAllowlist] instead. This field will be removed in a future release.
	ParseErrorsWhitelist ParseErrorsAllowlist

	name              string // FlagSet名称
	parsed            bool   // FlagSet是否已解析
	actual            map[NormalizedName]*Flag      // 用户实际设置过的 flag
	orderedActual     []*Flag
	sortedActual      []*Flag
	formal            map[NormalizedName]*Flag      // 所有已定义的 flag的定义表， name -> Flag
	orderedFormal     []*Flag
	sortedFormal      []*Flag
	shorthands        map[byte]*Flag
	args              []string // 非flag参数(Parse后)
	argsLenAtDash     int      // len(args) when a '--' was located when parsing, or -1 if no --
	errorHandling     ErrorHandling
	output            io.Writer // nil means stderr; use Output() accessor
	interspersed      bool      // allow interspersed option/non-option args
	normalizeNameFunc func(f *FlagSet, name string) NormalizedName      // 名称规范化

	addedGoFlagSets []*goflag.FlagSet
}
```



# 一些注意点
- `pflag`只负责1.定义flag,2.解析flag,3.管理flag集合(FlagSet)
- `pflag`不参与子命令，子命令等不是`pflag`的能力，而是Cobra框架在`pflag`之上实现的抽象
- `FlagSet`不适命令，本质是`flag`的命名空间+解析上下文

## Q:Cobra是如何补齐子命令的呢？
A:
```text
Cobra 在 pflag 之上增加了4层能力:

           Cobra Command Tree
                 │
          argv → 路由到 command
                 │
         command.Flags() → pflag.FlagSet
                 │
          pflag 负责解析
```
Cobra Command本质:
```go
type Command struct {
    Use   string
    Run   func(...)
    Flags *pflag.FlagSet
    Children []*Command
}
```
Cobra 先解析命令，再把剩余 argv 交给对应的`FlagSet`

## pflag 与 Cobra 的职责对比
| 能力      | pflag | cobra       |
| ------- | ----- | ----------- |
| flag 定义 | ✅     | ✅（基于 pflag） |
| flag 解析 | ✅     | ✅           |
| FlagSet | ✅     | ✅           |
| 子命令     | ❌     | ✅           |
| 命令树     | ❌     | ✅           |
| help 层级 | ❌     | ✅           |
| 执行流     | ❌     | ✅           |


## 为什么 pflag不内建子命令?(设计哲学)
这是一个刻意的设计选择：

1️⃣ 单一职责

pflag = 参数解析

Cobra = CLI 框架

2️⃣ 低耦合

pflag 可被任意框架复用

Kubernetes / Docker / etcd 都直接用 pflag

3️⃣ 组合优于继承（Go 风格）

子命令是“组合出来的能力”
