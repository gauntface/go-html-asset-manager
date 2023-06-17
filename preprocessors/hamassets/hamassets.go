package hamassets

import (
	"github.com/gauntface/go-html-asset-manager/v5/assets/assetmanager"
	"github.com/gauntface/go-html-asset-manager/v5/embedassets"
	"github.com/gauntface/go-html-asset-manager/v5/preprocessors"
)

func Preprocessor(runtime preprocessors.Runtime) error {
	relDir := runtime.Assets.StaticDir()
	if relDir == "" {
		return nil
	}

	files, err := embedassets.CopyAssets(relDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		l, err := assetmanager.NewLocalAsset(relDir, f)
		if err != nil {
			return err
		}
		runtime.Assets.AddLocal(l)
	}

	return nil
}
