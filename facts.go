package facts

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// FactsLib - main object og the library,
// will returned after linrary initialization
type FactsLib struct {
	keywords [][]string
	facts    map[string][]string
}

// GetFact - return random fact by provided keyword
// Return errors
// 		- keyword's length validation (min length is 2)
//		- fact not found
func (fl *FactsLib) GetFact(keyword string) (string, error) {
	if len(keyword) < 2 {
		return "", fmt.Errorf("Keyword's length must be equal or more than 2 symbols")
	}

	for _, kw := range fl.keywords {
		for _, k := range kw {
			if strings.ToLower(k) == strings.ToLower(keyword) {
				hash := calculateHash(kw)
				return getRandomFact(fl.facts[hash]), nil
			}
		}
	}
	return "", fmt.Errorf("Fact not found")
}

// FindFact - return random fact if provided text message contains any of available keywords
// Retirn random fact for first keyword found
// Return errors
// 		- provided text length validation (cannot be empty)
//		- fact not found
func (fl *FactsLib) FindFact(text string) (string, error) {
	if len(text) == 0 {
		return "", fmt.Errorf("Provided text is empty")
	}

	for _, kw := range fl.keywords {
		if containgKeywords(text, kw) {
			hash := calculateHash(kw)
			return getRandomFact(fl.facts[hash]), nil
		}
	}
	return "", fmt.Errorf("Fact not found")
}

// Init the library with provided path to folder with files with facts
// Collect whole keywords and facts, create and return library object
func Init(path string) (*FactsLib, error) {
	files, err := getFilesList(path)
	if err != nil {
		return nil, err
	}

	keywords := [][]string{}
	facts := map[string][]string{}

	for _, file := range files {
		k, f, err := parseFile(path, file)
		if err == nil && len(k) > 0 && len(f) > 0 {
			hash := calculateHash(k)
			keywords = append(keywords, k)
			facts[hash] = f
		}
	}

	if len(keywords) == 0 || len(facts) == 0 {
		return nil, fmt.Errorf("Empty facts list")
	}

	return &FactsLib{
		keywords: keywords,
		facts:    facts}, nil
}

// Read files list in provided path to folder
// Return list of file names or error
func getFilesList(path string) ([]string, error) {
	const maxFileSize = 100000 // in bytes ~100kb
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	if len(files) == 0 {
		return []string{}, fmt.Errorf("Empty files list")
	}

	filePaths := []string{}
	for _, file := range files {
		if !file.IsDir() && file.Size() < maxFileSize {
			filePaths = append(filePaths, file.Name())
		}
	}
	return filePaths, nil
}

// Parse file located in `path + / + file` location:
// 		- parse first line as list of keywords
//		- parse lines 2-... as facts (one fact per line)
// Returns:
// 		- first return value - is slice of keywords parsed from file
// 		- second return value - is slice of facts parsed from file
//		- third return value - is parsing error
func parseFile(path string, file string) ([]string, []string, error) {
	const maxKeywordsNumber = 10
	const maxFactsNumber = 100

	keywords := []string{}
	facts := []string{}

	filePath := path + "/" + file
	f, err := os.Open(filePath)
	if err != nil {
		return keywords, facts, fmt.Errorf("Error opening file %s: %s", filePath, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	firstLineParsed := false
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if len(strings.Trim(line, " ")) > 0 {
			if firstLineParsed && len(facts) < maxFactsNumber {
				facts = append(facts, line) // Add line as fact
			} else {
				// parse keywords from first line
				line = strings.Replace(line, "keywords:", "", 1) //remove field name keywords:
				keywords = strings.Split(line, ",")
				for idx, keyword := range keywords {
					keywords[idx] = strings.Trim(keyword, " ")
				}
			}
			firstLineParsed = true
		}
	}
	// fmt.Printf("Scan file %s. Got key words %d and %d facts.\n", filePath, len(keywords), len(facts))
	return keywords, facts, nil
}

// Calculate hash for list of keywords
// Usage to save and retrieve facts by keywords list
func calculateHash(s []string) string {
	return strings.ToLower(strings.Join(s[:], "_"))
}

// Search provided keywords as whole words in the provided message
func containgKeywords(message string, keywords []string) bool {
	words := strings.ToLower(strings.Join(keywords[:], "|"))
	r := regexp.MustCompile(`\b(` + words + `)\b.*`)
	return len(r.FindString(strings.ToLower(message))) > 0
}

// Get random fact from provided slice of facts
func getRandomFact(facts []string) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	if facts != nil && len(facts) > 0 {
		i := r1.Intn(len(facts) - 1)
		return facts[i]
	}
	return ""
}
