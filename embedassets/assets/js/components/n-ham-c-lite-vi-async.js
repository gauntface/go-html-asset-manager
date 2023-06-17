var l={[0]:"\u{1F41B}",[1]:"\u2139\uFE0F",[2]:"\u{1F4AC}",[3]:"\u26A0\uFE0F",[4]:"\u2620\uFE0F",[5]:"\u{1F9F5}"},c=class{constructor(e){this.opts=e||{},this.setPrefix(this.opts.prefix),this.currentLogLevel=this.getDefaultLogLevel()}setPrefix(e){if(!e)this.opts.prefix=l;else if(typeof e=="object"){this.opts.prefix={};for(let t of Object.keys(l))this.opts.prefix[t]=e[t]||l[t]}else this.opts.prefix=e}setLogLevel(e){this.currentLogLevel=e}debug(...e){this.print(console.debug,0,e)}info(...e){this.print(console.info,1,e)}log(...e){this.print(console.log,2,e)}warn(...e){this.print(console.warn,3,e)}error(...e){this.print(console.error,4,e)}group(...e){this.print(console.group,5,e)}groupCollapsed(...e){this.print(console.groupCollapsed,5,e)}groupEnd(){console.groupEnd()}print(e,t,o){this.currentLogLevel>t||e(...this.getArgs(t,o))}getArgs(e,t){let o=this.getPrefix(e);return o?[...o,...t]:t}getPrefix(e){let t=this.opts.prefix;t||(t=l);let o=t;return typeof o=="object"&&(o=t[e]),this.colorPrefix(e,o)}getDefaultLogLevel(){return 0}};var m={red:99,green:110,blue:114},v={red:72,green:126,blue:176},R={red:76,green:209,blue:55},x={red:225,green:177,blue:44},E={red:231,green:76,blue:60},O={red:0,green:168,blue:255};var d=class extends c{colorPrefix(e,t){return[`%c${t}`,this.getLevelCSS(e)]}getLevelCSS(e){function t(o){return`background: rgb(${o.red}, ${o.green}, ${o.blue}); color: white; padding: 2px 0.5em; border-radius: 0.5em`}switch(e){case 0:return t(m);case 1:return t(v);case 3:return t(x);case 4:return t(E);case 5:return t(O);case 2:default:return t(R)}}};var u=new d;function b(r){window.addEventListener("load",r),document.readyState=="complete"&&r()}u.setPrefix("ham/lite-vimeo");var h=class r{element;anchor;videoID;preconnected;constructor(e){this.element=e,this.anchor=e.querySelector(`${r.selector()}__link`);let t=e.getAttribute("videoid");if(!t){u.warn("Failed to get the 'videoid' attribute from element: ",e);return}this.videoID=encodeURIComponent(t),this.preconnected=!1,this.setup()}setup(){this.anchor.addEventListener("pointerover",()=>this.warmConnections(),{once:!0}),this.anchor.addEventListener("click",e=>{e.preventDefault(),this.addIframe()})}addIframe(){let e=`<iframe
    allow="autoplay; picture-in-picture" allowfullscreen
    src="https://player.vimeo.com/video/${this.videoID}?color=ffffff&title=0&byline=0&portrait=0&autoplay=1"
    style="width:100%;height:100%;border:none;"
    ></iframe>`;this.element.removeChild(this.anchor),this.element.insertAdjacentHTML("beforeend",e),this.element.classList.add(`${r.classname()}--activated`)}addPrefetch(e,t,o){let s=document.createElement("link");s.rel=e,s.href=t,o&&(s.as=o),s.crossOrigin="",document.head.append(s)}warmConnections(){this.preconnected||(this.preconnected=!0,this.addPrefetch("preconnect","https://player.vimeo.com"),this.addPrefetch("preconnect","https://i.vimeocdn.com"),this.addPrefetch("preconnect","https://f.vimeocdn.com"),this.addPrefetch("preconnect","https://fresnel.vimeocdn.com"))}static classname(){return"n-ham-c-lite-vi"}static selector(){return`.${r.classname()}`}};b(function(){let r=document.querySelectorAll(h.selector());for(let e of r)new h(e)});
