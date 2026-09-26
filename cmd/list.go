package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/khpalwatan/portswarden/internal/ports"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	listPortFilter string
	listJSONOutput bool
	listSortBy     string
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List listening ports",
	RunE: func(c *cobra.Command, _ []string) error {
		data, err := ports.List()
		if err != nil {
			return err
		}

		if listPortFilter != "" {
			var filtered []ports.PortInfo
			for _, p := range data {
				if fmt.Sprint(p.Port) == listPortFilter {
					filtered = append(filtered, p)
				}
			}
			data = filtered
		}

		switch listSortBy {
		case "pid":
			sort.Slice(data, func(i, j int) bool { return data[i].PID < data[j].PID })
		case "process":
			sort.Slice(data, func(i, j int) bool { return data[i].Process < data[j].Process })
		default:
			sort.Slice(data, func(i, j int) bool { return data[i].Port < data[j].Port })
		}

		if listJSONOutput {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(data)
		}

		table := tablewriter.NewTable(os.Stdout)
		table.Header([]string{"PORT", "PROTO", "STATE", "PID", "PROCESS", "PATH"})

		for _, p := range data {
			_ = table.Append([]any{
				fmt.Sprint(p.Port),
				p.Proto,
				p.State,
				fmt.Sprint(p.PID),
				p.Process,
				p.Path,
			})
		}
		_ = table.Render()
		return nil
	},
}

func init() {
	listCmd.Flags().StringVarP(&listPortFilter, "port", "p", "", "filter by port")
	listCmd.Flags().BoolVar(&listJSONOutput, "json", false, "output as JSON")
	listCmd.Flags().StringVarP(&listSortBy, "sort", "s", "port", "sort by: port, pid, process")
	rootCmd.AddCommand(listCmd)
}