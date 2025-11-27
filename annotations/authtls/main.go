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

package authtls

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	annotationAuthTLSSecret             = "auth-tls-secret" //#nosec G101
	annotationAuthTLSVerifyClient       = "auth-tls-verify-client"
	annotationAuthTLSVerifyDepth        = "auth-tls-verify-depth"
	annotationAuthTLSErrorPage          = "auth-tls-error-page"
	annotationAuthTLSPassCertToUpstream = "auth-tls-pass-certificate-to-upstream" //#nosec G101
	annotationAuthTLSMatchCN            = "auth-tls-match-cn"
)

var (
	authVerifyClientRegex = regexp.MustCompile(`^(on|off|optional|optional_no_ca)$`)
	redirectRegex         = regexp.MustCompile(`^(@[A-Za-z0-9_-]+|((https?://)?[A-Za-z0-9\-.]+(:\d+)?)?(/[A-Za-z0-9\-_.]+)*/?)$`)
)

var AuthTLSAnnotations = parser.Annotation{
	Group: "authentication",
	Annotations: parser.AnnotationFields{
		annotationAuthTLSSecret: {
			Validator:     parser.ValidateRegex(parser.BasicCharsRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium, // Medium as it allows a subset of chars
			Documentation: `This annotation defines the secret that contains the certificate chain of allowed certs`,
		},
		annotationAuthTLSVerifyClient: {
			Validator:     parser.ValidateRegex(authVerifyClientRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium, // Medium as it allows a subset of chars
			Documentation: `This annotation enables verification of client certificates. Can be "on", "off", "optional" or "optional_no_ca"`,
		},
		annotationAuthTLSVerifyDepth: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation defines validation depth between the provided client certificate and the Certification Authority chain.`,
		},
		annotationAuthTLSErrorPage: {
			Validator:     parser.ValidateRegex(redirectRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskHigh,
			Documentation: `This annotation defines the URL/Page that user should be redirected in case of a Certificate Authentication Error`,
		},
		annotationAuthTLSPassCertToUpstream: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation defines if the received certificates should be passed or not to the upstream server in the header "ssl-client-cert"`,
		},
		annotationAuthTLSMatchCN: {
			Validator:     parser.CommonNameAnnotationValidator,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskHigh,
			Documentation: `This annotation adds a sanity check for the CN of the client certificate that is sent over using a string / regex starting with "CN="`,
		},
	},
}
