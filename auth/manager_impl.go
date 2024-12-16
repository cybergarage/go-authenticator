// Copyright (C) 2019 The go-sasl Authors. All rights reserved.
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
	"errors"
)

type manager struct {
	certAuthenticator CertificateAuthenticator
}

// NewManager returns a new auth manager instance.
func NewManager() Manager {
	mgr := &manager{
		certAuthenticator: NewDefaultCertificateAuthenticator(),
	}
	return mgr
}

// SetCertificateAuthenticator sets the certificate authenticator.
func (mgr *manager) SetCertificateAuthenticator(auth CertificateAuthenticator) {
	mgr.certAuthenticator = auth
}

// VerifyCertificate verifies the client certificate.
func (mgr *manager) VerifyCertificate(conn Conn, state tls.ConnectionState) (bool, error) {
	if mgr.certAuthenticator == nil {
		return false, errors.New("no certificate authenticator")
	}
	return mgr.certAuthenticator.VerifyCertificate(conn, state)
}
