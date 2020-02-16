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

// Note: the example only works with the code within the same release/branch.

package main

import (
	"fmt"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"github.com/davidwalter0/go-cfg"
	"github.com/davidwalter0/tools/x/k8s/kubeconfig"
)

// ServerCfg runtime options config struct
type ServerCfg struct {
	Debug      bool   `json:"debug"       doc:"increase verbosity"                               default:"false"`
	Kubeconfig string `json:"kubeconfig"  doc:"kubernetes auth secrets / configuration file"     default:"cluster/auth/kubeconfig"`
	Kubernetes bool   `json:"kubernetes"  doc:"use kubernetes dynamic endpoints from service/ns" default:"true"`
}

// Read from env variables or command line flags
func (envCfg *ServerCfg) Read() {
	var err error
	if err = cfg.Init(envCfg); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

var envCfg = &ServerCfg{}

func init() {
	// Load the configuration
	envCfg.Read()
}

// Cfg exposes common configuration item
func Cfg() *ServerCfg {
	return envCfg
}

func main() {
	clientset := kubeconfig.NewClientset(envCfg.Kubeconfig)
	if clientset == nil {
		log.Fatal("Kubernetes connection failed")
	}
	var selector = map[string]string{
		"LabelSelector": "node-role.kubernetes.io/worker",
	}

	for i := 0; i < 2; i++ {
		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{LabelSelector: selector["LabelSelector"]})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d %s nodes in the cluster\n", len(nodes.Items), selector["LabelSelector"])
		for i, item := range nodes.Items {
			fmt.Printf("node %d of %d nodes %s\n", i, len(nodes.Items), item.ObjectMeta.Name)
		}

		// all nodes
		nodes, err = clientset.CoreV1().Nodes().List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
		for i, item := range nodes.Items {
			fmt.Printf("node %d of %d nodes %s\n", i, len(nodes.Items), item.ObjectMeta.Name)
		}

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		time.Sleep(10 * time.Second)
	}

}
