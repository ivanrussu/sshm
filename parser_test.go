package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	config, err := os.ReadFile("config_test")
	if err != nil {
		t.Fatal("cannot read config: " + err.Error())
	}

	hosts, err := parseSshConfig(config)
	if err != nil {
		t.Fatal("cannot parse config: " + err.Error())
	}

	expected := []string{
		"Red",
		"Purple",
		"orange",
		"White",
		"😎 Dark Black",
		"Green",
		"Pink",
	}

	if !reflect.DeepEqual(expected, hosts) {
		t.Errorf("slices are not equal: expected %v, got %v", expected, hosts)
	}
}
