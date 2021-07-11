package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	yaml "github.com/goccy/go-yaml"
	jamf "jamf-pro-manager/jamf-pro-go"
)

type DeployConfig struct {
	Policy jamf.Policy   `yaml:"policy,omitempty"`
	Script []jamf.Script `yaml:"script,omitempty"`
	Category []jamf.Category `yaml:"category,omitempty"`
	ComputerGroup []jamf.ComputerGroup `yaml:"computer_group,omitempty"`
}

type WriteConfig struct {
	Policy      jamf.Policy   `yaml:"policy,omitempty"`
	Script      []WriteScript `yaml:"script,omitempty"`
	Category    []jamf.Category `yaml:"category,omitempty"`
	ComputerGroup []jamf.ComputerGroup `yaml:"computer_group,omitempty"`
}

type WriteScript struct {
	ID             uint32 `yaml:"id" json:"id"`
	Name           string `yaml:"name" json:"name"`
	Info           string `yaml:"info" json:"info"`
	Notes          string `yaml:"notes" json:"notes"`
	Priority       string `yaml:"priority" json:"priority"`
	CategoryID     uint32 `yaml:"categoryId" json:"categoryId"`
	CategoryName   string `yaml:"categoryName" json:"categoryName"`
	Parameter4     string `yaml:"parameter4" json:"parameter4"`
	Parameter5     string `yaml:"parameter5" json:"parameter5"`
	Parameter6     string `yaml:"parameter6" json:"parameter6"`
	Parameter7     string `yaml:"parameter7" json:"parameter7"`
	Parameter8     string `yaml:"parameter8" json:"parameter8"`
	Parameter9     string `yaml:"parameter9" json:"parameter9"`
	Parameter10    string `yaml:"parameter10" json:"parameter10"`
	Parameter11    string `yaml:"parameter11" json:"parameter11"`
	OsRequirements string `yaml:"osRequirements" json:"osRequirements"`
	ScriptContents string `yaml:"scriptContents" json:"scriptContents"`
}

func main() {
	// server connection information
	url := os.Getenv("JAMF_URL")
	userName := os.Getenv("JAMF_USER")
	password := os.Getenv("JAMF_USER_PASSWORD")

	// Specify the target directory to be deployed
	targetDir := os.Getenv("TARGET_DIR")
	fmt.Printf("[info] TARGET_DIR: %s\n\n", targetDir)

	// read deployConfig.yml
	ymlConfig, err := ioutil.ReadFile(targetDir + "/deployConfig.yml")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	var config DeployConfig
	err = yaml.Unmarshal([]byte(ymlConfig), &config)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Printf("deployConfig:\n%+v\n", config)

	// create http client
	conf, err := jamf.NewConfig(url, userName, password)
	if err != nil {
		fmt.Println(err.Error())
	}
	client := jamf.NewClient(conf)

	var w WriteConfig

	// deploy Categories
	if config.Category != nil{
		categories := deployCategory(*client, config.Category)
		w.Category = *categories
	} else {
		// deploy Scripts
		scripts := deployScript(*client, config.Script, targetDir)
		w.Script = *scripts

		// Scripts - Policy.PolicyScripts の紐付け (名寄せ)
		m := map[string]uint32{}
		if len(w.Script) > 0 {
			for i, _ := range w.Script {
				m[w.Script[i].Name] = w.Script[i].ID
			}
			if len(config.Policy.Scripts.PolicyScript) > 0 {
				for i, _ := range config.Policy.Scripts.PolicyScript {
					if config.Policy.Scripts.PolicyScript[i].ID == 0 {
						config.Policy.Scripts.PolicyScript[i].ID = m[config.Policy.Scripts.PolicyScript[i].Name]
					}
				}
			}
		}

		computerGroups := deployComputerGroup(*client, config.ComputerGroup)
		w.ComputerGroup = *computerGroups

		mComputerGroup := map[string]uint32{}
		if len(w.ComputerGroup) > 0 {
			for i, _ := range w.ComputerGroup {
				mComputerGroup[w.ComputerGroup[i].Name] = w.ComputerGroup[i].ID
			}
			if len(config.Policy.Scope.ComputerGroups.ComputerGroup) > 0 {
				for i, _ := range config.Policy.Scope.ComputerGroups.ComputerGroup {
					if config.Policy.Scope.ComputerGroups.ComputerGroup[i].ID == 0 {
						config.Policy.Scope.ComputerGroups.ComputerGroup[i].ID =
							mComputerGroup[config.Policy.Scope.ComputerGroups.ComputerGroup[i].Name]
						fmt.Printf("map id: %v\n", mComputerGroup[config.Policy.Scope.ComputerGroups.ComputerGroup[i].Name])
					}
				}
			}
		}

		fmt.Printf("map:\n%+v\n", mComputerGroup)

		// deploy Policy
		policy := deployPolicy(*client, config.Policy)
		w.Policy = *policy
	}
	writeConfig(w, targetDir)
}

