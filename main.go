package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"log"
	"os"
	"sakura-vps-cli/lib"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "sv"
	app.Usage = "fight the loneliness!"
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
						client := lib.NewAPI("https://secure.sakura.ad.jp/vps/api/v6.5/", "e0b0880e-9859-49ee-a07c-3cd2c8d0e6d7")
						var servers = client.GetServers()
						var data [][]string
						for _, server := range servers {
							data = append(data, []string{
								strconv.Itoa(server.ID),
								server.Name,
								server.Ipv4[0].Hostname,
								server.Ipv4[0].Address,
								server.Ipv6[0].Address,
								server.GetZoneName(),
								strconv.Itoa(server.CPUCores),
								server.GetMemorySize(),
								server.GetDiskSize(),
								server.GetStorageType(),
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
	//token := lib.GetToken()
	//fmt.Println(token)
	//client := lib.NewAPI("https://secure.sakura.ad.jp/vps/api/v6.5/", "e0b0880e-9859-49ee-a07c-3cd2c8d0e6d7")
	//fmt.Println(client.GetMonitoring(client.GetMonitorings()[0].ID))
}
