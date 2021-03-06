package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func main() {
	var in, out, dir string
	var isappend bool
	app := cli.NewApp()
	app.Usage = "jcbのCSVファイルをExcelに変換"
	app.UsageText = filepath.Base(os.Args[0]) + " {-f file | -d dir} -o out [-a]"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "f",
			Usage:       "JCBの明細ファイル(CSV)",
			Destination: &in,
		},
		cli.StringFlag{
			Name:        "d",
			Usage:       "JCBの明細ファイル(CSV)の格納ディレクトリ",
			Destination: &dir,
		},
		cli.StringFlag{
			Name:        "o",
			Usage:       "出力ファイル名",
			Destination: &out,
		},
		cli.BoolFlag{
			Name:        "a",
			Usage:       "追加書き",
			Destination: &isappend,
		},
	}
	app.HideVersion = true

	app.Action = func(c *cli.Context) error {
		if err := argCheck(in, dir, out); err != nil {
			return cli.NewExitError(app.UsageText, 1)
		}
		return doTask(in, dir, out, isappend)
	}

	app.Run(os.Args)
}

func argCheck(in, dir, out string) error {
	if (in == "" && dir == "") || out == "" {
		return errors.New("Invalid Argment")
	}
	if in != "" && dir != "" {
		return errors.New("-fか-dのどちらかのみを指定")
	}
	return nil
}

func inputList(in, dir string) ([]string, error) {
	var paths []string
	if in != "" {
		return append(paths, in), nil
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".csv" {
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths, nil
}

func doTask(in, dir, out string, isappend bool) error {
	paths, err := inputList(in, dir)
	if err != nil {
		return err
	}

	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	if isappend {
		mode = os.O_APPEND | os.O_WRONLY
	}
	ofile, err := os.OpenFile(out, mode, 0666)
	if err != nil {
		return err
	}
	defer ofile.Close()
	writer := csv.NewWriter(transform.NewWriter(ofile, japanese.ShiftJIS.NewEncoder()))
	defer writer.Flush()
	writer.UseCRLF = true

	for _, f := range paths {
		rows, err := readIn(f)
		if err != nil {
			fmt.Println(err)
			return err
		}
		for _, row := range rows {
			if len(row) < 12 || row[0] == "ご利用者" {
				continue
			}
			//ご利用日,ご利用先など,ご利用金額
			writer.Write([]string{row[2], row[3], row[4]})
		}
	}
	return nil
}

func readIn(file string) ([][]string, error) {
	csvFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(transform.NewReader(csvFile, japanese.ShiftJIS.NewDecoder()))
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}
