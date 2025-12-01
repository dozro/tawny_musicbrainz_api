package musicbrainz_api

import (
	"fmt"

	"gitlab.com/rye_tawny/api_commons"
	musicbrainztypes "gitlab.com/rye_tawny/musicbrainz/types"
)

func ArtistLookupByMbid(mbid string, includeAliases bool) (*musicbrainztypes.Artist, error) {
	if mbid == "" {
		return nil, fmt.Errorf("musicbrainz_api: malformed artist lookup")
	}
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/artist/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=aliases"
	}
	data, err := api_commons.FetchXML[musicbrainztypes.WrappedArtistLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data.Artist, nil
}
