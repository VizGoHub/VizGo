(self.webpackChunk=self.webpackChunk||[]).push([[997],{5100:function(__unused_webpack_module,__webpack_exports__,__webpack_require__){"use strict";var react__WEBPACK_IMPORTED_MODULE_0__=__webpack_require__(67294),echarts__WEBPACK_IMPORTED_MODULE_1__=__webpack_require__(64389),react_jsx_runtime__WEBPACK_IMPORTED_MODULE_2__=__webpack_require__(85893),ChartView=function ChartView(_ref){var chartID=_ref.chartID,scripts=_ref.scripts,data=_ref.data,chartRef=(0,react__WEBPACK_IMPORTED_MODULE_0__.useRef)(null);return(0,react__WEBPACK_IMPORTED_MODULE_0__.useEffect)(function(){var chartDom=chartRef.current,myChart=echarts__WEBPACK_IMPORTED_MODULE_1__.S1(chartDom),option={};try{try{eval(scripts),myChart.setOption(option)}catch(j){console.log(j)}}catch(j){}var resizeObserver=new ResizeObserver(function(){myChart.resize()});return resizeObserver.observe(chartDom),function(){myChart.dispose(),resizeObserver.disconnect()}},[chartID,scripts,data]),(0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_2__.jsx)("div",{ref:chartRef,style:{width:"100%",height:"100%"}})};__webpack_exports__.Z=ChartView},37939:function(j,g,a){"use strict";a.r(g),a.d(g,{default:function(){return te}});var y=a(9783),m=a.n(y),T=a(97857),u=a.n(T),I=a(15009),p=a.n(I),b=a(99289),M=a.n(b),P=a(5574),f=a.n(P),c=a(67294),C=a(69229),v=a(92317),d=a(18370),E=a(87973),s=a(82601),r=a(71230),l=a(15746),K=a(24435),X=a(72269),J=a(2244),k=a(96074),q=a(5100),ee=a(63764),W=a(67066),B=a(45803),ae=a(50942),e=a(85893),z=C.Z.Content,_e=C.Z.Sider,Z=v.Z.Title,V=v.Z.Text,re=d.Z.TextArea,x=E.Z.Option;function te(){var ne=(0,c.useState)({}),S=f()(ne,2),D=S[0],O=S[1],se=(0,c.useState)({}),G=f()(se,2),le=G[0],ie=G[1],oe=(0,c.useState)(`// @ts-nocheck 
`),H=f()(oe,2),A=H[0],w=H[1],ue=(0,c.useState)(!1),$=f()(ue,2),de=$[0],ce=$[1],he=(0,ae.UO)(),N=he.id,R=0;if(N!==void 0){var Q=parseInt(N,10);isNaN(Q)||(R=Q)}var F=(0,c.useCallback)(function(){var i=M()(p()().mark(function _(n){var o,t;return p()().wrap(function(h){for(;;)switch(h.prev=h.next){case 0:return h.prev=0,h.next=3,W.fy(n);case 3:o=h.sent,console.log(o),o&&o.code===0&&(ie(o.data),t=o.data.chart,O(t),w(t.chart_code||`// @ts-nocheck 
`),ce(!0)),h.next=11;break;case 8:h.prev=8,h.t0=h.catch(0),console.error("Error loading chart data:",h.t0);case 11:case"end":return h.stop()}},_,null,[[0,8]])}));return function(_){return i.apply(this,arguments)}}(),[]);(0,c.useEffect)(function(){R!==0&&F(R)},[F]);var L=function(_){var n=_.target,o=n.name,t=n.value;O(function(Y){return u()(u()({},Y),{},m()({},o,t))})},me=function(_,n){O(function(o){return u()(u()({},o),{},m()({},n,_))})},U=function(_,n){O(function(o){return u()(u()({},o),{},m()({},n,_!=null?_:0))})},Ee=(0,c.useCallback)(function(i){w(i||""),O(function(_){return u()(u()({},_),{},{chart_code:i||""})})},[]),pe=function(){var i=M()(p()().mark(function _(){var n;return p()().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.prev=0,t.next=3,W.MP({data:D});case 3:n=t.sent,n&&n.code===0?(B.ZP.success("Chart updated successfully"),console.log("Chart updated successfully")):(B.ZP.error("Failed to update chart: "+n.message),console.error("Failed to update chart:",n.message)),t.next=11;break;case 7:t.prev=7,t.t0=t.catch(0),B.ZP.error("Error updating chart"),console.error("Error updating chart:",t.t0);case 11:case"end":return t.stop()}},_,null,[[0,7]])}));return function(){return i.apply(this,arguments)}}();return(0,e.jsxs)(C.Z,{style:{height:"100vh"},children:[(0,e.jsxs)(_e,{width:"15%",style:{background:"#fff",padding:"20px",overflowY:"auto"},children:[(0,e.jsx)(Z,{level:3,children:"Chart Editor"}),(0,e.jsxs)(s.Z,{layout:"vertical",children:[(0,e.jsx)(r.Z,{children:(0,e.jsx)(l.Z,{children:(0,e.jsx)(s.Z.Item,{label:(0,e.jsx)(V,{copyable:!0,children:"data.chart.chart_name"}),style:{marginBottom:"8px"},children:(0,e.jsx)(d.Z,{name:"chart_name",value:D.chart_name||"",onChange:L})})})}),(0,e.jsx)(r.Z,{children:(0,e.jsx)(l.Z,{children:(0,e.jsx)(s.Z.Item,{label:(0,e.jsx)(V,{copyable:!0,children:"data.chart.labels"}),style:{marginBottom:"8px"},children:(0,e.jsx)(d.Z,{name:"labels",value:D.labels||"",onChange:L,style:{width:"100%"}})})})}),(0,e.jsx)(r.Z,{children:(0,e.jsx)(l.Z,{children:(0,e.jsx)(s.Z.Item,{label:"[data.chart.chart_type]",style:{marginBottom:"8px"},children:(0,e.jsxs)(E.Z,{name:"chart_type",value:D.chart_type||"",onChange:function(_){return me(_,"chart_type")},style:{width:"100%"},children:[(0,e.jsx)(x,{value:"bar",children:"bar"}),(0,e.jsx)(x,{value:"line",children:"line"}),(0,e.jsx)(x,{value:"pie",children:"pie"}),(0,e.jsx)(x,{value:"treemap",children:"treemap"})]})})})}),(0,e.jsx)(r.Z,{children:(0,e.jsx)(l.Z,{children:(0,e.jsx)(s.Z.Item,{label:"[DataSource]",style:{marginBottom:"8px"},children:(0,e.jsx)(K.Z,{name:"data_source_id",value:D.data_source_id||0,onChange:function(_){return U(_,"data_source_id")},style:{width:"100%"}})})})}),(0,e.jsx)(r.Z,{children:(0,e.jsx)(l.Z,{children:(0,e.jsx)(s.Z.Item,{label:"[Query SQL]",style:{marginBottom:"8px"},children:(0,e.jsx)(re,{name:"query_sql",value:D.query_sql||"",onChange:L,rows:2})})})}),(0,e.jsx)(r.Z,{children:(0,e.jsx)(l.Z,{children:(0,e.jsx)(s.Z.Item,{label:"[Sort]",style:{marginBottom:"8px"},children:(0,e.jsx)(K.Z,{name:"sort",value:D.sort||0,onChange:function(_){return U(_,"sort")},style:{width:"100%"}})})})}),(0,e.jsx)(r.Z,{gutter:16,children:(0,e.jsx)(l.Z,{span:12,children:(0,e.jsx)(s.Z.Item,{style:{marginBottom:"8px"},children:(0,e.jsx)(X.Z,{checked:D.status===1,onChange:function(_){return U(_?1:0,"status")},checkedChildren:"Enable",unCheckedChildren:"Disable"})})})}),(0,e.jsx)(r.Z,{gutter:16,children:(0,e.jsx)(l.Z,{span:12,children:(0,e.jsxs)(s.Z.Item,{style:{marginBottom:"8px"},children:[(0,e.jsx)(J.ZP,{type:"primary",onClick:pe,style:{width:"100%",marginTop:"20px"},children:"Apply"}),(0,e.jsx)(k.Z,{})]})})})]})]}),(0,e.jsxs)(z,{style:{padding:"20px",background:"#fff",flex:1,display:"flex",flexDirection:"column"},children:[(0,e.jsx)(Z,{level:4,children:"TSX Code Editor"}),(0,e.jsx)("div",{style:{width:"100%",height:"800px",padding:"20px",border:"1px solid #d9d9d9",borderRadius:"8px",background:"#fff"},children:(0,e.jsx)(ee.ZP,{height:"calc(100% - 60px)",language:"typescript",value:A,theme:"vs-dark",onChange:Ee,options:{selectOnLineNumbers:!0,automaticLayout:!0,minimap:{enabled:!1}}})})]}),(0,e.jsxs)(z,{style:{padding:"20px",background:"#f0f2f5",flex:1,overflowY:"auto"},children:[(0,e.jsx)(Z,{level:4,children:"Rendered Output"}),de&&A&&(0,e.jsx)("div",{style:{width:"100%",height:"600px",padding:"20px",border:"1px solid #d9d9d9",borderRadius:"8px",background:"#fff"},children:(0,e.jsx)(q.Z,{chartID:R,scripts:A,data:le})})]})]})}},67066:function(j,g,a){"use strict";a.d(g,{GQ:function(){return M},MP:function(){return C},fy:function(){return f}});var y=a(15009),m=a.n(y),T=a(97857),u=a.n(T),I=a(99289),p=a.n(I),b=a(11238);function M(d){return P.apply(this,arguments)}function P(){return P=p()(m()().mark(function d(E){return m()().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return r.abrupt("return",(0,b.ZP)("/api/charts",u()({method:"GET"},E||{})));case 1:case"end":return r.stop()}},d)})),P.apply(this,arguments)}function f(d,E){return c.apply(this,arguments)}function c(){return c=p()(m()().mark(function d(E,s){return m()().wrap(function(l){for(;;)switch(l.prev=l.next){case 0:return l.abrupt("return",(0,b.ZP)("/api/chartData/".concat(E),u()({method:"GET"},s||{})));case 1:case"end":return l.stop()}},d)})),c.apply(this,arguments)}function C(d){return v.apply(this,arguments)}function v(){return v=p()(m()().mark(function d(E){return m()().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return r.abrupt("return",(0,b.ZP)("/api/updateChart",u()({method:"POST"},E||{})));case 1:case"end":return r.stop()}},d)})),v.apply(this,arguments)}},24654:function(){}}]);