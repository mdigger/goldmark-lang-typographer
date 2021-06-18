// Package typographer allows you to select the language typographic
// substitutions for Goldmark Typograph.
//
// This library is not a replacement of the standard Goldmark typography, and
// only expands its capabilities by adding quotation symbols for different
// languages. Quotes substitutions based on Wikipedia Page
// https://en.wikipedia.org/wiki/Quotation_mark.
// If the substitutions for this language are not specified, the library tries
// to use substitutions for the base language or default.
package typographer

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"golang.org/x/text/language"
)

// language configurations.
var (
	// substitutions contains all supported TypographicSubstitutions definitions.
	substitutions = [...]extension.TypographicSubstitutions{
		nil, // 0 (default)
		{
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		}, // 1 |„“
		{
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&rdquo;`),
		}, // 2 |„”
		{
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
		}, // 3 |«»
		{
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
			extension.EnDash:           []byte(`&mdash;`), // replace to mdash;
		}, // 4 |«» (+ ndash as mdash)
		{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&lsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		}, // 5 ’‘|„“
		{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&rsquo;`),
			extension.LeftDoubleQuote:  []byte(`&rdquo;`),
			extension.RightDoubleQuote: []byte(`&rdquo;`),
		}, // 6 ’’|””
		{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&rsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		}, // 7 ’’|„“
		{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&rsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&rdquo;`),
		}, // 8 ’’|„”
		{
			extension.LeftSingleQuote:  []byte(`&sbquo;`),
			extension.RightSingleQuote: []byte(`&lsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		}, // 9 ‚‘|„“
		{
			extension.LeftSingleQuote:  []byte(`&sbquo;`),
			extension.RightSingleQuote: []byte(`&lsquo;`),
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
		}, // 10 ‚‘|«»
		{
			extension.LeftSingleQuote:  []byte(`&lsaquo;`),
			extension.RightSingleQuote: []byte(`&rsaquo;`),
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
		}, // 11 ‹›|«»
		{
			extension.LeftSingleQuote:  []byte(`&lsaquo;&#8239;`),
			extension.RightSingleQuote: []byte(`&#8239;&rsaquo;`),
			extension.LeftDoubleQuote:  []byte(`&laquo;&#8239;`),
			extension.RightDoubleQuote: []byte(`&#8239;&raquo;`),
		}, // 12 ‹ ›|« » (with nnbsp spaces)
		{
			extension.LeftSingleQuote:  []byte(`&rsaquo;`),
			extension.RightSingleQuote: []byte(`&lsaquo;`),
			extension.LeftDoubleQuote:  []byte(`&raquo;`),
			extension.RightDoubleQuote: []byte(`&laquo;`),
		}, // 13 ›‹|»«
	}

	// configs содержит настройки для каждого языка, отличные от базового
	configs = map[language.Tag]extension.TypographicSubstitutions{
		// --------- 1
		language.Albanian: substitutions[1], // sq - Albanian
		language.Estonian: substitutions[1], // et - Estonian
		language.Georgian: substitutions[1], // ka - Georgian
		// --------- 2
		language.Croatian: substitutions[2], // hr - Croatian
		language.Polish:   substitutions[2], // pl - Polish
		language.Romanian: substitutions[2], // ro - Romanian
		// --------- 3
		language.Azerbaijani:          substitutions[3], // az - Azerbaijani
		language.Catalan:              substitutions[3], // ca - Catalan
		language.Spanish:              substitutions[3], // es - Spanish
		language.MustParse("eu"):      substitutions[3], // eu - Basque
		language.MustParse("gl"):      substitutions[3], // gl - Galician
		language.Armenian:             substitutions[3], // hy - Armenian
		language.MustParse("io"):      substitutions[3], // io - Ido
		language.Italian:              substitutions[3], // it - Italian
		language.Kazakh:               substitutions[3], // kk - Kazakh
		language.MustParse("mn-Cyrl"): substitutions[3], // mn-Cyrl - Mongolian, Cyrillic script
		language.Norwegian:            substitutions[3], // no - Norwegian
		language.BrazilianPortuguese:  substitutions[3], // pt-BR - Portuguese, Brazilian
		language.MustParse("pt-PL"):   substitutions[3], // pt-PL - Portuguese, Portugal
		language.Arabic:               substitutions[3], // ar - Arabic
		language.Greek:                substitutions[3], // el - Greek
		language.Persian:              substitutions[3], // fa - Persian
		language.Khmer:                substitutions[3], // km - Khmer
		language.MustParse("oc"):      substitutions[3], // oc - Occitan
		language.MustParse("ps"):      substitutions[3], // ps - Pashto
		// --------- 4
		language.Russian:         substitutions[4], // ru - Russian
		language.Ukrainian:       substitutions[4], // uk - Ukrainian
		language.MustParse("be"): substitutions[4], // be - Belarusian
		// --------- 5
		language.Macedonian: substitutions[5], // mk - Macedonian
		// --------- 6
		language.MustParse("bs"): substitutions[6], // bs - Bosnian
		language.Finnish:         substitutions[6], // fi - Finnish
		language.Hebrew:          substitutions[6], // he - Hebrew
		language.Swedish:         substitutions[6], // sv - Swedish
		// --------- 7
		language.Bulgarian: substitutions[7], // bg - Bulgarian
		// --------- 8
		language.Hungarian: substitutions[8], // hu - Hungarian
		language.Serbian:   substitutions[8], // sr - Serbian
		// --------- 9
		language.Czech:            substitutions[9], // cs - Czech
		language.German:           substitutions[9], // de - German
		language.Icelandic:        substitutions[9], // is - Icelandic
		language.Lithuanian:       substitutions[9], // lt - Lithuanian
		language.Slovak:           substitutions[9], // sk - Slovak
		language.Slovenian:        substitutions[9], // sl - Slovene
		language.MustParse("wen"): substitutions[9], // wen - Sorbian
		// --------- 10
		language.Uzbek: substitutions[10], // uz - Uzbek
		// --------- 11
		language.Amharic:            substitutions[11], // am - Amharic
		language.MustParse("de-CH"): substitutions[11], // de-CH - German, Swiss
		language.MustParse("fr-CH"): substitutions[11], // fr-CH - French, Swiss
		language.MustParse("it-CH"): substitutions[11], // it-CH - Italian, Swiss
		language.MustParse("rm"):    substitutions[11], // rm - Romansh
		language.MustParse("ti"):    substitutions[11], // ti - Tigrinya
		language.MustParse("ug"):    substitutions[11], // ug - Uyghur
		// --------- 12
		language.French: substitutions[12], // fr - French
		// --------- 13
		language.Danish: substitutions[13], // da - Danish
	}

	// languages содержит список поддерживаемых языков в порядке важности.
	languages = []language.Tag{
		// -------- 0 (as default)
		language.English, // en - English
		// -------- 1
		language.Albanian, // sq - Albanian
		language.Estonian, // et - Estonian
		language.Georgian, // ka - Georgian
		// --------- 2
		language.Croatian, // hr - Croatian
		language.Polish,   // pl - Polish
		language.Romanian, // ro - Romanian
		// --------- 3
		language.Azerbaijani,          // az - Azerbaijani
		language.Catalan,              // ca - Catalan
		language.Spanish,              // es - Spanish
		language.MustParse("eu"),      // eu - Basque
		language.MustParse("gl"),      // gl - Galician
		language.Armenian,             // hy - Armenian
		language.MustParse("io"),      // io - Ido
		language.Italian,              // it - Italian
		language.Kazakh,               // kk - Kazakh
		language.MustParse("mn-Cyrl"), // mn-Cyrl - Mongolian, Cyrillic script
		language.Norwegian,            // no - Norwegian
		language.BrazilianPortuguese,
		language.MustParse("pt-PL"), // pt-PL - Portuguese, Portugal
		language.Arabic,             // ar - Arabic
		language.Greek,              // el - Greek
		language.Persian,            // fa - Persian
		language.Khmer,              // km - Khmer
		language.MustParse("oc"),    // oc - Occitan
		language.MustParse("ps"),    // ps - Pashto
		// --------- 4
		language.Russian,         // ru - Russian
		language.Ukrainian,       // uk - Ukrainian
		language.MustParse("be"), // be - Belarusian
		// --------- 5
		language.Macedonian, // mk - Macedonian
		// --------- 6
		language.MustParse("bs"), // bs - Bosnian
		language.Finnish,         // fi - Finnish
		language.Hebrew,          // he - Hebrew
		language.Swedish,         // sv - Swedish
		// --------- 7
		language.Bulgarian, // bg - Bulgarian
		// --------- 8
		language.Hungarian, // hu - Hungarian
		language.Serbian,   // sr - Serbian
		// --------- 9
		language.Czech,            // cs - Czech
		language.German,           // de - German
		language.Icelandic,        // is - Icelandic
		language.Lithuanian,       // lt - Lithuanian
		language.Slovak,           // sk - Slovak
		language.Slovenian,        // sl - Slovene
		language.MustParse("wen"), // wen - Sorbian
		// --------- 10
		language.Uzbek, // uz - Uzbek
		// --------- 11
		language.Amharic,            // am - Amharic
		language.MustParse("de-CH"), // de-CH - German, Swiss
		language.MustParse("fr-CH"), // fr-CH - French, Swiss
		language.MustParse("it-CH"), // it-CH - Italian, Swiss
		language.MustParse("rm"),    // rm - Romansh
		language.MustParse("ti"),    // ti - Tigrinya
		language.MustParse("ug"),    // ug - Uyghur
		// --------- 12
		language.French, // fr - French
		// --------- 13
		language.Danish, // da - Danish
		// --------- 0
		language.Afrikaans,           // af - Afrikaans
		language.MustParse("cy"),     // cy, Welsh
		language.AmericanEnglish,     // en-US
		language.BritishEnglish,      // en-GB
		language.MustParse("en-CA"),  // English, Canada
		language.MustParse("eo"),     // Esperanto (eo)
		language.Filipino,            // fil - Filipino
		language.MustParse("ga"),     // Irish
		language.MustParse("gd"),     // Scottish Gaelic
		language.Hindi,               // hi - Hindi
		language.MustParse("ia"),     // Interlingua
		language.Indonesian,          // id - Indonesian
		language.MustParse("ko-KR"),  // Korean, South Korea
		language.MustParse("mt"),     // Maltese
		language.Dutch,               // nl - Dutch
		language.BrazilianPortuguese, // pt-BR, Portuguese, Brazil
		language.Tamil,               // ta - Tamil
		language.Thai,                // th - Thai
		language.Turkish,             // tr - Turkish
		language.Urdu,                // ur - Urdu
		language.SimplifiedChinese,   // zh-Hans, Chinese, simplified
		language.Lao,                 // lo - Lao
		language.Latvian,             // lv - Latvian
		language.Vietnamese,          // vi - Vietnamese
	}

	// matcher используется для выбора подходящего языка из списка поддерживаемых
	matcher = language.NewMatcher(languages)
)

// Substitutions returns a substitutions of typographic characters for the
// specified language.
func Substitutions(lang ...string) extension.TypographicSubstitutions {
	_, i := language.MatchStrings(matcher, lang...)
	return configs[languages[i]]
}

// Options return an TypographerOption options for the TypographerParser.
func Options(lang ...string) extension.TypographerOption {
	return extension.WithTypographicSubstitutions(
		Substitutions(lang...))
}

// New returns a new goldmark.Extender that replaces punctuations with
// typographic entities for the specified language.
func New(lang ...string) goldmark.Extender {
	return extension.NewTypographer(Options(lang...))
}
