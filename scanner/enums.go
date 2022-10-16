package scanner

type PkgType string
type PkgUrl string

const (
	Godev    PkgType = "Go"
	GodevUrl PkgUrl  = "https://pkg.go.dev/"

	Npm    PkgType = "NPM"
	NpmUrl PkgUrl  = "https://www.npmjs.com/package/"

	Mvn    PkgType = "Maven Repository"
	MvnUrl PkgUrl  = "https://mvnrepository.com/artifact/"

	Pypi    PkgType = "PyPi"
	PypiUrl PkgUrl  = "https://pypi.org/project/"

	Ruby    PkgType = "Ruby Gems"
	RubyUrl PkgUrl  = "https://rubygems.org/gems/"
)
