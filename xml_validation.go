package epp

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/lestrrat-go/libxml2"
	xsd "github.com/lestrrat-go/libxml2/xsd"
)

// Validator represents a validator holding the XSD schema to calidate against.
type Validator struct {
	Schema *xsd.Schema
}

// NewValidator creates a new validator.
func NewValidator(rootXSD string) (*Validator, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = os.Chdir(cwd)
	}()

	// Change path to the root directory so include works. This assumes that the
	// path of included XSD files is always the same as the root XSD.
	path, file := path.Split(rootXSD)
	if err := os.Chdir(path); err != nil {
		return nil, err
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	schema, err := xsd.Parse(buf)
	if err != nil {
		return nil, err
	}

	return &Validator{
		Schema: schema,
	}, nil
}

// Validate will validate XML towards the XSD schema.
func (v *Validator) Validate(xml []byte) error {
	d, err := libxml2.Parse(xml)
	if err != nil {
		return err
	}

	if err := v.Schema.Validate(d); err != nil {
		return err
	}

	return nil
}
