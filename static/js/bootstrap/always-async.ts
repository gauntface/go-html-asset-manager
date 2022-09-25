import {OnLoad} from '../utils/_onload.js';

function addJSLoadedSignal() {
  document.body.classList.add('u-js-loaded');
}

function loadPreloadCSS() {
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

function loadPreloadJS() {
  const scripts = document.querySelectorAll('link[rel="preload"][as="script"]') as NodeListOf<HTMLLinkElement>;
  for (const s of scripts) {
    const script = document.createElement('script');
    script.src = s.href;
    document.body.appendChild(script);
  }
}

function run() {
  asyncDataSrc();
  loadPreloadCSS();
  loadPreloadJS();
  addJSLoadedSignal();
}

OnLoad(run);
