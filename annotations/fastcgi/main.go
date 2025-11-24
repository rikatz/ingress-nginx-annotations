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

package fastcgi

import (
	"regexp"

	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	fastCGIIndexAnnotation  = "fastcgi-index"
	fastCGIParamsAnnotation = "fastcgi-params-configmap" //#nosec G101
)

// fast-cgi valid parameters is just a single file name (like index.php)
var (
	regexValidIndexAnnotationAndKey = regexp.MustCompile(`^[A-Za-z0-9.\-\_]+$`)
	validFCGIValue                  = regexp.MustCompile(`^[A-Za-z0-9\-\_\$\{\}/.]*$`)
)

var fastCGIAnnotations = parser.Annotation{
	Group: "fastcgi",
	Annotations: parser.AnnotationFields{
		fastCGIIndexAnnotation: {
			Validator:     parser.ValidateRegex(regexValidIndexAnnotationAndKey, true),
			Scope:         parser.AnnotationScopeLocation,
			Risk:          parser.AnnotationRiskMedium,
			Documentation: `This annotation can be used to specify an index file`,
		},
		fastCGIParamsAnnotation: {
			Validator: parser.ValidateRegex(parser.BasicCharsRegex, true),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskMedium,
			Documentation: `This annotation can be used to specify a ConfigMap containing the fastcgi parameters as a key/value.
			Only ConfigMaps on the same namespace of ingress can be used. They key and value from ConfigMap are validated for unauthorized characters.`,
		},
	},
}
