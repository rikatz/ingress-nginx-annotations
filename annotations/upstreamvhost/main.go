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

package upstreamvhost

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

const (
	upstreamVhostAnnotation = "upstream-vhost"
)

var UpstreamVhostAnnotations = parser.Annotation{
	Group: "backend",
	Annotations: parser.AnnotationFields{
		upstreamVhostAnnotation: {
			Validator: parser.ValidateServerName,
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow, // Low, as it allows regexes but on a very limited set
			Documentation: `This configuration setting allows you to control the value for host in the following statement: proxy_set_header Host $host, which forms part of the location block. 
			This is useful if you need to call the upstream server by something other than $host`,
		},
	},
}
