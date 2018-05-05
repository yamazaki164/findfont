package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	test1 := &Config{
		Targets:    []string{},
		Extensions: []string{"otf"},
	}
	if test1.IsValid() == true {
		t.Error("IsValid: Targets length error")
	}

	test2 := &Config{
		Targets:    []string{"/a"},
		Extensions: []string{},
	}
	if test2.IsValid() == true {
		t.Error("IsValid: Extensions length error")
	}

	test3 := &Config{
		Targets:    []string{"/a"},
		Extensions: []string{"otf"},
	}
	if test3.IsValid() == false {
		t.Error("IsValid: Targets, Extensions validation error")
	}
}

func TestPatternString(t *testing.T) {
	test1 := &Config{
		Targets:    []string{"/a", "/b"},
		Extensions: []string{"otf"},
	}
	if test1.PatternString() != `\.(otf)$` {
		t.Error("pattern: one extension")
	}

	test2 := &Config{
		Targets:    []string{"/a", "/b"},
		Extensions: []string{"otf", "ttf"},
	}
	if test2.PatternString() != `\.(otf|ttf)$` {
		t.Error("pattern: two extensions")
	}
}

func TestExtensions2Regexp(t *testing.T) {
	data1 := &Config{
		Targets:    []string{"/a", "/b"},
		Extensions: []string{"otf"},
	}

	test1 := data1.Extensions2Regexp()
	if test1 == nil {
		t.Error("Extensions2Regexp: regexp compile error")
	}
}

func TestLoadConfigIsErrorOnConfigFile(t *testing.T) {
	oldValue := configFile
	configFile = "/path/to/error/toml"

	test1, err := LoadConfig()
	if test1 != nil {
		t.Error("LoadConfig: toml load error on invalid path with some Config value")
	}
	t.Log(err)
	if err == nil {
		t.Error("LoadConfig: toml load error on invalid path")
	}

	configFile = oldValue
}

func TestLoadConfigIsErrorOnParameters(t *testing.T) {
	oldValue := configFile
	configFile = "./sample/error.toml"

	_, err := LoadConfig()
	if err == nil {
		t.Error("LoadConfig: toml load error on invalid path")
	}

	if err.Error() != "Invalid toml file parameters" {
		t.Error("LoadConfig: toml parameters validation message is wrong")
	}

	configFile = oldValue
}

func TestLoadConfigIsSuccess(t *testing.T) {
	oldValue := configFile
	configFile = "./sample/success.toml"

	test1, _ := LoadConfig()
	if test1 == nil {
		t.Error("LoadConfig: toml parse error at success.toml")
	}

	configFile = oldValue
}
