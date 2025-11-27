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

package authreq

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	authReqURLAnnotation                = "auth-url"
	authReqMethodAnnotation             = "auth-method"
	authReqSigninAnnotation             = "auth-signin"
	authReqSigninRedirParamAnnotation   = "auth-signin-redirect-param"
	authReqSnippetAnnotation            = "auth-snippet"
	authReqCacheKeyAnnotation           = "auth-cache-key"
	authReqKeepaliveAnnotation          = "auth-keepalive"
	authReqKeepaliveShareVarsAnnotation = "auth-keepalive-share-vars"
	authReqKeepaliveRequestsAnnotation  = "auth-keepalive-requests"
	authReqKeepaliveTimeout             = "auth-keepalive-timeout"
	authReqCacheDuration                = "auth-cache-duration"
	authReqResponseHeadersAnnotation    = "auth-response-headers"
	authReqProxySetHeadersAnnotation    = "auth-proxy-set-headers"
	authReqRequestRedirectAnnotation    = "auth-request-redirect"
	authReqAlwaysSetCookieAnnotation    = "auth-always-set-cookie"

	// This should be exported as it is imported by other packages
	AuthSecretAnnotation = "auth-secret"
)

var (
	methodsRegex = regexp.MustCompile("(GET|HEAD|POST|PUT|PATCH|DELETE|CONNECT|OPTIONS|TRACE)")
)

var AuthReqAnnotations = parser.Annotation{
	Group: "authentication",
	Annotations: parser.AnnotationFields{
		authReqURLAnnotation: {
			Validator:     parser.ValidateRegex(parser.URLWithNginxVariableRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskHigh,
			Documentation: `This annotation allows to indicate the URL where the HTTP request should be sent`,
		},
		authReqMethodAnnotation: {
			Validator:     parser.ValidateRegex(methodsRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation allows to specify the HTTP method to use`,
		},
		authReqSigninAnnotation: {
			Validator:     parser.ValidateRegex(parser.URLWithNginxVariableRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskHigh,
			Documentation: `This annotation allows to specify the location of the error page`,
		},
		authReqSigninRedirParamAnnotation: {
			Validator:     parser.ValidateRegex(parser.URLIsValidRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation allows to specify the URL parameter in the error page which should contain the original URL for a failed signin request`,
		},
		authReqSnippetAnnotation: {
			Validator:     parser.ValidateNull,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskCritical,
			Documentation: `This annotation allows to specify a custom snippet to use with external authentication`,
		},
		authReqCacheKeyAnnotation: {
			Validator:     parser.ValidateRegex(parser.NGINXVariable, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation enables caching for auth requests.`,
		},
		authReqKeepaliveAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation specifies the maximum number of keepalive connections to auth-url. Only takes effect when no variables are used in the host part of the URL`,
		},
		authReqKeepaliveShareVarsAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation specifies whether to share Nginx variables among the current request and the auth request`,
		},
		authReqKeepaliveRequestsAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation defines the maximum number of requests that can be served through one keepalive connection`,
		},
		authReqKeepaliveTimeout: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation specifies a duration in seconds which an idle keepalive connection to an upstream server will stay open`,
		},
		authReqCacheDuration: {
			Validator:     parser.ValidateRegex(parser.ExtendedCharsRegex, false),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation allows to specify a caching time for auth responses based on their response codes, e.g. 200 202 30m`,
		},
		authReqResponseHeadersAnnotation: {
			Validator:     parser.ValidateRegex(parser.HeadersVariable, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation sets the headers to pass to backend once authentication request completes. They should be separated by comma.`,
		},
		authReqProxySetHeadersAnnotation: {
			Validator: parser.ValidateRegex(parser.BasicCharsRegex, true),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation sets the name of a ConfigMap that specifies headers to pass to the authentication service.
			Only ConfigMaps on the same namespace are allowed`,
		},
		authReqRequestRedirectAnnotation: {
			Validator:     parser.ValidateRegex(parser.URLIsValidRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation allows to specify the X-Auth-Request-Redirect header value`,
		},
		authReqAlwaysSetCookieAnnotation: {
			Validator: parser.ValidateBool,
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `This annotation enables setting a cookie returned by auth request. 
			By default, the cookie will be set only if an upstream reports with the code 200, 201, 204, 206, 301, 302, 303, 304, 307, or 308`,
		},
	},
}
