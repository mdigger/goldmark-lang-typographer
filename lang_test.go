package typographer

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"
)

func TestLang(t *testing.T) {
	for _, lang := range []string{
		"", "und", // Unknown
		"af",      // Afrikaans
		"cy",      // Welsh
		"en",      // English
		"en-CA",   // English, Canada
		"en-GB",   // English, UK
		"en-US",   // English, US
		"eo",      // Esperanto
		"fil",     // Filipino
		"ga",      // Irish
		"gd",      // Scottish Gaelic
		"hi",      // Hindi
		"ia",      // Interlingua
		"id",      // Indonesian
		"ko",      // Korean
		"ko-KR",   // Korean, South Korea
		"mt",      // Maltese
		"nl",      // Dutch
		"pt-BR",   // Portuguese, Brazil
		"ta",      // Tamil
		"th",      // Thai
		"tr",      // Turkish
		"ur",      // Urdu
		"zh-Hans", // Chinese, simplified
		"lo",      // Lao
		"lv",      // Latvian
		"vi",      // Vietnamese
		"",        // ---------
		"sq",      // Albanian
		"et",      // Estonian
		"ka",      // Georgian
		"",        // ---------
		"hr",      // Croatian
		"pl",      // Polish
		"ro",      // Romanian
		"",        // ---------
		"az",      // Azerbaijani
		"ca",      // Catalan
		"es",      // Spanish
		"eu",      // Basque
		"gl",      // Galician
		"hy",      // Armenian
		"io",      // Ido
		"it",      // Italian
		"kk",      // Kazakh
		"mn-Cyrl", // Mongolian, Cyrillic script
		"no",      // Norwegian
		"pt-PL",   // Portuguese, Portugal
		"ar",      // Arabic
		"el",      // Greek
		"fa",      // Persian
		"km",      // Khmer
		"oc",      // Occitan
		"ps",      // Pashto
		"",        // ---------
		"ru",      // Russian
		"uk",      // Ukrainian
		"be",      // Belarusian
		"",        // ---------
		"mk",      // Macedonian
		"",        // ---------
		"bs",      // Bosnian
		"fi",      // Finnish
		"he",      // Hebrew
		"sv",      // Swedish
		"",        // ---------
		"bg",      // Bulgarian
		"",        // ---------
		"hu",      // Hungarian
		"sr",      // Serbian
		"",        // ---------
		"cs",      // Czech
		"de",      // German
		"is",      // Icelandic
		"lt",      // Lithuanian
		"sk",      // Slovak
		"sl",      // Slovene
		"wen",     // Sorbian
		"",        // ---------
		"uz",      // Uzbek
		"",        // ---------
		"am",      // Amharic
		"de-CH",   // German, Swiss
		"fr-CH",   // French, Swiss
		"it-CH",   // Italian, Swiss
		"rm",      // Romansh
		"ti",      // Tigrinya
		"ug",      // Uyghur
		"",        // ---------
		"fr",      // French
		"",        // ---------
		"da",      // Danish
		"",        // ---------
		"en-Latn", "sh", "zh-cmn", "bjd", "iw-Latn-fonipa-u-cu-usd",
		"und-US", "und-NL", "und-419", "und-ZZ",
		"sr", "sr_Latn",
		"en-US", "en-GB", "en-UK", "en-BU", "en-RU",
		"de-1901-u-co-phonebk", "ja-JP", "fi-x-ing",
		"en-Latn-UK",
		"zh-CN", "en-AU", "zh-HK", "zh-Hant", "de-1994-u-co-phonebk", "de-Latn-LU",
		"iw-IL", "iw", "he",
		"fr-BE", "af-NA",
		"", // ---------
	} {
		if lang == "" {
			println("-------------------------------------------------------------------")
			continue
		}

		_, i := language.MatchStrings(matcher, lang)
		s := languages[i] // исходный язык из списка поддерживаемых
		// name := display.English.Tags().Name(s)
		name, _ := s.Base()
		fmt.Printf("%24s | %-7s | %2d | %s\n", lang, s, i, name)
	}
}
