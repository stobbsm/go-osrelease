package osrelease

import (
	"fmt"

	"github.com/spf13/viper"
)

// package level access methods, to be used as
// `osrelease.Id()` to retreive an Id, which calls the
// `Id()` method on the default osrelease object

const (
	printBasicEqualsFormat = "%s=%s\n"
	printShellEqualsFormat = "%s=\"%s\"\n"
)

var osr *OsRelease
var printEqualFormat string

func init() {
	osr = NewOsInfo()
	osr.Load()

}

// Default is used to access the default OsRelease object
// Shouldn't be needed regularly
func Default() *OsRelease { return osr }

// Returns values set in the os-release file
func IsLike(name string) bool { return osr.IsLike(name) }
func Pretty() string          { return osr.Pretty() }
func PrettyVariant() string   { return osr.PrettyVariant() }
func PrettyVersion() string   { return osr.PrettyVersion() }
func Name() string            { return osr.Name() }
func Id() string              { return osr.Id() }
func Version() string         { return osr.Version() }
func BuildId() string         { return osr.BuildId() }
func Arch() string            { return osr.Arch() }
func ImageId() string         { return osr.ImageId() }
func ImageVersion() string    { return osr.ImageVersion() }

// PrintSet prints all of the SET values, leaving out anything that is not set
// Will print in a shell consumable format if the arg for shell is true
func PrintSet() {
	if viper.GetBool("shell") {
		printEqualFormat = printShellEqualsFormat
	} else {
		printEqualFormat = printBasicEqualsFormat
	}

	printIfSet(osr.name, NAME)
	printIfSet(osr.id, ID)
	printIfSet(osr.idLike, ID_LIKE)
	printIfSet(osr.prettyName, PRETTY_NAME)
	printIfSet(osr.cpeName, CPE_NAME)
	printIfSet(osr.variant, VARIANT)
	printIfSet(osr.variantId, VARIANT_ID)
	printIfSet(osr.version, VERSION)
	printIfSet(osr.versionId, VERSION_ID)
	printIfSet(osr.versionCodename, VERSION_CODENAME)
	printIfSet(osr.buildId, BUILD_ID)
	printIfSet(osr.imageId, IMAGE_ID)
	printIfSet(osr.imageVersion, IMAGE_VERSION)
	printIfSet(osr.homeUrl, HOME_URL)
	printIfSet(osr.documentationUrl, DOCUMENTATION_URL)
	printIfSet(osr.supportUrl, SUPPORT_URL)
	printIfSet(osr.bugReportUrl, BUG_REPORT_URL)
	printIfSet(osr.privacyPolicyUrl, PRIVACY_POLICY_URL)
	printIfSet(osr.supportEnd, SUPPORT_END)
	printIfSet(osr.logo, LOGO)
	printIfSet(osr.ansiColor, ANSI_COLOR)
	printIfSet(osr.vendorName, VENDOR_NAME)
	printIfSet(osr.vendorUrl, VENDOR_URL)
	printIfSet(osr.defaultHostname, DEFAULT_HOSTNAME)
	printIfSet(osr.architecture, ARCHITECTURE)
	printIfSet(osr.sysextLevel, SYSEXT_LEVEL)
	printIfSet(osr.confextLevel, CONFEXT_LEVEL)
	printIfSet(osr.sysextScope, SYSEXT_SCOPE)
	printIfSet(osr.confextScope, CONFEXT_SCOPE)
	printIfSet(osr.portablePrefixes, PORTABLE_PREFIXES)
	fmt.Println("###\n### THE REST ARE NOT OFFICIALLY SUPPORTED ###\n###")
	printAllExtra(osr)
}

func printIfSet(v, n string) {
	if v != "" {
		fmt.Printf(printEqualFormat, n, v)
	}
}

func printAllExtra(i *OsRelease) {
	for k, v := range i.unknown {
		fmt.Printf(printEqualFormat, k, v)
	}
}
