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

// NewQuery returns a new query with options.
func NewQuery(opts ...QueryOptionFn) Query {
	return auth.NewQuery(opts...)
}

// WithGroup returns an option to set the group.
func WithGroup(group string) QueryOptionFn {
	return auth.WithQueryGroup(group)
}

// WithUsername returns an option to set the username.
func WithUsername(username string) QueryOptionFn {
	return auth.WithQueryUsername(username)
}

// WithPassword returns an option to set the password.
func WithPassword(password string) QueryOptionFn {
	return auth.WithQueryPassword(password)
}

// WithMechanism returns an option to set the mechanism.
func WithMechanism(mech string) QueryOptionFn {
	return auth.WithQueryMechanism(mech)
}
