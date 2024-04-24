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
 image: appropriate/curl
 要求包含 curl / tar / rsync
 — 2. 容器启动顺序控制
 需要配置如下的环境变量
 KRUISE_CONTAINER_PRIORITY = 1
 保证初始化时，配表的同步逻辑
 同步后启动主容器，不会导致主容器启动失败 ^4jW8Wa0p

GameServerSet-0 ^Y6te5I70

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
			"height": 104.10029506050445,
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
			"version": 3065,
			"versionNonce": 583261405,
			"isDeleted": false,
			"boundElements": [
				{
					"type": "text",
					"id": "6YTLpLBg"
				}
			],
			"updated": 1713939026985,
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
			"version": 2920,
			"versionNonce": 1100655475,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939026985,
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
			"version": 1208,
			"versionNonce": 994982205,
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
			"strokeColor": "#f08c00",
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
			"updated": 1713939026985,
			"link": null,
			"locked": false
		},
		{
			"type": "rectangle",
			"version": 813,
			"versionNonce": 314230035,
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
			"updated": 1713939026985,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 663,
			"versionNonce": 521013661,
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
			"updated": 1713939026985,
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
			"version": 1214,
			"versionNonce": 1937825789,
			"isDeleted": false,
			"id": "166-wHnnYJohQCGllhyGY",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1182.899552190497,
			"y": -544.457360902619,
			"strokeColor": "#ffffff",
			"backgroundColor": "#9775fa",
			"width": 237.76212700000013,
			"height": 47,
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
			"updated": 1713939039427,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 1258,
			"versionNonce": 791989757,
			"isDeleted": false,
			"id": "zISz1qKu",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1139.170832440497,
			"y": -530.1573609026191,
			"strokeColor": "#ffffff",
			"backgroundColor": "transparent",
			"width": 150.3046875,
			"height": 18.4,
			"seed": 1735831795,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939026985,
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
			"version": 1440,
			"versionNonce": 633658451,
			"isDeleted": false,
			"id": "m9e69gIjZJW2S35GvfCsR",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1183.1234812824543,
			"y": -471.9231668737764,
			"strokeColor": "#ffffff",
			"backgroundColor": "#da77f2",
			"width": 237.76212735593472,
			"height": 46.15096492463637,
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
			"updated": 1713939026985,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 1466,
			"versionNonce": 717489757,
			"isDeleted": false,
			"id": "I2Zg4Mtw",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1135.394761354487,
			"y": -458.0476844114582,
			"strokeColor": "#ffffff",
			"backgroundColor": "transparent",
			"width": 142.3046875,
			"height": 18.4,
			"seed": 82844989,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939026985,
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
			"version": 292,
			"versionNonce": 1479268851,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939026985,
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
			"y": -525.1379602766601,
			"width": 115.62542414868722,
			"height": 32.47869167711667,
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
			"version": 4804,
			"versionNonce": 82001811,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939070442,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					115.62542414868722,
					-32.47869167711667
				]
			],
			"lastCommittedPoint": null,
			"startBinding": {
				"elementId": "166-wHnnYJohQCGllhyGY",
				"focus": 0.5449815103524869,
				"gap": 6.335563185238584
			},
			"endBinding": {
				"elementId": "DxFxs4tXhJyjlf3fcP39M",
				"focus": 0.2133270615733646,
				"gap": 9.341184003487115
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
			"strokeColor": "#f08c00",
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
			"version": 3026,
			"versionNonce": 971480691,
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
			"updated": 1713939109227,
			"link": null,
			"locked": false
		},
		{
			"id": "YEBLzA7LV5RT6Jz8CHGb2",
			"type": "arrow",
			"x": -939.8737821349068,
			"y": -601.5282030335841,
			"width": 123.08215267286869,
			"height": 28.780336952265543,
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
			"version": 4338,
			"versionNonce": 432777821,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939067612,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					123.08215267286869,
					28.780336952265543
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
			"x": -1630.7330111472172,
			"y": -686.5527624701074,
			"width": 340.29799067632166,
			"height": 181.8800540021107,
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
			"version": 1511,
			"versionNonce": 1725791997,
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
			"updated": 1713939058716,
			"link": null,
			"locked": false
		},
		{
			"id": "4jW8Wa0p",
			"type": "text",
			"x": -1625.7330111472172,
			"y": -681.5527624701074,
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
			"version": 1295,
			"versionNonce": 519813971,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939058716,
			"link": null,
			"locked": false,
			"text": " configReload-sidecar Notes：\n — 1. 镜像选择\n image: appropriate/curl\n 要求包含 curl / tar / rsync\n — 2. 容器启动顺序控制\n 需要配置如下的环境变量\n KRUISE_CONTAINER_PRIORITY = 1\n 保证初始化时，配表的同步逻辑\n 同步后启动主容器，不会导致主容器启动失败",
			"rawText": " configReload-sidecar Notes：\n — 1. 镜像选择\n image: appropriate/curl\n 要求包含 curl / tar / rsync\n — 2. 容器启动顺序控制\n 需要配置如下的环境变量\n KRUISE_CONTAINER_PRIORITY = 1\n 保证初始化时，配表的同步逻辑\n 同步后启动主容器，不会导致主容器启动失败",
			"fontSize": 16,
			"fontFamily": 2,
			"textAlign": "left",
			"verticalAlign": "top",
			"baseline": 162,
			"containerId": "ajyFHx1Vez30Lk9Jg8Hcj",
			"originalText": " configReload-sidecar Notes：\n — 1. 镜像选择\n image: appropriate/curl\n 要求包含 curl / tar / rsync\n — 2. 容器启动顺序控制\n 需要配置如下的环境变量\n KRUISE_CONTAINER_PRIORITY = 1\n 保证初始化时，配表的同步逻辑\n 同步后启动主容器，不会导致主容器启动失败",
			"lineHeight": 1.15
		},
		{
			"id": "3IWBCKp51deBGzBIoti32",
			"type": "arrow",
			"x": -1186.5862400711735,
			"y": -528.2956068643422,
			"width": 91.23542752716821,
			"height": 59.243969733669815,
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
			"version": 365,
			"versionNonce": 451849629,
			"isDeleted": false,
			"boundElements": null,
			"updated": 1713939064549,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-91.23542752716821,
					-59.243969733669815
				]
			],
			"lastCommittedPoint": [
				-73.45627540653823,
				-38.84387610980673
			],
			"startBinding": {
				"elementId": "166-wHnnYJohQCGllhyGY",
				"focus": -0.7175224497576083,
				"gap": 3.6866878806766863
			},
			"endBinding": {
				"elementId": "ajyFHx1Vez30Lk9Jg8Hcj",
				"focus": -0.5491031919061399,
				"gap": 12.613352872553833
			},
			"startArrowhead": null,
			"endArrowhead": "arrow"
		},
		{
			"type": "text",
			"version": 294,
			"versionNonce": 1156009203,
			"isDeleted": false,
			"id": "Y6te5I70",
			"fillStyle": "solid",
			"strokeWidth": 4,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -1457.3404103541616,
			"y": -804.0368455250385,
			"strokeColor": "#ffffff",
			"backgroundColor": "#e6fcf5",
			"width": 239.11990356445312,
			"height": 35,
			"seed": 596697949,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939111528,
			"link": null,
			"locked": false,
			"fontSize": 28,
			"fontFamily": 1,
			"text": "GameServerSet-0",
			"rawText": "GameServerSet-0",
			"textAlign": "left",
			"verticalAlign": "top",
			"containerId": null,
			"originalText": "GameServerSet-0",
			"lineHeight": 1.25,
			"baseline": 25
		},
		{
			"id": "jPRtZG1O36Rgo4R9Kn04K",
			"type": "line",
			"x": -750.2952212356925,
			"y": -550.2207097098242,
			"width": 204.5235738439228,
			"height": 10.417884004084499,
			"angle": 0,
			"strokeColor": "#1971c2",
			"backgroundColor": "#ffc9c9",
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
			"seed": 719934141,
			"version": 266,
			"versionNonce": 1955410621,
			"isDeleted": true,
			"boundElements": null,
			"updated": 1713939026985,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-204.5235738439228,
					10.417884004084499
				]
			],
			"lastCommittedPoint": null,
			"startBinding": null,
			"endBinding": null,
			"startArrowhead": null,
			"endArrowhead": null
		},
		{
			"id": "WUehjwOct3fWEJQliq6N3",
			"type": "line",
			"x": -746.4042532528804,
			"y": -464.86716883425737,
			"width": 207.61035428957757,
			"height": 67.05624362862733,
			"angle": 0,
			"strokeColor": "#1971c2",
			"backgroundColor": "#ffc9c9",
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
			"seed": 1225774419,
			"version": 278,
			"versionNonce": 1277305747,
			"isDeleted": true,
			"boundElements": null,
			"updated": 1713939026985,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-207.61035428957757,
					-67.05624362862733
				]
			],
			"lastCommittedPoint": null,
			"startBinding": null,
			"endBinding": null,
			"startArrowhead": null,
			"endArrowhead": null
		},
		{
			"id": "FmZ-LnbV29-e0IVYVBaef",
			"type": "line",
			"x": -752.1310432902133,
			"y": -380.4071696666433,
			"width": 200.99002254429195,
			"height": 142.80421114370324,
			"angle": 0,
			"strokeColor": "#1971c2",
			"backgroundColor": "#ffc9c9",
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
			"seed": 828239165,
			"version": 193,
			"versionNonce": 422695709,
			"isDeleted": true,
			"boundElements": null,
			"updated": 1713939026985,
			"link": null,
			"locked": false,
			"points": [
				[
					0,
					0
				],
				[
					-200.99002254429195,
					-142.80421114370324
				]
			],
			"lastCommittedPoint": null,
			"startBinding": null,
			"endBinding": null,
			"startArrowhead": null,
			"endArrowhead": null
		},
		{
			"type": "rectangle",
			"version": 737,
			"versionNonce": 702429491,
			"isDeleted": true,
			"id": "_I-EqLUwkYzx_O7Qh2Ijd",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -749.8644056498041,
			"y": -652.8024868313053,
			"strokeColor": "#1971c2",
			"backgroundColor": "#b2f2bb",
			"width": 172.25801319082905,
			"height": 61.10026160168893,
			"seed": 129608029,
			"groupIds": [],
			"frameId": null,
			"roundness": {
				"type": 3
			},
			"boundElements": [
				{
					"type": "text",
					"id": "RYvSIVyk"
				}
			],
			"updated": 1713939026985,
			"link": null,
			"locked": false
		},
		{
			"type": "text",
			"version": 617,
			"versionNonce": 354641789,
			"isDeleted": true,
			"id": "RYvSIVyk",
			"fillStyle": "solid",
			"strokeWidth": 2,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 100,
			"angle": 0,
			"x": -709.9814928043896,
			"y": -631.4523560304608,
			"strokeColor": "#1971c2",
			"backgroundColor": "transparent",
			"width": 92.4921875,
			"height": 18.4,
			"seed": 148272573,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"boundElements": [],
			"updated": 1713939026985,
			"link": null,
			"locked": false,
			"fontSize": 16,
			"fontFamily": 2,
			"text": "Init-container",
			"rawText": "Init-container",
			"textAlign": "center",
			"verticalAlign": "middle",
			"containerId": "_I-EqLUwkYzx_O7Qh2Ijd",
			"originalText": "Init-container",
			"lineHeight": 1.15,
			"baseline": 15
		},
		{
			"id": "GVxvZWnV",
			"type": "text",
			"x": -726.1243631284522,
			"y": -588.7496199088453,
			"width": 10,
			"height": 25,
			"angle": 0,
			"strokeColor": "#f08c00",
			"backgroundColor": "#f3f0ff",
			"fillStyle": "hachure",
			"strokeWidth": 1,
			"strokeStyle": "solid",
			"roughness": 1,
			"opacity": 80,
			"groupIds": [],
			"frameId": null,
			"roundness": null,
			"seed": 629812819,
			"version": 2653,
			"versionNonce": 2062637075,
			"isDeleted": true,
			"boundElements": null,
			"updated": 1713939109231,
			"link": null,
			"locked": false,
			"text": "",
			"rawText": "",
			"fontSize": 20,
			"fontFamily": 1,
			"textAlign": "center",
			"verticalAlign": "middle",
			"baseline": 18,
			"containerId": "DxFxs4tXhJyjlf3fcP39M",
			"originalText": "",
			"lineHeight": 1.25
		},
		{
			"id": "kYzyV3zt",
			"type": "text",
			"x": -1445.7755606597714,
			"y": -593.5256900813197,
			"width": 7.779296875,
			"height": 32.199999999999996,
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
			"seed": 240142141,
			"version": 15,
			"versionNonce": 741250909,
			"isDeleted": true,
			"boundElements": null,
			"updated": 1713939058716,
			"link": null,
			"locked": false,
			"text": "",
			"rawText": "",
			"fontSize": 28,
			"fontFamily": 2,
			"textAlign": "left",
			"verticalAlign": "top",
			"baseline": 26,
			"containerId": "ajyFHx1Vez30Lk9Jg8Hcj",
			"originalText": "",
			"lineHeight": 1.15
		}
	],
	"appState": {
		"theme": "light",
		"viewBackgroundColor": "#fff9db",
		"currentItemStrokeColor": "#f08c00",
		"currentItemBackgroundColor": "#12b886",
		"currentItemFillStyle": "solid",
		"currentItemStrokeWidth": 1,
		"currentItemStrokeStyle": "solid",
		"currentItemRoughness": 2,
		"currentItemOpacity": 90,
		"currentItemFontFamily": 2,
		"currentItemFontSize": 16,
		"currentItemTextAlign": "left",
		"currentItemStartArrowhead": null,
		"currentItemEndArrowhead": "arrow",
		"scrollX": 1707.4323892207105,
		"scrollY": 990.9295302733902,
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