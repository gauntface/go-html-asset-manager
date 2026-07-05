package jsonassets

import (
	"errors"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/assets"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetstubs"
	"github.com/gauntface/go-html-asset-manager/v5/preprocessors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

func Test_Preprocessor(t *testing.T) {
	tests := []struct {
		description string
		jsonAssets  []assetmanager.Asset
		wantError   error
		want        []*assetmanager.RemoteAsset
	}{
		{
			description: "does nothing if there are no JSON assets",
		},
		{
			description: "returns an error if reading contents fails",
			jsonAssets: []assetmanager.Asset{
				&assetstubs.Asset{ContentsError: errInjected},
			},
			wantError: errInjected,
		},
		{
			description: "returns an error if the JSON cannot be parsed",
			jsonAssets: []assetmanager.Asset{
				&assetstubs.Asset{ContentsReturn: `not json`},
			},
			wantError: errJSONParseFailed,
		},
		{
			description: "registers a remote asset per entry, tagged with the right asset type",
			jsonAssets: []assetmanager.Asset{
				&assetstubs.Asset{
					IDReturn: "example",
					ContentsReturn: `{
						"css": {
							"preload": [{"src": "https://example.com/preload.css"}]
						},
						"js": {
							"sync": [{"src": "https://example.com/sync.js", "attributes": [{"Key": "defer", "Value": ""}]}],
							"async": [{"src": "https://example.com/async.js"}]
						}
					}`,
				},
			},
			want: []*assetmanager.RemoteAsset{
				assetmanager.NewRemoteAsset("example", "https://example.com/preload.css", []html.Attribute{}, assets.PreloadCSS),
				assetmanager.NewRemoteAsset("example", "https://example.com/sync.js", []html.Attribute{{Key: "defer", Val: ""}}, assets.SyncJS),
				assetmanager.NewRemoteAsset("example", "https://example.com/async.js", []html.Attribute{}, assets.AsyncJS),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			manager := &assetstubs.Manager{
				WithTypeReturn: map[assets.Type][]assetmanager.Asset{
					assets.JSON: tt.jsonAssets,
				},
			}

			err := Preprocessor(preprocessors.Runtime{Assets: manager})
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			// Preprocessor iterates a map of asset types, so remote assets
			// aren't registered in a deterministic order; sort by URL before
			// comparing.
			sortByURL := cmpopts.SortSlices(func(a, b *assetmanager.RemoteAsset) bool {
				aURL, _ := a.URL()
				bURL, _ := b.URL()
				return aURL < bURL
			})
			if diff := cmp.Diff(manager.AddRemoteCalls, tt.want, cmp.AllowUnexported(assetmanager.RemoteAsset{}), sortByURL); diff != "" {
				t.Fatalf("Unexpected remote assets registered; diff %v", diff)
			}
		})
	}
}
