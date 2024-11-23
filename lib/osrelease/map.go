package osrelease

// constants of os-release official entries
const (
	// Official fields
	NAME               = "NAME"
	ID                 = "ID"
	ID_LIKE            = "ID_LIKE"
	PRETTY_NAME        = "PRETTY_NAME"
	CPE_NAME           = "CPU_NAME"
	VARIANT            = "VARIANT"
	VARIANT_ID         = "VARIANT_ID"
	VERSION            = "VERSION"
	VERSION_ID         = "VERSION_ID"
	VERSION_CODENAME   = "VERSION_CODENAME"
	BUILD_ID           = "BUILD_ID"
	IMAGE_ID           = "IMAGE_ID"
	IMAGE_VERSION      = "IMAGE_VERSION"
	HOME_URL           = "HOME_URL"
	DOCUMENTATION_URL  = "DOCUMENTATION_URL"
	SUPPORT_URL        = "SUPPORT_URL"
	BUG_REPORT_URL     = "BUG_REPORT_URL"
	PRIVACY_POLICY_URL = "PRIVACY_POLICY_URL"
	SUPPORT_END        = "SUPPORT_END"
	LOGO               = "LOGO"
	ANSI_COLOR         = "ANSI_COLOR"
	VENDOR_NAME        = "VENDOR_NAME"
	VENDOR_URL         = "VENDOR_URL"
	DEFAULT_HOSTNAME   = "DEFAULT_HOSTNAME"
	ARCHITECTURE       = "ARCHITECTURE"
	SYSEXT_LEVEL       = "SYSEXT_LEVEL"
	CONFEXT_LEVEL      = "CONFEXT_LEVEL"
	SYSEXT_SCOPE       = "SYSEXT_SCOPE"
	CONFEXT_SCOPE      = "CONFEXT_SCOPE"
	PORTABLE_PREFIXES  = "PORTABLE_PREFIXES"

	// File paths to check
	ETC_OS_RELEASE     = `/etc/os-release`
	USR_LIB_OS_RELEASE = `/usr/lib/os-release`
)

// Names to match in IsLike to determine what is compatible
const (
	ARCH     = `arch`
	DEBIAN   = `debian`
	FEDORA   = `fedora`
	OPENSUSE = `opensuse`
	RHEL     = `rhel`
	SUSE     = `suse`
	UBUNTU   = `ubuntu`
)
