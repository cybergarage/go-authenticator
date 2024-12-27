# go-authenticator

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-authenticator)
[![test](https://github.com/cybergarage/go-authenticator/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-authenticator/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-authenticator.svg)](https://pkg.go.dev/github.com/cybergarage/go-authenticator)
 [![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/go-authenticator) 
 [![codecov](https://codecov.io/gh/cybergarage/go-authenticator/graph/badge.svg?token=OCU5V0H3OX)](https://codecov.io/gh/cybergarage/go-authenticator)


The `go-authenticator` is an open-source framework for user authentication in Go applications. It supports multiple authentication methods, including username and password authentication, SASL (Simple Authentication and Security Layer) authentication, and certificate-based authentication.


![](doc/img/framework.png)

The `go-authenticator` framework is designed to be flexible and extensible, making it easy to integrate into existing applications. Its goal is to provide robust authentication functionality through a unified interface. The framework is actively used by the following projects:

- [PuzzleDB](https://github.com/cybergarage/puzzledb-go)
- [go-mysql](https://github.com/cybergarage/go-mysql)
- [go-postgresql](https://github.com/cybergarage/go-postgresql)
- [go-redis](https://github.com/cybergarage/go-redis)
- [go-mongo](https://github.com/cybergarage/go-mongo)

## Features

The `go-authenticator` framework provides the following features:

- **User Authentication**: Authenticate users using user and password.
- **SASL Authentication**: Authenticate users using SASL (Simple Authentication and Security Layer).
- **Certificate Authentication**: Authenticate users using certificate of TLS connection.

## Getting Started

The `go-authenticator` provides a authentication manager that manages the authentication process. The manager can be configured with different authentication methods, such as credential authentication, SASL authentication, and certificate authentication.

```go
type Manager interface {
	// SetCredentialAuthenticator sets the credential authenticator.
	SetCredentialAuthenticator(auth CredentialAuthenticator)
	// VerifyCredential verifies the client credential.
	VerifyCredential(conn auth.Conn, q auth.Query) (bool, error)
	// SetCredentialStore sets the credential store.
	SetCredentialStore(store CredentialStore)
	// CredentialStore returns the credential store.
	CredentialStore() CredentialStore
	// SetCertificateAuthenticator sets the certificate authenticator.
	SetCertificateAuthenticator(auth CertificateAuthenticator)
	// VerifyCertificate verifies the client certificate.
	VerifyCertificate(conn tls.Conn) (bool, error)
	// Mechanisms returns the mechanisms.
	Mechanisms() []sasl.Mechanism
	// Mechanism returns a mechanism by name.
	Mechanism(name string) (sasl.Mechanism, error)
}
```

### Credential Authentication

This section explains how to authenticate users based on credentials using the `CredentialAuthenticator` interface from the `go-authenticator` framework.

#### CredentialStore

The `go-authenticator` has a default credential authenticator which uses the CredentialStore. The CredentialStore to be used is set with the `Manager::SetCredentialStore`.

```go
type CredentialStore interface {
	LookupCredential(q Query) (Credential, bool, error)
}
```

The `CredentialStore::LookupCredential` should return true with the queried credential if it is found or false. Detailed information about the query failure can be returned with an error, while security information can be NULL as it may lead to vulnerabilities.

#### CredentialAuthenticator

The `go-authenticator` is configured with a standard authenticator, but the user can set their own authenticator.　The `CredentialAuthenticator` is a simple interface that verifies users based on their credentials. The `VerifyCredential` method takes a connection, a query, and a credential as arguments and returns a boolean value indicating whether the user is authenticated.

```go
type CredentialAuthenticator interface {
	VerifyCredential(conn Conn, q Query, cred Credential) (bool, error)
}
```

The `VerifyCredential` should, as a minimum, return true or false if the queried credential is correct. Detailed information about the query failure can be returned with an error, but the 

#### Examples

To integrate the user authentication function into your application, refer to the example below.

- [go-postgresql](https://github.com/cybergarage/go-postgresql)
  - [Server::receive()](https://github.com/cybergarage/go-postgresql/blob/master/postgresql/protocol/server_impl.go)
- [go-redis](https://github.com/cybergarage/go-redis)
  - [Server::Auth()](https://github.com/cybergarage/go-redis/blob/main/redis/server_auth.go)

### SASL Authentication

The `go-authenticator` framework includes the `go-sasl` package, which provides a set of SASL (Simple Authentication and Security Layer) mechanisms that can be used to authenticate users in Go applications. For information on how to use the SASL API, see the go-sasl documentation.

- [go-sasl](https://github.com/cybergarage/go-sasl)

#### Examples

To integrate the SASL authentication into your application, refer to the example below.

- [go-mongo](https://github.com/cybergarage/go-mongo)
  - [BaseCommandExecutor::ExecuteCommand()](https://github.com/cybergarage/go-mongo/blob/master/mongo/command_base_executor.go)


### Certificate Authentication

This section explains how to authenticate users based on the certificate of a TLS connection using the `CredentialAuthenticator` interface from the `go-authenticator` framework.

#### CertificateAuthenticator

The `CertificateAuthenticator` is a simple interface that verifies users by examining the certificate of the TLS connection.

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
    auth.WithCommonNameRegexp("localhost"))
if err != nil {
    t.Error(err)
    return
}
mgr.SetCertificateAuthenticator(ca)
```

By following these steps, you can easily authenticate users through TLS certificate verification, enhancing the security of your application.

#### Examples

To integrate the certificate authentication function into your application, refer to the example below.

- [go-mysql](https://github.com/cybergarage/go-mysql)
  - [Server::receive](https://github.com/cybergarage/go-mysql/blob/main/mysql/protocol/server.go)
- [go-mongo](https://github.com/cybergarage/go-mongo)
  - [Server::serve()](https://github.com/cybergarage/go-mongo/blob/master/mongo/server.go)
- [go-redis](https://github.com/cybergarage/go-redis)
  - [Server::tlsServe()](https://github.com/cybergarage/go-redis/blob/main/redis/server_impl.go)
