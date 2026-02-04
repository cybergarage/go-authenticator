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
	"regexp"

	"github.com/cybergarage/go-authenticator/auth/tls"
)

type certificateAuthenticator struct {
	commonNameRegexp []*regexp.Regexp
}

// CertificateAuthenticatorOption is a function to set the certificate authenticator options.
type CertificateAuthenticatorOption = func(*certificateAuthenticator) error

// WithCommonNameRegexp sets the common name regular expressions to the certificate authenticator.
func WithCommonNameRegexp(regexps ...string) CertificateAuthenticatorOption {
	return func(ca *certificateAuthenticator) error {
		for _, re := range regexps {
			r, err := regexp.Compile(re)
			if err != nil {
				return err
			}
			ca.commonNameRegexp = append(ca.commonNameRegexp, r)
		}
		return nil
	}
}

// NewCertificateAuthenticator returns a new certificate authenticator with the options.
func NewCertificateAuthenticator(opts ...CertificateAuthenticatorOption) (CertificateAuthenticator, error) {
	ca := &certificateAuthenticator{
		commonNameRegexp: []*regexp.Regexp{},
	}
	for _, opt := range opts {
		if err := opt(ca); err != nil {
			return nil, err
		}
	}
	return ca, nil
}

// VerifyCertificate verifies the client certificate.
func (ca *certificateAuthenticator) VerifyCertificate(conn tls.Conn) (bool, error) {
	state := conn.ConnectionState()
	for _, cert := range state.PeerCertificates {
		for _, re := range ca.commonNameRegexp {
			if re.MatchString(cert.Subject.CommonName) {
				return true, nil
			}
		}
	}
	return false, nil
}
