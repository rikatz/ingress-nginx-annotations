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

package loadbalancing

import (
	"github.com/rikatz/ingress-nginx-annotations/parser"
)

// LB Algorithms are defined in https://github.com/kubernetes/ingress-nginx/blob/d3e75b056f77be54e01bdb18675f1bb46caece31/rootfs/etc/nginx/lua/balancer.lua#L28

const (
	loadBalanceAlgorithmAnnotation = "load-balance"
)

var loadBalanceAlgorithms = []string{"round_robin", "chash", "chashsubset", "sticky_balanced", "sticky_persistent", "ewma"}

var loadBalanceAnnotations = parser.Annotation{
	Group: "backend",
	Annotations: parser.AnnotationFields{
		loadBalanceAlgorithmAnnotation: {
			Validator: parser.ValidateOptions(loadBalanceAlgorithms, true, true),
			Scope:     parser.AnnotationScopeLocation,
			Risk:      parser.AnnotationRiskLow,
			Documentation: `This annotation allows setting the load balancing algorithm that should be used. If none is specified, defaults to
			the default configured by Ingress admin, otherwise to round_robin`,
		},
	},
}
