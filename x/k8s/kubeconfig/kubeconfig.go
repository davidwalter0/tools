/*

------------------------------------------------------------------------

Modified from the original example code to first connect with
incluster configuration then if unsuccessful external kubeconfig
format file

------------------------------------------------------------------------

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

// Note: the example only works with the code within the same
// release/branch.

package kubeconfig

import (
	"os"

	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"

	// Uncomment the following line to load the gcp plugin (only
	// required to authenticate against GKE clusters).

	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"k8s.io/client-go/tools/clientcmd"
)

// InCluster true when endpoints are accessible, when the service is
// running with the cluster's network namespaces
var InCluster bool

// CheckInCluster reports if the env variable is set for cluster
func CheckInCluster() bool {
	return len(os.Getenv("KUBERNETES_PORT")) > 0
}

func init() {
	InCluster = CheckInCluster()
}

// NewClientset returns a new handle to a kubernetes client takes
// kubeconfig path arg
func NewClientset(kubeconfig string) *kubernetes.Clientset {

	// kubeRestConfig kubernetes config object
	var kubeRestConfig *restclient.Config
	// clientset is a handle to execute kubernetes commands
	var clientset *kubernetes.Clientset
	var err error

	// creates the in-cluster configuration
	kubeRestConfig, err = restclient.InClusterConfig()
	if err != nil {
		// try with a kubeconfig file
		kubeRestConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		// InCluster = true
	}

	if err == nil {
		// creates the clientset
		clientset, err = kubernetes.NewForConfig(kubeRestConfig)
	}
	return clientset
}
