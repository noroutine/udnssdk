package udnssdk

import "net/http"

// RRSetsService provides access to RRSet resources
type ZonesService struct {
    client *Client
}

type ZoneProperties struct {
    Name                    string      `json:"name"`
    AccountName             string      `json:"accountName"`
    Type                    string      `json:"type"`
    DNSSecStatus            string      `json:"dnssecStatus"`
    Status                  string      `json:"status"`
    Owner                   string      `json:"owner"`
    ResourceRecordCount     int         `json:"resourceRecordCount"`
    LastModificationTime    string      `json:"lastModifiedDateTime"`
}

type ZoneRegistrarInfo struct {
    Registrar           string              `json:"registrar,omitempty"`
    WhoisExpiration     string              `json:"whoisExpiration,omitempty"`
    NameServers         ZoneNameServers     `json:"nameServers,omitempty"`
}

type ZoneNameServers struct {
    Ok          []string                `json:"ok,omitempty"`
    Unknown     []string                `json:"unknown,omitempty"`
}

type ZoneRestrictIP struct {
    SingleIP    string                  `json:"singleIP,omitempty"`
    Comment     string                  `json:"comment,omitempty"`
}

// Zone wraps an RRSet resource
type Zone struct {
    Properties          ZoneProperties          `json:"properties"`
    RegistrarInfo       ZoneRegistrarInfo       `json:"registrarInfo"`
    RestrictIPList      []ZoneRestrictIP        `json:"restrictIpList,omitempty"`
}

// RRSetListDTO wraps a list of RRSet resources
type ZoneListDTO struct {
    Zones      []Zone     `json:"zones"`
    Queryinfo  QueryInfo  `json:"queryInfo"`
    Resultinfo ResultInfo `json:"resultInfo"`
}

// EventsURI generates the URI for an RRSet
func ZonesURI() string {
    return "/zones"
}

func (s *ZonesService) Select() ([]Zone, error) {
    var zones []Zone
    var e error
    zones, _, _, e = s.SelectInternal()
    return zones, e
}

func (s *ZonesService) SelectInternal() ([]Zone, ResultInfo, *http.Response, error) {
    var zonesDTO ZoneListDTO

    uri := ZonesURI()
    res, err := s.client.get(uri, &zonesDTO)

    return zonesDTO.Zones, zonesDTO.Resultinfo, res, err
}