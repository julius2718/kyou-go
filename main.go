package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "today",
		Usage: "Print Today's Date",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "copy",
				Value: false,
				Usage: "Copy to clipboard",
			},
		},
		Action: func(cCtx *cli.Context) error {
			shouldCopy := cCtx.Bool("copy")

			fmt.Println(getAlfredJson(shouldCopy))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getAlfredJson(copy bool) string {
	// today := getCurrentDate()
	var alfred string

	if copy == true {
		alfred = getCurrentDate("2006-01-02")
	} else {
		dates := map[string]string{
			"YYYY-MM-DD":    getCurrentDate("2006-01-02"),
			"YYYY/MM/DD":    getCurrentDate("2006/01/02"),
			"YYYY年M月D日":     getCurrentDate("2006年1月2日"),
			"令和W年M月D日":      getCurrentDate("平成18年1月2日"),
			"D Month YYYY":  getCurrentDate("2 January 2006"),
			"D Mon. YYYY":   getCurrentDate("2 Jan. 2006"),
			"Month D, YYYY": getCurrentDate("January 2, 2006"),
			"Mon. D, YYYY":  getCurrentDate("Jan. 2, 2006"),
		}
		alfred = toAlfredJson(dates)
	}
	// fmt.Println(alfred)
	return alfred
}

func getCurrentDate(format string) string {
	var nowFmtd string

	now := time.Now()

	if format == "平成18年1月2日" {
		wyear := now.Year() - 2018
		fString := fmt.Sprintf("令和%d年1月2日", wyear)

		nowFmtd = now.Format(fString)
	} else {
		nowFmtd = now.Format(format)
	}

	// fmt.Println(nowFmtd)
	return nowFmtd
}

func toAlfredJson(results map[string]string) string {
	var itms []map[string]string
	var i int = 0

	for key, res := range results {
		itm := map[string]string{
			"uid":      fmt.Sprint(i),
			"title":    res,
			"subtitle": key,
			"arg":      res,
		}
		itms = append(itms, itm)
		i++
	}
	alfredDict := map[string][]map[string]string{
		"items": itms,
	}
	alfredJson, err := json.Marshal(alfredDict)
	if err != nil {
		panic(err)
	}

	return string(alfredJson)
}
