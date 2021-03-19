package tests

import (
	"regexp"
	"testing"

	tf "github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func getShortName(t *testing.T, inputName string, maxLength, suffixLength int, upper bool) string {
	tfOptions := tf.WithDefaultRetryableErrors(t, &tf.Options{
		TerraformDir: ".",
		Vars: map[string]interface{}{
			"input_name":    inputName,
			"max_length":    maxLength,
			"suffix_length": suffixLength,
			"upper":         upper,
		},
	})

	defer tf.Destroy(t, tfOptions)

	tf.InitAndApply(t, tfOptions)

	return tf.Output(t, tfOptions, "unique_name")
}

func TestShortNameShortName(t *testing.T) {
	name := "shorty"
	assert.Equal(t, name, getShortName(t, name, 10, 4, true))
}

func TestShortNameSanitize(t *testing.T) {
	name := "i-like$weird.names? :)"
	assert.Equal(t, "i-like-weird-names", getShortName(t, name, 32, 4, true))
}

func TestShortNameLongName(t *testing.T) {
	name := "this-is-a-very-long-name"
	expected := regexp.MustCompile(`this-[[:alnum:]]{4}`)
	assert.Regexp(t, expected, getShortName(t, name, 10, 4, true))
}

func TestShortNameReplaceUnderscores(t *testing.T) {
	name := "name_with_underscores"
	expected := "name-with-underscores"
	assert.Equal(t, expected, getShortName(t, name, 100, 4, true))
}

func TestShortNameRemoveEndingDashes(t *testing.T) {
	name := "ends-with-dashes----"
	assert.Equal(t, "ends-with-dashes", getShortName(t, name, 100, 4, true))

	name = "truncated----ends-with-dashes"
	expected := regexp.MustCompile(`truncated-[[:alnum:]]{4}`)
	assert.Regexp(t, expected, getShortName(t, name, 14, 4, true))
}

func TestShortNameLongNameNoUpper(t *testing.T) {
	name := "this-is-a-very-long-name"
	expected := regexp.MustCompile(`this-[0-9a-z]{4}`)
	assert.Regexp(t, expected, getShortName(t, name, 10, 4, false))
}
