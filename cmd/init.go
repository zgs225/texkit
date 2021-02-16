/*
Copyright © 2021 yuez <i@yuez.me>

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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zgs225/texkit/generator"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   `init {TITLE}`,
	Short: "Initialize a empty project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		author, _ := cmd.Flags().GetString("author")
		dir, _ := cmd.Flags().GetString("dir")

		if len(title) == 0 {
			log.Panic("title is empty")
		}

		if len(dir) == 0 {
			dir, _ = os.Getwd()
		}

		o := generator.Option{
			Title:     title,
			Author:    author,
			OutputDir: dir,
		}

		if err := generator.Run(&o); err != nil {
			log.Panic(err)
		}

		log.Println("project generated at ", dir)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("author", "YOUR NAME", "Author name for the project(default \"YOUR NAME\")")
	initCmd.Flags().String("dir", "", "Output dir for the project(default \".\")")
}
