package jamf_pro_go

import (
    "encoding/xml"
    "fmt"
    "io"
    "net/http"
    "path"
)

const (
    APIVersionCategories = "classic"
    APIPathCategories    = "categories"
)

type Categoris struct {
    size     uint32     `yaml:"size,omitempty" xml:"size,omitempty"`
    Category []Category `yaml:"category,omitempty" xml:"category,omitempty"`
}

type Category struct {
    XMLName  *xml.Name `yaml:"-" xml:"category,omitempty"`
    ID       uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
    Name     string `yaml:"name,omitempty" xml:"name,omitempty"`
    Priority uint32 `yaml:"priority,omitempty" xml:"priority,omitempty"`
}

func (c *Client) GetCategories() (*Categoris, error) {
    var result Categoris

    err := c.call(APIPathCategories, http.MethodGet,
        APIVersionCategories, nil, nil, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

func (c *Client) GetCategory(categoryID uint32) (*Category, error) {
    var result Category

    err := c.call(path.Join(APIPathCategories, "id", fmt.Sprint(categoryID)), http.MethodGet,
        APIVersionCategories, nil, nil, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

type CategoryResult struct {
    XMLName    *xml.Name  `yaml:"-" xml:"category,omitempty"`
    ID uint32 `yaml:"id,omitempty" xml:"id,omitempty"`
}

func (c *Client) CreateCategory (params Category) (*CategoryResult, error) {
    var result CategoryResult

    err := c.call(path.Join(APIPathCategories, "id", "0"), http.MethodPost,
        APIVersionCategories, nil, params, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

func (c *Client) UpdateCategory (categoryID uint32, params Category) (*CategoryResult, error) {
    var result CategoryResult

    err := c.call(path.Join(APIPathCategories, "id", fmt.Sprint(categoryID)), http.MethodPut,
        APIVersionCategories, nil, params, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

func (c *Client) DeleteCategory (categoryID uint32) (*CategoryResult, error) {
    var result CategoryResult

    err := c.call(path.Join(APIPathCategories, "id", fmt.Sprint(categoryID)), http.MethodDelete,
        APIVersionCategories, nil, nil, nil)
    if err != io.EOF {
        return nil, err
    }

    return &result, nil
}
