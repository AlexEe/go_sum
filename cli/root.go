package cmd

import (
	"context"
	"fmt"
	"goSum/pkg/proto"
	"log"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	addressDefault = "localhost:8080"
)

var (
	cfgFile string
	numbers []int32
	address string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "sum",
	Short:   "Add numbers entered on the Command Line",
	Example: "sum -n 4,1,-2 -u localhost:8080",
	Run: func(cmd *cobra.Command, args []string) {
		// If address has been set via flag use input
		// Else use default address
		if address == "" {
			address = addressDefault
		}

		// Set up a connection to the server
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Did not connect: %v", err)
		}
		defer conn.Close()
		client := proto.NewSumServiceClient(conn)

		// Contact the server and print out its response
		if len(numbers) < 1 {
			log.Fatalf("No numbers were entered. Example command: 'sum -n 1,3,4'")
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		result, err := client.Sum(ctx, &proto.SumRequest{Numbers: numbers})
		if err != nil {
			log.Fatalf("Could not sum: %v", err)
		}
		fmt.Printf("The sum of ")
		for _, v := range numbers {
			fmt.Print(v, " ")
		}
		fmt.Printf("is %v.\n", result.GetResult())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.omniactl.yaml)")
	rootCmd.Flags().Int32SliceVarP(&numbers, "numbers", "n", []int32{}, "Numbers to be added up")
	rootCmd.Flags().StringVarP(&address, "url", "u", "", "")
	// sumCmd.AddSubCommands(rootCmd)
}

// initConfig reads in config file and ENV variables if set.
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
