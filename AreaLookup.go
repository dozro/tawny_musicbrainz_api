package musicbrainz_api

import (
	"fmt"

	"gitlab.com/rye_tawny/api_commons"
	musicbrainztypes "gitlab.com/rye_tawny/musicbrainz/types"
)

func AreaLookupByMbid(mbid string, includeAliases bool) (*musicbrainztypes.AreaLookupResult, error) {
	if mbid == "" {
		return nil, fmt.Errorf("musicbrainz_api: malformed area lookup")
	}
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/area/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainztypes.AreaLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
