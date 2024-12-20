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
	"github.com/cybergarage/go-authenticator/auth/tls"
	"github.com/cybergarage/go-sasl/sasl"
)

type manager struct {
	sasl.Server
	certAuthenticator CertificateAuthenticator
}

// NewManager returns a new manager.
func NewManager() Manager {
	return &manager{
		certAuthenticator: nil,
		Server:            sasl.NewServer(),
	}
}

// SetCertificateAuthenticator sets the certificate authenticator.
func (mgr *manager) SetCertificateAuthenticator(auth CertificateAuthenticator) {
	mgr.certAuthenticator = auth
}

// VerifyCertificate verifies the client certificate. If the certificate authenticator is not set, it returns true.
func (mgr *manager) VerifyCertificate(conn tls.Conn) (bool, error) {
	if mgr.certAuthenticator == nil {
		return true, nil
	}
	return mgr.certAuthenticator.VerifyCertificate(conn)
}
