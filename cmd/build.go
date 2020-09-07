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
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ds := make([]string, 0)
		r := GetFlagValue(cmd, "registry", "registry.cn-shanghai.aliyuncs.com")
		o := GetFlagValue(cmd, "origin", "qingmuio")
		u := GetFlagValue(cmd, "username", viper.GetString("DOCKER_USERNAME"))
		p := GetFlagValue(cmd, "password", viper.GetString("DOCKER_PASSWORD"))
		a := GetFlagValue(cmd, "application", "")
		v := GetFlagValue(cmd, "application-version", "")
		root := "./dockerfile"

		if a != "" {
			root += "/" + a
			if v != "" {
				root += "/" + v
			}
		}

		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				ds = append(ds, path)
			}
			return err
		})

		for _, filepath := range ds {
			//wg.Add(1)
			func() {
				//defer wg.Add(-1)
				split := strings.Split(filepath, "/")
				imageName := split[len(split)-3]
				version := split[len(split)-2]
				dockerImage := fmt.Sprintf("%v/%v/%v:%v", r, o, imageName, version)
				cli, err := client.NewEnvClient()
				if err != nil {
					log.Panic(err.Error())
				}
				log.Println(dockerImage)
				resp, err := cli.ImageBuild(context.Background(), GetContext(strings.ReplaceAll(filepath, "Dockerfile", "")), types.ImageBuildOptions{
					Tags:   []string{dockerImage},
					Remove: true,
				})
				if err != nil {
					log.Panic(err.Error())
				}
				defer resp.Body.Close()
				termFd, isTerm := term.GetFdInfo(os.Stderr)
				jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stderr, termFd, isTerm, nil)
				authConfig := types.AuthConfig{
					Username: u,
					Password: p,
				}
				encodedJSON, err := json.Marshal(authConfig)
				if err != nil {
					log.Panic(err.Error())
				}
				authStr := base64.URLEncoding.EncodeToString(encodedJSON)
				pushResp, err := cli.ImagePush(context.Background(), dockerImage, types.ImagePushOptions{
					RegistryAuth: authStr,
				})
				if err != nil {
					log.Panic(err.Error())
				}
				defer pushResp.Close()
				jsonmessage.DisplayJSONMessagesStream(pushResp, os.Stderr, termFd, isTerm, nil)
				log.Println(fmt.Sprintf("pushed image %v", dockerImage))
			}()
		}
		//wg.Wait()
	},
}

func GetContext(filePath string) io.Reader {
	// Use homedir.Expand to resolve paths like '~/repos/myrepo'
	filePath, _ = homedir.Expand(filePath)
	ctx, _ := archive.TarWithOptions(filePath, &archive.TarOptions{})
	return ctx
}
func init() {
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().StringP("registry", "r", "registry.cn-shanghai.aliyuncs.com", "docker registry")
	buildCmd.Flags().StringP("origin", "o", "qingmuio", "image origin")
	buildCmd.Flags().StringP("username", "u", "", "docker username")
	buildCmd.Flags().StringP("password", "p", "", "docker password")
	buildCmd.Flags().StringP("application", "a", "", "application name eg. tomcat openjre")
	buildCmd.Flags().StringP("application-version", "v", "", "application version eg. 8.252,9.0.37")
}
