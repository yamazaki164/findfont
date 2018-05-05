package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	Pattern  *regexp.Regexp
	Hostname string     = GetHostname()
	Buffer   [][]string = [][]string{}
)

func GetHostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "UnknownHost"
	}

	return strings.Split(host, ".")[0]
}

func WalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	if Pattern.MatchString(strings.ToLower(path)) {
		item := []string{
			Hostname,
			filepath.Dir(path),
			filepath.Base(path),
		}
		Buffer = append(Buffer, item)
	}
	return nil
}

func FindDirs(root string) error {
	return filepath.Walk(root, WalkFunc)
}

func CsvFilename(dir string) string {
	return filepath.Join(dir, Hostname+".csv")
}

func WriteToCSV(output string) error {
	fp, err := os.OpenFile(output, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	writer := csv.NewWriter(fp)
	writer.UseCRLF = true

	return writer.WriteAll(Buffer)
}

func WalkPaths(paths []string) error {
	for _, path := range paths {
		root, err := filepath.Abs(path)
		if err != nil {
			continue
		}
		if _, err = os.Stat(root); err != nil {
			continue
		}

		err = FindDirs(root)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetSaveDirectory() string {
	return filepath.Dir(os.Args[0])
}

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	Pattern = config.Extensions2Regexp()

	if err := WalkPaths(config.Targets); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := CsvFilename(GetSaveDirectory())
	if err := WriteToCSV(output); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
