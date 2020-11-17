package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Config struct {
	HelmInfo   HelmConfig     `yaml:"helm"`
	DeployInfo []DeployConfig `yaml:"deploy-list"`
}

func NewConfig(configFilePath string) *Config {
	config := &Config{}
	var filePath string
	if configFilePath != "" {
		filePath = configFilePath
	} else {
		homeDir, err := homedir.Dir()
		utils.CheckError(err)
		filePath = fmt.Sprintf("%s/.hdm/config.yaml", homeDir)
	}

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		configError := `No config file detected.
Please create "/var/lib/helm-deploy-manager/config.yaml" file

## Example
helm:
  command: helm
  version: 2 # 2 or 3

deploy-list:
  - name: amf-longtest-task30
    values:
      - /root/helm/custom-yamls/5g-amf.yaml
    chart: /root/helm/5g-amf
    namespace: amf-longtest-task30
  - name: smf-longtest-task30
    values:
      - /root/helm/custom-yamls/5g-smf.yaml
    chart: /root/helm/5g-smf
    namespace: smf-longtest-task30`
		fmt.Println(configError)
		log.Fatal(err.Error())
	}
	utils.CheckError(yaml.Unmarshal(yamlFile, config))
	return config
}

func (c *Config) ApplyCommand(index int) string {
	if c.isOutOfRangeInDeployInfo(index) {
		return ""
	}
	applyCmd := c.DeployInfo[index-1].applyCommand(c.HelmInfo)
	return applyCmd
}

func (c *Config) DeleteCommand(index int) string {
	if c.isOutOfRangeInDeployInfo(index) {
		return ""
	}
	deleteCmd := c.DeployInfo[index-1].deleteCommand(c.HelmInfo)
	return deleteCmd
}

func (c *Config) StatusCommand(index int) string {
	if c.isOutOfRangeInDeployInfo(index) {
		return ""
	}
	statusCmd := c.DeployInfo[index-1].statusCommand()
	return statusCmd
}

func (c *Config) UICommand(index int) string {
	if c.isOutOfRangeInDeployInfo(index) {
		return ""
	}
	uiCmd := c.DeployInfo[index-1].uiCommand()
	return uiCmd
}

func (c *Config) ListAllCommand() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Command", "Execution"})
	table.SetAutoMergeCells(true)
	table.SetAutoFormatHeaders(true)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoWrapText(true)
	table.SetColWidth(100)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	table.SetCaption(true, "HDM in golang v1.0 by m5.kim")
	for i := range c.DeployInfo {
		index := i + 1
		table.Append([]string{strconv.Itoa(index), "apply", c.ApplyCommand(index)})
		table.Append([]string{strconv.Itoa(index), "delete", c.DeleteCommand(index)})
		table.Append([]string{strconv.Itoa(index), "status", c.StatusCommand(index)})
		table.Append([]string{strconv.Itoa(index), "ui", c.UICommand(index)})
	}

	table.Render()
}

//func customWrap(raw string) {
//	newRaw := make([]string, 0, len(raw))
//
//}

func (c *Config) isOutOfRangeInDeployInfo(index int) bool {
	if index > len(c.DeployInfo) || index < 1 {
		log.Fatalf("index is out of range in deploy info: index=%d, length=%d\n", index, len(c.DeployInfo))
		return true
	}
	return false
}


