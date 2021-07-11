package jamf_pro_go

import (
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

const (
	APIVersionScripts = "v1"
	APIPathScripts    = "scripts"
)

type Scripts struct {
	TotalCount  uint32   `yaml:"totalCount" json:"totalCount"`
	Results     []Script `yaml:"results" json:"results"`
}

type Script struct {
	ID              string `yaml:"id" json:"id"`
	Name            string `yaml:"name" json:"name"`
	Info            string `yaml:"info" json:"info"`
	Notes           string `yaml:"notes" json:"notes"`
	Priority        string `yaml:"priority" json:"priority"` // [ BEFORE, AFTER, AT_REBOOT ] (default: BEFORE)
	CategoryID      string `yaml:"categoryId" json:"categoryId"`
	CategoryName    string `yaml:"categoryName" json:"categoryName"`
	Parameter4      string `yaml:"parameter4" json:"parameter4"`
	Parameter5      string `yaml:"parameter5" json:"parameter5"`
	Parameter6      string `yaml:"parameter6" json:"parameter6"`
	Parameter7      string `yaml:"parameter7" json:"parameter7"`
	Parameter8      string `yaml:"parameter8" json:"parameter8"`
	Parameter9      string `yaml:"parameter9" json:"parameter9"`
	Parameter10     string `yaml:"parameter10" json:"parameter10"`
	Parameter11     string `yaml:"parameter11" json:"parameter11"`
	OsRequirements  string `yaml:"osRequirements" json:"osRequirements"`
	ScriptContents  string `yaml:"scriptContents" json:"scriptContents"`
}


type GetScriptsOpts struct {
	Page      uint32   `yaml:"page" url:"page,omitempty"`
	PageSize  uint32   `yaml:"page-size" url:"page-size,omitempty"`
	Sort      []string `yaml:"sort" url:"sort,omitempty"`
	Filter    string   `yaml:"filter" url:"filter,omitempty"`
}

func (c *Client) GetScripts(opts GetScriptsOpts) (*Scripts, error) {
	var result Scripts

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	err = c.call(path.Join(APIVersionScripts, APIPathScripts), http.MethodGet,
		APIVersionScripts, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetScript(scriptID uint32) (*Script, error) {
	var result Script

	err := c.call(path.Join(APIVersionScripts, APIPathScripts, fmt.Sprint(scriptID)), http.MethodGet,
		APIVersionScripts, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type CreateScriptResult struct {
	ID    string `yaml:"id" json:"id"`
	Href  string `yaml:"href" json:"href"`
}

func (c *Client) CreateScript (params Script) (*CreateScriptResult, error) {
	var result CreateScriptResult

	err := c.call(path.Join(APIVersionScripts, APIPathScripts), http.MethodPost,
		APIVersionScripts, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateScript (scriptID uint32, params Script) (*Script, error) {
	var result Script

	err := c.call(path.Join(APIVersionScripts, APIPathScripts, fmt.Sprint(scriptID)), http.MethodPut,
		APIVersionScripts, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteScript (scriptID uint32) error {
	err := c.call(path.Join(APIVersionScripts, APIPathScripts, fmt.Sprint(scriptID)), http.MethodDelete,
		APIVersionScripts, nil, nil, nil)
	if err != io.EOF {
		return err
	}
	fmt.Println("[jamf-pro-go] Script (ID: ", scriptID, ") is deleted")

	return nil
}