package locale

import (
	"encoding/json"
	"sync"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
)

const (
	EnLanguage = "en"
	ViLanguage = "vi"
	Fallback   = EnLanguage
)

var (
	once      sync.Once
	localizer *Localizer
	bundle    *i18n.Bundle
	Locales   = map[string]language.Tag{
		EnLanguage: language.English,
		ViLanguage: language.Vietnamese,
	}
)

type Localizer struct {
	*i18n.Localizer
}

func NewLocalizer() *Localizer {
	once.Do(func() {
		bundle = i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
		for lang := range Locales {
			_, err := LoadMessageFile(lang)
			if err != nil {
				logx.Error(err)
			}
		}
		localizer = &Localizer{i18n.NewLocalizer(bundle, EnLanguage)}
	})
	return localizer
}
func (l *Localizer) Register() {
	// Implement core service interface
}
