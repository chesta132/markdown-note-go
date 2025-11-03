package grammar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
)

var (
	AllowedLanguages = []string{
		"en-US", // English (US)
		"en-GB", // English (UK)
		"id",    // Indonesian
		"de-DE", // German
		"fr",    // French
		"es",    // Spanish
		"pt-BR", // Portuguese (Brazil)
		"pt-PT", // Portuguese (Portugal)
		"nl",    // Dutch
		"pl-PL", // Polish
		"ru-RU", // Russian
		"it",    // Italian
		"ca-ES", // Catalan
		"zh-CN", // Chinese
		"ja",    // Japanese
		"sv",    // Swedish
		"uk-UA", // Ukrainian
		"ro-RO", // Romanian
		"sk-SK", // Slovak
		"da-DK", // Danish
		"el-GR", // Greek
	}
)

type Match struct {
	Message      string `json:"message"`
	Offset       int    `json:"offset"`
	Length       int    `json:"length"`
	Replacements []struct {
		Value string `json:"value"`
	} `json:"replacements"`
}

type LTResponse struct {
	Matches []Match `json:"matches"`
}

func IsLanguageAllowed(lang string) bool {
	for _, allowed := range AllowedLanguages {
		if allowed == lang {
			return true
		}
	}
	return false
}

func FixGrammar(text string, lang string) (string, error) {
	if !IsLanguageAllowed(lang) {
		return "", fmt.Errorf("language '%s' not allowed. supported languages: %v", lang, AllowedLanguages)
	}

	apiURL := "https://api.languagetool.org/v2/check"

	data := url.Values{}
	data.Set("text", text)
	data.Set("language", lang)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return "", fmt.Errorf("error calling API: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var ltResp LTResponse
	if err := json.Unmarshal(body, &ltResp); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	sort.Slice(ltResp.Matches, func(i, j int) bool {
		return ltResp.Matches[i].Offset > ltResp.Matches[j].Offset
	})

	corrected := text
	for _, match := range ltResp.Matches {
		if len(match.Replacements) > 0 {
			replacement := match.Replacements[0].Value
			start := match.Offset
			end := start + match.Length

			corrected = corrected[:start] + replacement + corrected[end:]
		}
	}

	return corrected, nil
}
