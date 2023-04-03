package i18n

import (
	"embed"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed languages/*
var fs embed.FS

func NewBundle() (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(language.Chinese)
	englishFilename := "active.en.toml"
	chineseFilename := "active.zh.toml"
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	_, err := bundle.LoadMessageFileFS(fs, "languages/"+englishFilename)
	if err != nil {
		log.Fatalf("load languages pack failed: %s", err)
		return nil, err
	}
	_, err = bundle.LoadMessageFileFS(fs, "languages/"+chineseFilename)
	if err != nil {
		log.Fatalf("load languages pack failed: %s", err)
		return nil, err
	}
	return bundle, nil
}

func NewLocalizer(bundle *i18n.Bundle, lang string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, lang)
}
