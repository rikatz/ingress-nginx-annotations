/*
Copyright 2022 The Kubernetes Authors.

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

package opentelemetry

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	enableOpenTelemetryAnnotation = "enable-opentelemetry"
	otelTrustSpanAnnotation       = "opentelemetry-trust-incoming-span"
	otelOperationNameAnnotation   = "opentelemetry-operation-name"
)

var regexOperationName = regexp.MustCompile(`^[A-Za-z0-9_\-]*$`)

var otelAnnotations = parser.Annotation{
	Group: "opentelemetry",
	Annotations: parser.AnnotationFields{
		enableOpenTelemetryAnnotation: {
			Validator: parser.ValidateBool,
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `This annotation defines if Open Telemetry collector should be enable for this location. OpenTelemetry should 
			already be configured by Ingress administrator`,
		},
		otelTrustSpanAnnotation: {
			Validator:     parser.ValidateBool,
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskLow,
			Documentation: `This annotation enables or disables using spans from incoming requests as parent for created ones`,
		},
		otelOperationNameAnnotation: {
			Validator:     parser.ValidateRegex(regexOperationName, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation defines what operation name should be added to the span`,
		},
	},
}
