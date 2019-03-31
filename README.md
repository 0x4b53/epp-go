# EPP Server

This is an implementation of how to handle EPP requests concurrently.

## NOTE

**This is a work in progress and a long way from completed. This repository is
created to allow collaborations and inspire other people. This probject is a
private project and an experiment to work with XSD files and XML with Go.**

## Client

A good way to test the implementation of this server is to use an existing and
tested client to ensure communication works as epxected.

Example clients:

* [Domainr EPP client in Go (WIP)](https://github.com/domainr/epp)
* [python-epp-client](https://github.com/Darkfish/python-epp-client)

Even this package includes a test client which can send XML to the EPP server
although no validation or verificate is made.

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
to install the [`libxml2`](http://xmlsoft.org/downloads.html) C bindings.

## Installation macOS

Since macOS 10.14 [brew](https://brew.sh/) won't link packages and libraries
bundlede with maCOS. This includes `libxml2` and it's header files.

```sh
$ brew install libxml2

$ PKG_CONFIG_PATH="/usr/local/opt/libxml2/lib/pkgconfig" \
    go get github.com/lestrrat-go/libxml2/...
```
