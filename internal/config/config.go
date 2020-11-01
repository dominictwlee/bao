package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dominictwlee/bao/internal/pkgjson"
	"github.com/evanw/esbuild/pkg/api"
	"io/ioutil"
)

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
	TargetNames = getTargetNames(BuildTargets)

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

	EngineNames = getEngineNames(BuildEngines)

	BuildFormats = map[string]api.Format{
		"iife": api.FormatIIFE,
		"cjs":  api.FormatCommonJS,
		"esm":  api.FormatESModule,
	}

	SourceMapTypes = map[string]api.SourceMap{
		"linked": api.SourceMapLinked,
		"inline": api.SourceMapInline,
	}

	FormatNames = getFormatNames(BuildFormats)
)

func getTargetNames(m map[string]api.Target) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getFormatNames(m map[string]api.Format) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getEngineNames(m map[string]api.EngineName) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type BuildOptions struct {
	Entry     string
	Target    string
	Engines   map[string]string
	Bundle    *bool
	Define    map[string]interface{}
	Format    string
	Minify    *bool
	SourceMap string
}

func ConfigureBuild(cfg *BuildOptions) (*api.BuildOptions, error) {
	opts := api.BuildOptions{
		Color: api.ColorAlways,
		//Loader: map[string]api.Loader{
		//	".js": api.LoaderJSX,
		//},
		Write: true,
	}

	if cfg.Entry != "" {
		opts.EntryPoints = []string{cfg.Entry}
	} else {
		opts.EntryPoints = []string{"src/index.js"}
	}

	pkgJson, err := readPkgJSON()
	if err != nil {
		return nil, err
	}
	if pkgJson.Main == "" {
		return nil, errors.New("main field must be specified in package.json. it is used for your build output")
	}
	opts.Outfile = pkgJson.Main

	if cfg.Bundle != nil {
		opts.Bundle = *cfg.Bundle
	}
	if cfg.Minify != nil {
		opts.MinifyIdentifiers = *cfg.Minify
		opts.MinifySyntax = *cfg.Minify
		opts.MinifyWhitespace = *cfg.Minify
	}

	if len(cfg.Engines) > 0 {
		opts.Engines = []api.Engine{}
		for name, ver := range cfg.Engines {
			engine, ok := BuildEngines[name]
			if !ok {
				return nil, fmt.Errorf("invalid engine name. Must be one of: %v", EngineNames)
			}
			opts.Engines = append(opts.Engines, api.Engine{
				Name:    engine,
				Version: ver,
			})
		}
	}

	if cfg.Format != "" {
		format, ok := BuildFormats[cfg.Format]
		if !ok {
			return nil, fmt.Errorf("invalid build format. Must be one of: %v", FormatNames)
		}
		opts.Format = format
	}

	if cfg.SourceMap != "" {
		format, ok := SourceMapTypes[cfg.SourceMap]
		if !ok {
			return nil, errors.New("invalid sourcemap type. Must be one of: [linked, inline]")
		}
		opts.Sourcemap = format
	}

	if cfg.Target != "" {
		target, ok := BuildTargets[cfg.Target]
		if !ok {
			return nil, fmt.Errorf("invalid build target. Must be one of: %v", TargetNames)
		}
		opts.Target = target
	}

	return &opts, nil
}

func readPkgJSON() (pkgjson.PackageJSON, error) {
	var pkgJson pkgjson.PackageJSON
	pkgJsonData, err := ioutil.ReadFile("package.json")
	if err != nil {
		return pkgJson, err
	}
	if err := json.Unmarshal(pkgJsonData, &pkgJson); err != nil {
		return pkgJson, err
	}
	return pkgJson, nil
}
