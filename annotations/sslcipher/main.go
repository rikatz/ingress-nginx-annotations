/*
Copyright 2017 The Kubernetes Authors.

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

package sslcipher

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	sslPreferServerCipherAnnotation = "ssl-prefer-server-ciphers"
	sslCipherAnnotation             = "ssl-ciphers"
)

// Should cover something like "ALL:!aNULL:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP"
// (?:@STRENGTH) is included twice so it can appear before or after @SECLEVEL=n
var regexValidSSLCipher = regexp.MustCompile(`^(?:(?:[A-Za-z0-9!:+\-_])*(?:@STRENGTH)*(?:@SECLEVEL=[0-5])*(?:@STRENGTH)*)*$`)

var sslCipherAnnotations = parser.Annotation{
	Group: "backend",
	Annotations: parser.AnnotationFields{
		sslPreferServerCipherAnnotation: {
			Validator: parser.ValidateBool,
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `The following annotation will set the ssl_prefer_server_ciphers directive at the server level. 
			This configuration specifies that server ciphers should be preferred over client ciphers when using the TLS protocols.`,
		},
		sslCipherAnnotation: {
			Validator:     parser.ValidateRegex(regexValidSSLCipher, true),
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `Using this annotation will set the ssl_ciphers directive at the server level. This configuration is active for all the paths in the host.`,
		},
	},
}
