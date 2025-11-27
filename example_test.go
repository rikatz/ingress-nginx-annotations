/*
Copyright 2025 The Kubernetes Authors.

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

package annotations_test

import (
	"testing"

	annotations "github.com/rikatz/ingress-nginx-annotations"
	networking "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// How to use this library
func TestExample(t *testing.T) {
	factory := annotations.NewAnnotationFactory()

	ing := &networking.Ingress{
		ObjectMeta: v1.ObjectMeta{
			Name:      "something",
			Namespace: "default",
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/disable-proxy-intercept-errors": "not-a-bool",
				"nginx.ingress.kubernetes.io/default-backend":                "-invalid-name",
			},
		},
	}

	err := factory.Validate(ing)
	t.Errorf("an error happened: %s", err)

}
