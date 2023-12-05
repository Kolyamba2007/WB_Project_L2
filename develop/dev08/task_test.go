package main

import (
	"testing"
)

func Test_cd(t *testing.T) {
	cd("D:\\Downloads")
	if pwd() != "D:\\Downloads" {
		t.Errorf("error")
	}
}

func Test_ps(t *testing.T) {
	if len(ps()) == 0 {
		t.Errorf("error")
	}
}

func Test_kill(t *testing.T) {
	err := kill(ps()["Discord.exe"])
	if err != nil {
		t.Errorf("error")
	}
}
