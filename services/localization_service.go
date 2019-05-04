package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Translation - localization structure
type translation struct {
	locales      []string
	translations map[string]map[string]string
}

var trans *translation

// InitLocales - initiate locales from the folder
func InitLocales(trPath_ string) error {
	trans = &translation{translations: make(map[string]map[string]string)}
	return loadTranslations(trPath_)
}

// Tr - translate for current locale
func Translate(locale_ string, trKey_ string) string {
	trValue, ok := trans.translations[locale_][trKey_]
	if ok {
		return trValue
	}
	trValue, ok = trans.translations["en"][trKey_]
	if ok {
		return trValue
	}
	return trKey_
}

// DetectLanguage - parse to find the most preferable language
func DetectLanguage(acceptLanguage_ string) string {

	langStrs := strings.Split(acceptLanguage_, ",")
	for _, langStr := range langStrs {
		lang := strings.Split(strings.Trim(langStr, " "), ";")
		if checkLocale(lang[0]) {
			return lang[0]
		}
	}

	return "en"
}

// LoadTranslations - load translations files from the folder
func loadTranslations(trPath_ string) error {
	files, err := filepath.Glob(trPath_ + "/*.json")
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("No translations found")
	}

	for _, file := range files {
		err := loadFileToMap(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadFileToMap(filename_ string) error {
	var objmap map[string]string

	localName := strings.Replace(filepath.Base(filename_), ".json", "", 1)

	content, err := ioutil.ReadFile(filename_)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &objmap)
	if err != nil {
		return err
	}
	trans.translations[localName] = objmap
	trans.locales = append(trans.locales, localName)
	return nil
}

func checkLocale(localeName_ string) bool {
	for _, locale := range trans.locales {
		if locale == localeName_ {
			return true
		}
	}
	return false
}