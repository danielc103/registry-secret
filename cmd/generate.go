/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/danielc103/registry-secret/secret"
	"github.com/spf13/cobra"
)

// fileName set as constant
const fileName = ".dockerconfigjson"

// Username flag variable
var Username string

// Password flag variable
var Password string

// Email flag variable
var Email string

// Server flag variable
var Server string

// SetAuth flag bool
var SetAuth bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates .dockerconfigjson file",
	Run: func(cmd *cobra.Command, args []string) {
		j, err := secret.CreateDockerJSON(Username, Password, Email, Server, SetAuth, nil)
		if err != nil {
			println(err)
		}

		secret.WriteToFile(j, fileName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&Username, "username", "u", "", "registry username")
	generateCmd.MarkFlagRequired("username")
	generateCmd.Flags().StringVarP(&Password, "password", "p", "", "registry password")
	generateCmd.MarkFlagRequired("password")
	generateCmd.Flags().StringVarP(&Server, "registry", "r", "", "regsitry name - ex: gitlab.com/myrepo/")
	generateCmd.MarkFlagRequired("registry")
	generateCmd.Flags().StringVarP(&Email, "email", "e", "", "user email address")

	generateCmd.Flags().BoolVarP(&SetAuth, "auth", "a", false, "encrytped auth - defaults to false")

}
