package server

import (
	"context"
	"fmt"
	"goSum/pkg/calc/sum"
	"goSum/pkg/proto"
	"log"
	"net"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	portDefault = "8080"
)

var (
	port    string
	cfgFile string
)

type server struct{}

func (s *server) Sum(ctx context.Context, request *proto.SumRequest) (*proto.SumResult, error) {
	// Receive array of ints from request and add them up
	result, err := sum.Calculate(request.Numbers)
	if err != nil {
		os.Exit(1)
	}
	// Send back result in response
	return &proto.SumResult{Result: result}, nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start new server on default port 8080 or specify your own port using -p flag",
	Example: "-p 8081",
	Run: func(cmd *cobra.Command, args []string) {
		if port == "" {
			port = portDefault
		}
		log.Printf("Starting new server on port %v.\n", port)

		port = fmt.Sprintf(":%v", port)
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		proto.RegisterSumServiceServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.omniactl.yaml)")
	rootCmd.Flags().StringVarP(&port, "port", "p", "", "Set port for server")
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".gosum")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
