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

package redirect

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	fromToWWWRedirAnnotation        = "from-to-www-redirect"
	temporalRedirectAnnotation      = "temporal-redirect"
	temporalRedirectAnnotationCode  = "temporal-redirect-code"
	permanentRedirectAnnotation     = "permanent-redirect"
	permanentRedirectAnnotationCode = "permanent-redirect-code"
	relativeRedirectsAnnotation     = "relative-redirects"
)

var RedirectAnnotations = parser.Annotation{
	Group: "redirect",
	Annotations: parser.AnnotationFields{
		fromToWWWRedirAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `In some scenarios, it is required to redirect from www.domain.com to domain.com or vice versa, which way the redirect is performed depends on the configured host value in the Ingress object.`,
			GatewayAPI:    "Supported by HTTPRoute 'spec.rules[].filters[].requestRedirect'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		temporalRedirectAnnotation: {
			Validator: parser.ValidateRegex(parser.URLIsValidRegex, false),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskMedium, // Medium, as it allows arbitrary URLs that needs to be validated
			Documentation: `This annotation allows you to return a temporal redirect (Return Code 302) instead of sending data to the upstream. 
			For example setting this annotation to https://www.google.com would redirect everything to Google with a Return Code of 302 (Moved Temporarily).`,
			GatewayAPI:    "Supported by HTTPRoute 'spec.rules[].filters[].requestRedirect'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		temporalRedirectAnnotationCode: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `This annotation allows you to modify the status code used for temporal redirects.`,
			GatewayAPI:    "Supported by HTTPRoute 'spec.rules[].filters[].requestRedirect'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		permanentRedirectAnnotation: {
			Validator: parser.ValidateRegex(parser.URLIsValidRegex, false),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskMedium, // Medium, as it allows arbitrary URLs that needs to be validated
			Documentation: `This annotation allows to return a permanent redirect (Return Code 301) instead of sending data to the upstream. 
			For example setting this annotation https://www.google.com would redirect everything to Google with a code 301`,
			GatewayAPI:    "Supported by HTTPRoute 'spec.rules[].filters[].requestRedirect'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		permanentRedirectAnnotationCode: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `This annotation allows you to modify the status code used for permanent redirects.`,
			GatewayAPI:    "Supported by HTTPRoute 'spec.rules[].filters[].requestRedirect'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#httproutefilter",
		},
		relativeRedirectsAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `If enabled, redirects issued by nginx will be relative. See https://nginx.org/en/docs/http/ngx_http_core_module.html#absolute_redirect`,
		},
	},
}
