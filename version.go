package udnssdk

// Version Info response
type VersionInfo struct {
    Version string `json:"version"`
}

func VersionURI() string {
    return "version"
}

func (client *Client) GetVersion() (string, error) {
    var verInfo VersionInfo
    _, err := client.get(VersionURI(), &verInfo)
    return verInfo.Version, err
}