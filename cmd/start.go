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
	"HttpBox/router"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

const pidFile = "./pid.lock"

func IsPidExist() bool {
	_, err := os.Stat(pidFile) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func init() {
	var daemon bool
	var port int
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start App Server",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				if IsPidExist() {
					fmt.Println("HttpBox is already running...")
					os.Exit(0)
				}
				var command *exec.Cmd
				switch runtime.GOOS {
				case "windows":
					command = exec.Command("HttpBox.exe", "start")
				case "linux":
					command = exec.Command("./HttpBox", "start")
				default:
					command = exec.Command("./HttpBox", "start")
				}

				if err := command.Start(); err != nil {
					fmt.Println("start failed.", err.Error())
					os.Exit(0)
				}
				fmt.Printf("app start, [PID] %d running...\n", command.Process.Pid)
				pid := []byte(fmt.Sprintf("%d", command.Process.Pid))
				if err := ioutil.WriteFile(pidFile, pid, 0666); err != nil {
					fmt.Println("write lock failed.", err.Error())
				}
				daemon = false
				os.Exit(0)
			} else {
				router.StartServer(port)
			}
		},
	}
	startCmd.Flags().IntVarP(&port, "port", "p", 8080, "listen on ?")
	startCmd.Flags().BoolVarP(&daemon, "daemon", "d", false, "is daemon?")

	rootCmd.AddCommand(startCmd)
}
