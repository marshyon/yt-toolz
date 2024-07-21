/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets an RSS feed from a YouTube channel",
	Long: `gets a YouTube channel's RSS feed
using the channel's handle name

examples:

yt-toolz get --name @somechannel

`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		//toggle, _ := cmd.Flags().GetBool("toggle")
		rssURL := grabRSS(name)
		fmt.Println(rssURL)
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

	//getCmd.flags().StringP("name", "n", "World", "A name to say hello to");

	getCmd.Flags().StringP("name", "n", "World", "A name to say hello to")

	getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func grabRSS(name string) string {

	channelURL := "https://www.youtube.com/" + name

	resp, err := http.Get(channelURL)
	if err != nil {
		fmt.Println("Error getting channel page:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	pageContent := string(body)

	var channelID string = ""
	re := regexp.MustCompile(`https://www.youtube.com/channel/([^/]+)">`)
	matches := re.FindStringSubmatch(pageContent)
	if len(matches) > 1 {
		channelID = matches[1]
	} else {
		fmt.Println("Channel ID not found")
		return ""
	}

	return fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?channel_id=%s", channelID)
}