func deployCategory(client jamf.Client, categories []jamf.Category) *[]jamf.Category {
	var categoryParams jamf.Category

	for i := 0; i < len(categories); i++ {
		categoryParams = categories[i]

		// create category
		if categoryParams.ID == 0 {
			createCategoryResult, err := client.CreateCategory(categoryParams)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("%+v\n", createCategoryResult)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			createdCategoryID := createCategoryResult.ID
			fmt.Printf("[info] Category ID is %v\n\n", createdCategoryID)

			categories[i].ID = createdCategoryID
		} else {
			// update category
			newCategory, err := client.UpdateCategory(categoryParams.ID, categoryParams)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			fmt.Printf("[info] Configure the updated Category =====================\n%+v\n\n", newCategory)
		}
	}

	return &categories
}

func deployScript(client jamf.Client, scripts []jamf.Script, targetDir string) *[]WriteScript {
	// get GitHub Actions Env
	gitHubUrl := os.Getenv("GITHUB_SERVER_URL")
	gitHubRepository := os.Getenv("GITHUB_REPOSITORY")
	gitHubSha := os.Getenv("GITHUB_SHA")
	gitHubActionsRunNumber := os.Getenv("GITHUB_RUN_NUMBER")

	var scriptParams jamf.Script
	var arrWS []WriteScript

	for i := 0; i < len(scripts); i++ {

		// read script file
		scriptFile := targetDir + "/" + scripts[i].Name
		scriptFileContent, err := ioutil.ReadFile(scriptFile)
		if err != nil {
			fmt.Println("err: ", err)
			os.Exit(1)
		}
		scripts[i].ScriptContents = string(scriptFileContent)

		// set deploy information
		scripts[i].Notes =
			gitHubUrl + "/" + gitHubRepository + "/commit/" + gitHubSha +
				"\n - GitHub Actions Run Number: #" + gitHubActionsRunNumber

		scriptParams = scripts[i]
		fmt.Printf("scriptParams: %+v\n\n", scriptParams)

		// create script
		if scriptParams.ID == "0" {
			createScriptResult, err := client.CreateScript(scriptParams)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("[info] Script ID is %s\n\n", createScriptResult.ID)
			fmt.Printf("[info] Create Script \n%+v\n", createScriptResult)
			scripts[i].ID = createScriptResult.ID
		} else {
			intScriptID, _ := strconv.Atoi(scriptParams.ID)
			/*
			// get exist script
			script, err := client.GetScript(uint32(intScriptID))
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			fmt.Printf("[info] Current Script Configuration =====================\n%+v\n\n", script)

			 */

			// update script
			newScript, err := client.UpdateScript(uint32(intScriptID), scriptParams)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			fmt.Printf("[info] Configure the updated Script =====================\n%+v\n\n", newScript)
		}
		wsParams := wsConv(scripts[i])
		arrWS = append(arrWS, wsParams)
	}

	return &arrWS
}

