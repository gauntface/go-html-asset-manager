import {OnLoad} from '../utils/_onload.js';

function addJSClass() {
  document.body.classList.add('u-js-loaded');
}

function loadStylesheet(s: HTMLLinkElement) {
  const l = document.createElement('link');
    l.href = s.href;
    l.rel = 'stylesheet';
    if (s.attributes['ham-media']) {
      l.media = s.attributes['ham-media'];
    }
    document.head.appendChild(l);
}

function loadStylesheets(selector) {
  const styles = document.querySelectorAll(selector) as NodeListOf<HTMLLinkElement>;
  for (const s of styles) {
    loadStylesheet(s);
  }
}

function loadPreloadCSS() {
  loadStylesheets('link[rel="preload"][as="style"]');
}

function loadAsyncCSS() {
  loadStylesheets('link[rel="stylesheet"][media="ham-async"]');
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
  asyncDataSrc();
  loadPreloadCSS();
  loadAsyncCSS();
  addJSClass();
}

OnLoad(run);