/*
Copyright © 2022 xiexianbin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logger

import (
	"fmt"
	"github.com/spf13/viper"
)

func Println(a ...interface{}) {
	fmt.Println(a...)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	fmt.Println()
}

func Debug(a ...interface{}) {
	if viper.GetBool("verbose") {
		Println(a...)
	}
}

func Debugf(format string, a ...interface{}) {
	if viper.GetBool("verbose") {
		Printf(format, a...)
	}
}
