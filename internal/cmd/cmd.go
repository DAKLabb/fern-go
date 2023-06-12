package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fern-api/fern-go"
	"github.com/fern-api/fern-go/internal/fern/generatorexec"
	"github.com/fern-api/fern-go/internal/generator"
	"github.com/fern-api/fern-go/internal/goexec"
	"github.com/fern-api/fern-go/internal/writer"
)

// defaultImports specify the default imports used in the generated
// go.mod (if any).
var defaultImports = map[string]string{
	"github.com/gofrs/uuid/v5": "v5.0.0",
}

// Config represents the common configuration required from all of
// the commands (e.g. fern-go-{client,model}).
type Config struct {
	DryRun     bool
	IrFilepath string
	ImportPath string
	Module     *generator.ModuleConfig
	Writer     *writer.Config
}

// GeneratorFunc is a function that generates files.
type GeneratorFunc func(*Config) ([]*generator.File, error)

// Run runs the given command that produces generated files.
func Run(usage string, fn GeneratorFunc) {
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Fprintln(os.Stdout, fern.Version)
		os.Exit(0)
	}
	if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Fprintln(os.Stdout, usage)
		os.Exit(0)
	}
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	config, err := newConfig(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	files, err := fn(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err := writeFiles(config.Writer, config.Module, files); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// newConfig returns the *Config found at the given configFilename.
func newConfig(configFilename string) (*Config, error) {
	config, err := readConfig(configFilename)
	if err != nil {
		return nil, err
	}
	customConfig, err := customConfigFromConfig(config)
	if err != nil {
		return nil, err
	}
	outputMode, err := outputModeFromConfig(config)
	if err != nil {
		return nil, err
	}
	writerConfig, err := writer.NewConfig(outputMode)
	if err != nil {
		return nil, err
	}
	moduleConfig, err := moduleConfigFromCustomConfig(customConfig, outputMode)
	if err != nil {
		return nil, err
	}
	importPath := customConfig.ImportPath
	if moduleConfig != nil {
		if customConfig.ImportPath == "" {
			importPath = moduleConfig.Path
		}
		if importPath != moduleConfig.Path {
			return nil, fmt.Errorf(
				"both module path (%q) and import path (%q) are specified, but not equal; please remove import path",
				moduleConfig.Path,
				importPath,
			)
		}
	}
	return &Config{
		DryRun:     config.DryRun,
		IrFilepath: config.IrFilepath,
		ImportPath: importPath,
		Module:     moduleConfig,
		Writer:     writerConfig,
	}, nil
}

// writeFiles writes the given files according to the configuration.
func writeFiles(
	writerConfig *writer.Config,
	moduleConfig *generator.ModuleConfig,
	files []*generator.File,
) error {
	writer, err := writer.New(writerConfig)
	if err != nil {
		return err
	}
	if err := writer.WriteFiles(files); err != nil {
		return err
	}
	// If a module file was written, we still need to generate the go.sum.
	if moduleConfig != nil {
		return goexec.RunTidy(writer.Root())
	}
	return nil
}

// readConfig returns the generator configuration from the given filename.
func readConfig(configFilename string) (*generatorexec.GeneratorConfig, error) {
	bytes, err := os.ReadFile(configFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to read generator configuration: %v", err)
	}
	config := new(generatorexec.GeneratorConfig)
	if err := json.Unmarshal(bytes, config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal generator configuration: %v", err)
	}
	return config, nil
}

type customConfig struct {
	ImportPath string        `json:"importPath,omitempty"`
	Module     *moduleConfig `json:"module,omitempty"`
}

type moduleConfig struct {
	Path    string            `json:"path,omitempty"`
	Version string            `json:"version,omitempty"`
	Imports map[string]string `json:"imports,omitempty"`
}

func customConfigFromConfig(c *generatorexec.GeneratorConfig) (*customConfig, error) {
	if c.CustomConfig == nil {
		return &customConfig{}, nil
	}
	configBytes, err := json.Marshal(c.CustomConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize custom configuration: %v", err)
	}
	// We use a custom decoder here to validate the custom configuration fields.
	decoder := json.NewDecoder(bytes.NewReader(configBytes))
	decoder.DisallowUnknownFields()

	config := new(customConfig)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("failed to read custom configuration: %v", err)
	}
	return config, nil
}

func moduleConfigFromCustomConfig(customConfig *customConfig, outputMode writer.OutputMode) (*generator.ModuleConfig, error) {
	githubConfig, ok := outputMode.(*writer.GithubConfig)
	if !ok && customConfig.Module == nil || customConfig.Module == (&moduleConfig{}) {
		// Neither a GitHub configuration nor a module configuration were provided.
		return nil, nil
	}
	var modulePath string
	if ok {
		modulePath = strings.TrimPrefix(githubConfig.RepoURL, "https://")
	}
	if customConfig.Module == nil || customConfig.Module == (&moduleConfig{}) {
		// A GitHub configuration was provided, so the module config should use
		// the GitHub configuration's repository url.
		return &generator.ModuleConfig{
			Path:    modulePath,
			Imports: defaultImports,
		}, nil
	}
	if customConfig.Module.Path != "" {
		// The user explicitly specified a module path, so we should use it.
		modulePath = customConfig.Module.Path
	}
	imports := customConfig.Module.Imports
	if imports == nil {
		imports = defaultImports
	}
	return &generator.ModuleConfig{
		Path:    modulePath,
		Version: customConfig.Module.Version, // If not specified, defaults to the Go command's current version.
		Imports: imports,
	}, nil
}

func outputModeFromConfig(c *generatorexec.GeneratorConfig) (writer.OutputMode, error) {
	switch outputConfigMode := c.Output.Mode; outputConfigMode.Type {
	case "github":
		return writer.NewGithubConfig(c.Output.Path, outputConfigMode.Github.RepoUrl)
	case "downloadFiles":
		return writer.NewLocalConfig(c.Output.Path)
	default:
		return nil, fmt.Errorf("unrecognized output configuration mode: %T", outputConfigMode)
	}
}