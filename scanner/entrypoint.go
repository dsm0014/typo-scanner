package scanner

import "github.com/dsm0014/typo-scanner/typo"

type scanFn func(string, typo.GeneratorFlags) ([]string, error)

func TypoScanEntrypoint(fn scanFn, args []string, genFlags typo.GeneratorFlags) error {
	for _, pkg := range args {
		_, err := fn(pkg, genFlags)
		if err != nil {
			return err
		}
	}
	return nil
}
