package scanner

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ScanGo(original string, typos []string) ([]string, error) {
	return Scan(Godev, GodevUrl, original, typos)
}

func ScanMvn(original string, typos []string) ([]string, error) {
	return Scan(Mvn, MvnUrl, original, typos)
}

func ScanNpm(original string, typos []string) ([]string, error) {
	return Scan(Npm, NpmUrl, original, typos)
}

func ScanPypi(original string, typos []string) ([]string, error) {
	return Scan(Pypi, PypiUrl, original, typos)
}

func ScanRuby(original string, typos []string) ([]string, error) {
	return Scan(Ruby, RubyUrl, original, typos)
}

func Scan(pkgType PkgType, pkgUrl PkgUrl, original string, typos []string) ([]string, error) {
	log.Printf("Scanning %s for typos of: %s", pkgType, original)
	var matches []string
	for _, typo := range typos {
		resp, err := http.Get(fmt.Sprintf("%s%s", pkgUrl, typo))
		if err != nil {
			log.Printf("Error looking up NPM package: %s", err)
		}
		defer resp.Body.Close()
		
		// Early exit if the page 404's
		if resp.StatusCode == http.StatusNotFound {
			log.Printf("No results in %s for: %s", pkgType, typo)
			continue
		}
		
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		bodyString := string(bodyBytes)

		// Additional checks which may indicate/disprove TypoSquatting
		switch pkgType {
		case Npm:
			// Sometimes packages that were previously tagged as malicious will have the URL lying around but not be active
			if strings.Contains(bodyString, "You may adopt this package by contacting") {
				log.Printf("Previous TypoSquatter found in %s using the name: %s\n", pkgType, typo)
				continue
			}
		case Mvn:
			// Maven is weird
			if resp.StatusCode == http.StatusForbidden {
				log.Printf("No results in %s for: %s", pkgType, typo)
				continue
			}
		default:
		}

		log.Printf("Active TypoSquatter found in %s using the name: %s\n", pkgType, typo)
		matches = append(matches, typo)
	}
	if len(matches) > 0 {
		log.Printf("%d matches found in %s", len(matches), pkgType)
		log.Fatal(fmt.Sprintf("Potential TypoSquatters: %s", matches))
		return matches, nil
	}
	log.Printf("SUCCESS: %d TypoSquatters found in %s!", len(matches), pkgType)
	return matches, nil
}
