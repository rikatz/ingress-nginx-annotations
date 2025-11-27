/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package auth

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	authSecretTypeAnnotation = "auth-secret-type" //#nosec G101
	authRealmAnnotation      = "auth-realm"
	authTypeAnnotation       = "auth-type"
	// This should be exported as it is imported by other packages
	AuthSecretAnnotation = "auth-secret" //#nosec G101
)

var (
	authTypeRegex       = regexp.MustCompile(`basic|digest`)
	authSecretTypeRegex = regexp.MustCompile(`auth-file|auth-map`)

	// AuthDirectory default directory used to store files
	// to authenticate request
	AuthDirectory = "/etc/ingress-controller/auth"
)

var AuthSecretConfig = parser.AnnotationConfig{
	Validator:     parser.ValidateRegex(parser.BasicCharsRegex, true),
	Scope:         parser.AnnotationScopeLocation,
	Risk:          parser.AnnotationRiskMedium, // Medium as it allows a subset of chars
	Documentation: `This annotation defines the name of the Secret that contains the usernames and passwords which are granted access to the paths defined in the Ingress rules. `,
}

var AuthSecretAnnotations = parser.Annotation{
	Group: "authentication",
	Annotations: parser.AnnotationFields{
		AuthSecretAnnotation: AuthSecretConfig,
		authSecretTypeAnnotation: {
			Validator: parser.ValidateRegex(authSecretTypeRegex, true),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `This annotation what is the format of auth-secret value. Can be "auth-file" that defines the content of an htpasswd file, or "auth-map" where each key
			is a user and each value is the password.`,
		},
		authRealmAnnotation: {
			Validator:     parser.ValidateRegex(parser.CharsWithSpace, false),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium, // Medium as it allows a subset of chars
			Documentation: `This annotation defines the realm (message) that should be shown to user when authentication is requested.`,
		},
		authTypeAnnotation: {
			Validator:     parser.ValidateRegex(authTypeRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation defines the basic authentication type. Should be "basic" or "digest"`,
		},
	},
}
