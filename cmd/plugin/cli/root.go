package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/k8s/kubectl-test/pkg/logger"
	"github.com/k8s/kubectl-test/pkg/plugin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kubectl-test",
		Short:         "test",
		Long:          `test long`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: test,
	}

	cobra.OnInitialize(initConfig)

	KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	KubernetesConfigFlags.AddFlags(cmd.Flags())

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}

func test(cmd *cobra.Command, args []string) error {
	log := logger.NewLogger()
	if len(args) == 0 {
		_ = cmd.Help()
		return nil
	}
	var err error
	ns, err := cmd.Flags().GetString("namespace")
	if err != nil {
		fmt.Println("parse flag err:", err)
	}
	for _, arg := range args {
		if err != nil {
			return err
		}
		switch arg {
		case "ns":
			err = plugin.RunNs(KubernetesConfigFlags)
		case "pod":
			err = plugin.RunPod(KubernetesConfigFlags, ns)
		case "svc":
			err = plugin.RunSvc(KubernetesConfigFlags, ns)
		default:
			fmt.Println("input resource")
		}
	}
	log.Info("")

	return nil
}
