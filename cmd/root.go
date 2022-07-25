package cmd

import (
	"fmt"
	"os"
	"strings"
	"wagger/config"
	"wagger/pkg/file"

	"github.com/spf13/cobra"
)

var cfgFile string
var onlyHighlight = false
var ignore = false
var filePath string
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wagger",
	Short: "A tool for tailing files",
	Long:  `A tool for tailing files`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Config()
		ign := cfg.GetStringSlice("ignore")
		highlight := cfg.GetStringMap("highlight")
		dat, err := file.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		d := string(dat)
		ds := strings.Split(d, "\n")
		file.HighlightAndIgnore(ds, highlight, ign, onlyHighlight, ignore)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&onlyHighlight, "only-highlight", "o", false, "Display only highlighted strings")
	rootCmd.PersistentFlags().BoolVarP(&ignore, "ignore", "i", false, "Ignore strings that match the ignore list")
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Path to file to tail")
	// TODO: I want to get this to work
	// rootCmd.PersistentFlags().StringVarP(&log.LogLevel, "log-level", "l", "info", "Set the log level")
	// rootCmd.PersistentFlags().BoolVarP(&log.JSONLogs, "json-logs", "j", false, "Output logs in JSON format")
	// rootCmd.PersistentFlags().BoolVarP(&log.PrettyPrint, "pretty-print", "p", false, "Pretty print logs")
}
