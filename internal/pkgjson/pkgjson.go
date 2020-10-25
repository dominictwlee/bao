package pkgjson

//var (
//	BaseDevDeps = [...]string{"bao"}
//)

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
