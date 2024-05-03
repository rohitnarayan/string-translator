package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Translator handles translation operations
type Translator struct {
	bundle     *i18n.Bundle
	localizer  *i18n.Localizer
	language   string
	configPath string
}

// NewTranslator creates a new Translator instance
func NewTranslator(language, configPath string) (*Translator, error) {
	bundle := i18n.NewBundle(languageTag(language))
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	if _, err := bundle.LoadMessageFile(configPath); err != nil {
		return nil, err
	}

	localizer := i18n.NewLocalizer(bundle, languageTag(language).String())

	return &Translator{
		bundle:     bundle,
		localizer:  localizer,
		language:   language,
		configPath: configPath,
	}, nil
}

// languageTag converts language code to language tag
func languageTag(l string) language.Tag {
	return language.Make(l)
}

// Translate translates a messageID to the configured language
func (t *Translator) Translate(messageID string, templateData map[string]interface{}) (string, error) {
	return t.localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
}

func main() {
	// Initialize translator

	locale := "id"

	trans, err := NewTranslator(fmt.Sprintf("%s", locale), fmt.Sprintf("./resources/%s.json", locale))
	if err != nil {
		fmt.Println("Failed to initialize translator:", err)
		os.Exit(1)
	}

	// Example usage
	name := "Rohit"
	welcomeMessage, err := trans.Translate("welcome_message", map[string]interface{}{"Name": name})
	if err != nil {
		fmt.Println("Failed to translate welcome message:", err)
		os.Exit(1)
	}
	fmt.Println(welcomeMessage)
}
