/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/RokibulHasan7/book-api/api"
	"github.com/spf13/cobra"
)

// startserverCmd represents the startserver command
var port string
var startserverCmd = &cobra.Command{
	Use:   "startserver",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.StartServer(port)
	},
}

func init() {
	rootCmd.AddCommand(startserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startserverCmd.PersistentFlags().StringVarP(&port, "port", "p", "3333", "Sets the port of server.")
}
