package types

// ContactCheck represents a check for domain(s).
type ContactCheck struct {
	Names []string `xml:"command>check>contact:check>name"`
}
