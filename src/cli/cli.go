package cli

import (
	"errors"
	"log"
	"time"

	"github.com/kalbhor/tracesite/src/tracesite"
	"github.com/urfave/cli"
)

const DEFAULT_PORT = 33434
const DEFAULT_PACKET_SIZE = 52
const DEFAULT_TIMEOUT_MS = 4000
const DEFAULT_RETRIES = 3
const DEFAULT_DEST = "kalbhor.xyz"
const DEFAULT_MAX_HOPS = 24
const DEFAULT_START_TTL = 1

func Run(args []string) {
	app := cli.NewApp()
	app.Name = "tracesite"
	app.Compiled = time.Now()
	app.Usage = "trace the route to a site"
	app.UsageText = `Trace a site: tracesite --hop=3 --timeout=2000 kalbhor.xyz`
	app.Flags = []cli.Flag{
		cli.IntFlag{Name: "hop", Value: DEFAULT_START_TTL, Usage: "start from a custom hop number"},
		cli.IntFlag{Name: "maxhops", Value: DEFAULT_MAX_HOPS, Usage: "custom max hops"},
		cli.IntFlag{Name: "port", Value: DEFAULT_PORT, Usage: "custom port number"},
		cli.IntFlag{Name: "timeout", Value: DEFAULT_TIMEOUT_MS, Usage: "custom timeout in ms"},
		cli.IntFlag{Name: "retries", Value: DEFAULT_RETRIES, Usage: "custom retries"},
		cli.IntFlag{Name: "packetsize", Value: DEFAULT_PACKET_SIZE, Usage: "custom packet size"},
	}
	app.Action = func(c *cli.Context) error {
		if c.NArg() >= 1 {
			return tracesite.Tracesite(c)
		} else {
			return errors.New("No host specified. Check --help for usage")
		}
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
