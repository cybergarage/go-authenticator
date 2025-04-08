// Copyright (C) 2020 The go-mongo Authors. All rights reserved.
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

package tls

import (
	"crypto/tls"
)

// CertConfig represents a TLS configuration interface.
type CertConfig interface {
	// SetClientAuthType sets a client authentication type.
	SetClientAuthType(authType tls.ClientAuthType)
	// SetServerKey sets a SSL server key.
	SetServerKey(key []byte)
	// SetServerCert sets a SSL server certificate.
	SetServerCert(cert []byte)
	// SetRootCerts sets a SSL root certificates.
	SetRootCerts(certs ...[]byte)
	// SetServerKeyFile loads a SSL server key file and sets it.
	SetServerKeyFile(file string) error
	// SetServerCertFile loads a SSL server certificate file and sets it.
	SetServerCertFile(file string) error
	// SetRootCertFile loads SSL root certificate files and sets them.
	SetRootCertFiles(files ...string) error
	// SetTLSConfig sets a TLS configuration directly.
	// If the provided configuration is nil, TLS will be disabled.
	SetTLSConfig(tlsConfig *tls.Config)
	// TLSConfig returns a TLS configuration from the configuration.
	TLSConfig() (*tls.Config, error)
}
