import {logger} from '@gauntface/logger';
logger.setPrefix('go-html-asset-manager/lite-vimeo');

class LiteVimeoEmbed {

  private element: HTMLElement;
  private anchor: HTMLElement;
  private videoID: string;
  private preconnected: boolean;

  constructor(e: HTMLElement) {
      this.element = e;
      this.anchor = e.querySelector(`${LiteVimeoEmbed.selector()}__link`) as HTMLElement;

      const vid = e.getAttribute('videoid');
      if (!vid) {
        logger.warn(`Failed to get the 'videoid' attribute from element: `, e);
        return;
      }
      this.videoID = encodeURIComponent(vid);
      this.preconnected = false;

      this.setup();
  }

  setup() {
      // On hover (or tap), warm up the TCP connections we're (likely) about to use.
      this.anchor.addEventListener('pointerover', () => this.warmConnections(), {once: true});

      // Once the user clicks, add the real iframe
      this.anchor.addEventListener('click', (e) => {
        e.preventDefault();
        this.addIframe();
      });
  }

  addIframe(){
    const iframeHTML = `<iframe
    allow="autoplay; picture-in-picture" allowfullscreen
    src="https://player.vimeo.com/video/${this.videoID}?color=ffffff&title=0&byline=0&portrait=0&autoplay=1"
    style="width:100%;height:100%;border:none;"
    ></iframe>`;
    this.element.removeChild(this.anchor);
    this.element.insertAdjacentHTML('beforeend', iframeHTML);
    this.element.classList.add(`${LiteVimeoEmbed.classname()}--activated`);
  }

  /**
   * Add a <link rel={preload | preconnect} ...> to the head
   */
  addPrefetch(kind: string, url: string, as?: string) {
      const linkElem = document.createElement('link');
      linkElem.rel = kind;
      linkElem.href = url;
      if (as) {
          linkElem.as = as;
      }
      linkElem.crossOrigin = "";
      document.head.append(linkElem);
  }

  /**
   * Begin pre-connecting to warm up the iframe load
   * Since the embed's network requests load within its iframe,
   *   preload/prefetch'ing them outside the iframe will only cause double-downloads.
   * So, the best we can do is warm up a few connections to origins that are in the critical path.
   *
   * Maybe `<link rel=preload as=document>` would work, but it's unsupported: http://crbug.com/593267
   * But TBH, I don't think it'll happen soon with Site Isolation and split caches adding serious complexity.
   */
  warmConnections() {
      if (this.preconnected) return;
      this.preconnected = true;

      // The iframe document and most of its subresources come right off player.vimeo.com
      this.addPrefetch('preconnect', 'https://player.vimeo.com');
      // Images
      this.addPrefetch('preconnect', 'https://i.vimeocdn.com');
      // Files .js, .css
      this.addPrefetch('preconnect', 'https://f.vimeocdn.com');
      // Metrics
      this.addPrefetch('preconnect', 'https://fresnel.vimeocdn.com');
  }

  static classname() {
    return 'n-ham-c-lite-vi'
  }
  static selector() {
    return `.${LiteVimeoEmbed.classname()}`
  }
}

(function () {
  function prepViLite() {
    const ytElements = document.querySelectorAll<HTMLElement>(LiteVimeoEmbed.selector());
    for (const e of ytElements) {
      new LiteVimeoEmbed(e);
    }
  }

  window.addEventListener('load', prepViLite)
  if (document.readyState == 'complete') {
    prepViLite();
  }
})()
