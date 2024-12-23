# go-authenticator

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-authenticator)
[![test](https://github.com/cybergarage/go-authenticator/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-authenticator/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-authenticator.svg)](https://pkg.go.dev/github.com/cybergarage/go-authenticator)
 [![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/go-authenticator) 
 [![codecov](https://codecov.io/gh/cybergarage/go-authenticator/graph/badge.svg?token=OCU5V0H3OX)](https://codecov.io/gh/cybergarage/go-authenticator)


The `go-authenticator` is a open source framework for authenticating users in Go applications. It supports multiple authentication methods, including user and password authentication, SASL (Simple Authentication and Security Layer) authentication, and certificate authentication.

![](doc/img/framework.png)

The `go-authenticator` framework is designed to be flexible and extensible, allowing you to easily integrate it into your existing applications. 

## Features

The `go-authenticator` framework provides the following features:

- **User Authentication**: Authenticate users using user and password.
- **SASL Authentication**: Authenticate users using SASL (Simple Authentication and Security Layer).
- **Certificate Authentication**: Authenticate users using certificate of TLS connection.

## Getting Started

### User Authentication

### SASL Authentication

The `go-authenticator` framework includes the `go-sasl` package, which provides a set of SASL (Simple Authentication and Security Layer) mechanisms that can be used to authenticate users in Go applications. For information on how to use the SASL API, see the go-sasl documentation.

- [go-sasl](https://github.com/cybergarage/go-sasl)

#### Examples

To integrate the SASL authentication into your application, refer to the example below.

- [go-mongo](https://github.com/cybergarage/go-mongo)
  - [BaseCommandExecutor::ExecuteCommand()](https://github.com/cybergarage/go-mongo/blob/master/mongo/command_base_executor.go)


### Certificate Authentication

This section explains how to authenticate users based on the certificate of a TLS connection using the `CredentialAuthenticator` interface from the `go-authenticator` framework.

#### CertificateAuthenticator Overview

The `CertificateAuthenticator` is a simple interface that verifies users by examining the certificate of the TLS connection.

##### Interface Definition
```go
type CertificateAuthenticator interface {
	VerifyCertificate(conn tls.Conn) (bool, error)
}
```

##### Creating a CertificateAuthenticator

To create an instance of `CertificateAuthenticator`, use the `NewCertificateAuthenticator()` function provided by `go-authenticator`. This instance authenticates users based on the common names (CN) found in the TLS connection certificate.

##### Enabling Certificate Authentication

To enable certificate authentication, set the `CertificateAuthenticator` instance to the manager by using the `SetCertificateAuthenticator` method.

```go
mgr := auth.NewManager()
ca, err := auth.NewCertificateAuthenticator(
    auth.WithCertificateAuthenticatorCommonNameRegexp("localhost"))
if err != nil {
    t.Error(err)
    return
}
mgr.SetCertificateAuthenticator(ca)
```

By following these steps, you can easily authenticate users through TLS certificate verification, enhancing the security of your application.

#### Examples

To integrate the certificate authentication into your application, refer to the example below.

- [go-mongo](https://github.com/cybergarage/go-mongo)
  - [Server::serve()](https://github.com/cybergarage/go-mongo/blob/master/mongo/server.go)
