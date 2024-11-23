package osrelease

// package level access methods, to be used as
// `osrelease.Id()` to retreive an Id, which calls the
// `Id()` method on the default osrelease object

var osr *OsRelease

func init() {
	osr = NewOsInfo()
	osr.Load()
}

// Default is used to access the default OsRelease object
// Shouldn't be needed regularly
func Default() *OsRelease { return osr }

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
