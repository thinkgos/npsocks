package deployed

// 设置viper默认值回调
var defaultValueFuncs []func()

// RegisterViperDefaultFunc 增加设置viper默认值回调
func RegisterViperDefault(f ...func()) {
	defaultValueFuncs = append(defaultValueFuncs, f...)
}

// ViperInitDefault 运行注册了的初始化viper默认值的所有回调
func ViperInitDefault() {
	for _, f := range defaultValueFuncs {
		f()
	}
}
