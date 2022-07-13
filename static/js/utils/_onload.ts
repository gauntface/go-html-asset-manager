function OnLoad(f: () => void) {
  window.addEventListener('load', f)
  if (document.readyState == 'complete') {
    f();
  }
}

export {OnLoad};