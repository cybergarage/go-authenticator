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
	"github.com/cybergarage/go-sasl/sasl/auth"
)

// Query represents a query.
type Query = auth.Query

// QueryOptionFn represents an option function for a query.
type QueryOptionFn = auth.QueryOptionFn

// QueryOption represents an option for a query.
type QueryOption = auth.QueryOption

// EncryptFunc represents an encrypt function.
type EncryptFunc = auth.EncryptFunc

// NewQuery returns a new query with options.
func NewQuery(opts ...QueryOptionFn) (Query, error) {
	return auth.NewQuery(opts...)
}

// WithQueryGroup returns an option to set the group.
func WithQueryGroup(group string) QueryOptionFn {
	return auth.WithQueryGroup(group)
}

// WithQueryUsername returns an option to set the username.
func WithQueryUsername(username string) QueryOptionFn {
	return auth.WithQueryUsername(username)
}

// WithQueryPassword returns an option to set the password.
func WithQueryPassword(password string) QueryOptionFn {
	return auth.WithQueryPassword(password)
}

// WithQueryMechanism returns an option to set the mechanism.
func WithQueryMechanism(mech string) QueryOptionFn {
	return auth.WithQueryMechanism(mech)
}

// WithQueryEncryptFunc returns an option to set the encrypt function.
func WithQueryEncryptFunc(encryptFunc EncryptFunc) QueryOptionFn {
	return auth.WithQueryEncryptFunc(encryptFunc)
}

// WithQueryOptions returns an option to set the options.
func WithQueryOptions(opts ...any) QueryOptionFn {
	return auth.WithQueryOptions(opts)
}
