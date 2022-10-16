# typo-scanner
`typo-scanner` is a quick and simple CLI to help you scan public registries for TypoSquatters trying to attack your package's userbase.

## What is TypoSquatiing? 
TypoSquatting is when malicious developers upload copies of software packages with typos in the package name. 

These developers then include nefarious payloads in the otherwise functional copy, and clumsy fingers everywhere 
are vulnerable when attempting to use open source software.

## Supported Registries
The following public registries are currently supported:
- go.dev
- mvnrepository
- npmjs
- pypi
- rubygems

## Getting Started
Install from source
```
git clone https://github.com/dsm0014/typo-scanner.git
go build
```

Scan some package registries!<br>
`./typo-scanner npm react-dom -dr`<br>
The `-dr` specified above will search for [d]uplicate and [r]eversed character typos.

For more details on flags and command options explore the detailed help commands.
```
./typo-scanner --help
./typo-scanner [command] --help
```
