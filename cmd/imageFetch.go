/*
Copyright © 2021 DJ Patterson

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/davidPatters0n/imgurcli/api"
	"github.com/spf13/cobra"
)

// imageFetchCmd represents the imageFetch command
var imageFetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Supply an image ID and get information about about it.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		image := api.Fetch(args[0])
		res, err := json.MarshalIndent(image, "", "  ")
		fmt.Println(string(res), err)
	},
}

func init() {
	imageCmd.AddCommand(imageFetchCmd)
}
