(()=>{"use strict";var t={995:function(t,e){var n=this&&this.__awaiter||function(t,e,n,r){return new(n||(n=Promise))((function(o,a){function s(t){try{c(r.next(t))}catch(t){a(t)}}function u(t){try{c(r.throw(t))}catch(t){a(t)}}function c(t){var e;t.done?o(t.value):(e=t.value,e instanceof n?e:new n((function(t){t(e)}))).then(s,u)}c((r=r.apply(t,e||[])).next())}))},r=this&&this.__generator||function(t,e){var n,r,o,a,s={label:0,sent:function(){if(1&o[0])throw o[1];return o[1]},trys:[],ops:[]};return a={next:u(0),throw:u(1),return:u(2)},"function"==typeof Symbol&&(a[Symbol.iterator]=function(){return this}),a;function u(u){return function(c){return function(u){if(n)throw new TypeError("Generator is already executing.");for(;a&&(a=0,u[0]&&(s=0)),s;)try{if(n=1,r&&(o=2&u[0]?r.return:u[0]?r.throw||((o=r.return)&&o.call(r),0):r.next)&&!(o=o.call(r,u[1])).done)return o;switch(r=0,o&&(u=[2&u[0],o.value]),u[0]){case 0:case 1:o=u;break;case 4:return s.label++,{value:u[1],done:!1};case 5:s.label++,r=u[1],u=[0];continue;case 7:u=s.ops.pop(),s.trys.pop();continue;default:if(!((o=(o=s.trys).length>0&&o[o.length-1])||6!==u[0]&&2!==u[0])){s=0;continue}if(3===u[0]&&(!o||u[1]>o[0]&&u[1]<o[3])){s.label=u[1];break}if(6===u[0]&&s.label<o[1]){s.label=o[1],o=u;break}if(o&&s.label<o[2]){s.label=o[2],s.ops.push(u);break}o[2]&&s.ops.pop(),s.trys.pop();continue}u=e.call(t,s)}catch(t){u=[6,t],r=0}finally{n=o=0}if(5&u[0])throw u[1];return{value:u[0]?u[1]:void 0,done:!0}}([u,c])}}};Object.defineProperty(e,"__esModule",{value:!0}),e.default=function(){setInterval(s,a)};var o="https://https://wisdom.tkstorm.com/api/wisdom",a=1e4;function s(){console.log("Refreshing wisdom..."),function(){return n(this,void 0,void 0,(function(){var t,e,n;return r(this,(function(r){switch(r.label){case 0:return r.trys.push([0,3,,4]),[4,fetch(o)];case 1:if(!(t=r.sent()).ok)throw new Error("HTTP error! Status: ".concat(t.status));return[4,t.json()];case 2:return e=r.sent(),console.log("Fetched wisdom:",e),[2,e];case 3:return n=r.sent(),console.error("Failed to fetch wisdom:",n),[3,4];case 4:return[2]}}))}))}().then((function(t){var e,n;t&&"success"===t.status&&t.data&&t.data.PageData?(e=t.data.PageData.sentence,(n=document.querySelector(".content"))?n.textContent=e:console.error("Content div not found!")):console.error("Invalid data structure:",t)})).catch((function(t){console.error("Got error: "+t)}))}},926:function(t,e,n){var r=this&&this.__importDefault||function(t){return t&&t.__esModule?t:{default:t}};Object.defineProperty(e,"__esModule",{value:!0}),(0,r(n(995)).default)(),console.log("wisdom amazing!")}},e={};!function n(r){var o=e[r];if(void 0!==o)return o.exports;var a=e[r]={exports:{}};return t[r].call(a.exports,a,a.exports,n),a.exports}(926)})();