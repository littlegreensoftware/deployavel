package resources

import (
	"net/http"
	"strconv"

	"github.com/littlegreensoftware/deployavel/api"
)

// SiteInput models a payload to create a site
type SiteInput struct {
	Domain      string `yaml:"Domain" json:"domain"`
	ProjectType string `yaml:"ProjectType" json:"project_type"`
	Directory   string `yaml:"Directory" json:"directory"`
}

// SiteRequest models input from yml files
type SiteRequest struct {
	Site SiteInput `yaml:"site"`
}

// CreatedSite models a site created response
type CreatedSite struct {
	Site Site `json:"site"`
}

// Site models a site on a server
type Site struct {
	BaseResource
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	Directory          string      `json:"directory"`
	Wildcards          bool        `json:"wildcards"`
	Status             string      `json:"status"`
	Repository         interface{} `json:"repository"`
	RepositoryProvider interface{} `json:"repository_provider"`
	RepositoryBranch   interface{} `json:"repository_branch"`
	RepositoryStatus   interface{} `json:"repository_status"`
	QuickDeploy        bool        `json:"quick_deploy"`
	ProjectType        string      `json:"project_type"`
	App                interface{} `json:"app"`
	AppStatus          interface{} `json:"app_status"`
	HipchatRoom        interface{} `json:"hipchat_room"`
	SlackChannel       interface{} `json:"slack_channel"`
	CreatedAt          string      `json:"created_at"`
}

// SiteCreate creates a site on forge for a particular server
func SiteCreate(r api.Request, id int, data SiteRequest) (*Site, error) {
	var created CreatedSite

	resp, err := r.Post("servers/"+strconv.Itoa(id)+"/sites", data)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, api.RequestError(resp.StatusCode)
	}

	err = ParseResource(resp, &created)

	return &created.Site, err
}
