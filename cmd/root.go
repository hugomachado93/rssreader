package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"rssreader/service"
	"strings"
	"strconv"
)

var serviceCmd service.RssReaderService

func NewCmd(service service.RssReaderService) {
	serviceCmd = service
}

var rootCmd = &cobra.Command{
	Use:   "rssreader",
	Short: "List all urls",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// addurlCmd represents the addurl command
var addurlCmd = &cobra.Command{
	Use:   "addurl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serviceCmd.SaveUrl(strings.Join(args, " "))
	},
}

// listurlCmd represents the listurl command
var listurlCmd = &cobra.Command{
	Use:   "listurl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serviceCmd.GetUrls()
		//for _, element := range serviceCmd.GetUrls() {
		//	fmt.Printf("%v - %s\n", element.Index, element.Url)
		//}
	},
}


var selurlCmd = &cobra.Command{
	Use:   "selurl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showDesc, err := cmd.Flags().GetBool("show_desc")
		allArticle, err := cmd.Flags().GetBool("show_all")
		if err != nil {
			fmt.Println(err)
		}
		val, err:= strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
		}
		serviceCmd.ShowArticle(uint(val), showDesc, allArticle)
	},
}

var urlFlag string

func init() {

	selurlCmd.Flags().BoolP("show_desc", "s", false, "show rss description")
	selurlCmd.Flags().BoolP("show_all", "a", false, "show all article")

	rootCmd.AddCommand(addurlCmd)
	rootCmd.AddCommand(listurlCmd)
	rootCmd.AddCommand(selurlCmd)
}