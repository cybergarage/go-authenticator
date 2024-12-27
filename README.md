# go-authenticator

`go-authenticator` is an open-source framework for user authentication in Go applications. It supports multiple authentication methods, including username and password authentication, SASL (Simple Authentication and Security Layer) authentication, and certificate-based authentication.

![](doc/img/framework.png)

The `go-authenticator` framework is designed to be flexible and extensible, allowing seamless integration into existing applications. Its goal is to provide robust authentication functionality through a unified interface. The framework is actively used by the following projects:

- [PuzzleDB](https://github.com/cybergarage/puzzledb-go)
- [go-mysql](https://github.com/cybergarage/go-mysql)
- [go-postgresql](https://github.com/cybergarage/go-postgresql)
- [go-redis](https://github.com/cybergarage/go-redis)
- [go-mongo](https://github.com/cybergarage/go-mongo)

## Features

The `go-authenticator` framework offers the following features:

- **User Authentication**: Authenticate users using a username and password.
- **SASL Authentication**: Authenticate users via SASL (Simple Authentication and Security Layer).
- **Certificate Authentication**: Authenticate users based on the TLS connection certificate.

## Getting Started

`go-authenticator` provides an authentication manager to handle the authentication process. The manager can be configured with different authentication methods, such as credential authentication, SASL authentication, and certificate authentication.

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

`go-authenticator` includes a default credential authenticator that uses `CredentialStore`. You can set the `CredentialStore` by calling `Manager::SetCredentialStore`.

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

- [go-postgresql](https://github.com/cybergarage/go-postgresql)
  - [Server::receive()](https://github.com/cybergarage/go-postgresql/blob/master/postgresql/protocol/server_impl.go)
- [go-redis](https://github.com/cybergarage/go-redis)
  - [Server::Auth()](https://github.com/cybergarage/go-redis/blob/main/redis/server_auth.go)

### SASL Authentication

`go-authenticator` includes the `go-sasl` package, providing SASL mechanisms for authentication. For details on using the SASL API, see the `go-sasl` documentation.

- [go-sasl](https://github.com/cybergarage/go-sasl)

#### Examples

For SASL authentication integration, refer to the examples below:

- [go-mongo](https://github.com/cybergarage/go-mongo)
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

- [go-postgresql](https://github.com/cybergarage/go-postgresql)
  - [Server::receive()](https://github.com/cybergarage/go-postgresql/blob/master/postgresql/protocol/server_impl.go)
- [go-mysql](https://github.com/cybergarage/go-mysql)
  - [Server::receive()](https://github.com/cybergarage/go-mysql/blob/main/mysql/protocol/server.go)
- [go-mongo](https://github.com/cybergarage/go-mongo)
  - [Server::serve()](https://github.com/cybergarage/go-mongo/blob/master/mongo/server.go)
- [go-redis](https://github.com/cybergarage/go-redis)
  - [Server::tlsServe()](https://github.com/cybergarage/go-redis/blob/main/redis/server_impl.go)
