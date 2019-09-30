package main

import (
	cli "goSum/cli"
)

func main() {
	cli.Execute()
}

// const (
// 	address = "localhost:8080"
// )

// var cfgFile string

// // rootCmd represents the base command when called without any subcommands
// var rootCmd = &cobra.Command{
// 	Use:   "GoSum",
// 	Short: "Adds numbers entered on the Command Line",
// 	Example: "GoSum -n 1 2 3"
// 	Run: func(cmd *cobra.Command, args []string) {

// 	 },
// }

// // Execute adds all child commands to the root command and sets flags appropriately.
// // This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	cobra.OnInitialize(initConfig)
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.
// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.omniactl.yaml)")

// 	rootCmd.Flags().Int32SliceVarP(&numbers, "numbers", "n", []int32{}, "Numbers to be added up")

// 	sumCmd.AddSubCommands(rootCmd)
// }

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// Search config in home directory with name ".omniactl" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".gosum")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
