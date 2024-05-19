package locale

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func LoadMessageFile(locale string) (*i18n.MessageFile, error) {
	if !validLocale(locale) {
		return nil, fmt.Errorf("locale %s does not exist", locale)
	}
	return bundle.LoadMessageFile(fmt.Sprintf("locale/%s.json", locale))
}
func validLocale(locale string) bool {
	_, ok := Locales[locale]
	return ok
}
