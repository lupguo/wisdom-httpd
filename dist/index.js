/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./assets/src/autoload.ts":
/*!********************************!*\
  !*** ./assets/src/autoload.ts ***!
  \********************************/
/***/ (function(__unused_webpack_module, exports) {

eval("\nvar __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {\n    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }\n    return new (P || (P = Promise))(function (resolve, reject) {\n        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }\n        function rejected(value) { try { step(generator[\"throw\"](value)); } catch (e) { reject(e); } }\n        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }\n        step((generator = generator.apply(thisArg, _arguments || [])).next());\n    });\n};\nvar __generator = (this && this.__generator) || function (thisArg, body) {\n    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;\n    return g = { next: verb(0), \"throw\": verb(1), \"return\": verb(2) }, typeof Symbol === \"function\" && (g[Symbol.iterator] = function() { return this; }), g;\n    function verb(n) { return function (v) { return step([n, v]); }; }\n    function step(op) {\n        if (f) throw new TypeError(\"Generator is already executing.\");\n        while (g && (g = 0, op[0] && (_ = 0)), _) try {\n            if (f = 1, y && (t = op[0] & 2 ? y[\"return\"] : op[0] ? y[\"throw\"] || ((t = y[\"return\"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;\n            if (y = 0, t) op = [op[0] & 2, t.value];\n            switch (op[0]) {\n                case 0: case 1: t = op; break;\n                case 4: _.label++; return { value: op[1], done: false };\n                case 5: _.label++; y = op[1]; op = [0]; continue;\n                case 7: op = _.ops.pop(); _.trys.pop(); continue;\n                default:\n                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }\n                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }\n                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }\n                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }\n                    if (t[2]) _.ops.pop();\n                    _.trys.pop(); continue;\n            }\n            op = body.call(thisArg, _);\n        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }\n        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };\n    }\n};\nObject.defineProperty(exports, \"__esModule\", ({ value: true }));\nexports[\"default\"] = autoRefreshWisdom;\nvar WISDOM_URI = 'http://127.0.0.1:1666/api/wisdom';\nvar REFRESH_INTERVAL = 10000; // 刷新间隔时间（毫秒）\n// 发送请求以获取智慧内容\nfunction fetchWisdom() {\n    return __awaiter(this, void 0, void 0, function () {\n        var response, data, error_1;\n        return __generator(this, function (_a) {\n            switch (_a.label) {\n                case 0:\n                    _a.trys.push([0, 3, , 4]);\n                    return [4 /*yield*/, fetch(WISDOM_URI)];\n                case 1:\n                    response = _a.sent();\n                    if (!response.ok) {\n                        throw new Error(\"HTTP error! Status: \".concat(response.status));\n                    }\n                    return [4 /*yield*/, response.json()];\n                case 2:\n                    data = _a.sent();\n                    console.log(\"Fetched wisdom:\", data);\n                    return [2 /*return*/, data];\n                case 3:\n                    error_1 = _a.sent();\n                    console.error(\"Failed to fetch wisdom:\", error_1);\n                    return [3 /*break*/, 4];\n                case 4: return [2 /*return*/];\n            }\n        });\n    });\n}\n// 更新模板中的内容\nfunction updateTemplate(sentence) {\n    var contentDiv = document.querySelector('.content');\n    if (contentDiv) {\n        contentDiv.textContent = sentence; // 填充句子到模板\n    }\n    else {\n        console.error(\"Content div not found!\");\n    }\n}\n// 自动请求wisdomURI地址\nfunction refreshWisdom() {\n    console.log(\"Refreshing wisdom...\");\n    fetchWisdom().then(function (data) {\n        // 从数据中提取句子并更新模板\n        if (data && data.status === \"success\" && data.data && data.data.PageData) {\n            var sentence = data.data.PageData.sentence;\n            updateTemplate(sentence);\n        }\n        else {\n            console.error(\"Invalid data structure:\", data);\n        }\n    }).catch(function (err) {\n        console.error(\"Got error: \" + err);\n    });\n}\n// 启动自动刷新\nfunction autoRefreshWisdom() {\n    setInterval(refreshWisdom, REFRESH_INTERVAL);\n}\n\n\n//# sourceURL=webpack://wisdom-httpd/./assets/src/autoload.ts?");

/***/ }),

/***/ "./assets/src/index.ts":
/*!*****************************!*\
  !*** ./assets/src/index.ts ***!
  \*****************************/
/***/ (function(__unused_webpack_module, exports, __webpack_require__) {

eval("\nvar __importDefault = (this && this.__importDefault) || function (mod) {\n    return (mod && mod.__esModule) ? mod : { \"default\": mod };\n};\nObject.defineProperty(exports, \"__esModule\", ({ value: true }));\nvar autoload_1 = __importDefault(__webpack_require__(/*! ./autoload */ \"./assets/src/autoload.ts\"));\n// 自动加载\n(0, autoload_1.default)();\nconsole.log(\"wisdom amazing!\");\n\n\n//# sourceURL=webpack://wisdom-httpd/./assets/src/index.ts?");

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module is referenced by other modules so it can't be inlined
/******/ 	var __webpack_exports__ = __webpack_require__("./assets/src/index.ts");
/******/ 	
/******/ })()
;