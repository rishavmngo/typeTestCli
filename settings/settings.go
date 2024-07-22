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
	Mode            int
	CursorCharacter string
}

var settings *Settings

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

	return settings

}

func (settings *Settings) Load() *Settings {
	data, err := os.ReadFile(settingsPath)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(data), &settings)
	return settings
}

func (settings *Settings) Write() {

	data, err := json.Marshal(settings)
	if err = os.WriteFile(settingsPath, data, 0644); err != nil {
		log.Fatalf("failed to write  file: %s", err)
	}
}

func (settings *Settings) GetWords() []string {

	data, err := os.ReadFile(wordsPath)
	if err != nil {
		panic(err)

	}
	var samples []SampleWord
	json.Unmarshal(data, &samples)
	words := strings.Split(samples[settings.Mode].Text, samples[settings.Mode].Delimiter)
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}
	return words
}
