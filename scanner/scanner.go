package scanner

import (
	"fmt"
	"github.com/dsm0014/typo-scanner/typo"
	"io"
	"log"
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
	log.Printf("Scanning %s for typos of: %s", pkgType, original)
	var matches []string
	typos := typo.TypoGenerator(original, flags)
	var wg sync.WaitGroup
	wg.Add(len(typos))
	for _, t := range typos {
		go ScanTypoRoutine(&wg, pkgType, pkgUrl, &matches, t)
	}
	wg.Wait()
	log.Printf("%d typos scanned in %s", len(typos), pkgType)
	if len(matches) > 0 {
		log.Printf("%d matches found in %s", len(matches), pkgType)
		log.Fatal(fmt.Sprintf("Potential TypoSquatters: %s", matches))
		return matches, nil
	}
	log.Printf("SUCCESS: %d TypoSquatters found in %s!", len(matches), pkgType)
	return matches, nil
}

func ScanTypoRoutine(wg *sync.WaitGroup, pkgType PkgType, pkgUrl PkgUrl, matches *[]string, t string) {
	defer wg.Done()
	resp, err := http.Get(fmt.Sprintf("%s%s", pkgUrl, t))
	if err != nil {
		log.Fatalf("Error looking up NPM package: %s", err)
	}
	defer resp.Body.Close()

	// Early exit if the page 404's
	if resp.StatusCode == http.StatusNotFound {
		log.Printf("No results in %s for: %s", pkgType, t)
		return
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	// Additional checks which may indicate/disprove TypoSquatting
	switch pkgType {
	case Npm:
		// Sometimes packages that were previously tagged as malicious will have the URL lying around but not be active
		if strings.Contains(bodyString, "You may adopt this package by contacting") {
			log.Printf("Previous TypoSquatter found in %s using the name: %s\n", pkgType, t)
			return
		}
	case Mvn:
		// Maven is weird
		if resp.StatusCode == http.StatusForbidden {
			log.Printf("No results in %s for: %s", pkgType, t)
			return
		}
	default:
	}

	log.Printf("Active TypoSquatter found in %s using the name: %s\n", pkgType, t)
	*matches = append(*matches, t)
}
