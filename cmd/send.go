/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send the torrent link and media type to the server",
	Long: `
	Requires a -l and -m flag


blackbeard send -l avengers.torrent -m Movies`,
	Run: func(cmd *cobra.Command, args []string) {
		link, _ := cmd.Flags().GetString("l")
		media, _ := cmd.Flags().GetString("m")

		if media != "Movies" && media != "Anime" && media != "TvShows" {
			fmt.Println("You selected something other than Movies, Anime, or TvShows")
		} else {
			fmt.Println(link, media)

			const url = "http://192.168.1.5:7123"
			jsonData := fmt.Sprintf(`{"Magnet": "%s", "Media": "%s"}`, link, media)
			requestBody := strings.NewReader(jsonData)
			resp, err := http.Post(url, "application/json", requestBody)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			content, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(content)
		}

	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")
	sendCmd.PersistentFlags().String("l", "", "Specify a magent or torrent link (required)")
	sendCmd.MarkFlagRequired("l")
	sendCmd.PersistentFlags().String("m", "", "Specify media type. Ex: Movies,TvShows or Anime (required)")
	sendCmd.MarkFlagRequired("m")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
