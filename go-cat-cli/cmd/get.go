/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get desired cat",
	Long: `This get command will call github repo in order to return the cat`,
	Run: func(cmd *cobra.Command, args []string) {
		var catName = "sadCatto.gif"

		if len(args) >= 1 && args[0] != "" {
			catName = args[0]
		}

		URL := "https://github.com/Pandahoro/cats/raw/main/" + catName + ".gif"
		
		fmt.Println("try to get '" + catName + "' cat...")

		//get data

		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			//create the file
			out, err := os.Create(catName + ".gif")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()

			//Writer the body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println ("Purrfect! just saved in " + out.Name() + "!")

		} else {
			fmt.Println("Error: " + catName + " no exist! :< ")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
