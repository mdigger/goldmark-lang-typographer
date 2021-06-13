# Goldmark Typographer Language Extension

[![Go Reference](https://pkg.go.dev/badge/github.com/mdigger/goldmark-lang-typographer.svg)](https://pkg.go.dev/github.com/mdigger/goldmark-lang-typographer)

Package typographer allows you to select the language typographic substitutions 
for Goldmark Typographer.

This library is not a replacement of the standard Goldmark typography, and only 
expands its capabilities by adding quotation symbols for different languages.

```golang
import typographer "github.com/mdigger/goldmark-lang-typographer"

markdown := goldmark.New(
    goldmark.WithExtensions(
        typographer.New("ru"),
    ),
)
```

## Quotation marks substitution

Quotes substitutions based on [Wikipedia](https://en.wikipedia.org/wiki/Quotation_mark). 
If the substitutions for this language are not specified, the library tries to 
use substitutions for the base language or default.

| Code      | Lang                       | Double | Single | Dash |
|-----------|----------------------------|:------:|:------:|:----:|
| `af`      | Afrikaans                  |        |        |      |
| `sq`      | Albanian                   |  `„“`  |        |      |
| `am`      | Amharic                    |  `«»`  |  `‹›`  |      |
| `ar`      | Arabic                     |  `«»`  |        |      |
| `hy`      | Armenian                   |  `«»`  |        |      |
| `az`      | Azerbaijani                |  `«»`  |        |      |
| `eu`      | Basque                     |  `«»`  |        |      |
| `be`      | Belarusian                 |  `«»`  |        | `--` |
| `bs`      | Bosnian                    |  `””`  |  `’’`  |      |
| `bg`      | Bulgarian                  |  `„“`  |  `’’`  |      |
| `ca`      | Catalan                    |  `«»`  |        |      |
| `zh-Hans` | Chinese, simplified        |        |        |      |
| `hr`      | Croatian                   |  `„”`  |        |      |
| `cs`      | Czech                      |  `„“`  |  `‚‘`  |      |
| `da`      | Danish                     |  `»«`  |  `›‹`  |      |
| `nl`      | Dutch                      |        |        |      |
| `en`      | English                    |        |        |      |
| `en-GB`   | English, UK                |        |        |      |
| `en-US`   | English, US                |        |        |      |
| `en-CA`   | English, Canada            |        |        |      |
| `eo`      | Esperanto                  |        |        |      |
| `et`      | Estonian                   |  `„“`  |        |      |
| `fil`     | Filipino                   |        |        |      |
| `fi`      | Finnish                    |  `””`  |  `’’`  |      |
| `fr`      | French                     | `« »`  | `‹ ›`  |      |
| `fr-CH`   | French, Swiss              |  `«»`  |  `‹›`  |      |
| `gl`      | Galician                   |  `«»`  |        |      |
| `ka`      | Georgian                   |  `„“`  |        |      |
| `de`      | German                     |  `„“`  |  `‚‘`  |      |
| `de-CH`   | German, Swiss              |  `«»`  |  `‹›`  |      |
| `el`      | Greek                      |  `«»`  |        |      |
| `he`      | Hebrew                     |  `””`  |  `’’`  |      |
| `hi`      | Hindi                      |        |        |      |
| `hu`      | Hungarian                  |  `„”`  |  `’’`  |      |
| `is`      | Icelandic                  |  `„“`  |  `‚‘`  |      |
| `io`      | Ido                        |  `«»`  |        |      |
| `id`      | Indonesian                 |        |        |      |
| `ia`      | Interlingua                |        |        |      |
| `ga`      | Irish                      |        |        |      |
| `it`      | Italian                    |  `«»`  |        |      |
| `it-CH`   | Italian, Swiss             |  `«»`  |  `‹›`  |      |
| `kk`      | Kazakh                     |  `«»`  |        |      |
| `km`      | Khmer                      |  `«»`  |        |      |
| `ko`      | Korean, South Korea        |        |        |      |
| `lo`      | Lao                        |        |        |      |
| `lv`      | Latvian                    |        |        |      |
| `lt`      | Lithuanian                 |  `„“`  |  `‚‘`  |      |
| `mk`      | Macedonian                 |  `„“`  |  `’‘`  |      |
| `mt`      | Maltese                    |        |        |      |
| `mn-Cyrl` | Mongolian, Cyrillic script |  `«»`  |        |      |
| `no`      | Norwegian                  |  `«»`  |        |      |
| `oc`      | Occitan                    |  `«»`  |        |      |
| `ps`      | Pashto                     |  `«»`  |        |      |
| `fa`      | Persian                    |  `«»`  |        |      |
| `pl`      | Polish                     |  `„”`  |        |      |
| `pt-BR`   | Portuguese, Brazil         |        |        |      |
| `pt-PL`   | Portuguese, Portugal       |  `«»`  |        |      |
| `ro`      | Romanian                   |  `„”`  |        |      |
| `rm`      | Romansh                    |  `«»`  |  `‹›`  |      |
| `ru`      | Russian                    |  `«»`  |        | `--` |
| `sr`      | Serbian                    |  `„”`  |  `’’`  |      |
| `gd`      | Scottish Gaelic            |        |        |      |
| `sk`      | Slovak                     |  `„“`  |  `‚‘`  |      |
| `sl`      | Slovene                    |  `„“`  |  `‚‘`  |      |
| `wen`     | Sorbian                    |  `„“`  |  `‚‘`  |      |
| `es`      | Spanish                    |  `«»`  |        |      |
| `sv`      | Swedish                    |  `””`  |  `’’`  |      |
| `ta`      | Tamil                      |        |        |      |
| `ti`      | Tigrinya                   |  `«»`  |  `‹›`  |      |
| `th`      | Thai                       |        |        |      |
| `tr`      | Turkish                    |        |        |      |
| `uk`      | Ukrainian                  |  `«»`  |        | `--` |
| `ur`      | Urdu                       |        |        |      |
| `ug`      | Uyghur                     |  `«»`  |  `‹›`  |      |
| `uz`      | Uzbek                      |  `«»`  |  `‚‘`  |      |
| `vi`      | Vietnamese                 |        |        |      |
| `cy`      | Welsh                      |        |        |      |
