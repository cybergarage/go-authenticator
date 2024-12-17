// Copyright (C) 2024 The go-authenticator Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"crypto/tls"

	"github.com/cybergarage/go-sasl/sasl"
	"github.com/cybergarage/go-sasl/sasl/auth"
)

type CredentialStore = auth.CredentialStore

type Manager interface {
	// Version returns the version.
	Version() string
	// Mechanisms returns the mechanisms.
	Mechanisms() []sasl.Mechanism
	// Mechanism returns a mechanism by name.
	Mechanism(name string) (sasl.Mechanism, error)
	// SetCredentialAuthenticator sets the credential authenticator.
	SetCredentialAuthenticator(auth CredentialAuthenticator)
	// SetCredentialStore sets the credential store.
	SetCredentialStore(store CredentialStore)
	// CredentialStore returns the credential store.
	CredentialStore() CredentialStore
	// VerifyCredential verifies the client credential.
	VerifyCredential(conn auth.Conn, q auth.Query) (bool, error)
	// SetCertificateAuthenticator sets the certificate authenticator.
	SetCertificateAuthenticator(auth CertificateAuthenticator)
	// VerifyCertificate verifies the client certificate.
	VerifyCertificate(conn Conn, state tls.ConnectionState) (bool, error)
}
