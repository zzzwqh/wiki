---

excalidraw-plugin: parsed
tags: [excalidraw]

---
==⚠  Switch to EXCALIDRAW VIEW in the MORE OPTIONS menu of this document. ⚠==


# Text Elements
Game-container ^zv3jv4X1

configReload-sidecar ^zISz1qKu

codeReload-sidecar ^I2Zg4Mtw

GameServerSet-0 ^7G67mrDv

容器的本地目录，利用镜像传输配表压缩包，并做同步到容器 emptydir 共享目录即可 ^6YTLpLBg

 configReload-sidecar Notes：
 — 1. 镜像选择
 image: appropriate/curl （apk add rsync）
 要求包含 curl / tar / rsync
 — 2. 容器启动顺序控制
 需要配置如下的环境变量
 KRUISE_CONTAINER_PRIORITY = 1
 保证初始化时，配表的同步逻辑
 同步后启动主容器，不会导致主容器启动失败 ^4jW8Wa0p

Emptydir ^Y6te5I70

%%
# Drawing
```json
{
	"type": "excalidraw",
	"version": 2,
	"source": "https://github.com/zsviczian/obsidian-excalidraw-plugin/releases/tag/2.0.17",
	"elements": [
		{
			"id": "YopnGtPxVxWkr1jQnn23d",
			"type": "rectangle",
			"x": -808.8249939055324,
			"y": -557.2183369172159,
			"width": 174.08316243345507,
			"height": 85.64935738743357,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#b2f2bb",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 0,
			"opacity": 90,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 2087708339,
			"version": 3104,
			"versionNonce": 197243539,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "6YTLpLBg"
				}
			],
			"updated": 1713939329567,
			"link": null,
			"locked": false
		},
		{
			"id": "6YTLpLBg",
			"type": "text",
			"x": -803.8249939055324,
			"y": -552.2183369172159,
			"width": 161.796875,
			"height": 73.6,
			"angle": 0,
			"strokeColor": "#1e1e1e",
			"backgroundColor": "#b2f2bb",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 0,
			"opacity": 90,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 1010486803,
			"version": 2925,
			"versionNonce": 1347859485,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939329567,
			"link": null,
			"locked": false,
			"text": "容器的本地目录，利用\n镜像传输配表压缩包，\n并做同步到容器\nemptydir 共享目录即可",
			"rawText": "容器的本地目录，利用镜像传输配表压缩包，并做同步到容器 emptydir 共享目录即可",
			"fontSize": 16,
			"fontFamily": 2,
			"textAlign": "left",
			"verticalAlign": "top",
			"baseline": 70,
			"containerId": "YopnGtPxVxWkr1jQnn23d",
			"originalText": "容器的本地目录，利用镜像传输配表压缩包，并做同步到容器 emptydir 共享目录即可",
			"lineHeight": 1.15
		},
		{
			"type": "rectangle",
			"version": 1213,
			"versionNonce": 213116979,
			"isDeleted": false,
			"id": "qFtxZpH6TZkAaUO95Ss9-",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 70,
			"angle": 0,
			"x": -1238.7177432238827,
			"y": -753.5567874452358,
			"strokeColor": "#ffffff",
			"backgroundColor": "#12b886",
			"width": 346.67179311929544,
			"height": 406.8929026932842,
			"seed": 164441149,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [],
			"updated": 1713939329567,
			"link": null,
			"locked": false
		},
		{
			"type": "rectangle",
			"version": 829,
			"versionNonce": 923019389,
			"isDeleted": false,
			"id": "wsi_BxcKEJyiBtFso0rix",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1183.123481,
			"y": -631.5396724941447,
			"strokeColor": "#ffffff",
			"backgroundColor": "#4dabf7",
			"width": 237.76212700000008,
			"height": 61.10026160168893,
			"seed": 476089149,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"type": "text",
					"id": "zv3jv4X1"
				},
				{
					"id": "YEBLzA7LV5RT6Jz8CHGb2",
					"type": "arrow"
				}
			],
			"updated": 1713939329567,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 668,
			"versionNonce": 1225681363,
			"isDeleted": false,
			"id": "zv3jv4X1",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1121.6017925,
			"y": -610.1895416933003,
			"strokeColor": "#ffffff",
			"backgroundColor": "transparent",
			"width": 114.71875,
			"height": 18.4,
			"seed": 1771852189,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939329567,
			"link": null,
			"locked": false,
			"fontSize": 16,
			"fontFamily": 2,
			"text": "Game-container",
			"rawText": "Game-container",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "wsi_BxcKEJyiBtFso0rix",
			"originalText": "Game-container",
			"lineHeight": 1.15,
			"baseline": 15
		},
		{
			"type": "rectangle",
			"version": 1228,
			"versionNonce": 1690570973,
			"isDeleted": false,
			"id": "166-wHnnYJohQCGllhyGY",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1182.8995521904967,
			"y": -544.457360902619,
			"strokeColor": "#ffffff",
			"backgroundColor": "#9775fa",
			"width": 237.76212700000013,
			"height": 50.358927762203734,
			"seed": 1726586419,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"type": "text",
					"id": "zISz1qKu"
				},
				{
					"id": "1SrL3CYyzMEpUxzdMHhKl",
					"type": "arrow"
				},
				{
					"id": "3IWBCKp51deBGzBIoti32",
					"type": "arrow"
				}
			],
			"updated": 1713939329567,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 1272,
			"versionNonce": 1172813683,
			"isDeleted": false,
			"id": "zISz1qKu",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1139.1708324404967,
			"y": -528.4778970215172,
			"strokeColor": "#ffffff",
			"backgroundColor": "transparent",
			"width": 150.3046875,
			"height": 18.4,
			"seed": 1735831795,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939329567,
			"link": null,
			"locked": false,
			"fontSize": 16,
			"fontFamily": 2,
			"text": "configReload-sidecar",
			"rawText": "configReload-sidecar",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "166-wHnnYJohQCGllhyGY",
			"originalText": "configReload-sidecar",
			"lineHeight": 1.15,
			"baseline": 15
		},
		{
			"type": "rectangle",
			"version": 1701,
			"versionNonce": 925363517,
			"isDeleted": false,
			"id": "m9e69gIjZJW2S35GvfCsR",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1182.8995516810367,
			"y": -468.1163826911807,
			"strokeColor": "#ffffff",
			"backgroundColor": "#da77f2",
			"width": 237.76212735593472,
			"height": 47.90764553590387,
			"seed": 1406288093,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"type": "text",
					"id": "I2Zg4Mtw"
				}
			],
			"updated": 1713939329568,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 1727,
			"versionNonce": 562450707,
			"isDeleted": false,
			"id": "I2Zg4Mtw",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1135.1708317530693,
			"y": -453.3625599232288,
			"strokeColor": "#ffffff",
			"backgroundColor": "transparent",
			"width": 142.3046875,
			"height": 18.4,
			"seed": 82844989,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939329568,
			"link": null,
			"locked": false,
			"fontSize": 16,
			"fontFamily": 2,
			"text": "codeReload-sidecar",
			"rawText": "codeReload-sidecar",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "m9e69gIjZJW2S35GvfCsR",
			"originalText": "codeReload-sidecar",
			"lineHeight": 1.15,
			"baseline": 15
		},
		{
			"id": "7G67mrDv",
			"type": "text",
			"x": -1187.1040388216948,
			"y": -712.9916110443893,
			"width": 239.11990356445312,
			"height": 35,
			"angle": 0,
			"strokeColor": "#ffffff",
			"backgroundColor": "#e6fcf5",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 734939037,
			"version": 296,
			"versionNonce": 1569799581,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939329568,
			"link": null,
			"locked": false,
			"text": "GameServerSet-0",
			"rawText": "GameServerSet-0",
			"fontSize": 28,
			"fontFamily": 1,
			"textAlign": "left",
			"verticalAlign": "top",
			"baseline": 25,
			"containerId": null,
			"originalText": "GameServerSet-0",
			"lineHeight": 1.25
		},
		{
			"id": "1SrL3CYyzMEpUxzdMHhKl",
			"type": "arrow",
			"x": -938.801862005258,
			"y": -523.1618985178395,
			"width": 115.62542414868722,
			"height": 33.63323063406767,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "#b2f2bb",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 80,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 160704115,
			"version": 4828,
			"versionNonce": 1835054771,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939329568,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					115.62542414868722,
					-33.63323063406767
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "166-wHnnYJohQCGllhyGY",
				"gap": 6.335563185238584,
				"focus": 0.5449815103524869
			},
			"endBinding": {
				"elementId": "DxFxs4tXhJyjlf3fcP39M",
				"gap": 9.341184003487115,
				"focus": 0.2133270615733646
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "DxFxs4tXhJyjlf3fcP39M",
			"type": "rectangle",
			"x": -813.8352538530837,
			"y": -597.2387221117617,
			"width": 185.42178144926305,
			"height": 41.978204405833004,
			"angle": 0,
			"strokeColor": "#ffffff",
			"backgroundColor": "#12b886",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 70,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"seed": 617600157,
			"version": 3031,
			"versionNonce": 523338237,
			"isDeleted": false,
			"boundElements": [
				{
					"id": "1SrL3CYyzMEpUxzdMHhKl",
					"type": "arrow"
				},
				{
					"id": "YEBLzA7LV5RT6Jz8CHGb2",
					"type": "arrow"
				}
			],
			"updated": 1713939329568,
			"link": null,
			"locked": false
		},
		{
			"id": "YEBLzA7LV5RT6Jz8CHGb2",
			"type": "arrow",
			"x": -939.8737821349068,
			"y": -601.6030495630259,
			"width": 123.08215267286869,
			"height": 28.837016005619944,
			"angle": 0,
			"strokeColor": "#2f9e44",
			"backgroundColor": "#b2f2bb",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 80,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 2
			},
			"seed": 1833583901,
			"version": 4344,
			"versionNonce": 1209886803,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939329568,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					123.08215267286869,
					28.837016005619944
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "wsi_BxcKEJyiBtFso0rix",
				"gap": 5.48757186509323,
				"focus": -0.5095871187441112
			},
			"endBinding": {
				"elementId": "DxFxs4tXhJyjlf3fcP39M",
				"gap": 2.9563756089545166,
				"focus": -0.6063520844122358
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"id": "ajyFHx1Vez30Lk9Jg8Hcj",
			"type": "rectangle",
			"x": -1595.726806526641,
			"y": -743.2105382291188,
			"width": 340.29799067632166,
			"height": 176,
			"angle": 0,
			"strokeColor": "#ffffff",
			"backgroundColor": "#9775fa",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 2,
			"opacity": 90,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 1936756381,
			"version": 1834,
			"versionNonce": 1628060253,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "4jW8Wa0p"
				},
				{
					"id": "3IWBCKp51deBGzBIoti32",
					"type": "arrow"
				}
			],
			"updated": 1713939329568,
			"link": null,
			"locked": false
		},
		{
			"id": "4jW8Wa0p",
			"type": "text",
			"x": -1590.726806526641,
			"y": -738.2105382291188,
			"width": 324.4453125,
			"height": 165.6,
			"angle": 0,
			"strokeColor": "#ffffff",
			"backgroundColor": "#9775fa",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 0,
			"opacity": 90,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 997061565,
			"version": 1671,
			"versionNonce": 932374003,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939329568,
			"link": null,
			"locked": false,
			"text": " configReload-sidecar Notes：\n — 1. 镜像选择\n image: appropriate/curl （apk add rsync）\n 要求包含 curl / tar / rsync\n — 2. 容器启动顺序控制\n 需要配置如下的环境变量\n KRUISE_CONTAINER_PRIORITY = 1\n 保证初始化时，配表的同步逻辑\n 同步后启动主容器，不会导致主容器启动失败",
			"rawText": " configReload-sidecar Notes：\n — 1. 镜像选择\n image: appropriate/curl （apk add rsync）\n 要求包含 curl / tar / rsync\n — 2. 容器启动顺序控制\n 需要配置如下的环境变量\n KRUISE_CONTAINER_PRIORITY = 1\n 保证初始化时，配表的同步逻辑\n 同步后启动主容器，不会导致主容器启动失败",
			"fontSize": 16,
			"fontFamily": 2,
			"textAlign": "left",
			"verticalAlign": "top",
			"baseline": 162,
			"containerId": "ajyFHx1Vez30Lk9Jg8Hcj",
			"originalText": " configReload-sidecar Notes：\n — 1. 镜像选择\n image: appropriate/curl （apk add rsync）\n 要求包含 curl / tar / rsync\n — 2. 容器启动顺序控制\n 需要配置如下的环境变量\n KRUISE_CONTAINER_PRIORITY = 1\n 保证初始化时，配表的同步逻辑\n 同步后启动主容器，不会导致主容器启动失败",
			"lineHeight": 1.15
		},
		{
			"id": "3IWBCKp51deBGzBIoti32",
			"type": "arrow",
			"x": -1190.2308697349897,
			"y": -525.9975847043497,
			"width": 53.040171950752665,
			"height": 77.9576745252017,
			"angle": 0,
			"strokeColor": "#6741d9",
			"backgroundColor": "#da77f2",
			"fillStyle": "solid",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 2,
			"opacity": 90,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 202134963,
			"version": 1059,
			"versionNonce": 1810687677,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939329568,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-53.040171950752665,
					-77.9576745252017
				]
			],
			"lastCommittedPoint": [
				-73.45627540653823,
				-38.84387610980673
			],
			"startBinding": {
				"elementId": "166-wHnnYJohQCGllhyGY",
				"gap": 7.33131754449289,
				"focus": -0.8945328658142531
			},
			"endBinding": {
				"elementId": "ajyFHx1Vez30Lk9Jg8Hcj",
				"gap": 12.15777416457695,
				"focus": -0.6409576036260821
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "text",
			"version": 505,
			"versionNonce": 1120382867,
			"isDeleted": false,
			"id": "Y6te5I70",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -761.9226776462308,
			"y": -589.7215942286296,
			"strokeColor": "#ffffff",
			"backgroundColor": "#e6fcf5",
			"width": 81.0399169921875,
			"height": 25,
			"seed": 596697949,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939329568,
			"link": null,
			"locked": false,
			"fontSize": 20,
			"fontFamily": 1,
			"text": "Emptydir",
			"rawText": "Emptydir",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "Emptydir",
			"lineHeight": 1.25,
			"baseline": 18
		}
	],
	"appState": {
		"theme": "light",
		"viewBackgroundColor": "#fff9db",
		"currentItemStrokeColor": "#ffffff",
		"currentItemBackgroundColor": "#40c057",
		"currentItemFillStyle": "solid",
		"currentItemStrokeWidth": 1,
		"currentItemStrokeStyle": "solid",
		"currentItemRoughness": 2,
		"currentItemOpacity": 90,
		"currentItemFontFamily": 2,
		"currentItemFontSize": 20,
		"currentItemTextAlign": "left",
		"currentItemStartArrowhead": null,
		"currentItemEndArrowhead": "arrow",
		"scrollX": 1707.4323892207105,
		"scrollY": 938.6296737231,
		"zoom": {
			"value": 1.0117626041979357
		},
		"currentItemRoundness": "sharp",
		"gridSize": null,
		"gridColor": {
			"Bold": "#FFE770FF",
			"Regular": "#FFF3B7FF"
		},
		"currentStrokeOptions": null,
		"previousGridSize": null,
		"frameRendering": {
			"enabled": true,
			"clip": true,
			"name": true,
			"outline": true
		}
	},
	"files": {}
}
```
%%