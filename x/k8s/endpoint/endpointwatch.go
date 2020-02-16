/*

Copyright 2018 David Walter.

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

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/davidwalter0/go-cfg"
	"github.com/davidwalter0/loadbalancer/kubeconfig"
	"github.com/davidwalter0/loadbalancer/watch"
)

// ServerCfg runtime options config struct
type ServerCfg struct {
	Debug      bool   `json:"debug"       doc:"increase verbosity"                               default:"false"`
	Kubeconfig string `json:"kubeconfig"  doc:"kubernetes auth secrets / configuration file"     default:"cluster/auth/kubeconfig"`
	Kubernetes bool   `json:"kubernetes"  doc:"use kubernetes dynamic endpoints from endpoint/ns" default:"true"`
}

// Read from env variables or command line flags
func (envCfg *ServerCfg) Read() {
	var err error
	if err = cfg.AddStruct(envCfg); err != nil {
		log.Fatalf("Error: %v", err)
	}
	cfg.Finalize()
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
	var err error
	endpointWatcher := watch.NewQueueMgr(watch.EndpointAPIName, clientset)
	go endpointWatcher.Run(1, 1)

	time.Sleep(10 * time.Second)
	for i := 0; i < 100; i++ {
		select {
		case item := <-endpointWatcher.QueueItems:
			endpoint := item.Interface.(*v1.Endpoints)
			fmt.Println(endpoint.Name)
			for key, value := range endpoint.ObjectMeta.Labels {
				fmt.Printf("  %-32s %v\n", key, value)
			}
			for _, value := range endpoint.Subsets {
				var service *v1.Service
				namespace := endpoint.ObjectMeta.Namespace
				name := endpoint.ObjectMeta.Name
				service, err = clientset.CoreV1().Services(endpoint.ObjectMeta.Namespace).Get(name, metav1.GetOptions{})
				if errors.IsNotFound(err) {
					fmt.Printf("Service %s in namespace %s not found\n", name, namespace)
				} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
					fmt.Printf("Error getting service %s in namespace %s: %v\n", name, namespace, statusError.ErrStatus.Message)
				} else if err != nil {
					panic(err.Error())
				} else {
					if service != nil && service.Spec.Type == v1.ServiceTypeLoadBalancer {
						fmt.Printf("Found service %s in namespace %s\n", name, namespace)
						// if jsonText, err := json.MarshalIndent(service.Spec, "", "  "); err == nil {
						// 	fmt.Println(string(jsonText))
						// } else {
						// 	fmt.Println(err)
						// }
						if jsonText, err := json.MarshalIndent(value, "", "  "); err == nil {
							fmt.Println(string(jsonText))
						} else {
							fmt.Println(err)
						}
					}
				}
				fmt.Printf("  %-32s %v\n", name, value.Addresses)
			}
			// if jsonText, err := json.MarshalIndent(endpoint, "", "  "); err == nil {
			// 	fmt.Println(string(jsonText))
			// } else {
			// 	fmt.Println(err)
			// }

		default:
			// non blocking channel, choose default and break
			break
		}
	}
}
