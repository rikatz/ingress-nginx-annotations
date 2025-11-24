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

package proxy

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	proxyConnectTimeoutAnnotation      = "proxy-connect-timeout"
	proxySendTimeoutAnnotation         = "proxy-send-timeout"
	proxyReadTimeoutAnnotation         = "proxy-read-timeout"
	proxyBuffersNumberAnnotation       = "proxy-buffers-number"
	proxyBufferSizeAnnotation          = "proxy-buffer-size"
	proxyBusyBuffersSizeAnnotation     = "proxy-busy-buffers-size"
	proxyCookiePathAnnotation          = "proxy-cookie-path"
	proxyCookieDomainAnnotation        = "proxy-cookie-domain"
	proxyBodySizeAnnotation            = "proxy-body-size"
	proxyNextUpstreamAnnotation        = "proxy-next-upstream"
	proxyNextUpstreamTimeoutAnnotation = "proxy-next-upstream-timeout"
	proxyNextUpstreamTriesAnnotation   = "proxy-next-upstream-tries"
	proxyRequestBufferingAnnotation    = "proxy-request-buffering"
	proxyRedirectFromAnnotation        = "proxy-redirect-from"
	proxyRedirectToAnnotation          = "proxy-redirect-to"
	proxyBufferingAnnotation           = "proxy-buffering"
	proxyHTTPVersionAnnotation         = "proxy-http-version"
	proxyMaxTempFileSizeAnnotation     = "proxy-max-temp-file-size" //#nosec G101
)

var validUpstreamAnnotation = regexp.MustCompile(`^((error|timeout|invalid_header|http_500|http_502|http_503|http_504|http_403|http_404|http_429|non_idempotent|off)\s?)+$`)

var proxyAnnotations = parser.Annotation{
	Group: "backend",
	Annotations: parser.AnnotationFields{
		proxyConnectTimeoutAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation allows setting the timeout in seconds of the connect operation to the backend.`,
		},
		proxySendTimeoutAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation allows setting the timeout in seconds of the send operation to the backend.`,
		},
		proxyReadTimeoutAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation allows setting the timeout in seconds of the read operation to the backend.`,
		},
		proxyBuffersNumberAnnotation: {
			Validator: parser.ValidateInt,
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `This annotation sets the number of the buffers in proxy_buffers used for reading the first part of the response received from the proxied server. 
			By default proxy buffers number is set as 4`,
		},
		proxyBufferSizeAnnotation: {
			Validator: parser.ValidateRegex(parser.SizeRegex, true),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `This annotation sets the size of the buffer proxy_buffer_size used for reading the first part of the response received from the proxied server. 
			By default proxy buffer size is set as "4k".`,
		},
		proxyBusyBuffersSizeAnnotation: {
			Validator:     parser.ValidateRegex(parser.SizeRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation limits the total size of buffers that can be busy sending a response to the client while the response is not yet fully read.`,
		},
		proxyCookiePathAnnotation: {
			Validator:     parser.ValidateRegex(parser.URLIsValidRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation sets a text that should be changed in the path attribute of the "Set-Cookie" header fields of a proxied server response.`,
		},
		proxyCookieDomainAnnotation: {
			Validator:     parser.ValidateRegex(parser.BasicCharsRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation ets a text that should be changed in the domain attribute of the "Set-Cookie" header fields of a proxied server response.`,
		},
		proxyBodySizeAnnotation: {
			Validator:     parser.ValidateRegex(parser.SizeRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation allows setting the maximum allowed size of a client request body.`,
		},
		proxyNextUpstreamAnnotation: {
			Validator: parser.ValidateRegex(validUpstreamAnnotation, false),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation defines when the next upstream should be used. 
			This annotation reflect the directive https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_next_upstream 
			and only the allowed values on upstream are allowed here.`,
		},
		proxyNextUpstreamTimeoutAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation limits the time during which a request can be passed to the next server`,
		},
		proxyNextUpstreamTriesAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation limits the number of possible tries for passing a request to the next server`,
		},
		proxyRequestBufferingAnnotation: {
			Validator:     parser.ValidateOptions([]string{"on", "off"}, true, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables or disables buffering of a client request body.`,
		},
		proxyRedirectFromAnnotation: {
			Validator:     parser.ValidateRegex(parser.URLIsValidRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `The annotations proxy-redirect-from and proxy-redirect-to will set the first and second parameters of NGINX's proxy_redirect directive respectively`,
		},
		proxyRedirectToAnnotation: {
			Validator:     parser.ValidateRegex(parser.URLIsValidRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `The annotations proxy-redirect-from and proxy-redirect-to will set the first and second parameters of NGINX's proxy_redirect directive respectively`,
		},
		proxyBufferingAnnotation: {
			Validator:     parser.ValidateOptions([]string{"on", "off"}, true, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables or disables buffering of responses from the proxied server. It can be "on" or "off"`,
		},
		proxyHTTPVersionAnnotation: {
			Validator:     parser.ValidateOptions([]string{"1.0", "1.1"}, true, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotations sets the HTTP protocol version for proxying. Can be "1.0" or "1.1".`,
		},
		proxyMaxTempFileSizeAnnotation: {
			Validator:     parser.ValidateRegex(parser.SizeRegex, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation defines the maximum size of a temporary file when buffering responses.`,
		},
	},
}
