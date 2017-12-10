package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	target  string = "path.txt"
	suffix  string = ".wn.oro.co.jp"
	fontExt string = `\.(otf|ttf|ttc|fon)$`
	Pattern  *regexp.Regexp = regexp.MustCompile(fontExt)
	Hostname string         = GetHostname()
	Buffer   [][]string     = [][]string{}
)

func ScanPath(fp *os.File) []string {
	paths := []string{}
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			paths = append(paths, s)
		}
	}

	return paths
}

func GetPaths(filename string) ([]string, error) {
	fp, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	paths := ScanPath(fp)
	if len(paths) > 0 {
		return paths, nil
	} else {
		return nil, errors.New("none directory list")
	}
}

func GetHostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "unknown host"
	}

	return strings.Replace(host, suffix, "", 1)
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
	paths, err := GetPaths(target)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := WalkPaths(paths); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := CsvFilename(GetSaveDirectory())
	if err := WriteToCSV(output); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
