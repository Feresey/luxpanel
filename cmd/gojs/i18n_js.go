//go:build js

package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

//go:embed locales/*.yaml
var localeFiles embed.FS

type yamlCatalogEntry struct {
	ID string `yaml:"id"`
}

func (r *Runtime) InitI18n(ctx context.Context) error {
	_ = ctx
	ruBytes, err := localeFiles.ReadFile("locales/active.ru.yaml")
	if err != nil {
		return fmt.Errorf("read locales/active.ru.yaml: %w", err)
	}
	var ruEntries []yamlCatalogEntry
	if err := yaml.Unmarshal(ruBytes, &ruEntries); err != nil {
		return fmt.Errorf("parse active.ru.yaml: %w", err)
	}
	ids := make([]string, 0, len(ruEntries))
	for _, e := range ruEntries {
		if e.ID != "" {
			ids = append(ids, e.ID)
		}
	}
	r.uiMessageIDs = ids

	enBytes, err := localeFiles.ReadFile("locales/active.en.yaml")
	if err != nil {
		return fmt.Errorf("read locales/active.en.yaml: %w", err)
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	if _, err := bundle.ParseMessageFileBytes(ruBytes, "active.ru.yaml"); err != nil {
		return fmt.Errorf("i18n load ru: %w", err)
	}
	if _, err := bundle.ParseMessageFileBytes(enBytes, "active.en.yaml"); err != nil {
		return fmt.Errorf("i18n load en: %w", err)
	}
	r.i18nBundle = bundle
	return nil
}

func (r *Runtime) marshalUITranslationsJSON(lang string) (string, error) {
	if r.i18nBundle == nil || len(r.uiMessageIDs) == 0 {
		return "{}", nil
	}
	loc := i18n.NewLocalizer(r.i18nBundle, normalizeUILangTag(lang))
	out := make(map[string]string, len(r.uiMessageIDs))
	for _, id := range r.uiMessageIDs {
		s, err := loc.Localize(&i18n.LocalizeConfig{MessageID: id})
		if err != nil || s == "" {
			s = id
		}
		out[id] = s
	}
	b, err := json.Marshal(out)
	if err != nil {
		return "{}", fmt.Errorf("marshal ui translations: %w", err)
	}
	return string(b), nil
}

func (r *Runtime) registerI18nBindings(ctx context.Context) {
	r.register(ctx, "getUITranslationsJSON", func(_ js.Value, args []js.Value) any {
		lang := "en"
		if len(args) >= 1 && args[0].Type() == js.TypeString {
			lang = args[0].String()
		}
		s, err := r.marshalUITranslationsJSON(lang)
		if err != nil {
			return "{}"
		}
		return s
	}, i18nLangPerfExtras)

	r.register(ctx, "setUILanguage", func(_ js.Value, args []js.Value) any {
		lang := "en"
		if len(args) >= 1 && args[0].Type() == js.TypeString {
			lang = args[0].String()
		}
		r.SetUILang(lang)
		return nil
	}, i18nLangPerfExtras)
}

func i18nLangPerfExtras(args []js.Value, _ *any) []any {
	lang := "en"
	if len(args) >= 1 && args[0].Type() == js.TypeString {
		lang = args[0].String()
	}
	return []any{"lang", lang}
}
