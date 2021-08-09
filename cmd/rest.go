/*
Copyright © 2021 Coding4u <contato@coding4u.com.br>

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

	"github.com/c-4u/employee-service/application/grpc/pb"
	"github.com/c-4u/employee-service/application/rest"
	"github.com/c-4u/employee-service/infrastructure/external"
	"github.com/spf13/cobra"
)

// NewRestCmd represents the rest command
func NewRestCmd() *cobra.Command {
	var restPort int

	restCmd := &cobra.Command{
		Use:   "rest",
		Short: "Run rest Service",

		Run: func(cmd *cobra.Command, args []string) {
			authServiceAddr := os.Getenv("AUTH_SERVICE_ADDR")
			conn, err := external.ConnectAuthService(authServiceAddr)
			if err != nil {
				log.Fatal(err)
			}

			defer conn.Close()
			authService := pb.NewAuthKeycloakAclClient(conn)

			keycloak := external.ConnectKeycloak()
			rest.StartRestServer(keycloak, authService, restPort)
		},
	}

	restCmd.Flags().IntVarP(&restPort, "port", "p", 8080, "rest server port")

	return restCmd
}

func init() {
	rootCmd.AddCommand(NewRestCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
