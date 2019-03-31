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
bundlede with maCOS. This includes `libxml2` and it's header files.

```sh
$ brew install libxml2

$ PKG_CONFIG_PATH="/usr/local/opt/libxml2/lib/pkgconfig" \
    go get github.com/lestrrat-go/libxml2/...
```
