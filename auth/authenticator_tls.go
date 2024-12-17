// Copyright (C) 2024 The go-mysql Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"crypto/tls"
)

type certificateAuthenticator struct {
	commonName string
}

// CertificateAuthenticatorOption is a function to set the certificate authenticator options.
type CertificateAuthenticatorOption = func(*certificateAuthenticator)

// WithCertificateAuthenticatorCommonName sets the common name.
func WithCertificateAuthenticatorCommonName(name string) func(*certificateAuthenticator) {
	return func(ca *certificateAuthenticator) {
		ca.commonName = name
	}
}

// NewDefaultCertificateAuthenticator creates a new defaultTLSAuthenticator.
func NewDefaultCertificateAuthenticator() CertificateAuthenticator {
	return NewCertificateAuthenticatorWith()
}

// NewCertificateAuthenticator creates a new certificate authenticator.
func NewCertificateAuthenticatorWith(opts ...CertificateAuthenticatorOption) CertificateAuthenticator {
	ca := &certificateAuthenticator{
		commonName: "",
	}
	for _, opt := range opts {
		opt(ca)
	}
	return ca
}

// VerifyCertificate verifies the client certificate.
func (ca *certificateAuthenticator) VerifyCertificate(conn Conn, state *tls.ConnectionState) (bool, error) {
	if len(ca.commonName) == 0 {
		return true, nil
	}
	for _, cert := range state.PeerCertificates {
		if cert.Subject.CommonName == ca.commonName {
			return true, nil
		}
	}
	return false, nil
}
