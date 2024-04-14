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
		link, _ := cmd.Flags().GetString("link")
		media, _ := cmd.Flags().GetString("media")

		if media != "Movies" && media != "TvShows" && media != "Anime" {
			fmt.Println("-m must be 'Movies' or 'TvShows' or 'Anime'")
		} else {
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
	sendCmd.PersistentFlags().StringP("link", "l", "", "Specify a magent or torrent link (required)")
	sendCmd.PersistentFlags().StringP("media", "m", "", "Specify media type. Ex: Movies,TvShows or Anime (required)")
	sendCmd.MarkPersistentFlagRequired("link")
	sendCmd.MarkPersistentFlagRequired("media")
	// sendCmd.PersistentFlags().String("l", "", "Specify a magent or torrent link (required)")
	// sendCmd.MarkPersistentFlagRequired("l")
	// sendCmd.PersistentFlags().String("m", "", "Specify media type. Ex: Movies,TvShows or Anime (required)")
	// sendCmd.MarkPersistentFlagRequired("m")

}
