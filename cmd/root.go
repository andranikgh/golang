/*
Copyright © 2022 Andranik Gharakeshishyan
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "examples",
	Short: "This is a example of cobra code",
	Long:  `This application will cover almost all features that cobra gaves`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: rootFunction,
}

// This function should print all arguments of example command
func rootFunction(cmd *cobra.Command, args []string) {

	for _, arg := range args {
		fmt.Println(arg)
	}

	fmt.Println("Test if toggle flad is set")
	region, _ := cmd.Flags().GetString("region")
	fmt.Println("Hallo regio  = " + region)

	toggle, _ := cmd.Flags().GetBool("toggle")
	fmt.Println("Hello toggle = ")
	fmt.Println(toggle)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.examples.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	viper.AddConfigPath("$HOME/.example")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	Region := " "
	rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	rootCmd.MarkFlagRequired("region")
	//To disable completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}
