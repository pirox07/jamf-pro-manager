package jamf_pro_go

import (
    "encoding/xml"
    "fmt"
    "net/http"
    "path"
)

const (
    APIVersionComputerGroups = "classic"
    APIPathComputerGroups    = "computergroups"
)

type ComputerGroup struct {
    XMLName    *xml.Name  `yaml:"-" xml:"computer_group,omitempty"`
    ID         uint32     `yaml:"id,omitempty" xml:"id,omitempty"`
    Name       string     `yaml:"name,omitempty" xml:"name,omitempty"`
    IsSmart    bool       `yaml:"is_smart,omitempty" xml:"is_smart,omitempty"`
    Site       *Site      `yaml:"site,omitempty,omitempty" xml:"site,omitempty"`
    Criteria   *Criteria  `yaml:"criteria,omitempty" xml:"criteria,omitempty"`
    Computers  *Computers `yaml:"computers,omitempty" xml:"computers,omitempty"`
}

type Site struct {
    ID    int32  `yaml:"id,omitempty" xml:"id,omitempty"`
    Name  string `yaml:"name,omitempty" xml:"name,omitempty"`
}

type Criteria struct {
    Size       uint32       `yaml:"size,omitempty" xml:"size,omitempty"`
    Criterion  *[]Criterion `yaml:"criterion,omitempty" xml:"criterion,omitempty"`
}

type Criterion struct {
    Name          string `yaml:"name,omitempty" xml:"name,omitempty"`
    Priority      uint32 `yaml:"priority,omitempty" xml:"priority,omitempty"`
    AndOr         string `yaml:"and_or,omitempty" xml:"and_or,omitempty"`
    SearchType    string `yaml:"search_type,omitempty" xml:"search_type,omitempty"`
    Value         string `yaml:"value,omitempty" xml:"value,omitempty"`
    OpeningParen  bool `yaml:"opening_paren,omitempty" xml:"opening_paren,omitempty"`
    ClosingParen  bool `yaml:"closing_paren,omitempty" xml:"closing_paren,omitempty"`
}

type Computers struct {
    Size      uint32 `yaml:"size,omitempty" xml:"size,omitempty"`
    Computer  *[]Computer `yaml:"computer,omitempty" xml:"computer,omitempty"`
}

type Computer struct {
    ID             uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
    Name           string `yaml:"name,omitempty" xml:"name,omitempty"`
    MacAddress     string `yaml:"mac_address,omitempty" xml:"mac_address,omitempty"`
    AltMacAddress  string `yaml:"alt_mac_address,omitempty" xml:"alt_mac_address,omitempty"`
    SerialNumber   string `yaml:"serial_number,omitempty" xml:"serial_number,omitempty"`
}

type ComputerGroups struct {
    Size          uint32 `yaml:"size,omitempty" xml:"size,omitempty"`
    ComputerGroup []ComputerGroup `yaml:"computer_group,omitempty" xml:"computer_group,omitempty"`
}


func (c *Client) GetComputerGroups() (*ComputerGroups, error) {
    var result ComputerGroups

    err := c.call(APIPathComputerGroups, http.MethodGet,
        APIVersionComputerGroups, nil, nil, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

func (c *Client) GetComputerGroup(computerGroupID uint32) (*ComputerGroup, error) {
    var result ComputerGroup

    err := c.call(path.Join(APIPathComputerGroups, "id", fmt.Sprint(computerGroupID)), http.MethodGet,
        APIVersionComputerGroups, nil, nil, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

type ComputerGroupResult struct {
    XMLName    *xml.Name  `yaml:"-" xml:"computer_group,omitempty"`
    ID uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
}


func (c *Client) CreateComputerGroup (params ComputerGroup) (*ComputerGroupResult, error) {
    var result ComputerGroupResult
    fmt.Printf("%+v\n", params)

    err := c.call(path.Join(APIPathComputerGroups, "id", "0"), http.MethodPost,
        APIVersionComputerGroups, nil, params, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

func (c *Client) UpdateComputerGroup (computerGroupID uint32, params ComputerGroup) (*ComputerGroupResult, error) {
    var result ComputerGroupResult

    err := c.call(path.Join(APIPathComputerGroups, "id", fmt.Sprint(computerGroupID)), http.MethodPut,
        APIVersionComputerGroups, nil, params, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

func (c *Client) DeleteComputerGroup (computerGroupID uint32) (*ComputerGroupResult, error) {
    var result ComputerGroupResult

    err := c.call(path.Join(APIPathComputerGroups, "id", fmt.Sprint(computerGroupID)), http.MethodDelete,
        APIVersionComputerGroups, nil, nil, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}
