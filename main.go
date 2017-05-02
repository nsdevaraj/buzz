/*
 * Buggy, (C) 2016,2017 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

var token = ""

// TomlConfig - holds all the repo names
type tomlConfig struct {
	RepoNames []string `toml:"repoNames"`
}

var config tomlConfig

func main() {

	token = os.Getenv("GIT_TOKEN")
	fmt.Println(token)
	if token == "" {
		fmt.Println("Git Token is not set")
		os.Exit(101)

	}
	if _, err := toml.DecodeFile("repo.toml", &config); err != nil {
		exitOnErr(err)
	}

	auth()
	http.HandleFunc("/getIssues", getIssues)
	http.HandleFunc("/getPRs", getPRs)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.ListenAndServe(":7000", nil)

}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func auth() {
	postURL := `https://api.github.com/?access_token=` + token
	_, err := http.Get(postURL)
	exitOnErr(err)

}