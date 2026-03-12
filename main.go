package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Host struct {
	Name        string
	Group       string
	Hostname    string
	User        string
	Port        string
	Description string
}

type Group struct {
	Name     string
	Hosts    []*Host
	Expanded bool
}

var groups []*Group

func parseConfig(path string) []*Group {
	file, err := os.Open(path)
	if err != nil {
		log.Println("Nie można otworzyć configu:", err)
		return nil
	}
	defer file.Close()

	var currentGroup *Group
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		key := strings.ToLower(parts[0])
		switch key {
		case "group":
			currentGroup = &Group{
				Name:     strings.Join(parts[1:], " "),
				Hosts:    []*Host{},
				Expanded: false,
			}
			groups = append(groups, currentGroup)
		case "host":
			if currentGroup == nil {
				continue
			}
			host := &Host{
				Name:  parts[1],
				Group: currentGroup.Name,
				Port:  "22",
			}
			currentGroup.Hosts = append(currentGroup.Hosts, host)
		case "hostname":
			if currentGroup != nil && len(currentGroup.Hosts) > 0 {
				currentGroup.Hosts[len(currentGroup.Hosts)-1].Hostname = parts[1]
			}
		case "user":
			if currentGroup != nil && len(currentGroup.Hosts) > 0 {
				currentGroup.Hosts[len(currentGroup.Hosts)-1].User = parts[1]
			}
		case "port":
			if currentGroup != nil && len(currentGroup.Hosts) > 0 {
				currentGroup.Hosts[len(currentGroup.Hosts)-1].Port = parts[1]
			}
		case "description":
			if currentGroup != nil && len(currentGroup.Hosts) > 0 {
				currentGroup.Hosts[len(currentGroup.Hosts)-1].Description = strings.Join(parts[1:], " ")
			}
		}
	}

	return groups
}

func findItemByIndex(index int) (interface{}, *Group) {
	count := 0
	for _, g := range groups {
		if count == index {
			return g, g
		}
		count++
		if g.Expanded {
			for _, h := range g.Hosts {
				if count == index {
					return h, g
				}
				count++
			}
		}
	}
	return nil, nil
}

func runSSH(host *Host) {
	cmdArgs := []string{}
	if host.Port != "" && host.Port != "22" {
		cmdArgs = append(cmdArgs, "-p", host.Port)
	}
	target := host.Hostname
	if host.User != "" {
		target = host.User + "@" + host.Hostname
	}
	cmdArgs = append(cmdArgs, target)

	c := exec.Command("ssh", cmdArgs...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		fmt.Println("Błąd SSH:", err)
	}
}

func main() {
	groups = parseConfig("ssh_connections.conf")
	if len(groups) == 0 {
		fmt.Println("Brak hostów w konfiguracji")
		return
	}

	app := tview.NewApplication()
	table := tview.NewTable().SetSelectable(true, false)

	details := tview.NewTextView()
	details.SetDynamicColors(true)
	details.SetBorder(true)
	details.SetTitle("Details")
	details.SetBorderColor(tcell.ColorGreen)
	details.SetTitleColor(tcell.ColorGreen)

	pingBox := tview.NewTextView()
	pingBox.SetDynamicColors(true)
	pingBox.SetBorder(true)
	pingBox.SetTitle("Ping Output")
	pingBox.SetBorderColor(tcell.ColorYellow)
	pingBox.SetTitleColor(tcell.ColorYellow)
	pingBox.SetScrollable(true)

	refreshTable := func() {
		table.Clear()
		row := 0
		for _, g := range groups {
			symbol := "▼"
			if !g.Expanded {
				symbol = "▶"
			}
			groupCell := tview.NewTableCell(fmt.Sprintf("[blue]%s", symbol+" "+g.Name)).
				SetSelectable(true).
				SetExpansion(0)
			table.SetCell(row, 0, groupCell)
			row++
			if g.Expanded {
				for _, h := range g.Hosts {
					hostDisplay := h.Name
					if h.Hostname != "" {
						hostDisplay += " (" + h.Hostname + ")"
					}
					hostCell := tview.NewTableCell(fmt.Sprintf("  [green]%s", hostDisplay)).
						SetSelectable(true).
						SetExpansion(0)
					table.SetCell(row, 0, hostCell)
					row++
				}
			}
		}
	}

	refreshTable()

	table.SetSelectedFunc(func(row int, column int) {
		item, _ := findItemByIndex(row)
		if h, ok := item.(*Host); ok {
			app.Stop()
			runSSH(h)
		}
	})

	table.SetSelectionChangedFunc(func(row, column int) {
		item, _ := findItemByIndex(row)
		switch v := item.(type) {
		case *Host:
			details.SetText(fmt.Sprintf(
				"[green]Name: [white]%s\n[cyan]Host: [white]%s\n[red]User: [white]%s\n[yellow]Port: [white]%s\n[blue]Group: [white]%s\n[darkgreen]Description: [white]%s",
				v.Name, v.Hostname, v.User, v.Port, v.Group, v.Description))
		case *Group:
			details.SetText(fmt.Sprintf("[green]Group: [white]%s", v.Name))
		}
	})

	// Obsługa klawiszy
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		row, _ := table.GetSelection()
		item, _ := findItemByIndex(row)

		switch event.Key() {
		case tcell.KeyRight:
			if g, ok := item.(*Group); ok && !g.Expanded {
				g.Expanded = true
				refreshTable()
				if len(g.Hosts) > 0 {
					table.Select(row+1, 0)
				} else {
					table.Select(row, 0)
				}
			}
			return nil
		case tcell.KeyLeft:
			if g, ok := item.(*Group); ok && g.Expanded {
				g.Expanded = false
				refreshTable()
				table.Select(row, 0)
			}
			return nil
		}

		// Nasłuchiwanie liter (np. "p" dla ping)
		if event.Rune() == 'p' || event.Rune() == 'P' {
			if h, ok := item.(*Host); ok && h.Hostname != "" {
				pingBox.Clear()
				fmt.Fprintf(pingBox, "[yellow]Pinging %s...\n", h.Hostname)
				go func(hostname string) {
					cmd := exec.Command("ping", "-c", "4", hostname) // 4 pakiety
					stdout, _ := cmd.StdoutPipe()
					if err := cmd.Start(); err != nil {
						app.QueueUpdateDraw(func() {
							fmt.Fprintf(pingBox, "Błąd uruchomienia ping: %v\n", err)
						})
						return
					}
					scanner := bufio.NewScanner(stdout)
					for scanner.Scan() {
						line := scanner.Text()
						app.QueueUpdateDraw(func() {
							fmt.Fprintln(pingBox, line)
						})
					}
					cmd.Wait()
				}(h.Hostname)
			}
			return nil
		}

		return event
	})

	table.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorDarkCyan).Foreground(tcell.ColorWhite))

	// Layout: tabela po lewej, detale + pingBox po prawej
	rightPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(details, 0, 1, false).
		AddItem(pingBox, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(table, 0, 2, true).
		AddItem(rightPane, 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}