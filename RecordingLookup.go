package musicbrainz_api

import (
	"fmt"

	"gitlab.com/rye_tawny/api_commons"
	musicbrainztypes "gitlab.com/rye_tawny/musicbrainz/types"
)

func RecordingLookupByMbid(mbid string, includeAliases bool) (*musicbrainztypes.Recording, error) {
	if mbid == "" {
		return nil, fmt.Errorf("musicbrainz_api: malformed artist lookup")
	}
	apiUrl := fmt.Sprintf("https://musicbrainz.org/ws/2/recording/%s", mbid)
	if includeAliases {
		apiUrl += "?inc=artist-credits+isrcs+releases"
	}
	data, err := api_commons.FetchXML[musicbrainztypes.WrappedRecordingLookupResult](apiUrl)
	if err != nil {
		return nil, err
	}
	return &data.Recording, nil
}
