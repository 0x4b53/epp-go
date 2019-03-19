package types

// Login represents the data passed to login.
type Login struct {
	ClientID    string        `xml:"command>login>clID,omitempty"`
	Password    string        `xml:"command>login>pw,omitempty"`
	NewPassword string        `xml:"command>login>newPW,omitempty"`
	Options     LoginOptions  `xml:"command>login>options,omitempty"`
	Services    LoginServices `xml:"command>login>svcs,omitempty"`
}

// LoginOptions represents options that belongs to the login command.
type LoginOptions struct {
	Version  string `xml:"version"`
	Language string `xml:"lang"`
}

// LoginServices represents services used while logging in
type LoginServices struct {
	ObjectURI        []string              `xml:"objURI"`
	ServiceExtension LoginServiceExtension `xml:"svcExtension"`
}

// LoginServiceExtension represents extension URIs.
type LoginServiceExtension struct {
	ExtensionURI []string `xml:"extURI"`
}
