package repo

import "github.com/BeardyC/draft/pkg/version"

// Builtin contains metadata to the built-in packs. Used to install/uninstall a pack.
type Builtin struct {
	Name    string
	URL     string
	Version string
}

// Builtins fetches all built-in pack repositories as a map of url=>ver.
func Builtins() []*Builtin {
	var ver string
	// canary draft releases should always test the latest version of the packs.
	if version.Release != "canary" {
		ver = version.Release
	}
	return []*Builtin{
		{
			Name:    "github.com/BeardyC/draft",
			URL:     "https://github.com/BeardyC/draft",
			Version: ver,
		},
	}
}
