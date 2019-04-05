# EPP Go - Extensible Provisioning Protocol Server and Client

This is an implementation of how to handle EPP requests concurrently, both as a
client and as a server. The main focus lays implementing types that may be used
both as a client and as a server. These types should be easy to use and support
all the allowed ways of setting up name spaces, attributes and tags. With the
types implemented the initial focus will be to ensure a complete server
implementation may be created. Since this is registry specific there will
probably only be minor helpers and wrappers.

## NOTE

**This is a work in progress and a long way from completed. This repository is
created to allow collaborations and inspire other people. This probject is a
private project and an experiment to work with XSD files and XML with Go.**

## Types

A big motivation behind this project is to define all the available EPP types
and even some extensions. This is so that even if you don't use this server och
client you should be able to use the types to marshal or unmarshal your XML to
your desired system.

There are a lot of knowns problems with this, especially since EPP is so heavily
dependent on namespaces. [This issue](https://github.com/golang/go/issues/13400)
in the `golang/go` project summarize the most common issues with namespaces,
aliases and attributes.

The way this is handled in this project is to first define all types as an outer
tag for each part with the RFC namespace defined. The code then uses
[`go-libxml`](https://github.com/alexrsagen/go-libxml) to figure out where the
namespaces should be, adds an alias and `xmlns` tag and then uses the aliases on
all the child elements.

Sadly, this does not solve the issue that the XML should be able to be
unmarshalled to the defined types despite the namespace or alias. To handle this
a codegen binary is bundled in this project which can generate a copy of all
types without the namespace.

Example usage and installment.

```sh
$ go install ./cmd/type-generator/...
$ type-generator
Generated file: contact_auto_generated.go
Generated file: domain_auto_generated.go
Generated file: host_auto_generated.go
...
```

To generate XML to be used for a client, use the specified type for this.

```go
domainInfo := types.DomainInfoType{
    Info: types.DomainInfo{
        Name: types.DomainInfoName{
            Name: "example.se",
        },
    },
}

bytes, err := Encode(
    domainInfo,
    ClientXMLAttributes(),
)

if err != nil {
    panic(err)
}
```

The above code will generate the following XML.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<epp xmlns="urn:ietf:params:xml:ns:epp-1.0">
  <command>
    <info>
      <domain:info xmlns:domain="urn:ietf:params:xml:ns:domain-1.0" xmlns="urn:ietf:params:xml:ns:domain-1.0">
        <domain:name hosts="">example.se</domain:name>
        <domain:authInfo />
      </domain:info>
    </info>
  </command>
</epp>
```

To unmarshal already created XML no matter the namespace or alias, use the auto
genrated types. The XML listed above could be unmarshaled like this.

```go
domainInfoRequest := DomainInfoTypeIn{}

if err := xml.Unmarshal(inData, &domainInfoRequest); err != nil {
    panic(err)
}

fmt.Println(domainInfoRequest.Info.Name.Name) // Prints `example.se`
```

## Client

To quickly get up and running and support testing of the server the repository
contains a set of real world examples of EPP commands foudn in
[xml/commands](xml/commands).

Inside the [example](example/client) folder there's a client utilizing a few of
the types and all of the read/writering confirming to EPP RFC. This client reads
from STDIN so it's just to copy and paste any of the example XML file contents
to test changes.

## References

### XSD files

All XSD schemas can be found att [IANA web
page](https://www.iana.org/assignments/xml-registry/xml-registry.xhtml). XSD
files from this repository linked below.

* [contact-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/contact-1.0.xsd)
* [domain-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/domain-1.0.xsd)
* [epp-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/epp-1.0.xsd)
* [eppcom-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/eppcom-1.0.xsd)
* [host-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/host-1.0.xsd)
* [secDNS-1.0.xsd](https://www.iana.org/assignments/xml-registry/schema/secDNS-1.0.xsd)
* [secDNS-1.1.xsd](https://www.iana.org/assignments/xml-registry/schema/secDNS-1.1.xsd)

### EPP RFC

* [RFC 5730 Extensible Provisioning Protocol (EPP)](http://www.rfc-editor.org/rfc/rfc5730.txt)
* [RFC 5731 Extensible Provisioning Protocol (EPP) Domain Name Mapping](http://www.rfc-editor.org/rfc/rfc5731.txt)
* [RFC 5732 Extensible Provisioning Protocol (EPP) Host Mapping](http://www.rfc-editor.org/rfc/rfc5732.txt)
* [RFC 5733 Extensible Provisioning Protocol (EPP) Contact Mapping](http://www.rfc-editor.org/rfc/rfc5733.txt)
* [RFC 5734 Extensible Provisioning Protocol (EPP) Transport over TCP](http://www.rfc-editor.org/rfc/rfc5734.txt)
* [RFC 5910 Domain Name System (DNS) Security Extensions Mapping for the Extensible Provisioning Protocol (EPP)](http://www.rfc-editor.org/rfc/rfc5910.txt)

### TLD specific (.SE)

* EPP extension [iis-1.2.xsd](https://registrar.iis.se/files/iis-1.2.xml)

## Development and testing

XML files are linted with [`xmllint`](http://xmlsoft.org/xmllint.html).

To validate XML [`libxml2` (bindings for
Go)](https://github.com/lestrrat-go/libxml2/) is used. This package requires you
to install the [`libxml2`](http://xmlsoft.org/downloads.html) C libraries.

### Installation macOS

Since macOS 10.14 [brew](https://brew.sh/) won't link packages and libraries
bundlede with macOS. This includes `libxml2` and it's header files.

```sh
$ brew install libxml2

$ PKG_CONFIG_PATH="/usr/local/opt/libxml2/lib/pkgconfig" \
    go get github.com/lestrrat-go/libxml2/...
```
