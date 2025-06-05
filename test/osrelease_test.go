package test

import (
	"testing"

	"github.com/stobbsm/go-osrelease/lib/osrelease"
)

var f *osrelease.OsRelease

func Test_FromFile(t *testing.T) {
	f = osrelease.NewOsInfo()
	_, err := f.FromFile("test_os_release")
	if err != nil {
		t.Errorf("Error while opening file %s: %s", "test_os_release", err.Error())
	}
}

func TestValidValuesFromRelease(t *testing.T) {
	if f.Name() != "AlmaLinux" {
		t.Errorf("parsed name `%s` should be `AlmaLinux`", f.Name())
	}
}