func wsConv(scriptParams jamf.Script) WriteScript{
	intScriptID, _ := strconv.Atoi(scriptParams.ID)
	uint32ScriptID := uint32(intScriptID)
	intCatID, _ := strconv.Atoi(scriptParams.CategoryID)
	uint32CatID := uint32(intCatID)

	var wsParams = WriteScript{
		ID:             uint32ScriptID,
		Name:           scriptParams.Name,
		Info:           scriptParams.Info,
		Notes:          scriptParams.Notes,
		Priority:       scriptParams.Priority,
		CategoryID:     uint32CatID,
		CategoryName:   scriptParams.CategoryName,
		Parameter4:     scriptParams.Parameter4,
		Parameter5:     scriptParams.Parameter5,
		Parameter6:     scriptParams.Parameter6,
		Parameter7:     scriptParams.Parameter7,
		Parameter8:     scriptParams.Parameter8,
		Parameter9:     scriptParams.Parameter9,
		Parameter10:    scriptParams.Parameter10,
		Parameter11:    scriptParams.Parameter11,
		OsRequirements: scriptParams.OsRequirements,
		ScriptContents: "(look at script file)",
	}
	return wsParams
}

func deployComputerGroup(client jamf.Client, computerGroups []jamf.ComputerGroup) *[]jamf.ComputerGroup {
	var computerGroupParams jamf.ComputerGroup

	for i := 0; i < len(computerGroups); i++ {
		computerGroupParams = computerGroups[i]

		// create ComputerGroup
		if computerGroupParams.ID == 0 {
			createComputerGroupResult, err := client.CreateComputerGroup(computerGroupParams)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("%+v\n", createComputerGroupResult)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			fmt.Printf("[info] Computer Group: ID is %v\n\n", createComputerGroupResult.ID)
			computerGroups[i].ID = createComputerGroupResult.ID
		} else {
			// update ComputerGroup
			newComputerGroup, err := client.UpdateComputerGroup(computerGroupParams.ID, computerGroupParams)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			fmt.Printf("[info] Configure the updated Computer Group =====================\n%+v\n\n", newComputerGroup)
		}
	}

	return &computerGroups
}


func deployPolicy(client jamf.Client, policyParams jamf.Policy) *jamf.Policy {

	fmt.Printf("%+v", policyParams)

	if policyParams.General != nil{
		if policyParams.General.ID == 0 {
			// create policy
			fmt.Printf("policyParams: %+v\n\n", policyParams.General)

			createPolicyResult, err := client.CreatePolicy(&policyParams)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			fmt.Printf("v: %+v\n", createPolicyResult)
			createdPolicyID := createPolicyResult.ID
			fmt.Printf("[info] Policy ID: %v\n\n", createdPolicyID)

			policyParams.General.ID = createdPolicyID

		} else {
			/*
				// get exist policy
				policyID := policyParams.General.ID
				policy, err := client.GetPolicy(policyID)
				if err != nil {
					fmt.Println("err: ", err.Error())
				}
				fmt.Printf("[info] Current Policy Configuration =====================\n%+v\n\n", policy)

			*/
			// update policy
			newPolicy, err := client.UpdatePolicy(policyParams.General.ID, &policyParams)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}
			fmt.Printf("[info] Configure the updated Policy =====================\n%+v\n\n", newPolicy)
		}
	}

	return &policyParams
}

func writeConfig(writeConfig interface{}, path string)  {
	fmt.Printf("wirteconfig:\n%+v\n", writeConfig)
	f2, err := os.Create(path + "/deployConfig.yml")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	defer f2.Close()

	d2 := yaml.NewEncoder(f2)
	if err := d2.Encode(writeConfig); err != nil {
		log.Fatal(err)
	}
	d2.Close()
}
