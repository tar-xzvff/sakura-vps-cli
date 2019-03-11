package main

import (
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"log"
	"os"
	"./lib"
	"strconv"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "sv"
	app.Usage = "CLI client for SakuraVps"
	app.Commands = []cli.Command{
		{
			Name:	"auth",
			Usage:	"Authentication commands",
			Subcommands: []cli.Command{
				{
					Name:	"login",
					Usage:	"Login SAKURA VPS",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "Your MemberID",
						},
						cli.StringFlag{
							Name:  "password",
							Usage: "Your Password",
						},
					},
					Action: func(c *cli.Context) error {
						if lib.LogIn(c.String("id"), c.String("password")) {
							fmt.Println(lib.Message("SUCCESS_LOGIN"))
						} else {
							fmt.Println(lib.Message("FAILED_LOGIN"))
						}
						return nil
					},
				},{
					Name:	"logout",
					Usage:	"Logout SAKURA VPS",
					Action: func(c *cli.Context) error {
						if lib.LogOut() {
							fmt.Println(lib.Message("SUCCESS_LOGOUT"))
						} else {
							fmt.Println(lib.Message("FAILED_LOGOUT"))
						}
						return nil
					},
				},
			},
		},
		{
			Name:	"server",
			Usage:	"A manage commands of Server",
			Subcommands: []cli.Command{
				{
					Name:  "ls",
					Usage: "Display of server list",
					Action: func(c *cli.Context) error {
						client := lib.NewAPI()
						var servers = client.GetServers()
						var data [][]string
						for _, server := range servers {
							data = append(data, []string{
								strconv.Itoa(server.ID),
								server.Name,
								server.Ipv4[0].Hostname,
								server.Ipv4[0].Address,
								server.Ipv6[0].Address,
								strconv.Itoa(server.Zone),
								strconv.Itoa(server.CPUCores),
								strconv.Itoa(server.MemoryBytes),
								strconv.Itoa(server.Storage[0].SizeBytes),
								strconv.Itoa(server.Storage[0].Type),
								server.HostType,
								server.Power,
							})
						}

						table := tablewriter.NewWriter(os.Stdout)
						table.SetHeader([]string{"ID", "Name", "Hostname", "IPv4", "IPv6", "Zone", "CPU", "Memory", "Disk", "Storage_Type", "Server_Type", "Status"})

						for _, v := range data {
							table.Append(v)
						}
						table.Render() // Send output
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
