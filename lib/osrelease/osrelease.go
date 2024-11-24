package osrelease

import (
	"os"
	"strings"
)

const SUPPORTED_SYSTEMD_VERSION = 256

// OsRelease uses the standard defined by [https://www.freedesktop.org/software/systemd/man/latest/os-release.html]
type OsRelease struct {
	name             string
	id               string
	idLike           string
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

// Get returns the specified value defined by the given key
// This method checks official and unknown keys
// Returns the found value and nil; or
// an empty string and an error
func (i *OsRelease) Get(key string) (string, error) {
	return i.get(key, false)
}

// IsLike returns true if any of the given values are in 'ID_LIKE'
func (i *OsRelease) IsLike(name string) bool {
	return strings.Contains(i.idLike, name)
}

// Pretty access methods, to access things that are considered
// the "pretty" version of the more programmatic ones

// Pretty returns the value of PRETTY_NAME
func (i *OsRelease) Pretty() string { return i.prettyName }

// PrettyVariant returns the value of VARIANT
func (i *OsRelease) PrettyVariant() string { return i.variant }

// PrettyVersion returns the value of VERSION
func (i *OsRelease) PrettyVersion() string { return i.version }

// Name returns the value of NAME
func (i *OsRelease) Name() string { return i.name }

// Id returns the value of ID
func (i *OsRelease) Id() string { return i.id }

// Version returns the value of VERSION_ID
func (i *OsRelease) Version() string { return i.versionId }

// Codename returns the value of VERSION_CODENAME
func (i *OsRelease) Codename() string { return i.versionCodename }

// Variant returns the value of VARIANT_ID
func (i *OsRelease) Variant() string { return i.variantId }

// BuildId returns the value of BUILD_ID
func (i *OsRelease) BuildId() string { return i.buildId }

// Arch returns the value of ARCHITECTURE
func (i *OsRelease) Arch() string { return i.architecture }

// ImageId returns the value of IMAGE_ID
func (i *OsRelease) ImageId() string { return i.imageId }

// ImageVersion returns the value of IMAGE_VERSION
func (i *OsRelease) ImageVersion() string { return i.imageVersion }
