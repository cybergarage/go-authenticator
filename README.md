# go-authenticator

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-authenticator)
[![test](https://github.com/cybergarage/go-authenticator/actions/workflows/make.yml/badge.svg)](https://github.com/cybergarage/go-authenticator/actions/workflows/make.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/cybergarage/go-authenticator.svg)](https://pkg.go.dev/github.com/cybergarage/go-authenticator)
 [![Go Report Card](https://img.shields.io/badge/go%20report-A%2B-brightgreen)](https://goreportcard.com/report/github.com/cybergarage/go-authenticator) 
 [![codecov](https://codecov.io/gh/cybergarage/go-authenticator/graph/badge.svg?token=OCU5V0H3OX)](https://codecov.io/gh/cybergarage/go-authenticator)
 
**go-authenticator** is an open-source framework for user authentication in Go applications. It supports multiple authentication methods, including username and password authentication, SASL (Simple Authentication and Security Layer) authentication, and certificate-based authentication.

![](doc/img/framework.png)

Built for flexibility and extensibility, [**go-authenticator**](https://github.com/cybergarage/go-authenticator) allows developers to easily integrate custom authentication methods. It provides a unified interface for handling authentication requests and responses and is used as the authentication component in the following projects:

- [**PuzzleDB**](https://github.com/cybergarage/puzzledb-go) ![](https://img.shields.io/github/v/tag/cybergarage/puzzledb-go)
- [**go-sqlserver**](https://github.com/cybergarage/go-sqlserver) ![](https://img.shields.io/github/v/tag/cybergarage/go-sqlserver)  

- [**go-mysql**](https://github.com/cybergarage/go-mysql) ![](https://img.shields.io/github/v/tag/cybergarage/go-mysql)
- [**go-postgresql**](https://github.com/cybergarage/go-postgresql) ![](https://img.shields.io/github/v/tag/cybergarage/go-postgresql)
- [**go-redis**](https://github.com/cybergarage/go-redis) ![](https://img.shields.io/github/v/tag/cybergarage/go-redis)
- [**go-mongo**](https://github.com/cybergarage/go-mongo) ![](https://img.shields.io/github/v/tag/cybergarage/go-mongo)

## Features

[**go-authenticator**](https://github.com/cybergarage/go-authenticator) is an open-source framework designed to simplify user authentication in Go applications. It supports various authentication methods, including:

- **User Authentication** with username and password  
- **SASL Authentication** for secure, extensible mechanisms  
- **Certificate Authentication** via TLS certificates  

[**go-authenticator**](https://github.com/cybergarage/go-authenticator) is a powerful and extensible framework for managing user authentication in Go applications. Its support for multiple authentication methods and seamless integration makes it an excellent choice for building secure, scalable systems.

## Getting Started

**go-authenticator** provides an authentication manager to handle the authentication process. The manager can be configured with different authentication methods, such as credential authentication, SASL authentication, and certificate authentication.

```go
type Manager interface {
    SetCredentialAuthenticator(auth CredentialAuthenticator)
    VerifyCredential(conn auth.Conn, q auth.Query) (bool, error)
    SetCredentialStore(store CredentialStore)
    CredentialStore() CredentialStore
    SetCertificateAuthenticator(auth CertificateAuthenticator)
    VerifyCertificate(conn tls.Conn) (bool, error)
    Mechanisms() []sasl.Mechanism
    Mechanism(name string) (sasl.Mechanism, error)
}
```

### Credential Authentication

This section explains how to authenticate users based on credentials using the `CredentialAuthenticator` interface.

#### CredentialStore

**go-authenticator** includes a default credential authenticator that uses `CredentialStore`. You can set the `CredentialStore` by calling `Manager::SetCredentialStore`.

```go
type CredentialStore interface {
    LookupCredential(q Query) (Credential, bool, error)
}
```

`LookupCredential` returns `true` if the queried credential is found. If not, it returns `false`. Detailed failure information can be returned via an error.

#### CredentialAuthenticator

The default authenticator can be replaced by a custom one. `CredentialAuthenticator` verifies users based on their credentials. The `VerifyCredential` method takes a connection, a query, and a credential, returning a boolean indicating successful authentication.

```go
type CredentialAuthenticator interface {
    VerifyCredential(conn Conn, q Query, cred Credential) (bool, error)
}
```

The `VerifyCredential` method should return `true` or `false` based on credential validity. Detailed failure information can be returned via an error.

#### Examples

To integrate user authentication into your application, refer to the examples below:

- [go-postgresql](https://github.com/cybergarage/go-postgresql) ![](https://img.shields.io/github/v/tag/cybergarage/go-postgresql)
  - [Server::receive()](https://github.com/cybergarage/go-postgresql/blob/master/postgresql/protocol/server_impl.go)
- [go-mysql](https://github.com/cybergarage/go-mysql) ![](https://img.shields.io/github/v/tag/cybergarage/go-mysql)
  - [Server::receive()](https://github.com/cybergarage/go-mysql/blob/main/mysql/protocol/server.go)
- [go-redis](https://github.com/cybergarage/go-redis) ![](https://img.shields.io/github/v/tag/cybergarage/go-redis)
  - [Server::Auth()](https://github.com/cybergarage/go-redis/blob/main/redis/server_auth.go)

### SASL Authentication

**go-authenticator** includes the `go-sasl` package, providing SASL mechanisms for authentication. For details on using the SASL API, see the `go-sasl` documentation.

- [go-sasl](https://github.com/cybergarage/go-sasl) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cybergarage/go-sasl)


#### Examples

For SASL authentication integration, refer to the examples below:

- [go-mongo](https://github.com/cybergarage/go-mongo) ![](https://img.shields.io/github/v/tag/cybergarage/go-mongo)
  - [BaseCommandExecutor::ExecuteCommand()](https://github.com/cybergarage/go-mongo/blob/master/mongo/command_base_executor.go)

### Certificate Authentication

This section explains how to authenticate users via TLS certificates using the `CertificateAuthenticator` interface.

#### CertificateAuthenticator

`CertificateAuthenticator` verifies users by inspecting the TLS connection certificate.

```go
type CertificateAuthenticator interface {
    VerifyCertificate(conn tls.Conn) (bool, error)
}
```

##### Creating a CertificateAuthenticator

To create a `CertificateAuthenticator`, use the `NewCertificateAuthenticator` function. This authenticates users based on common names (CN) in TLS certificates.

##### Enabling Certificate Authentication

Enable certificate authentication by setting the `CertificateAuthenticator` instance via `SetCertificateAuthenticator`.

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

By following these steps, you can enhance application security through TLS certificate verification.

#### Examples

For certificate authentication integration, refer to the examples below:

- [go-postgresql](https://github.com/cybergarage/go-postgresql) ![](https://img.shields.io/github/v/tag/cybergarage/go-postgresql)
  - [Server::receive()](https://github.com/cybergarage/go-postgresql/blob/master/postgresql/protocol/server_impl.go)
- [go-mysql](https://github.com/cybergarage/go-mysql) ![](https://img.shields.io/github/v/tag/cybergarage/go-mysql)
  - [Server::receive()](https://github.com/cybergarage/go-mysql/blob/main/mysql/protocol/server.go)
- [go-mongo](https://github.com/cybergarage/go-mongo) ![](https://img.shields.io/github/v/tag/cybergarage/go-mongo)
  - [Server::serve()](https://github.com/cybergarage/go-mongo/blob/master/mongo/server.go)
- [go-redis](https://github.com/cybergarage/go-redis) ![](https://img.shields.io/github/v/tag/cybergarage/go-redis)
  - [Server::tlsServe()](https://github.com/cybergarage/go-redis/blob/main/redis/server_impl.go)
