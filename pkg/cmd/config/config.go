package config

import (
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/olekukonko/tablewriter"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/utils"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/version"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var conf *Config

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
Please create "$HOME/.helm/config.yaml" file

## Example
helm:
  command: helm
  version: 2 # 2 or 3

deploy-list:
  - name: amf-longtest-task30
    values:
      - /root/helm/custom-yamls/5g-amf.yaml
    alias:
      - amf
    chart: /root/helm/5g-amf
    namespace: amf-longtest-task30
  - name: smf-longtest-task30
    values:
      - /root/helm/custom-yamls/5g-smf.yaml
    alias:
      - smf
    chart: /root/helm/5g-smf
    namespace: smf-longtest-task30`
		log.Fatal(errors.New(configError))
	}
	utils.CheckError(yaml.Unmarshal(yamlFile, config))
	conf = config

	err = conf.checkAliasValidation()
	if err != nil {
		log.Fatal(err)
	}

	return conf
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
	table.SetHeader([]string{"Index (Alias)", "Command", "Execution"})
	table.SetAutoMergeCells(true)
	table.SetAutoFormatHeaders(true)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoWrapText(true)
	table.SetColWidth(100)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_CENTER, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	table.SetCaption(true, fmt.Sprintf("HDM %s", version.Version))

	for i, di := range c.DeployInfo {
		index := i + 1
		table.Append([]string{buildIndexAndAlias(index, di.Alias), "apply", c.ApplyCommand(index)})
		table.Append([]string{buildIndexAndAlias(index, di.Alias), "delete", c.DeleteCommand(index)})
		table.Append([]string{buildIndexAndAlias(index, di.Alias), "status", c.StatusCommand(index)})
		table.Append([]string{buildIndexAndAlias(index, di.Alias), "ui", c.UICommand(index)})
	}

	table.Render()
}

func buildIndexAndAlias(index int, aliasList []string) string {
	indexStr := strconv.Itoa(index)
	aliases := strings.Join(aliasList, ",")
	var ret string
	if aliases != "" {
		ret = fmt.Sprintf("%s (%s)", indexStr, aliases)
	} else {
		ret = fmt.Sprintf("%s", indexStr)
	}

	return ret
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

func GetConfig() *Config {
	return conf
}

func GetTargetsToCompletion() []string {
	length := len(conf.DeployInfo)
	var ret []string
	for i := 0; i < length; i++ {
		ret = append(ret, strconv.Itoa(i+1))
	}

	for _, di := range conf.DeployInfo {
		for _, alias := range di.Alias {
			ret = append(ret, alias)
		}
	}

	return ret
}

func (c *Config) checkAliasValidation() error {
	aliasMap := make(map[string]int)
	for i, di := range c.DeployInfo {
		for _, alias := range di.Alias {
			if t, err := strconv.Atoi(alias); err == nil {
				return errors.New(fmt.Sprintf("alias cannot be integer: %d", t))
			}

			_, exists := aliasMap[alias]
			if exists == false {
				aliasMap[alias] = i
			} else {
				return errors.New("duplicate alias error")
			}
		}
	}
	return nil
}
