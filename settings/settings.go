package settings

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SampleWord struct {
	Text      string
	Delimiter string
}
type Settings struct {
	Duration        int
	Mode            string
	CursorCharacter string
}

var settings *Settings

var wordMap map[string]SampleWord

func getPaths() (string, string) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return "", ""
	}

	configDir := filepath.Join(homeDir, ".config/typeTest-go")
	wordsPath := filepath.Join(configDir, "words.json")

	settingsPath := filepath.Join(configDir, "settings.json")
	return settingsPath, wordsPath
}

var settingsPath, wordsPath = getPaths()

func Get() *Settings {

	if settings != nil {

		return settings

	}

	settings = &Settings{}
	settings.Load()

	return settings

}

func (settings *Settings) Load() *Settings {
	data, err := os.ReadFile(settingsPath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(data), &settings)
	settings.LoadWordMap()
	return settings
}

func (settings *Settings) Save() {

	data, err := json.Marshal(settings)
	if err = os.WriteFile(settingsPath, data, 0644); err != nil {
		log.Fatalf("failed to write  file: %s", err)
	}
}

func (settings *Settings) LoadWordMap() {

	data, err := os.ReadFile(wordsPath)

	if err != nil {
		panic(err)

	}

	json.Unmarshal(data, &wordMap)
}

func (settings *Settings) GetModeList() []string {

	keys := make([]string, 0, len(wordMap))

	for k := range wordMap {
		keys = append(keys, k)
	}

	return keys
}

func (settings *Settings) GetWords() []string {

	words := strings.Split(wordMap[settings.Mode].Text, wordMap[settings.Mode].Delimiter)
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}
	return words
}
