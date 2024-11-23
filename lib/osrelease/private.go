package osrelease

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/stobbsm/go-osrelease/lib/checkfile"
)

// set sets the given key to the given value
// if the key has an official entry in the standard
// it is set with a named struct entry,
// while all other keys are set in the 'unknown' map
// Return true if the key is officially supported
func (i *OsRelease) set(key, value string) bool {
	switch key {
	case NAME:
		i.name = value
		return true
	case ID:
		i.id = value
		return true
	case ID_LIKE:
		i.idLike = value
		return true
	case PRETTY_NAME:
		i.prettyName = value
		return true
	case CPE_NAME:
		i.cpeName = value
		return true
	case VARIANT:
		i.variant = value
		return true
	case VARIANT_ID:
		i.variantId = value
		return true
	case VERSION:
		i.version = value
		return true
	case VERSION_ID:
		i.versionId = value
		return true
	case VERSION_CODENAME:
		i.versionCodename = value
		return true
	case BUILD_ID:
		i.buildId = value
		return true
	case IMAGE_ID:
		i.imageId = value
		return true
	case IMAGE_VERSION:
		i.imageVersion = value
		return true
	case HOME_URL:
		i.homeUrl = value
		return true
	case DOCUMENTATION_URL:
		i.documentationUrl = value
		return true
	case SUPPORT_URL:
		i.supportUrl = value
		return true
	case BUG_REPORT_URL:
		i.bugReportUrl = value
		return true
	case PRIVACY_POLICY_URL:
		i.privacyPolicyUrl = value
		return true
	case SUPPORT_END:
		i.supportEnd = value
		return true
	case LOGO:
		i.logo = value
		return true
	case ANSI_COLOR:
		i.ansiColor = value
		return true
	case VENDOR_NAME:
		i.vendorName = value
		return true
	case VENDOR_URL:
		i.vendorUrl = value
		return true
	case DEFAULT_HOSTNAME:
		i.defaultHostname = value
		return true
	case ARCHITECTURE:
		i.architecture = value
		return true
	case SYSEXT_LEVEL:
		i.sysextLevel = value
		return true
	case CONFEXT_LEVEL:
		i.confextLevel = value
		return true
	case SYSEXT_SCOPE:
		i.sysextScope = value
		return true
	case CONFEXT_SCOPE:
		i.confextScope = value
		return true
	case PORTABLE_PREFIXES:
		i.portablePrefixes = value
		return true
	default:
		i.unknown[key] = value
		return false
	}
}

// get takes the key to retrieve and if only the official list should
// be checked
// returns the found value and nil, or an empty string and an error
func (i *OsRelease) get(key string, officialOnly bool) (string, error) {
	switch strings.ToUpper(key) {
	case NAME:
		return i.name, nil
	case ID:
		return i.id, nil
	case ID_LIKE:
		return i.idLike, nil
	case PRETTY_NAME:
		return i.prettyName, nil
	case CPE_NAME:
		return i.cpeName, nil
	case VARIANT:
		return i.variant, nil
	case VARIANT_ID:
		return i.variantId, nil
	case VERSION:
		return i.version, nil
	case VERSION_ID:
		return i.versionId, nil
	case VERSION_CODENAME:
		return i.versionCodename, nil
	case BUILD_ID:
		return i.buildId, nil
	case IMAGE_ID:
		return i.imageId, nil
	case IMAGE_VERSION:
		return i.imageVersion, nil
	case HOME_URL:
		return i.homeUrl, nil
	case DOCUMENTATION_URL:
		return i.documentationUrl, nil
	case SUPPORT_URL:
		return i.supportUrl, nil
	case BUG_REPORT_URL:
		return i.bugReportUrl, nil
	case PRIVACY_POLICY_URL:
		return i.privacyPolicyUrl, nil
	case SUPPORT_END:
		return i.supportEnd, nil
	case ANSI_COLOR:
		return i.ansiColor, nil
	case VENDOR_NAME:
		return i.vendorName, nil
	case VENDOR_URL:
		return i.vendorUrl, nil
	case DEFAULT_HOSTNAME:
		return i.defaultHostname, nil
	case ARCHITECTURE:
		return i.architecture, nil
	case SYSEXT_LEVEL:
		return i.sysextLevel, nil
	case CONFEXT_LEVEL:
		return i.confextLevel, nil
	case SYSEXT_SCOPE:
		return i.sysextScope, nil
	case CONFEXT_SCOPE:
		return i.confextScope, nil
	case PORTABLE_PREFIXES:
		return i.portablePrefixes, nil

	default:
		if officialOnly {
			return "", errors.New("unknown key")
		}
		if v, ok := i.unknown[key]; ok {
			return v, nil
		}
	}
	return "", errors.New("undefined key")
}

// load checks if readFrom is already set, and then reads
// and parses the file using loadFile
// If readFrom isn't set (normal case), checks
// ETC_OS_RELEASE then USR_LIB_OS_RELEASE.
// if neither of those exist, panic
func (i *OsRelease) load() error {
	if i.readFrom == "" {
		if checkfile.Exists(ETC_OS_RELEASE) {
			i.readFrom = ETC_OS_RELEASE
		} else if checkfile.Exists(USR_LIB_OS_RELEASE) {
			i.readFrom = USR_LIB_OS_RELEASE
		} else {
			panic(errors.New("unable to read os-release from any given file"))
		}
	}
	return i.loadFile(i.readFrom)
}

// osRelease reads all lines from ETC_OS_RELEASE and parses
// them info an OsInfo struct, containing the parsed details
// identifying the OS in use
func (i *OsRelease) loadFile(filename string) error {
	var lines []string
	var err error

	lines, err = readFile(filename)
	if err != nil {
		return err
	}

	for _, l := range lines {
		k, v, err := parseLine(l)
		if err != nil {
			// errors here are not considered fatal, skip and parse the next line
			continue
		}
		i.Set(k, v)
	}

	return nil
}

// readFile reads the contents of the given file and returns
// a slice of the lines read
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// parseLinux takes a line extracted by readFile
// and attempts to parse it, returning the name, value, or an error
// if parsing failed
func parseLine(line string) (string, string, error) {
	if len(line) == 0 {
		return "", "", errors.New("skipping zero-length")
	}

	if line[0] == '#' {
		return "", "", errors.New("skipping comment")
	}

	splitString := strings.SplitN(line, "=", 2)
	if len(splitString) != 2 {
		return "", "", errors.New("can not extract key=value")
	}

	key := strings.Trim(splitString[0], " ")
	value := strings.Trim(splitString[1], " ")

	// remove double quotes in value
	if strings.ContainsAny(value, `"'`) {
		value = strings.TrimPrefix(value, `'`)
		value = strings.TrimPrefix(value, `"`)
		value = strings.TrimSuffix(value, `'`)
		value = strings.TrimSuffix(value, `"`)
	}

	// expand escaped values
	value = strings.Replace(value, `\"`, `"`, -1)
	value = strings.Replace(value, `\$`, `$`, -1)
	value = strings.Replace(value, `\\`, `\`, -1)
	value = strings.Replace(value, "\\`", "`", -1)

	return key, value, nil
}

// parseLike takes the space separated list of 'alike'
// distributions (parents like rhel, debian, etc) and
// puts them into an easy to use map
func parseLike(like string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range strings.Split(like, " ") {
		m[v] = struct{}{}
	}
	return m
}
