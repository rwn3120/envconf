package envconf

import (
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

type nestedTestConf struct {
	Text string `env:"TEXT"`
}

type testConf struct {
	PtrStringParameter *string `env:"PTR_STRING_PARAMETER"`
	StringParameter    string  `env:"STRING_PARAMETER"`
	IntParameter       int     `env:"INT_PARAMETER"`
	Int32Parameter     int32   `env:"INT32_PARAMETER"`
	Int64Parameter     int64   `env:"INT64_PARAMETER"`
	UintParameter      uint    `env:"UINT_PARAMETER"`
	Uint32Parameter    uint32  `env:"UINT32_PARAMETER"`
	Uint64Parameter    uint64  `env:"UINT64_PARAMETER"`
	Float32Parameter   float32 `env:"FLOAT32_PARAMETER"`
	Float64Parameter   float64 `env:"FLOAT64_PARAMETER"`
	BoolParameter      bool    `env:"BOOL_PARAMETER"`
	UnknownParameter   string

	StructParameter    nestedTestConf
	PtrStructParameter *nestedTestConf
}

func TestLoadFromEnv(t *testing.T) {
	someString := "some string"
	os.Clearenv()
	os.Setenv("PTR_STRING_PARAMETER", someString)
	os.Setenv("STRING_PARAMETER", someString)
	os.Setenv("INT_PARAMETER", "0")
	os.Setenv("INT32_PARAMETER", "-1")
	os.Setenv("INT64_PARAMETER", "-2")
	os.Setenv("UINT_PARAMETER", "0")
	os.Setenv("UINT32_PARAMETER", "1")
	os.Setenv("UINT64_PARAMETER", "2")
	os.Setenv("FLOAT32_PARAMETER", "0.1")
	os.Setenv("FLOAT64_PARAMETER", ".2")
	os.Setenv("BOOL_PARAMETER", "true")
	os.Setenv("TEXT", someString)

	c := testConf{}
	assert.NilError(t, FromEnv(&c))
	assert.DeepEqual(t, c, testConf{
		PtrStringParameter: &someString,
		StringParameter:    someString,
		IntParameter:       0,
		Int32Parameter:     -1,
		Int64Parameter:     -2,
		UintParameter:      0,
		Uint32Parameter:    1,
		Uint64Parameter:    2,
		Float32Parameter:   0.1,
		Float64Parameter:   .2,
		BoolParameter:      true,
		UnknownParameter:   "",
		StructParameter: nestedTestConf{
			Text: someString,
		},
		PtrStructParameter: &nestedTestConf{
			Text: someString,
		},
	})
}

func TestFromYAML(t *testing.T) {
	someString := "some string"
	yaml := `ptrstringparameter: some string
stringparameter: some string
intparameter: 0
int32parameter: -1
int64parameter: -2
uintparameter: 0
uint32parameter: 1
uint64parameter: 2
float32parameter: 0.1
float64parameter: 0.2
boolparameter: true
unknownparameter: ""
structparameter:
  text: some text
ptrstructparameter:
  text: some text
`
	c := testConf{}
	assert.NilError(t, FromYAML(strings.NewReader(yaml), &c))
	assert.Equal(t, yaml, ToYAML(c))
	assert.DeepEqual(t, c, testConf{
		PtrStringParameter: &someString,
		StringParameter:    someString,
		IntParameter:       0,
		Int32Parameter:     -1,
		Int64Parameter:     -2,
		UintParameter:      0,
		Uint32Parameter:    1,
		Uint64Parameter:    2,
		Float32Parameter:   0.1,
		Float64Parameter:   .2,
		BoolParameter:      true,
		UnknownParameter:   "",
		StructParameter: nestedTestConf{
			Text: "some text",
		},
		PtrStructParameter: &nestedTestConf{
			Text: "some text",
		},
	})
}

func TestLoad(t *testing.T) {
	someString := "some string"
	yaml := `ptrstringparameter: some string
stringparameter: some string
intparameter: 0
int32parameter: -1
int64parameter: -2
uintparameter: 0
uint32parameter: 1
uint64parameter: 2
float32parameter: 0.1
float64parameter: 0.2
boolparameter: true
unknownparameter: ""
structparameter:
  text: some text
ptrstructparameter:
  text: some text
`

	os.Clearenv()
	os.Setenv("STRING_PARAMETER", "my string")
	os.Setenv("INT_PARAMETER", "31337")

	c := testConf{}
	assert.NilError(t, Load(strings.NewReader(yaml), &c))
	assert.DeepEqual(t, c, testConf{
		PtrStringParameter: &someString,
		StringParameter:    "my string",
		IntParameter:       31337,
		Int32Parameter:     -1,
		Int64Parameter:     -2,
		UintParameter:      0,
		Uint32Parameter:    1,
		Uint64Parameter:    2,
		Float32Parameter:   0.1,
		Float64Parameter:   .2,
		BoolParameter:      true,
		UnknownParameter:   "",
		StructParameter: nestedTestConf{
			Text: "some text",
		},
		PtrStructParameter: &nestedTestConf{
			Text: "some text",
		},
	})
}
