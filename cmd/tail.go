package cmd

import (
	"time"
	"wagger/config"
	"wagger/pkg/file"

	"github.com/spf13/cobra"
)

var seconds int

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config()
		ign := cfg.GetStringSlice("ignore")
		highlight := cfg.GetStringMap("highlight")
		duration := time.Second * time.Duration(seconds)
		file.TailFile(filePath, highlight, ign, onlyHighlight, ignore, duration)
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tailCmd.Flags().IntVarP(&seconds, "seconds", "s", 1, "Poll interval, how many seconds to wait before next poll")
}
