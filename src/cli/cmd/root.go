/*
Copyright © 2020 Zach Schulze

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
package cmd

import (
  "fmt"
  "log"
  "os"

  aw "github.com/deanishe/awgo"
  "github.com/deanishe/awgo/update"
  "github.com/spf13/cobra"
)

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
  wf *aw.Workflow

  doCheck bool
  query   string

  iconAvailable = &aw.Icon{Value: "update-available.png"}
  repo          = "deanishe/alfred-ssh" // GitHub repo
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "manipulator",
  Short: "Manipulator - A string manipulation utility",
  // Uncomment the following line if your bare application
  // has an action associated with it:
  Run: func(cmd *cobra.Command, args []string) {
    doCheck, err := cmd.Flags().GetBool("check")
    if err != nil {
      wf.FatalError(err)
    }
    // Alternate action: Get available releases from remote.
    if doCheck {
      wf.Configure(aw.TextErrors(true))
      log.Println("Checking for updates...")
      if err := wf.CheckForUpdate(); err != nil {
        wf.FatalError(err)
      }
      return
    }
  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  wf.Run(func() {
    if err := rootCmd.Execute(); err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
  })
}

func init() {
  wf = aw.New(update.GitHub(repo))
  wf.Args()

  // Cobra also supports local flags, which will only run
  // when this action is called directly.

  rootCmd.PersistentFlags().BoolP("check", "c", false, "Check for workflow update")
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
