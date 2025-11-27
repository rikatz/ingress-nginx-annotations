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

package parser

import (
	"fmt"
	"strings"
)

// DefaultAnnotationsPrefix defines the common prefix used in the nginx ingress controller
const (
	DefaultAnnotationsPrefix          = "nginx.ingress.kubernetes.io"
	DefaultEnableAnnotationValidation = true
)

var (
	// AnnotationsPrefix is the mutable attribute that the controller explicitly refers to
	AnnotationsPrefix = DefaultAnnotationsPrefix
	// Enable is the mutable attribute for enabling or disabling the validation functions
	EnableAnnotationValidation = DefaultEnableAnnotationValidation
)

// AnnotationGroup defines the group that this annotation may belong
// eg.: Security, Snippets, Rewrite, etc
type AnnotationGroup string

// AnnotationScope defines which scope this annotation applies. May be to the whole
// ingress, per location, etc
type AnnotationScope string

var (
	AnnotationScopeLocation AnnotationScope = "location"
	AnnotationScopeIngress  AnnotationScope = "ingress"
)

type GatewayAPICompatibility string

var (
	GatewayAPICompatible   GatewayAPICompatibility = "Compatible"
	GatewayAPIPartial      GatewayAPICompatibility = "Partial"
	GatewayAPIIncompatible GatewayAPICompatibility = "Incompatible"
)

// AnnotationRisk is a subset of risk that an annotation may represent.
// Based on the Risk, the admin will be able to allow or disallow users to set it
// on their ingress objects
type AnnotationRisk int

type AnnotationFields map[string]AnnotationConfig

// AnnotationConfig defines the configuration that a single annotation field
// has, with the Validator and the documentation of this field.
type AnnotationConfig struct {
	// Validator defines a function to validate the annotation value
	Validator AnnotationValidator
	// Documentation defines a user facing documentation for this annotation. This
	// field will be used to auto generate documentations
	Documentation string
	// Risk defines a risk of this annotation being exposed to the user. Annotations
	// with bool fields, or to set timeout are usually low risk. Annotations that allows
	// string input without a limited set of options may represent a high risk
	Risk AnnotationRisk

	// Scope defines which scope this annotation applies, may be to location, to an Ingress object, etc
	Scope AnnotationScope

	// AnnotationAliases defines other names this annotation may have.
	AnnotationAliases []string
}

// Annotation defines an annotation feature an Ingress may have.
// It should contain the internal resolver, and all the annotations
// with configs and Validators that should be used for each Annotation
type Annotation struct {
	// Annotations contains all the annotations that belong to this feature
	Annotations AnnotationFields
	// Group defines which annotation group this feature belongs to
	Group AnnotationGroup
}

// GetAnnotationWithPrefix returns the prefix of ingress annotations
func GetAnnotationWithPrefix(suffix string) string {
	return fmt.Sprintf("%v/%v", AnnotationsPrefix, suffix)
}

func TrimAnnotationPrefix(annotation string) string {
	return strings.TrimPrefix(annotation, AnnotationsPrefix+"/")
}

func (a AnnotationRisk) ToString() string {
	switch a {
	case AnnotationRiskCritical:
		return "Critical"
	case AnnotationRiskHigh:
		return "High"
	case AnnotationRiskMedium:
		return "Medium"
	case AnnotationRiskLow:
		return "Low"
	default:
		return "Unknown"
	}
}
