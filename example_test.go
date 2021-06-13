package typographer_test

import (
	"os"

	typographer "github.com/mdigger/goldmark-lang-typographer"
	"github.com/yuin/goldmark"
)

var source = []byte(
	`"Хруст французской булки" -- так заканчивается сказка про колобка братьев Гримм.`)

func ExampleNew() {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			typographer.New("ru"),
		),
	)

	err := markdown.Convert(source, os.Stdout)
	if err != nil {
		panic(err)
	}

	// Output:
	// <p>&laquo;Хруст французской булки&raquo; &mdash; так заканчивается сказка про колобка братьев Гримм.</p>
}
