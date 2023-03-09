/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

// Strat server on given port by using child command
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RokibulHasan7/book-api/api"
	"github.com/spf13/cobra"
)

var portNum string

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is 3333
		portChange(args)
	},
}

func init() {
	startserverCmd.AddCommand(portCmd)

	//portCmd.Flags().BoolP("port", "p", false, "Add port number")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func portChange(args []string) {
	api.Init()          // Initialize DB
	api.HandleRequest() // Expose Routers

	portNum = args[0]

	portNum = ":" + portNum
	fmt.Println(portNum)
	// Server start
	sigs := make(chan os.Signal, 1) // Channel created to get the notification of Interrupt
	signal.Notify(sigs, os.Interrupt)

	go func() {
		if err := http.ListenAndServe(portNum, api.Router); err != nil {
			log.Printf("Shutting down, reason: %s", err.Error())
			return
		}
	}()
	log.Printf("Server is listening on port %v", portNum)
	<-sigs

	time.Sleep(2 * time.Second)
	log.Println("Server is shutting down")
}
