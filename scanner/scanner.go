package scanner

import (
	"fmt"
	"github.com/dsm0014/typo-scanner/typo"
	"io"
	"net/http"
	"strings"
	"sync"
)

func ScanGo(original string, flags typo.GeneratorFlags) ([]string, error) {
	return Scan(Godev, GodevUrl, original, flags)
}

func ScanMvn(original string, flags typo.GeneratorFlags) ([]string, error) {
	return Scan(Mvn, MvnUrl, original, flags)
}

func ScanNpm(original string, flags typo.GeneratorFlags) ([]string, error) {
	return Scan(Npm, NpmUrl, original, flags)
}

func ScanPypi(original string, flags typo.GeneratorFlags) ([]string, error) {
	return Scan(Pypi, PypiUrl, original, flags)
}

func ScanRuby(original string, flags typo.GeneratorFlags) ([]string, error) {
	return Scan(Ruby, RubyUrl, original, flags)
}

func Scan(pkgType PkgType, pkgUrl PkgUrl, original string, flags typo.GeneratorFlags) ([]string, error) {
	logger.Printf("Scanning %s for typos of: %s", pkgType, original)
	var matches []string
	typos, err := typo.TypoGenerator(original, flags)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(typos))
	for _, t := range typos {
		go ScanTypoRoutine(&wg, pkgType, pkgUrl, &matches, t)
	}
	wg.Wait()
	logger.Printf("%d typos scanned in %s", len(typos), pkgType)
	if len(matches) > 0 {
		logger.Printf("%d matches found in %s", len(matches), pkgType)
		logger.Printf(fmt.Sprintf("Potential TypoSquatters: %s", matches))
		return matches, nil
	}
	logger.Printf("SUCCESS: %d TypoSquatters found in %s!", len(matches), pkgType)
	return matches, nil
}

func ScanTypoRoutine(wg *sync.WaitGroup, pkgType PkgType, pkgUrl PkgUrl, matches *[]string, t string) {
	defer wg.Done()
	resp, err := http.Get(fmt.Sprintf("%s%s", pkgUrl, t))
	if err != nil {
		logger.Printf("Error looking up NPM package: %s", err)
		return
	}
	defer resp.Body.Close()

	// Early exit if the page 404's
	if resp.StatusCode == http.StatusNotFound {
		logger.Printf("No results in %s for: %s", pkgType, t)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err)
		return
	}
	bodyString := string(bodyBytes)

	// Additional checks which may indicate/disprove TypoSquatting
	switch pkgType {
	case Npm:
		// Sometimes packages that were previously tagged as malicious will have the URL lying around but not be active
		if strings.Contains(bodyString, "You may adopt this package by contacting") {
			logger.Printf("Previous TypoSquatter found in %s using the name: %s\n", pkgType, t)
			return
		}
	case Mvn:
		// Maven is weird
		if resp.StatusCode == http.StatusForbidden {
			logger.Printf("No results in %s for: %s", pkgType, t)
			return
		}
	default:
	}

	logger.Printf("Active TypoSquatter found in %s using the name: %s\n", pkgType, t)
	*matches = append(*matches, t)
}
