package opengraphimg

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/gauntface/go-html-asset-manager/v5/assets/genimgs"
	"github.com/gauntface/go-html-asset-manager/v5/manipulations"
	"github.com/gauntface/go-html-asset-manager/v5/utils/config"
	"github.com/gauntface/go-html-asset-manager/v5/utils/html/htmlparsing"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

var errInjected = errors.New("injected error")

var reset func()

func TestMain(m *testing.M) {
	origGenimgsLookupSizes := genimgsLookupSizes

	reset = func() {
		genimgsLookupSizes = origGenimgsLookupSizes
	}

	os.Exit(m.Run())
}

func TestManipulator(t *testing.T) {
	tests := []struct {
		description        string
		doc                *html.Node
		findNodes          func(tag string, node *html.Node) []*html.Node
		genimgsLookupSizes func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error)
		wantError          error
		wantHTML           string
	}{
		{
			description: "do nothing when no property",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta diff-attr="og:image" />`),
			wantHTML:    `<html><head><meta diff-attr="og:image"/></head><body></body></html>`,
		},
		{
			description: "do nothing when not og:image property",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:other" />`),
			wantHTML:    `<html><head><meta property="og:other"/></head><body></body></html>`,
		},
		{
			description: "do nothing when no content",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:image" />`),
			wantHTML:    `<html><head><meta property="og:image"/></head><body></body></html>`,
		},
		{
			description: "do nothing if getting images fails",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:image" content="/images/default-social.png" />`),
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return nil, errInjected
			},
			wantHTML: `<html><head><meta property="og:image" content="/images/default-social.png"/></head><body></body></html>`,
		},
		{
			description: "do nothing when no images",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:image" content="/images/default-social.png" />`),
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{}, nil
			},
			wantHTML: `<html><head><meta property="og:image" content="/images/default-social.png"/></head><body></body></html>`,
		},
		{
			description: "do nothing when no basic images",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:image" content="/images/default-social.png" />`),
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Type: "image/webp",
					},
				}, nil
			},
			wantHTML: `<html><head><meta property="og:image" content="/images/default-social.png"/></head><body></body></html>`,
		},
		{
			description: "do nothing when no images on the right size",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:image" content="/images/default-social.png" />`),
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Size: RECOMMENDED_OG_IMG_WIDTH + 1,
					},
				}, nil
			},
			wantHTML: `<html><head><meta property="og:image" content="/images/default-social.png"/></head><body></body></html>`,
		},
		{
			description: "update the image with the correct size",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:image" content="/images/default-social.png" />`),
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Size: RECOMMENDED_OG_IMG_WIDTH,
						URL:  "/images/default-social.1200xabc.png",
					},
				}, nil
			},
			wantHTML: `<html><head><meta content="http://base-url.com/images/default-social.1200xabc.png" property="og:image"/></head><body></body></html>`,
		},

		{
			description: "use basic image and not other image",
			findNodes:   htmlparsing.FindNodesByTag,
			doc:         MustGetNode(t, `<meta property="og:image" content="/images/default-social.png" />`),
			genimgsLookupSizes: func(s3 genimgs.ListObjectsV2APIClient, conf *config.Config, imgPath string) ([]genimgs.GenImg, error) {
				return []genimgs.GenImg{
					{
						Size: RECOMMENDED_OG_IMG_WIDTH,
						URL:  "/images/default-social.1200xabc.png",
					},
					{
						Size: RECOMMENDED_OG_IMG_WIDTH,
						Type: "image/webp",
						URL:  "/images/default-social.1200xabc.webp",
					},
					{
						Size: RECOMMENDED_OG_IMG_WIDTH,
						Type: "image/avif",
						URL:  "/images/default-social.1200xabc.avif",
					},
				}, nil
			},
			wantHTML: `<html><head><meta content="http://base-url.com/images/default-social.1200xabc.png" property="og:image"/></head><body></body></html>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer reset()

			genimgsLookupSizes = tt.genimgsLookupSizes

			r := manipulations.Runtime{
				Config: &config.Config{
					BaseURL: "http://base-url.com",
				},
			}

			err := Manipulator(r, tt.doc)
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("Unexpected error; got %v, want %v", err, tt.wantError)
			}

			if err != nil {
				return
			}

			if diff := cmp.Diff(MustRenderNode(t, tt.doc), tt.wantHTML); diff != "" {
				t.Fatalf("Unexpected HTML files; diff %v", diff)
			}
		})
	}
}

func MustGetNode(t *testing.T, input string) *html.Node {
	t.Helper()

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}
	return doc
}

func MustRenderNode(t *testing.T, n *html.Node) string {
	t.Helper()

	if n == nil {
		return ""
	}

	var buf bytes.Buffer
	err := html.Render(&buf, n)
	if err != nil {
		t.Fatalf("failed to render html node to string: %v", err)
	}

	return buf.String()
}
