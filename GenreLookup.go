package musicbrainz_api

import (
	"fmt"

	"gitlab.com/rye_tawny/api_commons"
	musicbrainztypes "gitlab.com/rye_tawny/musicbrainz/types"
)

func GenreLookupByMbid(mbid string, includeAliases bool) (*musicbrainztypes.GenreLookupResult, error) {
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/genre/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainztypes.GenreLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
