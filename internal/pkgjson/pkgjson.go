package pkgjson

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

const (
	invalidLeadingCharRe = "^\\.|_"
)

var (
	scopedPkgRe = regexp.MustCompile("^(?:@([^/]+?)[/])?([^/]+?)$")
	BaseDevDeps = [...]string{"bao"}
	builtins    = [...]string{
		"assert",
		"buffer",
		"child_process",
		"cluster",
		"console",
		"constants",
		"crypto",
		"dgram",
		"dns",
		"domain",
		"events",
		"fs",
		"http",
		"https",
		"module",
		"net",
		"os",
		"path",
		"punycode",
		"querystring",
		"readline",
		"repl",
		"stream",
		"string_decoder",
		"sys",
		"timers",
		"tls",
		"tty",
		"url",
		"util",
		"vm",
		"zlib",
		"perf_hooks",
		"http2",
		"async_hooks",
		"process",
		"v8",
		"freelist",
	}
)

type PackageJSON struct {
	Name    string
	Author  string
	Version string
	License string
	Main    string
	Module  string
	Typings string
	Files   []string
	Engines struct {
		Node string
	}
	Scripts struct {
		Start string
		Build string
	}
}

func IsValidName(name string) (bool, error) {
	if len(name) == 0 {
		return false, errors.New("name length must be greater than zero")
	}

	if len(name) > 214 {
		return false, errors.New("name must be less than or equal to 214 characters")
	}

	if strings.ToLower(name) != name {
		return false, errors.New("name must not have uppercase letters")
	}

	for _, n := range builtins {
		if name == n {
			return false, errors.New("cannot use same name as a core Node module")
		}
	}

	matched, err := regexp.Match(invalidLeadingCharRe, []byte(name))
	if err != nil {
		return false, err
	}
	if matched {
		return false, errors.New("name cannot start with a dot or an underscore")
	}

	subMatches := scopedPkgRe.FindStringSubmatch(name)
	urlErrMsg := "scoped package can only contain URL-friendly characters"
	if len(subMatches) == 3 {
		scope := subMatches[1]
		pkg := subMatches[2]
		if url.PathEscape(scope) != scope && url.PathEscape(pkg) != pkg {
			return false, errors.New(urlErrMsg)
		}
	} else if url.PathEscape(name) != name {
		return false, errors.New(urlErrMsg)
	}

	return true, nil
}
