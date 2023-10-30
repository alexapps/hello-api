package translation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexapps/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "English",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
	}

	for _, test := range tt {
		// Act
		res := translation.Translate(test.Word, test.Language)

		// Assert
		assert.Equal(t, test.Translation, res, "they should be equal")

	}
}
