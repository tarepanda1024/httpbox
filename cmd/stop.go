// Copyright 2019 HttpBox Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "stop the server",
		Run: func(cmd *cobra.Command, args []string) {
			if !IsPidExist() {
				fmt.Println("app is not running")
				os.Exit(0)
			}
			pid, _ := ioutil.ReadFile(pidFile)
			var command *exec.Cmd
			switch runtime.GOOS {
			case "windows":
				command = exec.Command("taskkill", "/F", "/PID", string(pid))
			case "linux":
				command = exec.Command("kill", string(pid))
			default:
				command = exec.Command("kill", string(pid))
			}

			if err := command.Start(); err != nil {
				fmt.Println("fail to stop server.", err.Error())
			} else {
				fmt.Println("app server has stop.")
			}
			if err := os.Remove(pidFile); err != nil {
				fmt.Println("del pid file failed.")
			}
		},
	}
	rootCmd.AddCommand(stopCmd)
}
