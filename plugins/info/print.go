package info

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/dmux/pkg/sot"
)

const noResult = "No results found in the backend SoT"

// EmptyTable func will render a table with a single row that says "Query did not match any results in NautobotV1"
func EmptyTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.AppendRow(table.Row{noResult})
	t.Render()
}

// PrintDevices func will take a structured type of response data and render a table
// output to stdout.
func PrintDevices(data []sot.Device) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	h := text.FgHiMagenta.Sprint

	if len(data) == 0 {
		t.AppendRow(table.Row{noResult})
		t.Render()
		return
	}

	header := table.Row{
		h("Device"),
		h("IPAddress"),
		h("Status"),
		h("Serial"),
		h("Location"),
		h("Model"),
		h("Console"),
	}

	t.AppendHeader(header)
	for _, d := range data {
		c := status(d.Status)
		row := table.Row{
			c(d.Hostname),
			d.IP,
			c(d.Status),
			d.Serial,
			d.Location,
			d.Platform,
			func() string {
				if d.Console.Hostname == "" {
					return ""
				}
				return fmt.Sprintf("%s:%s", d.Console.Hostname, d.Console.Port)
			}(),
		}
		t.AppendRow(row)
	}
	t.Render()
}

// status func will return a function that will colorize the status of a site.
func status(s string) func(a ...interface{}) string {
	switch {
	case strings.EqualFold(s, "Active"):
		return text.FgGreen.Sprint
	case strings.EqualFold(s, "Failed"):
		return text.FgRed.Sprint
	case strings.EqualFold(s, "Decommissioning"):
		return text.FgYellow.Sprint
	case strings.EqualFold(s, "Inventory"):
		return text.FgBlue.Sprint
	case strings.EqualFold(s, "Planned"):
		return text.FgCyan.Sprint
	default:
		return text.Italic.Sprint
	}
}
