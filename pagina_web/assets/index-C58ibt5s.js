var te=Object.defineProperty;var ne=(e,t,n)=>t in e?te(e,t,{enumerable:!0,configurable:!0,writable:!0,value:n}):e[t]=n;var T=(e,t,n)=>(ne(e,typeof t!="symbol"?t+"":t,n),n);(function(){const t=document.createElement("link").relList;if(t&&t.supports&&t.supports("modulepreload"))return;for(const r of document.querySelectorAll('link[rel="modulepreload"]'))s(r);new MutationObserver(r=>{for(const l of r)if(l.type==="childList")for(const i of l.addedNodes)i.tagName==="LINK"&&i.rel==="modulepreload"&&s(i)}).observe(document,{childList:!0,subtree:!0});function n(r){const l={};return r.integrity&&(l.integrity=r.integrity),r.referrerPolicy&&(l.referrerPolicy=r.referrerPolicy),r.crossOrigin==="use-credentials"?l.credentials="include":r.crossOrigin==="anonymous"?l.credentials="omit":l.credentials="same-origin",l}function s(r){if(r.ep)return;r.ep=!0;const l=n(r);fetch(r.href,l)}})();function p(){}function Y(e){return e()}function K(){return Object.create(null)}function A(e){e.forEach(Y)}function Z(e){return typeof e=="function"}function P(e,t){return e!=e?t==t:e!==t||e&&typeof e=="object"||typeof e=="function"}let O;function U(e,t){return e===t?!0:(O||(O=document.createElement("a")),O.href=t,e===O.href)}function re(e){return Object.keys(e).length===0}function c(e,t){e.appendChild(t)}function L(e,t,n){e.insertBefore(t,n||null)}function x(e){e.parentNode&&e.parentNode.removeChild(e)}function se(e,t){for(let n=0;n<e.length;n+=1)e[n]&&e[n].d(t)}function d(e){return document.createElement(e)}function F(e){return document.createTextNode(e)}function v(){return F(" ")}function ie(e,t,n,s){return e.addEventListener(t,n,s),()=>e.removeEventListener(t,n,s)}function le(e){return function(t){return t.preventDefault(),e.call(this,t)}}function h(e,t,n){n==null?e.removeAttribute(t):e.getAttribute(t)!==n&&e.setAttribute(t,n)}function oe(e){return Array.from(e.childNodes)}function V(e,t){t=""+t,e.data!==t&&(e.data=t)}let H;function j(e){H=e}const y=[],G=[];let w=[];const J=[],ue=Promise.resolve();let z=!1;function ae(){z||(z=!0,ue.then(ee))}function D(e){w.push(e)}const B=new Set;let b=0;function ee(){if(b!==0)return;const e=H;do{try{for(;b<y.length;){const t=y[b];b++,j(t),ce(t.$$)}}catch(t){throw y.length=0,b=0,t}for(j(null),y.length=0,b=0;G.length;)G.pop()();for(let t=0;t<w.length;t+=1){const n=w[t];B.has(n)||(B.add(n),n())}w.length=0}while(y.length);for(;J.length;)J.pop()();z=!1,B.clear(),j(e)}function ce(e){if(e.fragment!==null){e.update(),A(e.before_update);const t=e.dirty;e.dirty=[-1],e.fragment&&e.fragment.p(e.ctx,t),e.after_update.forEach(D)}}function fe(e){const t=[],n=[];w.forEach(s=>e.indexOf(s)===-1?t.push(s):n.push(s)),n.forEach(s=>s()),w=t}const q=new Set;let de;function S(e,t){e&&e.i&&(q.delete(e),e.i(t))}function M(e,t,n,s){if(e&&e.o){if(q.has(e))return;q.add(e),de.c.push(()=>{q.delete(e)}),e.o(t)}}function Q(e){return(e==null?void 0:e.length)!==void 0?e:Array.from(e)}function R(e){e&&e.c()}function N(e,t,n){const{fragment:s,after_update:r}=e.$$;s&&s.m(t,n),D(()=>{const l=e.$$.on_mount.map(Y).filter(Z);e.$$.on_destroy?e.$$.on_destroy.push(...l):A(l),e.$$.on_mount=[]}),r.forEach(D)}function k(e,t){const n=e.$$;n.fragment!==null&&(fe(n.after_update),A(n.on_destroy),n.fragment&&n.fragment.d(t),n.on_destroy=n.fragment=null,n.ctx=[])}function me(e,t){e.$$.dirty[0]===-1&&(y.push(e),ae(),e.$$.dirty.fill(0)),e.$$.dirty[t/31|0]|=1<<t%31}function C(e,t,n,s,r,l,i=null,f=[-1]){const o=H;j(e);const u=e.$$={fragment:null,ctx:[],props:l,update:p,not_equal:r,bound:K(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(t.context||(o?o.$$.context:[])),callbacks:K(),dirty:f,skip_bound:!1,root:t.target||o.$$.root};i&&i(u.root);let $=!1;if(u.ctx=n?n(e,t.props||{},(a,g,..._)=>{const m=_.length?_[0]:g;return u.ctx&&r(u.ctx[a],u.ctx[a]=m)&&(!u.skip_bound&&u.bound[a]&&u.bound[a](m),$&&me(e,a)),g}):[],u.update(),$=!0,A(u.before_update),u.fragment=s?s(u.ctx):!1,t.target){if(t.hydrate){const a=oe(t.target);u.fragment&&u.fragment.l(a),a.forEach(x)}else u.fragment&&u.fragment.c();t.intro&&S(e.$$.fragment),N(e,t.target,t.anchor),ee()}j(o)}class I{constructor(){T(this,"$$");T(this,"$$set")}$destroy(){k(this,1),this.$destroy=p}$on(t,n){if(!Z(n))return p;const s=this.$$.callbacks[t]||(this.$$.callbacks[t]=[]);return s.push(n),()=>{const r=s.indexOf(n);r!==-1&&s.splice(r,1)}}$set(t){this.$$set&&!re(t)&&(this.$$.skip_bound=!0,this.$$set(t),this.$$.skip_bound=!1)}}const he="4";typeof window<"u"&&(window.__svelte||(window.__svelte={v:new Set})).v.add(he);function pe(e){let t;return{c(){t=d("nav"),t.innerHTML='<h1 class="svelte-1nm439z">Reseñas</h1>',h(t,"class","svelte-1nm439z")},m(n,s){L(n,t,s)},p,i:p,o:p,d(n){n&&x(t)}}}class ge extends I{constructor(t){super(),C(this,t,null,pe,P,{})}}function _e(e){let t,n,s,r,l,i,f,o,u,$,a,g;return{c(){t=d("div"),n=d("div"),s=v(),r=d("div"),l=d("details"),i=d("summary"),i.textContent="Subi tu reseña",f=v(),o=d("form"),o.innerHTML='<label for="titulo">Titulo:</label> <input type="text" name="titulo" id="titulo"/> <label for="link_imagen">Link a una imagen de la portada del libro:</label> <input type="text" name="link_img" id="link_imagen"/> <label for="parrafo">Reseña:</label> <input type="text" name="parrafo" id="parrafo"/> <button type="submit">Subir Reseña</button>',u=v(),$=d("div"),h(o,"method","post"),h(o,"class","svelte-n5lv8j"),h(r,"class","subida svelte-n5lv8j"),h(t,"class","container svelte-n5lv8j")},m(_,m){L(_,t,m),c(t,n),c(t,s),c(t,r),c(r,l),c(l,i),c(l,f),c(l,o),c(t,u),c(t,$),a||(g=ie(o,"submit",le(e[0])),a=!0)},p,i:p,o:p,d(_){_&&x(t),a=!1,g()}}}function $e(e){return[async n=>{const s=n.target,r=new FormData(s);fetch("http://192.168.1.12:8080/agregar_resenias",{method:"POST",body:r})}]}class ve extends I{constructor(t){super(),C(this,t,$e,_e,P,{})}}function W(e,t,n){const s=e.slice();return s[3]=t[n],s}function X(e){let t,n,s,r=e[3].titulo+"",l,i,f,o=e[3].parrafo+"",u,$,a,g,_;return{c(){t=d("article"),n=d("div"),s=d("h2"),l=F(r),i=v(),f=d("p"),u=F(o),$=v(),a=d("img"),_=v(),h(n,"class","svelte-90tqyq"),U(a.src,g=e[3].link_imagen)||h(a,"src",g),h(a,"alt","Imagen no disponible"),h(a,"class","svelte-90tqyq"),h(t,"class","svelte-90tqyq")},m(m,E){L(m,t,E),c(t,n),c(n,s),c(s,l),c(n,i),c(n,f),c(f,u),c(t,$),c(t,a),c(t,_)},p(m,E){E&1&&r!==(r=m[3].titulo+"")&&V(l,r),E&1&&o!==(o=m[3].parrafo+"")&&V(u,o),E&1&&!U(a.src,g=m[3].link_imagen)&&h(a,"src",g)},d(m){m&&x(t)}}}function be(e){let t,n=Q(e[0]),s=[];for(let r=0;r<n.length;r+=1)s[r]=X(W(e,n,r));return{c(){t=d("main");for(let r=0;r<s.length;r+=1)s[r].c();h(t,"class","svelte-90tqyq")},m(r,l){L(r,t,l);for(let i=0;i<s.length;i+=1)s[i]&&s[i].m(t,null)},p(r,[l]){if(l&1){n=Q(r[0]);let i;for(i=0;i<n.length;i+=1){const f=W(r,n,i);s[i]?s[i].p(f,l):(s[i]=X(f),s[i].c(),s[i].m(t,null))}for(;i<s.length;i+=1)s[i].d(1);s.length=n.length}},i:p,o:p,d(r){r&&x(t),se(s,r)}}}function ye(e,t,n){let s=[];return(async()=>{const i=await fetch("http://192.168.1.12:8080/leer_resenias");if(i.ok)return i.json()})().then(i=>{n(0,s=i)}),[s]}class we extends I{constructor(t){super(),C(this,t,ye,be,P,{})}}function xe(e){let t,n,s,r,l,i,f;return n=new ge({}),r=new ve({}),i=new we({}),{c(){t=d("div"),R(n.$$.fragment),s=v(),R(r.$$.fragment),l=v(),R(i.$$.fragment)},m(o,u){L(o,t,u),N(n,t,null),c(t,s),N(r,t,null),c(t,l),N(i,t,null),f=!0},p,i(o){f||(S(n.$$.fragment,o),S(r.$$.fragment,o),S(i.$$.fragment,o),f=!0)},o(o){M(n.$$.fragment,o),M(r.$$.fragment,o),M(i.$$.fragment,o),f=!1},d(o){o&&x(t),k(n),k(r),k(i)}}}class Ee extends I{constructor(t){super(),C(this,t,null,xe,P,{})}}new Ee({target:document.getElementById("app")});
