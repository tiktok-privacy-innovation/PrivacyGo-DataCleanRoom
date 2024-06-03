// Copyright 2024 TikTok Pte. Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	K8sClientSet     *kubernetes.Clientset
	RunningNameSpace string
)

func InitK8sClient() {
	var err error
	clusterConfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	K8sClientSet, err = kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		panic(err)
	}

	RunningNameSpaceByte, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err)
	}
	RunningNameSpace = string(RunningNameSpaceByte)
}
