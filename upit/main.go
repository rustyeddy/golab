// Example shows how to work with Upwork API
//
// Licensed under the Upwork's API Terms of Use;
// you may not use this file except in compliance with the Terms.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author::    Maksym Novozhylov (mnovozhilov@upwork.com)
// Copyright:: Copyright 2015(c) Upwork.com
// License::   See LICENSE.txt and TOS - https://developers.upwork.com/api-tos.html
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/upwork/golang-upwork/api"
	"github.com/upwork/golang-upwork/api/routers/auth"
)

// G global variables for this program
var G = struct {
	Client *api.Client
}{
	nil,
}

const cfgFile = "/Users/rusty/.config/upwork.json"

var (
	UpClient *api.ApiClient
	Cfg      *api.Config
	err      error // A catch all
)

func main() {

	// Parse command line arguments and prepare variables
	flag.Parse()

	// Get the upwork client
	if G.Client = getConfig(); G.Client == nil {
		log.Fatal("failed to get upwork client it is nil")
	}

	/*
	 * Figure out who the user is and proceed with authorization
	 *
	 * http.Response and []byte will be return, you can use any
	 * TODO - need to respond properly to various results
	 */
	resp, _ := auth.New(*client).GetUserInfo()
	if resp.StatusCode > 300 {
		log.Fatalf("failed to authorize editor %d", resp.StatusCode)
	}

	cmd := flag.Arg(0)
	if cmd == "" {
		log.Fatalf("require a command to do something, exiting")
	}

	switch cmd {
	case "jobs":

		err = Cmd_Joblist()
		var jobList *JobList
		jobList, err = getJobList()
	}

	if err != nil {
		log.Fatalf("failed running %s error => %v", cmd, err)
	}

	if jobList != nil {
		// params["q"] = "python"
		fmt.Printf("List jobs\n")
		for i, j := range jobList.Jobs {
			fmt.Printf("%3d - %s\n", i, j.Title)
		}
	}
}

func getHttpClient() {
	c := &http.Client{}
	config := api.ReadConfig(cfgFile)
	config.SetCustomHttpClient(c)
	*client = api.Setup(cfg)
}

func getAccessToken() {

	aurl := client.GetAuthorizationUrl("")

	// read verifier
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Visit the authorization url and provide oauth_verifier for ")
	fmt.Println("further authorization")
	fmt.Println(aurl)
	verifier, _ := reader.ReadString('\n')

	// get access token
	token := client.GetAccessToken(verifier)

	cfg.AccessToken = token.Token
	cfg.AccessSecret = token.Secret
	log.Printf("TOKEN: %s -> %s", token)
	err := saveConfig(cfgFile, cfg)
	if err != nil {
		log.Fatalf("failed to save config file %v", err)
	}
}
