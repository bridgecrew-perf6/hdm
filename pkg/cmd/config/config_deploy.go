package config

import "fmt"

type DeployConfig struct {
	ReleaseName    string   `yaml:"name"`
	Alias          []string `yaml:"alias"`
	ValueFiles     []string `yaml:"values"`
	ChartDirectory string   `yaml:"chart"`
	Namespace      string   `yaml:"namespace"`
}

func (dc *DeployConfig) applyCommand(hc HelmConfig) string {
	applyCmd := fmt.Sprintf("%s upgrade --install %s %s --namespace %s", hc.HelmCommand, dc.ReleaseName, dc.ChartDirectory, dc.Namespace)
	if len(dc.ValueFiles) > 0 {
		for _, valueFile := range dc.ValueFiles {
			applyCmd += fmt.Sprintf(" -f %s", valueFile)
		}
	}
	return applyCmd
}

func (dc *DeployConfig) deleteCommand(hc HelmConfig) string {
	deleteCmd := fmt.Sprintf("%s ", hc.HelmCommand)
	if hc.HelmVersion == 2 {
		deleteCmd += fmt.Sprintf("delete --purge ")
	} else {
		deleteCmd += fmt.Sprintf("uninstall --namespace %s ", dc.Namespace)
	}
	deleteCmd += dc.ReleaseName

	return deleteCmd
}

func (dc *DeployConfig) statusCommand() string {
	statusCmd := fmt.Sprintf("kubectl get pods -n %s -o wide", dc.Namespace)
	return statusCmd
}

func (dc *DeployConfig) uiCommand() string {
	uiCmd := fmt.Sprintf("kubectl exec -it -n %s svc/nf-userinterface -c ui -- /apps/pkg/sw/bin/CLI", dc.Namespace)
	return uiCmd
}
