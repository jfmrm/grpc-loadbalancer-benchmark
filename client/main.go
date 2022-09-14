package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func main() {
	e := flag.String("experiments", "", "list of experiment files separated by comma")

	flag.Parse()

	experiments := strings.Split(*e, ",")

	PrettyPrint(experiments)

	for i, e := range experiments {
		fmt.Printf("running experiment %s\n", e)

		var cfg runner.Config
		path := fmt.Sprintf("/configs/%s", e)
		err := runner.LoadConfig(path, &cfg)

		options := []runner.Option{runner.WithConfig(&cfg)}

		report, err := runner.Run(cfg.Call, cfg.Host, options...)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fName := strings.Split(e, ".")[0]

		f, _ := os.Create(fmt.Sprintf("/results/%s.csv", fName))

		printer := printer.ReportPrinter{
			Out:    f,
			Report: report,
		}

		printer.Print("csv")

		f.Close()

		if i < len(experiments)-1 {
			time.Sleep(2 * time.Minute)
		}
	}
}