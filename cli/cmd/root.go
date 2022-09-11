//: Copyright Herb Guo
//: Licensed under the terms of the Apache 2.0 License. See LICENSE file in the project root for terms.

package cmd

import (
	"fmt"
	"github.com/herbguo/kubectl-finalize/cli/cmd/data"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const (
	finalizeExamples = `
	# Finalize the test namespace in Terminating state
	%[1]s finalize test

`
)

type Options struct {
	configFlags *genericclioptions.ConfigFlags
	genericclioptions.IOStreams
}

func NewOptions(streams genericclioptions.IOStreams) *Options {
	return &Options{
		configFlags: genericclioptions.NewConfigFlags(false),
		IOStreams:   streams,
	}
}

func NewFinalizeCommand(streams genericclioptions.IOStreams) *cobra.Command {
	var (
		namespaceName string
	)

	options := NewOptions(streams)
	cmd := &cobra.Command{
		Use:                   "finalize [namespace-name]",
		DisableFlagsInUseLine: true,
		Short:                 "Finalize terminating ns.",
		Example:               fmt.Sprintf(finalizeExamples, "kubectl"),
		PersistentPreRun: func(c *cobra.Command, args []string) {
			c.SetOutput(streams.ErrOut)
		},
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				cmd.Help()
				return
			}
			namespaceName = args[0]
			if err := validateFlags(namespaceName); err != nil {
				fmt.Fprintln(streams.Out, err)
				os.Exit(1)
			}

			cfg := &data.FinalizeConfig{
				NamespaceName: namespaceName,
				ConfigFlags:   options.configFlags,
			}
			Finalize(cfg)
		},
	}

	//cmd.Flags().StringVar(&namespaceName, "target-namespace", "", "Finalize a namespace with the specified name")

	options.configFlags.AddFlags(cmd.Flags())

	return cmd
}

func validateFlags(namespaceName string) error {

	return nil
}
