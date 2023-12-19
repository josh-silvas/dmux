package shared

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/josh-silvas/nbot/shared/nautobot"
)

// EmptyTable func will render a table with a single row that says "Query did not match any results in Nautobot"
func EmptyTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.AppendRow(table.Row{"Query did not match any results in Nautobot"})
	t.Render()
}

// PrintSites func will take a structured type of response data and render a table
// output to stdout.
func PrintSites(data []nautobot.Site, verbose bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	h := text.FgHiBlue.Sprint

	headers := table.Row{
		h("Name"),
		h("Address"),
		h("Status"),
		h("Devices"),
		h("Circuits"),
		h("Prefixes"),
		h("ArcheType"),
	}
	if verbose {
		headers = append(headers, h("Tags"))
	}
	t.AppendHeader(headers)

	for _, d := range data {
		c := status(d.Status.Label)

		row := table.Row{
			c(d.Name),
			d.PhysicalAddress,
			c(d.Status.Label),
			d.DeviceCount,
			d.CircuitCount,
			d.PrefixCount,
			getCf("arch_type", d.CustomFields),
		}
		if verbose {
			row = append(row, getTags(d.Tags))
		}

		t.AppendRow(row)
		t.AppendSeparator()
	}
	t.Render()
}

// PrintSite func will take a structured type of response data and render a table
// output to stdout.
func PrintSite(site nautobot.Site) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	h := text.FgHiBlue.Sprint
	t.AppendHeader(table.Row{
		h("Site"),
		h("Address"),
		h("Status"),
		h("Devices"),
		h("Circuits"),
		h("Prefixes"),
		h("ArcheType"),
		h("Tags"),
	})
	c := status(site.Status.Label)
	t.AppendRow(table.Row{
		c(site.Name),
		site.PhysicalAddress,
		c(site.Status.Label),
		site.DeviceCount,
		site.CircuitCount,
		site.PrefixCount,
		getCf("arch_type", site.CustomFields),
		getTags(site.Tags),
	})
	t.Render()
}

// PrintDevices func will take a structured type of response data and render a table
// output to stdout.
func PrintDevices(data []DeviceAggregate, verbose bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	h := text.FgHiMagenta.Sprint

	if len(data) == 0 {
		t.AppendRow(table.Row{"Query did not match any devices in Nautobot"})
		t.Render()
		return
	}

	header := table.Row{
		h("Device"),
		h("AssetTag"),
		h("IPAddress"),
		h("Status"),
		h("Site"),
		h("Model"),
		h("Serial"),
		h("Tenant"),
		h("Console"),
	}
	if verbose {
		header = append(header, h("Manufacturer"))
		header = append(header, h("Tags"))
	}
	t.AppendHeader(header)
	for _, d := range data {
		c := status(d.Record.Status.Label)
		row := table.Row{
			c(d.Record.Name),
			d.Record.AssetTag,
			d.valuePrimaryIP(),
			c(d.Record.Status.Label),
			d.Record.Site.Name,
			d.Record.DeviceType.Model,
			d.Record.Serial,
			d.valueTenant(),
			d.valueConsole(),
		}
		if verbose {
			row = append(row, d.valueManufacturer())
			row = append(row, d.valueTags())
		}
		t.AppendRow(row)
	}
	t.Render()
}

// PrintPrefixes func will take a structured type of response data and render a table
func PrintPrefixes(data []nautobot.Prefix, verbose bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	h := text.FgHiBlue.Sprint

	headers := table.Row{
		h("Prefix"),
		h("Status"),
		h("Role"),
		h("Tenant"),
		h("VLAN"),
		h("VRF"),
		h("IsPool"),
	}
	if verbose {
		headers = append(headers, h("Tags"))
	}
	t.AppendHeader(headers)

	for _, d := range data {
		c := status(d.Status.Label)

		row := table.Row{
			c(d.Prefix),
			c(d.Status.Label),
			func() string {
				if d.Role == nil {
					return ""
				}
				return d.Role.Model
			}(),
			func() string {
				if d.Tenant == nil {
					return ""
				}
				return d.Tenant.Name
			}(),
			func() string {
				if d.VLAN == nil {
					return ""
				}
				return fmt.Sprintf("%d %s", d.VLAN.VID, d.VLAN.Name)
			}(),
			func() string {
				if d.VRF == nil {
					return ""
				}
				return d.VRF.Name
			}(),
			d.IsPool,
		}
		if verbose {
			row = append(row, getTags(d.Tags))
		}

		t.AppendRow(row)
		t.AppendSeparator()
	}
	t.Render()
}

// PrintCircuits pretty prints a list of Circuit models from Nautobot.
func PrintCircuits(data []nautobot.Circuit, verbose bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	h := text.FgHiBlue.Sprint

	headers := table.Row{
		h("Circuit"),
		h("Status"),
		h("Type"),
		h("Provider"),
		h("Tenant"),
		h("TermA"),
		h("TermZ"),
	}
	if verbose {
		headers = append(headers, h("Tags"))
	}
	t.AppendHeader(headers)

	for _, d := range data {
		c := status(d.Status.Label)

		row := table.Row{
			c(d.CircuitID),
			c(d.Status.Label),
			d.Type.Name,
			d.Provider.Name,
			d.Tenant.Name,
			d.TerminationA.Display,
			d.TerminationZ.Display,
		}
		if verbose {
			row = append(row, getTags(d.Tags))
		}

		t.AppendRow(row)
		t.AppendSeparator()
	}
	t.Render()
}

// getTags func will return a string of tags from a slice of tags.
func getTags(tags []nautobot.NestedTag) string {
	r := ""
	for i := range tags {
		r = fmt.Sprintf("%s%s\n", r, tags[i].Name)
	}
	return strings.TrimRight(r, "\n")
}

// getCf func will return a string of a custom field from a map of custom fields.
func getCf(cf string, cfs map[string]interface{}) string {
	if val, ok := cfs[cf]; ok {
		return fmt.Sprintf("%v", val)
	}
	return ""
}

// status func will return a function that will colorize the status of a site.
func status(s string) func(a ...interface{}) string {
	switch s {
	case "Active":
		return text.FgGreen.Sprint
	case "Failed":
		return text.FgRed.Sprint
	case "Decommissioning":
		return text.FgYellow.Sprint
	case "Inventory":
		return text.FgBlue.Sprint
	case "Planned":
		return text.FgCyan.Sprint
	default:
		return text.Italic.Sprint
	}
}
