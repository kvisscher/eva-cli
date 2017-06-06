package cmd

import (
	"fmt"

	"encoding/json"
	"os"

	"github.com/new-black/eva-cli/client"
	"github.com/new-black/eva-cli/messages"
	"github.com/urfave/cli"
)

func GenerateProductCommands(client client.Client) []cli.Command {
	return []cli.Command{
		cli.Command{
			Name: "get",
			Action: func(c *cli.Context) error {
				var response messages.GetProductDetailResponse

				if err := client.Get(fmt.Sprintf("api/v1/product/%s", c.Args().First()), &response); err != nil {
					return err
				}

				w := json.NewEncoder(os.Stdout)
				w.SetIndent("", " ")
				w.Encode(&response.Result)

				return nil
			},
		},
		cli.Command{
			Name: "barcode",
			Action: func(c *cli.Context) error {
				var response messages.GetProductDetailResponse

				if err := client.Send(messages.GetProductByBarcode{c.Args().First()}, &response); err != nil {
					return err
				}

				w := json.NewEncoder(os.Stdout)
				w.SetIndent("", " ")
				w.Encode(&response.Result)

				return nil
			},
		},
	}
}
