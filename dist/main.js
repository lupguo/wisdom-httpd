(()=>{"use strict";var n={711:(n,e,t)=>{t.d(e,{A:()=>c});var o=t(354),r=t.n(o),a=t(314),i=t.n(a)()(r());i.push([n.id,"/*首页*/\n.container {\n    display: flex;\n    flex-direction: column;\n    height: 100vh; /* 设置容器的高度为视口高度 */\n}\n\n.header {\n    background-color: #f0f0f0;\n    padding: 20px;\n    /* 导航栏样式 */\n}\n\n.content {\n    flex-grow: 1;\n    padding: 20px;\n    /* 主内容区域样式 */\n}\n\n.footer {\n    background-color: #e4cece;\n    padding: 20px;\n    /* 底部信息样式 */\n}\n\n/* append */\nbody {\n    display: flex;\n    flex-direction: column;\n    align-items: center;\n    justify-content: center;\n    height: 100vh;\n    font-family: sans-serif;\n    font-weight: 200;\n}\n\nh1 {\n    color: darkblue;\n}\n","",{version:3,sources:["webpack://./assets/src/css/index.css"],names:[],mappings:"AAAA,KAAK;AACL;IACI,aAAa;IACb,sBAAsB;IACtB,aAAa,EAAE,iBAAiB;AACpC;;AAEA;IACI,yBAAyB;IACzB,aAAa;IACb,UAAU;AACd;;AAEA;IACI,YAAY;IACZ,aAAa;IACb,YAAY;AAChB;;AAEA;IACI,yBAAyB;IACzB,aAAa;IACb,WAAW;AACf;;AAEA,WAAW;AACX;IACI,aAAa;IACb,sBAAsB;IACtB,mBAAmB;IACnB,uBAAuB;IACvB,aAAa;IACb,uBAAuB;IACvB,gBAAgB;AACpB;;AAEA;IACI,eAAe;AACnB",sourcesContent:["/*首页*/\n.container {\n    display: flex;\n    flex-direction: column;\n    height: 100vh; /* 设置容器的高度为视口高度 */\n}\n\n.header {\n    background-color: #f0f0f0;\n    padding: 20px;\n    /* 导航栏样式 */\n}\n\n.content {\n    flex-grow: 1;\n    padding: 20px;\n    /* 主内容区域样式 */\n}\n\n.footer {\n    background-color: #e4cece;\n    padding: 20px;\n    /* 底部信息样式 */\n}\n\n/* append */\nbody {\n    display: flex;\n    flex-direction: column;\n    align-items: center;\n    justify-content: center;\n    height: 100vh;\n    font-family: sans-serif;\n    font-weight: 200;\n}\n\nh1 {\n    color: darkblue;\n}\n"],sourceRoot:""}]);const c=i},314:n=>{n.exports=function(n){var e=[];return e.toString=function(){return this.map((function(e){var t="",o=void 0!==e[5];return e[4]&&(t+="@supports (".concat(e[4],") {")),e[2]&&(t+="@media ".concat(e[2]," {")),o&&(t+="@layer".concat(e[5].length>0?" ".concat(e[5]):""," {")),t+=n(e),o&&(t+="}"),e[2]&&(t+="}"),e[4]&&(t+="}"),t})).join("")},e.i=function(n,t,o,r,a){"string"==typeof n&&(n=[[null,n,void 0]]);var i={};if(o)for(var c=0;c<this.length;c++){var s=this[c][0];null!=s&&(i[s]=!0)}for(var u=0;u<n.length;u++){var l=[].concat(n[u]);o&&i[l[0]]||(void 0!==a&&(void 0===l[5]||(l[1]="@layer".concat(l[5].length>0?" ".concat(l[5]):""," {").concat(l[1],"}")),l[5]=a),t&&(l[2]?(l[1]="@media ".concat(l[2]," {").concat(l[1],"}"),l[2]=t):l[2]=t),r&&(l[4]?(l[1]="@supports (".concat(l[4],") {").concat(l[1],"}"),l[4]=r):l[4]="".concat(r)),e.push(l))}},e}},354:n=>{n.exports=function(n){var e=n[1],t=n[3];if(!t)return e;if("function"==typeof btoa){var o=btoa(unescape(encodeURIComponent(JSON.stringify(t)))),r="sourceMappingURL=data:application/json;charset=utf-8;base64,".concat(o),a="/*# ".concat(r," */");return[e].concat([a]).join("\n")}return[e].join("\n")}},742:(n,e,t)=>{t.r(e),t.d(e,{default:()=>m});var o=t(72),r=t.n(o),a=t(825),i=t.n(a),c=t(659),s=t.n(c),u=t(56),l=t.n(u),f=t(540),d=t.n(f),p=t(113),A=t.n(p),v=t(711),h={};h.styleTagTransform=A(),h.setAttributes=l(),h.insert=s().bind(null,"head"),h.domAPI=i(),h.insertStyleElement=d();r()(v.A,h);const m=v.A&&v.A.locals?v.A.locals:void 0},72:n=>{var e=[];function t(n){for(var t=-1,o=0;o<e.length;o++)if(e[o].identifier===n){t=o;break}return t}function o(n,o){for(var a={},i=[],c=0;c<n.length;c++){var s=n[c],u=o.base?s[0]+o.base:s[0],l=a[u]||0,f="".concat(u," ").concat(l);a[u]=l+1;var d=t(f),p={css:s[1],media:s[2],sourceMap:s[3],supports:s[4],layer:s[5]};if(-1!==d)e[d].references++,e[d].updater(p);else{var A=r(p,o);o.byIndex=c,e.splice(c,0,{identifier:f,updater:A,references:1})}i.push(f)}return i}function r(n,e){var t=e.domAPI(e);t.update(n);return function(e){if(e){if(e.css===n.css&&e.media===n.media&&e.sourceMap===n.sourceMap&&e.supports===n.supports&&e.layer===n.layer)return;t.update(n=e)}else t.remove()}}n.exports=function(n,r){var a=o(n=n||[],r=r||{});return function(n){n=n||[];for(var i=0;i<a.length;i++){var c=t(a[i]);e[c].references--}for(var s=o(n,r),u=0;u<a.length;u++){var l=t(a[u]);0===e[l].references&&(e[l].updater(),e.splice(l,1))}a=s}}},659:n=>{var e={};n.exports=function(n,t){var o=function(n){if(void 0===e[n]){var t=document.querySelector(n);if(window.HTMLIFrameElement&&t instanceof window.HTMLIFrameElement)try{t=t.contentDocument.head}catch(n){t=null}e[n]=t}return e[n]}(n);if(!o)throw new Error("Couldn't find a style target. This probably means that the value for the 'insert' parameter is invalid.");o.appendChild(t)}},540:n=>{n.exports=function(n){var e=document.createElement("style");return n.setAttributes(e,n.attributes),n.insert(e,n.options),e}},56:(n,e,t)=>{n.exports=function(n){var e=t.nc;e&&n.setAttribute("nonce",e)}},825:n=>{n.exports=function(n){if("undefined"==typeof document)return{update:function(){},remove:function(){}};var e=n.insertStyleElement(n);return{update:function(t){!function(n,e,t){var o="";t.supports&&(o+="@supports (".concat(t.supports,") {")),t.media&&(o+="@media ".concat(t.media," {"));var r=void 0!==t.layer;r&&(o+="@layer".concat(t.layer.length>0?" ".concat(t.layer):""," {")),o+=t.css,r&&(o+="}"),t.media&&(o+="}"),t.supports&&(o+="}");var a=t.sourceMap;a&&"undefined"!=typeof btoa&&(o+="\n/*# sourceMappingURL=data:application/json;base64,".concat(btoa(unescape(encodeURIComponent(JSON.stringify(a))))," */")),e.styleTagTransform(o,n,e.options)}(e,n,t)},remove:function(){!function(n){if(null===n.parentNode)return!1;n.parentNode.removeChild(n)}(e)}}}},113:n=>{n.exports=function(n,e){if(e.styleSheet)e.styleSheet.cssText=n;else{for(;e.firstChild;)e.removeChild(e.firstChild);e.appendChild(document.createTextNode(n))}}},42:function(n,e,t){var o=this&&this.__importDefault||function(n){return n&&n.__esModule?n:{default:n}};Object.defineProperty(e,"__esModule",{value:!0}),e.config=void 0,e.wisdom_url=a;var r=o(t(773));function a(n){var t=e.config.HOST.Wisdom,o=e.config.URI[n];if(!o)throw new Error("URI ".concat(n," not found in config"));return"".concat(t).concat(o)}e.config=r.default,console.log("global config: "+e.config);var i=a("GetRandomWisdom");console.log("Random Wisdom URL:",i);var c=a("SaveWisdom");console.log("Save Wisdom URL:",c)},507:function(n,e,t){var o=this&&this.__awaiter||function(n,e,t,o){return new(t||(t=Promise))((function(r,a){function i(n){try{s(o.next(n))}catch(n){a(n)}}function c(n){try{s(o.throw(n))}catch(n){a(n)}}function s(n){var e;n.done?r(n.value):(e=n.value,e instanceof t?e:new t((function(n){n(e)}))).then(i,c)}s((o=o.apply(n,e||[])).next())}))},r=this&&this.__generator||function(n,e){var t,o,r,a={label:0,sent:function(){if(1&r[0])throw r[1];return r[1]},trys:[],ops:[]},i=Object.create(("function"==typeof Iterator?Iterator:Object).prototype);return i.next=c(0),i.throw=c(1),i.return=c(2),"function"==typeof Symbol&&(i[Symbol.iterator]=function(){return this}),i;function c(c){return function(s){return function(c){if(t)throw new TypeError("Generator is already executing.");for(;i&&(i=0,c[0]&&(a=0)),a;)try{if(t=1,o&&(r=2&c[0]?o.return:c[0]?o.throw||((r=o.return)&&r.call(o),0):o.next)&&!(r=r.call(o,c[1])).done)return r;switch(o=0,r&&(c=[2&c[0],r.value]),c[0]){case 0:case 1:r=c;break;case 4:return a.label++,{value:c[1],done:!1};case 5:a.label++,o=c[1],c=[0];continue;case 7:c=a.ops.pop(),a.trys.pop();continue;default:if(!(r=a.trys,(r=r.length>0&&r[r.length-1])||6!==c[0]&&2!==c[0])){a=0;continue}if(3===c[0]&&(!r||c[1]>r[0]&&c[1]<r[3])){a.label=c[1];break}if(6===c[0]&&a.label<r[1]){a.label=r[1],r=c;break}if(r&&a.label<r[2]){a.label=r[2],a.ops.push(c);break}r[2]&&a.ops.pop(),a.trys.pop();continue}c=e.call(n,a)}catch(n){c=[6,n],o=0}finally{t=r=0}if(5&c[0])throw c[1];return{value:c[0]?c[1]:void 0,done:!0}}([c,s])}}};Object.defineProperty(e,"__esModule",{value:!0});var a=t(42);function i(){console.log("refreshWisdom..."),function(){return o(this,void 0,void 0,(function(){var n;return r(this,(function(e){switch(e.label){case 0:return[4,fetch((0,a.wisdom_url)("GetRandomWisdom"))];case 1:return(n=e.sent()).ok?[4,n.json()]:(console.error("HTTP error! Status: ".concat(n.status)),[2,null]);case 2:return[2,e.sent()]}}))}))}().then((function(n){var e=document.querySelector(".content");null!=e&&(e.textContent=n.sentence)}))}i(),setInterval(i,a.config.REFRESH_INTERVAL)},773:n=>{n.exports=JSON.parse('{"HOST":{"Wisdom":"http://127.0.0.1:1666"},"URI":{"GetRandomWisdom":"/api/wisdom?random=1","SaveWisdom":"/api/wisdom"},"REFRESH_INTERVAL":30000}')}},e={};function t(o){var r=e[o];if(void 0!==r)return r.exports;var a=e[o]={id:o,exports:{}};return n[o].call(a.exports,a,a.exports,t),a.exports}t.n=n=>{var e=n&&n.__esModule?()=>n.default:()=>n;return t.d(e,{a:e}),e},t.d=(n,e)=>{for(var o in e)t.o(e,o)&&!t.o(n,o)&&Object.defineProperty(n,o,{enumerable:!0,get:e[o]})},t.o=(n,e)=>Object.prototype.hasOwnProperty.call(n,e),t.r=n=>{"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(n,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(n,"__esModule",{value:!0})},t.nc=void 0;t(742),t(507)})();
//# sourceMappingURL=main.js.map