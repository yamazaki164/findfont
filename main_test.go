package main

import (
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	beforeAll()
	defer afterAll()

	os.Exit(m.Run())
}

func beforeAll() {
	Pattern = regexp.MustCompile(`\.(otf)$`)
	Buffer = [][]string{}
}

func afterAll() {

}

func TestGetHostname(t *testing.T) {
	result := GetHostname()
	if strings.Contains(result, ".") {
		t.Error("hostname parse error")
	}
}

func TestCsvFilename(t *testing.T) {
	Hostname = "test"
	if runtime.GOOS == "windows" {
		result := CsvFilename(`c:\path\to\dir`)
		if result != `c:\path\to\dir\test.csv` {
			t.Error("CsvFilename error")
		}
	} else {
		result := CsvFilename("/path/to/dir")
		if result != "/path/to/dir/test.csv" {
			t.Error("CsvFilename error")
		}
	}
}

func TestWriteToCSV(t *testing.T) {
	output1 := "./sample/test.csv"
	test1 := WriteToCSV(output1)
	if test1 != nil {
		t.Error(test1.Error())
	}

	output2 := "./sample/fail/fail.csv"
	test2 := WriteToCSV(output2)
	if test2 == nil {
		t.Error("TestWriteToCSV error on invalid path")
	}
}

func TestFindDirs(t *testing.T) {
	test1 := FindDirs("./sample")
	if test1 != nil {
		t.Error(test1.Error())
	}

	test2 := FindDirs("./sample/dummy")
	if test2 != nil {
		t.Error(test2.Error())
	}
}

func TestWalkPaths(t *testing.T) {
	Buffer = [][]string{}
	test1 := WalkPaths([]string{
		"./sample",
	})
	if test1 != nil {
		t.Error(test1.Error())
	}

	Buffer = [][]string{}
	p := []string{}
	test2 := WalkPaths(p)
	if test2 != nil {
		t.Error(test2.Error())
	}
	if len(Buffer) > 0 {
		t.Error("WalkPaths error when empty list")
	}
}

func TestGetSaveDirectory(t *testing.T) {
	old := os.Args
	os.Args[0] = "./sample/piyo"
	test1 := GetSaveDirectory()
	if test1 != "sample" {
		t.Error("GetSaveDirectory error")
	}

	os.Args = old
}
