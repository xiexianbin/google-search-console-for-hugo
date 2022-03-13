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

package utils

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strings"
)

func ReadFromCmd(tips string) (string, error) {
	fmt.Printf("%s", tips)
	reader := bufio.NewReader(os.Stdin)
	cmdStr, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	} else {
		cmdStr = strings.TrimSuffix(cmdStr, "\n")
		return cmdStr, nil
	}
}

func EId() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b8 := fmt.Sprintf("%x%x",
		b[4:6], b[6:8])
	return b8
}