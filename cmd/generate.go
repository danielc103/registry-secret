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
	"github.com/spf13/cobra"
)

var Username string
var Password string
var Email string
var Registry string

type DockerConfig struct {
	Auths DockerInfo `json:"auth"`
	// +optional
	HttpHeaders map[string]string `json:"HttpHeaders, omitempty"`
}

type DockerInfo map[string]DockerInfoValues

type DockerInfoValues struct {
	Username string `json:"username, omitempty"`
	Password string `json:"password, omitempty"`
	Email    string `json:"email, omitempty"`
	Auth     string `json:"auth, omitempty"`
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates .dockerconfigjson file",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&Username, "username", "u", "", "registry username")
	generateCmd.MarkFlagRequired("username")
	generateCmd.Flags().StringVarP(&Password, "password", "p", "", "registry password")
	generateCmd.MarkFlagRequired("password")
	generateCmd.Flags().StringVarP(&Registry, "registry", "r", "", "regsitry name - ex: gitlab.com/myrepo/")
	generateCmd.MarkFlagRequired("registry")
	generateCmd.Flags().StringVarP(&Email, "email", "e", "", "user email address")

}
