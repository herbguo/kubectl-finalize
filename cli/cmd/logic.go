//: Copyright Herb Guo
//: Licensed under the terms of the Apache 2.0 License. See LICENSE file in the project root for terms.

package cmd

import (
	"github.com/herbguo/kubectl-finalize/cli/cmd/data"
	"github.com/herbguo/kubectl-finalize/cli/cmd/kubernetes"
	"log"
)

func Finalize(cfg *data.FinalizeConfig) {
	err := kubernetes.Connect(cfg.ConfigFlags)
	if err != nil {
		log.Fatalf("Failed connecting to kubernetes cluster: %v\n", err)
	}
	log.Println("Finalize the Namespace: ", cfg.NamespaceName)
	err = kubernetes.Finalize(cfg.NamespaceName)
	if err != nil {
		log.Fatalf("Failed to execute Finalize: %v\n", err)
	}
}
