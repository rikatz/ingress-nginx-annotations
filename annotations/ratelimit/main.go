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

package ratelimit

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	limitRateAnnotation                = "limit-rate"
	limitRateAfterAnnotation           = "limit-rate-after"
	limitRateRPMAnnotation             = "limit-rpm"
	limitRateRPSAnnotation             = "limit-rps"
	limitRateConnectionsAnnotation     = "limit-connections"
	limitRateBurstMultiplierAnnotation = "limit-burst-multiplier"
	limitWhitelistAnnotation           = "limit-whitelist" // This annotation is an alias for limit-allowlist
	limitAllowlistAnnotation           = "limit-allowlist"
)

var rateLimitAnnotations = parser.Annotation{
	Group: "rate-limit",
	Annotations: parser.AnnotationFields{
		limitRateAnnotation: {
			Validator: parser.ValidateInt,
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `Limits the rate of response transmission to a client. The rate is specified in bytes per second. 
			The zero value disables rate limiting. The limit is set per a request, and so if a client simultaneously opens two connections, the overall rate will be twice as much as the specified limit.
			References: https://nginx.org/en/docs/http/ngx_http_core_module.html#limit_rate`,
		},
		limitRateAfterAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `Sets the initial amount after which the further transmission of a response to a client will be rate limited.`,
		},
		limitRateRPMAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `Requests per minute that will be allowed.`,
		},
		limitRateRPSAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `Requests per second that will be allowed.`,
		},
		limitRateConnectionsAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `Number of connections that will be allowed`,
		},
		limitRateBurstMultiplierAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation: `Burst multiplier for a limit-rate enabled location.`,
		},
		limitAllowlistAnnotation: {
			Validator:         parser.ValidateCIDRs,
			Scope:             parser.AnnotationScopeLocation,
			Risk:              parser.AnnotationRiskLow, // Low, as it allows just a set of options
			Documentation:     `List of CIDR/IP addresses that will not be rate-limited.`,
			AnnotationAliases: []string{limitWhitelistAnnotation},
		},
	},
}
