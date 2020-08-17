package main

import (
	"fmt"
	"os"

	"github.com/kalbhor/tracesite/src/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
