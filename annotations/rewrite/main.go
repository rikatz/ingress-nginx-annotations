/*
Copyright 2016 The Kubernetes Authors.

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

package rewrite

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	rewriteTargetAnnotation         = "rewrite-target"
	sslRedirectAnnotation           = "ssl-redirect"
	preserveTrailingSlashAnnotation = "preserve-trailing-slash"
	forceSSLRedirectAnnotation      = "force-ssl-redirect"
	useRegexAnnotation              = "use-regex"
	appRootAnnotation               = "app-root"
)

var RewriteAnnotations = parser.Annotation{
	Group: "rewrite",
	Annotations: parser.AnnotationFields{
		rewriteTargetAnnotation: {
			Validator: parser.ValidateRegex(parser.RegexPathWithCapture, false),
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation allows to specify the target URI where the traffic must be redirected. It can contain regular characters and captured 
			groups specified as '$1', '$2', etc.`,
			GatewayAPI:    "Partially supported by HTTPRoute 'spec.rules[].filters[].urlRewrite'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		sslRedirectAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation defines if the location section is only accessible via SSL`,
			GatewayAPI:    "Supported by HTTPRoute 'spec.rules[].filters[].requestRedirect'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		preserveTrailingSlashAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation defines if the trailing slash should be preserved in the URI with 'ssl-redirect'`,
		},
		forceSSLRedirectAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation forces the redirection to HTTPS even if the Ingress is not TLS Enabled`,
			GatewayAPI:    "Supported by HTTPRoute 'spec.rules[].filters[].requestRedirect'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		useRegexAnnotation: {
			Validator: parser.ValidateBool,
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `This annotation defines if the paths defined on an Ingress use regular expressions. To use regex on path
			the pathType should also be defined as 'ImplementationSpecific'.`,
		},
		appRootAnnotation: {
			Validator:     parser.ValidateRegex(parser.RegexPathWithCapture, false),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation defines the Application Root that the Controller must redirect if it's in / context`,
			GatewayAPI:    "Partially supported by HTTPRoute 'spec.rules[].filters[].urlRewrite'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
	},
}
