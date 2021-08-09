/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/c-4u/employee-service/application/grpc"
	"github.com/c-4u/employee-service/application/grpc/pb"
	"github.com/c-4u/employee-service/application/rest"
	"github.com/c-4u/employee-service/infrastructure/external"
	"github.com/spf13/cobra"
)

// NewAllCmd represents the all command
func NewAllCmd() *cobra.Command {
	var grpcPort int
	var restPort int

	allCmd := &cobra.Command{
		Use:   "all",
		Short: "Run both gRPC and rest servers",

		Run: func(cmd *cobra.Command, args []string) {
			authServiceAddr := os.Getenv("AUTH_SERVICE_ADDR")
			conn, err := external.ConnectAuthService(authServiceAddr)
			if err != nil {
				log.Fatal(err)
			}

			defer conn.Close()
			authService := pb.NewAuthKeycloakAclClient(conn)

			keycloak := external.ConnectKeycloak()
			go rest.StartRestServer(keycloak, authService, restPort)
			grpc.StartGrpcServer(keycloak, authService, grpcPort)
		},
	}

	allCmd.Flags().IntVarP(&grpcPort, "grpcPort", "g", 50051, "gRPC Server port")
	allCmd.Flags().IntVarP(&restPort, "restPort", "r", 8080, "rest server port")

	return allCmd
}

func init() {
	rootCmd.AddCommand(NewAllCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
