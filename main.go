package main
import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"errors"
	"strconv"
)

func main() {
	var token string
	var session Input
	var rootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "Adventofcode.com cli",
		PersistentPreRun: func(cmd *cobra.Command, args []string){
			// var err error
			session = Input{token: token}
		},
		// Run: func(cmd *cobra.Command, args []string) {
		  // Do Stuff Here
		// },
	}
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "aoc session token (required)")
	rootCmd.MarkFlagRequired("token")

	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true}) // no help command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "input [YYYY DD]",
		Short: "downloads input for specified date",
		Example: "  aoc input\n  aoc input 2021 1",
		Args:  func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 && len(args) != 2 {
				return errors.New("Need YYYY and DD args or no args to get today's input")
			}
			return nil
		},
		// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) { # validate year and day }
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Printf(session.GetTodaysInput())
			} else if len(args) == 2 {
				year, _ := strconv.Atoi(args[0])
				day, _ := strconv.Atoi(args[1])
				input, _ := session.GetInput(year, day)
				fmt.Printf(input)
			}
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "scenario [YYYY DD]",
		Short: "downloads scenario/problem for specified date",
		Example: "  aoc scenario\n  aoc scenario 2021 1",
		Args:  func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 && len(args) != 2 {
				return errors.New("Need YYYY and DD args or no args to get today's scenario")
			}
			return nil
		},
		// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) { # validate year and day }
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Printf(session.GetTodaysScenario())
			} else if len(args) == 2 {
				year, _ := strconv.Atoi(args[0])
				day, _ := strconv.Atoi(args[1])
				input, _ := session.GetScenario(year, day)
				fmt.Printf(input)
			}
		},
	})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}  
	  
}