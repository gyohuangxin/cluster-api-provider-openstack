// +build tools

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

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/a8m/envsubst"                             //nolint
	_ "github.com/golang/mock/mockgen"                      //nolint
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint" //nolint
	_ "k8s.io/code-generator"                               //nolint
	_ "sigs.k8s.io/cluster-api/hack/tools/release"          //nolint
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"     //nolint
	_ "sigs.k8s.io/kustomize/kustomize/v3"                  //nolint
	_ "sigs.k8s.io/testing_frameworks/integration"          //nolint
	_ "github.com/onsi/ginkgo/ginkgo"                       //nolint
)
