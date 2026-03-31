package i18n

import (
	"embed"
	"encoding/json"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var localeFS embed.FS
var Bundle *i18n.Bundle

func Init() {
	Bundle = i18n.NewBundle(language.English) // default language
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	loadLocale(Bundle, "locales/en.json")
	loadLocale(Bundle, "locales/id.json")
}

func loadLocale(bundle *i18n.Bundle, path string) {
	data, err := localeFS.ReadFile(path)
	if err != nil {
		panic(err)
	}

	bundle.MustParseMessageFileBytes(data, path)
}

func Translate(lang, messageID string, templateData ...map[string]any) string {
	localizer := i18n.NewLocalizer(Bundle, lang)

	config := &i18n.LocalizeConfig{
		MessageID: messageID,
	}

	// only set if provided
	if len(templateData) > 0 {
		config.TemplateData = templateData[0]
	}

	msg, err := localizer.Localize(config)
	if err == nil {
		return msg
	}

	parts := strings.Split(messageID, ".")
	messageDefault := parts[0] + ".DEFAULT"

	configDefault := &i18n.LocalizeConfig{
		MessageID: messageDefault,
	}

	msgDefault, err := localizer.Localize(configDefault)
	if err == nil {
		return msgDefault
	}

	return parts[1]
}
