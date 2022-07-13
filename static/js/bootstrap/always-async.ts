import {OnLoad} from '../utils/_onload.js';

function loadPreloadScripts() {
  const scripts = document.querySelectorAll('link[rel="preload"][as="script"]') as NodeListOf<HTMLLinkElement>;
  for (const s of scripts) {
    const script = document.createElement('script');
    script.src = s.href;
    document.body.appendChild(script);
  }
}

OnLoad(function() {
  loadPreloadScripts();
})