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

package cors

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	// Default values
	defaultCorsMethods = "GET, PUT, POST, DELETE, PATCH, OPTIONS"
	defaultCorsHeaders = "DNT,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization"
	defaultCorsMaxAge  = 1728000
)

var (
	// Regex are defined here to prevent information leak, if user tries to set anything not valid
	// that could cause the Response to contain some internal value/variable (like returning $pid, $upstream_addr, etc)
	// Origin must contain a http/s Origin (including or not the port) or the value '*'
	// This Regex is composed of the following:
	// * Sets a group that can be (https?://)?*?.something.com:port? OR null
	// * Allows this to be repeated as much as possible, and separated by comma
	// Otherwise it should be '*'
	corsOriginRegexValidator = regexp.MustCompile(`^((((([a-z]+://)?(\*\.)?[A-Za-z0-9\-.]*(:\d+)?,?)|null)+)|\*)?$`)
	// corsOriginRegex defines the regex for validation inside Parse
	corsOriginRegex = regexp.MustCompile(`^([a-z]+://(\*\.)?[A-Za-z0-9\-.]*(:\d+)?|\*|null)?$`)
	// Method must contain valid methods list (PUT, GET, POST, BLA)
	// May contain or not spaces between each verb
	corsMethodsRegex = regexp.MustCompile(`^([A-Za-z]+,?\s?)+$`)
	// Expose Headers must contain valid values only (*, X-HEADER12, X-ABC)
	// May contain or not spaces between each Header
	corsExposeHeadersRegex = regexp.MustCompile(`^(([A-Za-z0-9\-\_]+|\*),?\s?)+$`)
)

const (
	corsEnableAnnotation           = "enable-cors"
	corsAllowOriginAnnotation      = "cors-allow-origin"
	corsAllowHeadersAnnotation     = "cors-allow-headers"
	corsAllowMethodsAnnotation     = "cors-allow-methods"
	corsAllowCredentialsAnnotation = "cors-allow-credentials" //#nosec G101
	corsExposeHeadersAnnotation    = "cors-expose-headers"
	corsMaxAgeAnnotation           = "cors-max-age"
)

var CORSAnnotation = parser.Annotation{
	Group: "cors",
	Annotations: parser.AnnotationFields{
		corsEnableAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables Cross-Origin Resource Sharing (CORS) in an Ingress rule`,
		},
		corsAllowOriginAnnotation: {
			Validator: parser.ValidateRegex(corsOriginRegexValidator, true),
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation controls what's the accepted Origin for CORS.
			This is a multi-valued field, separated by ','. It must follow this format: protocol://origin-site.com, protocol://origin-site.com:port, null, or *.
			It also supports single level wildcard subdomains and follows this format: https://*.foo.bar, http://*.bar.foo:8080 or myprotocol://*.abc.bar.foo:9000
			Protocol can be any lowercase string, like http, https, or mycustomprotocol.`,
		},
		corsAllowHeadersAnnotation: {
			Validator: parser.ValidateRegex(parser.HeadersVariable, true),
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation controls which headers are accepted.
			This is a multi-valued field, separated by ',' and accepts letters, numbers, _ and -`,
		},
		corsAllowMethodsAnnotation: {
			Validator: parser.ValidateRegex(corsMethodsRegex, true),
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation controls which methods are accepted.
			This is a multi-valued field, separated by ',' and accepts only letters (upper and lower case)`,
		},
		corsAllowCredentialsAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation controls if credentials can be passed during CORS operations.`,
		},
		corsExposeHeadersAnnotation: {
			Validator: parser.ValidateRegex(corsExposeHeadersRegex, true),
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation controls which headers are exposed to response.
			This is a multi-valued field, separated by ',' and accepts letters, numbers, _, - and *.`,
		},
		corsMaxAgeAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation controls how long, in seconds, preflight requests can be cached.`,
		},
	},
}
