// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bketelsen/factor/factor/component"
	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(cwd)
		fmt.Println("dev called")

		dir := filepath.Join(cwd, "components")
		err = processComponents(dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		dir = filepath.Join(cwd, "routes")
		err = processComponents(dir)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func processComponents(base string) error {

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isHtml(info) {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c, _ := component.Parse(f, componentName(path))

			comp := componentName(path)
			gfn := filepath.Join(base, strings.ToLower(comp)+".go")
			_, err = os.Stat(gfn)
			if os.IsNotExist(err) {
				c.Struct = true
			}
			fmt.Printf("visited file: %q\n", path)
			gofile, err := os.Create(goFileName(base, componentName(path)))
			if err != nil {
				return err
			}
			defer gofile.Close()
			c.Transform(gofile)
			c.TransformStyle()
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", base, err)
	}
	return err
}
func isHtml(info os.FileInfo) bool {
	return filepath.Ext(info.Name()) == ".html"
}

func goFileName(base, comp string) string {
	return filepath.Join(base, strings.ToLower(comp)+"_generated.go")
}
func componentName(path string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	base := filepath.Base(path)
	base = strings.Replace(base, filepath.Ext(path), "", -1)
	return strings.Title(reg.ReplaceAllString(base, ""))
}

func init() {
	rootCmd.AddCommand(devCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
