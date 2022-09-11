package data

import "k8s.io/cli-runtime/pkg/genericclioptions"

type FinalizeConfig struct {
	NamespaceName string
	ConfigFlags   *genericclioptions.ConfigFlags
}
