# go-html-asset-manager

This is quite possibly the most terrible idea ever.

`go-html-asset-manager` is  library / set of CLI tools that look over a web project and perform
a set of operations on the source to optimize for web performance, this includes:

- Inlining critical CSS
- Loading CSS asynchronously
- Generating muliple sizes of images
- Switching out `<img>` tags for `<picture>`
- Replaces Vimeo and YouTube videos with still images and JS to load videos.

## Installation

```
go get -u github.com/gauntface/go-html-asset-manager/cmds/htmlassets/
go get -u github.com/gauntface/go-html-asset-manager/cmds/genimgs/
```

## Usage

I typically use it via [this node wrapper](https://github.com/gauntface/html-asset-manager#html-asset-manager) but you can run it like so:

```shell
htmlassets
```

Then add an assetmanager.json file.

```json
{
  "html-dir": "public/",
  "base-url": "https://www.gauntface.com",
  "assets": {
    "static-dir": "public/",
    "generated-dir": "public/generated/",
    "json-dir": "themes/gauntface/data/hopin/"
  },
  "gen-assets": {
    "static-dir": "static/",
    "output-dir": "static/generated/",
    "max-width": 800,
    "max-density": 3
  },
  "img-to-picture": [
    {
      "id": "l-blog",
      "max-width": 800,
      "source-sizes": [
        "(min-width: 800px) 800px",
        "100vw"
      ]
    },
    {
      "id": "c-blog-item__img",
      "max-width": 200,
      "source-sizes": [
        "(min-width: 800px) 200px",
        "20vw"
      ]
    }
  ],
  "ratio-wrapper": ["l-blog"]
}
```
