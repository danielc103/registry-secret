package secret

import (
	"encoding/base64"
	"encoding/json"
)

// DockerConfig stucture
type DockerConfig struct {
	Auths DockerInfo `json:"auth,omitempty"`
	// +optional
	HTTPHeaders map[string]string `json:"HttpHeaders,omitempty"`
}

// DockerInfo map for values
type DockerInfo map[string]DockerInfoValues

// DockerInfoValues structure containing username, password, email and auth
type DockerInfoValues struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Auth     string `json:"auth,omitempty"`
}

// CreateDockerJSON creates the JSON struct for registry secret
func CreateDockerJSON(username, password, email, server string, setAuth bool, headers map[string]string) ([]byte, error) {
	var auth string

	if setAuth {
		auth = encodeAuthValue(username, password)
	} else {
		auth = ""
	}

	dockInfoVals := DockerInfoValues{
		Username: username,
		Password: password,
		Email:    email,
		Auth:     auth,
	}

	dockerConfigJSON := DockerConfig{
		Auths:       map[string]DockerInfoValues{server: dockInfoVals},
		HTTPHeaders: headers,
	}

	return json.Marshal(dockerConfigJSON)
}

func encodeAuthValue(username, password string) string {
	value := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(value))
}
