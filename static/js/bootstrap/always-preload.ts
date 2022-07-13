import {OnLoad} from '../utils/_onload.js';

function addJSClass() {
  document.body.classList.add('u-js-loaded');
}

function asyncPreloadCSS() {
  const styles = document.querySelectorAll('link[rel="preload"][as="style"]') as NodeListOf<HTMLLinkElement>;
  for (const s of styles) {
    const l = document.createElement('link');
    l.href = s.href;
    l.rel = 'stylesheet';
    document.head.appendChild(l);
  }
}

function asyncDataSrc() {
  const elements = document.querySelectorAll("[data-src]") as NodeListOf<HTMLIFrameElement>;
  for (const e of elements) {
    if (e.dataset.src) {
      e.src = e.dataset.src;
    }
  }
}

function run() {
  addJSClass();
  asyncDataSrc();
  asyncPreloadCSS();
}

OnLoad(run);