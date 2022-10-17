package scanner

import (
	"fmt"
	"github.com/dsm0014/typo-scanner/typo"
	"github.com/dsm0014/typo-scanner/util"
	"log"
)

var logger *log.Logger

type scanFn func(string, typo.GeneratorFlags) ([]string, error)

func TypoScanEntrypoint(fn scanFn, args []string, genFlags typo.GeneratorFlags) error {
	logger = util.GetLogger(genFlags.SuppressLogs)
	for _, pkg := range args {
		matches, err := fn(pkg, genFlags)
		if err != nil {
			return err
		}
		fmt.Println(matches)
	}
	return nil
}
