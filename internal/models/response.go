package models

type Response []struct {
	Word     string `json:"word"`
	Phonetic string `json:"phonetic"`

	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`

		Definitions []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
			Synonyms   []string `json:"synonyms"`
			Antonyms   []string `json:"antonyms"`
		} `json:"definitions"`

	} `json:"meanings"`
}
