package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var path *string = flag.String("p", "", "/path/to/dir")
	var hostname *string = flag.String("h", "", "hostname")
	
	flag.Parse()

	pattern := regexp.MustCompile(`\.(otf|ttf|ttc|fon)$`)

	if *path == "" {
		fmt.Println("invalid -p parameter")
		os.Exit(1)
	}
	if *hostname == "" {
		fmt.Println("invaild -h parameter")
		os.Exit(1)
	}
	
	var err error
	var root string = ""
	
	root, err = filepath.Abs(*path)
	if err != nil {
		os.Exit(1)
	}
	
	if _, e := os.Stat(root); e != nil {
		fmt.Println("path not found")
		os.Exit(1)
	}
	
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		
		if pattern.MatchString(strings.ToLower(path)) {
			fmt.Printf("%s,%s,%s\r\n", *hostname, filepath.Dir(path), filepath.Base(path))
		}
		return nil
	})
	
	if err != nil {
		fmt.Println(1, err)
	}
}
