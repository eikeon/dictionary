package dictionary

import (
	"http"
	"log"
	"regexp"
	"io/ioutil"
	"path"
	"strings"
)


type Word struct {
	word        string
	wikitext    string
}

func getWikitext(word string) string {
	cache_file := path.Join("wiktionary", word)
	buf, err := ioutil.ReadFile(cache_file)
	if err!=nil {
		log.Print("Getting wikitext from wiktionary for: " + word)
		var _URL = "http://en.wiktionary.org/w/api.php"		
		URL := _URL + "?action=parse&page="+http.URLEscape(word)+"&prop=wikitext&format=json"
		r, err := http.Get(URL)
		if err != nil {
			log.Print(err)
		}
		b, err := ioutil.ReadAll(r.Body)
		write_err := ioutil.WriteFile(cache_file, b, 0666)
		if write_err != nil {
			log.Print("Error writing cache file for: " + word)
		}
		return string(b)
	}
	return string(buf)
}

func PartsOfSpeech(word string) {
	log.Print("Finding parts of speech for: " + word)
	s := getWikitext(word)
	template := regexp.MustCompile("{{([^}\\|]+)(\\|([^}]+))?}}")
	templates := template.FindAllStringSubmatch(s, -1)
	for _, t := range templates {
		if strings.HasPrefix(t[1], "en-") {
			log.Print(t[1])
		}
	}
}
