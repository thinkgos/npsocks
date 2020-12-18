// Copyright [2020] [thinkgos] thinkgo@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
