package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path"
    "strconv"
    "strings"

    yaml "github.com/goccy/go-yaml"
    jamf "jamf-pro-manager/jamf-pro-go"
)

const (
    dirOut = "out-conf"
    dirScriptContents = "out-conf/script-contents"
)

type Config struct {
    Policy      []jamf.Policy   `yaml:"policy,omitempty"`
    Script      []ScriptCnvt `yaml:"script,omitempty"`
    Category    []jamf.Category `yaml:"category,omitempty"`
    ComputerGroup []jamf.ComputerGroup `yaml:"computer_group,omitempty"`
}

type ScriptCnvt struct {
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
    url := os.Getenv("JAMF_BASE_URL")
    userName := os.Getenv("JAMF_USER")
    password := os.Getenv("JAMF_USER_PASSWORD")

    // create http client
    conf, err := jamf.NewConfig(url, userName, password)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    client := jamf.NewClient(conf)

    // create directory
    if _, err := os.Stat(dirOut); !os.IsNotExist(err) {
        // directory is exist
        if err = os.RemoveAll(dirOut); err != nil {
            fmt.Println(err)
        }
    }
    if err := os.MkdirAll(dirScriptContents, 0775); err != nil {
        fmt.Println(err)
    }

    var config Config
    var fileName string

    // get policies
    fmt.Println("getting policies ...")
    policies, err := client.GetPolicies()
    if err != nil {
        fmt.Println("err: ", err.Error())
    }

    for i, _ := range policies.Policy {
        policy, err := client.GetPolicy(policies.Policy[i].ID)
        if err != nil {
            fmt.Println("err: ", err.Error())
        }
        config.Policy = append(config.Policy, *policy)
    }
    fileName = "conf-policies.yml"
    err = WriteConfig(dirOut, fileName, config)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    config.Policy = nil

    // get scripts
    fmt.Println("getting scripts ...")
    queryOps := jamf.GetScriptsOpts{
      PageSize: 999999,
    }
    scripts, err := client.GetScripts(queryOps)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }

    for i, _ := range scripts.Results {
        tmp := convSctipt(scripts.Results[i])
        config.Script = append(config.Script, tmp)
        err = WriteScriptContent(dirScriptContents, scripts.Results[i].Name, scripts.Results[i].ScriptContents)
        if err != nil {
            fmt.Println("err: ", err.Error())
        }
    }
    fileName = "conf-scripts.yml"
    err = WriteConfig(dirOut, fileName, config)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    config.Script = nil

    // get computer groups
    fmt.Println("getting computer groups ...")
    computerGroups,err := client.GetComputerGroups()
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    for i, _ := range computerGroups.ComputerGroup {
        computerGroup, err := client.GetComputerGroup(computerGroups.ComputerGroup[i].ID)
        if err != nil {
            fmt.Println("err: ", err.Error())
        }
        config.ComputerGroup = append(config.ComputerGroup, *computerGroup)
    }
    fileName = "conf-computer-groups.yml"
    err = WriteConfig(dirOut, fileName, config)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    config.ComputerGroup = nil

    // get categories
    fmt.Println("getting categories ...")
    categories, err := client.GetCategories()
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    for i, _ := range categories.Category{
        category, err := client.GetCategory(categories.Category[i].ID)
        if err != nil {
            fmt.Println("err: ", err.Error())
        }
        config.Category = append(config.Category, *category)
    }
    fileName = "conf-categories.yml"
    err = WriteConfig(dirOut, fileName, config)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
}


func WriteConfig (dirName string, fileName string, config Config) error {
    // output YAML file
    f, err := os.OpenFile(path.Join(dirName, fileName) , os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
    if err != nil {
        fmt.Println("err: ", err.Error())
    }
    defer f.Close()
    d := yaml.NewEncoder(f)
    if err := d.Encode(config); err != nil {
        log.Fatal(err)
    }
    d.Close()

    return nil
}

func WriteScriptContent (dirName, scriptName, scriptContent string) error {
    // output YAML file
    arr := []string{}
    arr = strings.Split(scriptContent,"")

    b := []byte{}
    for _, line := range arr {
        ll := []byte(line)
        for _, l := range ll {
            b = append(b, l)
        }
    }

    err := ioutil.WriteFile(dirName + "/" + scriptName, b , 0666)
    if err != nil {
        fmt.Println(os.Stderr, err)
        os.Exit(1)
    }

    return nil
}

func convSctipt(scriptParams jamf.Script) ScriptCnvt{
    intScriptID, _ := strconv.Atoi(scriptParams.ID)
    uint32ScriptID := uint32(intScriptID)
    intCatID, _ := strconv.Atoi(scriptParams.CategoryID)
    uint32CatID := uint32(intCatID)

    var wsParams = ScriptCnvt{
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
