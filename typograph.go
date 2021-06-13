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

// Substitutions returns a substitutions of typographic characters for the
// specified language.
func Substitutions(lang string) extension.TypographicSubstitutions {
	// parse language definition string
	tag := language.Make(lang)
	if tag.IsRoot() {
		return nil // lang not recognized
	}
	lang = tag.String() // normalized language name

retry:
	// based on https://en.wikipedia.org/wiki/Quotation_mark
	switch lang {
	case
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
		"vi":      // Vietnamese
		return nil // ‘’|“” (default)

	case
		"sq", // Albanian
		"et", // Estonian
		"ka": // Georgian
		return extension.TypographicSubstitutions{
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		} // |„“

	case
		"hr", // Croatian
		"pl", // Polish
		"ro": // Romanian
		return extension.TypographicSubstitutions{
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&rdquo;`),
		} // |„”

	case
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
		"ps":      // Pashto
		return extension.TypographicSubstitutions{
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
		} // |«»

	case
		"ru", // Russian
		"uk", // Ukrainian
		"be": // Belarusian
		return extension.TypographicSubstitutions{
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
			extension.EnDash:           []byte(`&mdash;`), // replace to mdash;
		} // |«» (+ ndash as mdash)

	case
		"mk": // Macedonian
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&lsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		} // ’‘|„“

	case
		"bs", // Bosnian
		"fi", // Finnish
		"he", // Hebrew
		"sv": // Swedish
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&rsquo;`),
			extension.LeftDoubleQuote:  []byte(`&rdquo;`),
			extension.RightDoubleQuote: []byte(`&rdquo;`),
		} // ’’|””

	case
		"bg": // Bulgarian
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&rsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		} // ’’|„“

	case
		"hu", // Hungarian
		"sr": // Serbian
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&rsquo;`),
			extension.RightSingleQuote: []byte(`&rsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&rdquo;`),
		} // ’’|„”

	case
		"cs",  // Czech
		"de",  // German
		"is",  // Icelandic
		"lt",  // Lithuanian
		"sk",  // Slovak
		"sl",  // Slovene
		"wen": // Sorbian
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&sbquo;`),
			extension.RightSingleQuote: []byte(`&lsquo;`),
			extension.LeftDoubleQuote:  []byte(`&bdquo;`),
			extension.RightDoubleQuote: []byte(`&ldquo;`),
		} // ‚‘|„“

	case
		"uz": // Uzbek
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&sbquo;`),
			extension.RightSingleQuote: []byte(`&lsquo;`),
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
		} // ‚‘|«»

	case
		"am",    // Amharic
		"de-CH", // German, Swiss
		"fr-CH", // French, Swiss
		"it-CH", // Italian, Swiss
		"rm",    // Romansh
		"ti",    // Tigrinya
		"ug":    // Uyghur
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&lsaquo;`),
			extension.RightSingleQuote: []byte(`&rsaquo;`),
			extension.LeftDoubleQuote:  []byte(`&laquo;`),
			extension.RightDoubleQuote: []byte(`&raquo;`),
		} // ‹›|«»
	case
		"fr": // French
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&lsaquo;&#8239;`),
			extension.RightSingleQuote: []byte(`&#8239;&rsaquo;`),
			extension.LeftDoubleQuote:  []byte(`&laquo;&#8239;`),
			extension.RightDoubleQuote: []byte(`&#8239;&raquo;`),
		} // ‹ ›|« » (with nnbsp spaces)

	case
		"da": // Danish
		return extension.TypographicSubstitutions{
			extension.LeftSingleQuote:  []byte(`&rsaquo;`),
			extension.RightSingleQuote: []byte(`&lsaquo;`),
			extension.LeftDoubleQuote:  []byte(`&raquo;`),
			extension.RightDoubleQuote: []byte(`&laquo;`),
		} // ›‹|»«

	}

	// no definitions for this language
	if base, confidence := tag.Base(); confidence != language.No {
		if baseLang := base.String(); baseLang != lang {
			lang = baseLang
			goto retry // retry with base language
		}
	}

	return nil // default
}

// Options return an TypographerOption options for the TypographerParser.
func Options(lang string) extension.TypographerOption {
	return extension.WithTypographicSubstitutions(
		Substitutions(lang))
}

// New returns a new goldmark.Extender that replaces punctuations with
// typographic entities for the specified language.
func New(lang string) goldmark.Extender {
	return extension.NewTypographer(Options(lang))
}
