package main

import (
	"runtime"
	"strings"
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	beforeAll()
	defer afterAll()
	
	os.Exit(m.Run())
}

func beforeAll() {
	target = "path.txt"
	suffix = ".wn.oro.co.jp"
	fontExt = `\.(otf|ttf|ttc|fon)$`
	Buffer = [][]string{}
}

func afterAll() {
	
}

func TestGetHostname(t *testing.T) {
	result := GetHostname()
	if strings.Contains(result, suffix) {
		t.Error("suffix trimming error")
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

func TestGetPaths(t *testing.T) {
	test1, err1 := GetPaths("./sample/notfoundfile.txt")
	if err1 == nil {
		t.Error("GetPaths file open error")
	}
	if test1 != nil {
		t.Error("GetPaths file open error")
	}

	test2, err2 := GetPaths("./sample/path.txt")
	if err2 != nil {
		t.Error(err2.Error())
	}
	if len(test2) == 0 || len(test2) > 2 {
		t.Error("GetPaths line error")
	}

	test3, err3 := GetPaths("./sample/empty.txt")
	if err3 == nil {
		t.Error("GetPaths empty file error")
	} else if err3.Error() != "none directory list" {
		t.Error("GetPaths empty file error message")
	}
	if test3 != nil {
		t.Error("GetPaths empty file error")
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
