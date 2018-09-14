package utility

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//
// This contains miscellaneous functions for use in other packages in the Web project
//

// ------------------------------------------- Public ------------------------------------------- //

// Given a map src, copies all of its elements into dest
func CopyStringMap(src map[string]string, dest *map[string]string) *map[string]string {
	for key, element := range src {
		(*dest)[key] = element
	}
	return dest
}

// Reads a file and returns lines as array of strings
func ReadHttpFile(dir, file string) []string {

	arr := []string{}
	d := http.Dir(dir)
	f, err := d.Open(file)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(f)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("read file line error: %v", err)
				return []string{}
			}
		}
		arr = append(arr, string(line))
	}
	return arr
}

func DomainFromUrl(urlString string) string {
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
	return domain
}
