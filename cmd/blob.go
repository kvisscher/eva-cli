package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/new-black/eva-cli/client"
	"github.com/new-black/eva-cli/messages"
	"github.com/urfave/cli"
)

func GenerateBlobCommands(client client.Client) []cli.Command {
	return []cli.Command{
		cli.Command{
			Name: "upload",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "file, f"},
				cli.StringFlag{Name: "category, c", Value: "cli"},
			},
			Before: func(c *cli.Context) error {
				_, err := client.GetCurrentUser()

				return err
			},
			Action: func(c *cli.Context) error {
				fileName := c.String("file")

				f, err := os.Open(fileName)

				if err != nil {
					return err
				}

				b, err := ioutil.ReadAll(f)

				if err != nil {
					return err
				}

				var result messages.StoreBlobResponse

				err = client.Send(messages.StoreBlob{
					Category:     c.String("category"),
					Data:         b,
					OriginalName: path.Base(fileName),
					MimeType:     http.DetectContentType(b),
				}, &result)

				if err == nil {
					fmt.Println(result.Guid)
				}

				return err
			},
		},
	}
}
