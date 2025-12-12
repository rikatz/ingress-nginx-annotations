/*
Copyright 2019 The Kubernetes Authors.

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

package proxyssl

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

var (
	proxySSLOnOffRegex    = regexp.MustCompile(`^(on|off)$`)
	proxySSLProtocolRegex = regexp.MustCompile(`^(TLSv1\.2|TLSv1\.3| )*$`)
	proxySSLCiphersRegex  = regexp.MustCompile(`^[A-Za-z0-9\+:\_\-!]*$`)
)

const (
	proxySSLSecretAnnotation      = "proxy-ssl-secret"
	proxySSLCiphersAnnotation     = "proxy-ssl-ciphers"
	proxySSLProtocolsAnnotation   = "proxy-ssl-protocols"
	proxySSLNameAnnotation        = "proxy-ssl-name"
	proxySSLVerifyAnnotation      = "proxy-ssl-verify"
	proxySSLVerifyDepthAnnotation = "proxy-ssl-verify-depth"
	proxySSLServerNameAnnotation  = "proxy-ssl-server-name"
)

var ProxySSLAnnotation = parser.Annotation{
	Group: "proxy",
	Annotations: parser.AnnotationFields{
		proxySSLSecretAnnotation: {
			Validator: parser.ValidateRegex(parser.BasicCharsRegex, true),
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation specifies a Secret with the certificate tls.crt, key tls.key in PEM format used for authentication to a proxied HTTPS server. 
			It should also contain trusted CA certificates ca.crt in PEM format used to verify the certificate of the proxied HTTPS server. 
			This annotation expects the Secret name in the form "namespace/secretName"
			Just secrets on the same namespace of the ingress can be used.`,
			GatewayAPI:    "Partially supported by the BackendTLSPolicy '.spec.validations.caCertificateRefs'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#backendtlspolicyvalidation",
		},
		proxySSLCiphersAnnotation: {
			Validator: parser.ValidateRegex(proxySSLCiphersRegex, true),
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation Specifies the enabled ciphers for requests to a proxied HTTPS server. 
			The ciphers are specified in the format understood by the OpenSSL library.`,
		},
		proxySSLProtocolsAnnotation: {
			Validator:     parser.ValidateRegex(proxySSLProtocolRegex, true),
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables the specified protocols for requests to a proxied HTTPS server.`,
		},
		proxySSLNameAnnotation: {
			Validator: parser.ValidateServerName,
			Scope:     parser.AnnotationScopeIngress,
			Risk:      parser.AnnotationRiskHigh,
			Documentation: `This annotation allows to set proxy_ssl_name. This allows overriding the server name used to verify the certificate of the proxied HTTPS server. 
			This value is also passed through SNI when a connection is established to the proxied HTTPS server.`,
			GatewayAPI:    "Supported by the BackendTLSPolicy '.spec.validations.hostname'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#backendtlspolicyvalidation",
		},
		proxySSLVerifyAnnotation: {
			Validator:     parser.ValidateRegex(proxySSLOnOffRegex, true),
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables or disables verification of the proxied HTTPS server certificate. (default: off)`,
		},
		proxySSLVerifyDepthAnnotation: {
			Validator:     parser.ValidateInt,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation Sets the verification depth in the proxied HTTPS server certificates chain. (default: 1).`,
		},
		proxySSLServerNameAnnotation: {
			Validator:     parser.ValidateRegex(proxySSLOnOffRegex, true),
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables passing of the server name through TLS Server Name Indication extension (SNI, RFC 6066) when establishing a connection with the proxied HTTPS server.`,
			GatewayAPI:    "Supported by the BackendTLSPolicy '.spec.validations.hostname'",
			GatewayAPIRef: "https://gateway-api.sigs.k8s.io/reference/spec/#backendtlspolicyvalidation",
		},
	},
}
