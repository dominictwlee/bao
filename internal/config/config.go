package config

import "github.com/evanw/esbuild/pkg/api"

var (
	BuildTargets = map[string]api.Target{
		"es5":    api.ES5,
		"es2015": api.ES2015,
		"es2016": api.ES2016,
		"es2017": api.ES2017,
		"es2018": api.ES2018,
		"es2019": api.ES2019,
		"es2020": api.ES2020,
		"esnext": api.ESNext,
	}

	BuildEngines = map[string]api.EngineName{
		"chrome": api.EngineChrome,
		"ch":     api.EngineChrome,

		"firefox": api.EngineFirefox,
		"ff":      api.EngineFirefox,

		"edge": api.EngineEdge,

		"safari": api.EngineSafari,
		"sf":     api.EngineSafari,

		"node": api.EngineNode,
		"ios":  api.EngineIOS,
	}

	BuildFormats = map[string]api.Format{
		"iife": api.FormatIIFE,
		"cjs":  api.FormatCommonJS,
		"esm":  api.FormatESModule,
	}
)
