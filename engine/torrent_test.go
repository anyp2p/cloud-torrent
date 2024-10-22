package engine

import (
	"testing"

	"github.com/anacrolix/torrent/metainfo"
)

func TestMagnetURI(t *testing.T) {
	m, err := metainfo.ParseMagnetV2Uri(`magnet:?xt=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q`)
	if err != nil {
		panic(err.Error())
	}
	t.Logf(`%#v`, m)
}
