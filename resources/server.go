package resources

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/littlegreensoftware/deployavel/api"
)

// ServerRequest models input from yml files
type ServerRequest struct {
	Server struct {
		ID           int    `yaml:"ID" json:"id"`
		CredentialID int    `yaml:"CredentialID" json:"credential_id"`
		Provider     string `yaml:"Provider" json:"provider"`
		Database     string `yaml:"Database" json:"database"`
		Name         string `yaml:"Name" json:"name"`
		Size         string `yaml:"Size" json:"size"`
		Region       string `yaml:"Region" json:"region"`
		PhpVersion   string `yaml:"PhpVersion" json:"php_version"`
		RecipeID     int    `yaml:"RecipeID" json:"recipe_id"`
	} `yaml:"server"`
}

// ServerListResponse models list of servers from Forge
type ServerListResponse struct {
	All []Server `json:"servers"`
}

// Marshal the struct to a slice of bytes
func (s ServerListResponse) Marshal() ([]byte, error) {
	data, err := json.Marshal(s.All)
	return data, err
}

// ServerResponse models single server response from Forge
type ServerResponse struct {
	Server Server `json:"server"`
}

// CreatedServer models successful server creation
type CreatedServer struct {
	BaseResource
	Server       Server `json:"server"`
	SudoPassword string `json:"sudo_password"`
	DbPassword   string `json:"database_password"`
}

// Server represents a single forge server
type Server struct {
	BaseResource
	ID               int           `json:"id"`
	CredentialID     int           `json:"credential_id"`
	Name             string        `json:"name"`
	Size             string        `json:"size"`
	Region           string        `json:"region"`
	PhpVersion       string        `json:"php_version"`
	IPAddress        string        `json:"ip_address"`
	SSHPort          int           `json:"ssh_port"`
	PrivateIPAddress string        `json:"private_ip_address"`
	BlackfireStatus  interface{}   `json:"blackfire_status"`
	PapertrailStatus string        `json:"papertrail_status"`
	Revoked          bool          `json:"revoked"`
	CreatedAt        string        `json:"created_at"`
	IsReady          bool          `json:"is_ready"`
	Network          []interface{} `json:"network"`
}

// ServerCreate a server on forge
func ServerCreate(r api.Request, data interface{}) (*CreatedServer, error) {
	var created CreatedServer

	resp, err := r.Post("servers", data)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, api.RequestError(resp.StatusCode)
	}

	err = ParseResource(resp, &created)

	return &created, err
}

// ServerRead a single server from forge
func ServerRead(r api.Request, id int) (*Server, error) {
	var serv ServerResponse

	resp, err := r.Get("servers/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, api.RequestError(resp.StatusCode)
	}

	err = ParseResource(resp, &serv)

	return &serv.Server, err
}

// ServerUpdate a server from forge
func ServerUpdate(r api.Request, id int) error {
	return nil
}

// ServerList all servers on forge
func ServerList(r api.Request) (*ServerListResponse, error) {
	var list ServerListResponse

	resp, err := r.Get("servers")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, api.RequestError(resp.StatusCode)
	}

	err = ParseResource(resp, &list)

	return &list, err
}

// EnableOpCache enables opcache for a specific server
func EnableOpCache(r api.Request, id int) error {
	resp, err := r.Post("servers/"+strconv.Itoa(id)+"/php/opcache", nil)
	if resp.StatusCode != http.StatusOK {
		return api.RequestError(resp.StatusCode)
	}

	return err
}

// DisableOpCache disables opcache for a specific server
func DisableOpCache(r api.Request, id int) error {
	resp, err := r.Delete("servers/" + strconv.Itoa(id) + "/php/opcache")
	if resp.StatusCode != http.StatusOK {
		return api.RequestError(resp.StatusCode)
	}

	return err
}
