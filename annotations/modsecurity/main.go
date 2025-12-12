/*
Copyright 2018 The Kubernetes Authors.

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

package modsecurity

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	modsecEnableAnnotation          = "enable-modsecurity"
	modsecEnableOwaspCoreAnnotation = "enable-owasp-core-rules"
	modesecTransactionIDAnnotation  = "modsecurity-transaction-id"
	modsecSnippetAnnotation         = "modsecurity-snippet"
)

var ModsecurityAnnotation = parser.Annotation{
	Group: "modsecurity",
	Annotations: parser.AnnotationFields{
		modsecEnableAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables ModSecurity`,
			GatewayAPI:    "Not supported",
		},
		modsecEnableOwaspCoreAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables the OWASP Core Rule Set`,
			GatewayAPI:    "Not supported",
		},
		modesecTransactionIDAnnotation: {
			Validator:     parser.ValidateRegex(parser.NGINXVariable, true),
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskHigh,
			Documentation: `This annotation enables passing an NGINX variable to ModSecurity.`,
			GatewayAPI:    "Not supported",
		},
		modsecSnippetAnnotation: {
			Validator:     parser.ValidateNull,
			Scope:         parser.AnnotationScopeIngress,
			Risk:          parser.AnnotationRiskCritical,
			Documentation: `This annotation enables adding a specific snippet configuration for ModSecurity`,
			GatewayAPI:    "Not supported",
		},
	},
}
