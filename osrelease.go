package osrelease

import (
	"os"
)

// OsRelease uses the standard defined by [https://www.freedesktop.org/software/systemd/man/latest/os-release.html]
type OsRelease struct {
	name             string
	id               string
	idLike           map[string]struct{}
	idLike_raw       string
	prettyName       string
	cpeName          string
	variant          string
	variantId        string
	version          string
	versionId        string
	versionCodename  string
	buildId          string
	imageId          string
	imageVersion     string
	homeUrl          string
	documentationUrl string
	supportUrl       string
	bugReportUrl     string
	privacyPolicyUrl string
	supportEnd       string
	logo             string
	ansiColor        string
	vendorName       string
	vendorUrl        string
	defaultHostname  string
	architecture     string
	sysextLevel      string
	confextLevel     string
	sysextScope      string
	confextScope     string
	portablePrefixes string

	unknown  map[string]string
	readFrom string
}

// NewOsInfo initializes a new OsInfo struct
// with the defined defaults of os-release
func NewOsInfo() *OsRelease {
	return &OsRelease{
		name:       "Linux",
		id:         "linux",
		prettyName: "Linux",
		unknown:    make(map[string]string),
	}
}

// FromFile sets the readFrom attribute.
// Panics if the given file can't be opened/doesn't exist
func (i *OsRelease) FromFile(filename string) *OsRelease {
	if _, err := os.Stat(filename); err != nil {
		panic(err)
	}
	i.readFrom = filename
	return i
}

// Load reads and parses the os-release file
// that works
func (i *OsRelease) Load() error {
	return i.load()
}

// Set a value for the specified key
// returns true if the key is officially supported
func (i *OsRelease) Set(key, value string) bool {
	return i.set(key, value)
}

// GetOfficial only returns the value of the official keys
// returns an empty string if the key isn't set
func (i *OsRelease) GetOfficial(key string) string {
	if v, err := i.get(key, true); err == nil {
		return v
	}
	return ""
}

// GetUnknown only checks the unknown keys
// returns an empty string if nothing is found
func (i *OsRelease) GetUnknown(key string) string {
	if v, ok := i.unknown[key]; ok {
		return v
	}
	return ""
}

// GetAny returns the specified value defined by the given key.
// This method checks official and unknown keys
// Returns the found value and nil; or
// an empty string and an error
func (i *OsRelease) GetAny(key string) (string, error) {
	return i.get(key, false)
}

// Like returns true if any of the given values are in 'ID_LIKE'
func (i *OsRelease) Like(dname ...string) bool {
	for _, v := range dname {
		_, ok := i.idLike[v]
		if ok {
			return ok
		}
	}
	return false
}
