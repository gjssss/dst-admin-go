(self.webpackChunkgo_manager_web=self.webpackChunkgo_manager_web||[]).push([[4188,4870],{64870:(e,o,s)=>{"use strict";s.r(o)},24188:(e,o,s)=>{var m={"./editorBaseApi":41501,"./editorBaseApi.js":41501,"./editorSimpleWorker":71763,"./editorSimpleWorker.js":71763,"./editorWorker":20618,"./editorWorker.js":20618,"./editorWorkerHost":21264,"./editorWorkerHost.js":21264,"./findSectionHeaders":96377,"./findSectionHeaders.js":96377,"./getIconClasses":11074,"./getIconClasses.js":11074,"./languageFeatureDebounce":99042,"./languageFeatureDebounce.js":99042,"./languageFeatures":36188,"./languageFeatures.js":36188,"./languageFeaturesService":31545,"./languageFeaturesService.js":31545,"./languageService":97368,"./languageService.js":97368,"./languagesAssociations":75122,"./languagesAssociations.js":75122,"./languagesRegistry":31773,"./languagesRegistry.js":31773,"./markerDecorations":40148,"./markerDecorations.js":40148,"./markerDecorationsService":82605,"./markerDecorationsService.js":82605,"./model":8364,"./model.js":8364,"./modelService":86629,"./modelService.js":86629,"./resolverService":57440,"./resolverService.js":57440,"./semanticTokensDto":86330,"./semanticTokensDto.js":86330,"./semanticTokensProviderStyling":18046,"./semanticTokensProviderStyling.js":18046,"./semanticTokensStyling":11341,"./semanticTokensStyling.js":11341,"./semanticTokensStylingService":33754,"./semanticTokensStylingService.js":33754,"./textModelSync/textModelSync.impl":72399,"./textModelSync/textModelSync.impl.js":72399,"./textModelSync/textModelSync.protocol":64870,"./textModelSync/textModelSync.protocol.js":64870,"./textResourceConfiguration":81734,"./textResourceConfiguration.js":81734,"./treeSitterParserService":60346,"./treeSitterParserService.js":60346,"./treeViewsDnd":52037,"./treeViewsDnd.js":52037,"./treeViewsDndService":5906,"./treeViewsDndService.js":5906,"./unicodeTextModelHighlighter":52541,"./unicodeTextModelHighlighter.js":52541,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorBaseApi":41501,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorBaseApi.js":41501,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorSimpleWorker":71763,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorSimpleWorker.js":71763,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorWorker":20618,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorWorker.js":20618,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorWorkerHost":21264,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/editorWorkerHost.js":21264,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/findSectionHeaders":96377,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/findSectionHeaders.js":96377,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/getIconClasses":11074,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/getIconClasses.js":11074,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageFeatureDebounce":99042,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageFeatureDebounce.js":99042,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageFeatures":36188,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageFeatures.js":36188,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageFeaturesService":31545,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageFeaturesService.js":31545,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageService":97368,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languageService.js":97368,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languagesAssociations":75122,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languagesAssociations.js":75122,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languagesRegistry":31773,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/languagesRegistry.js":31773,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/markerDecorations":40148,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/markerDecorations.js":40148,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/markerDecorationsService":82605,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/markerDecorationsService.js":82605,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/model":8364,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/model.js":8364,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/modelService":86629,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/modelService.js":86629,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/resolverService":57440,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/resolverService.js":57440,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensDto":86330,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensDto.js":86330,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensProviderStyling":18046,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensProviderStyling.js":18046,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensStyling":11341,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensStyling.js":11341,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensStylingService":33754,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/semanticTokensStylingService.js":33754,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/textModelSync/textModelSync.impl":72399,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/textModelSync/textModelSync.impl.js":72399,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/textModelSync/textModelSync.protocol":64870,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/textModelSync/textModelSync.protocol.js":64870,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/textResourceConfiguration":81734,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/textResourceConfiguration.js":81734,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/treeSitterParserService":60346,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/treeSitterParserService.js":60346,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/treeViewsDnd":52037,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/treeViewsDnd.js":52037,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/treeViewsDndService":5906,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/treeViewsDndService.js":5906,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/unicodeTextModelHighlighter":52541,".pnpm/monaco-editor@0.52.0/node_modules/monaco-editor/esm/vs/editor/common/services/unicodeTextModelHighlighter.js":52541};function n(e){var o=r(e);return s(o)}function r(e){if(!s.o(m,e)){var o=new Error("Cannot find module '"+e+"'");throw o.code="MODULE_NOT_FOUND",o}return m[e]}n.keys=function(){return Object.keys(m)},n.resolve=r,e.exports=n,n.id=24188}}]);