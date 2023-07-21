package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vercel/turbo/cli/internal/env"
	"github.com/vercel/turbo/cli/internal/lockfile"
	"github.com/vercel/turbo/cli/internal/turbopath"
	"github.com/vercel/turbo/cli/internal/util"
)

// Code generated by capnpc-go. DO NOT EDIT.

func Test_CapnpHash(t *testing.T) {
	taskHashable := TaskHashable{
		GlobalHash:           "global_hash",
		TaskDependencyHashes: []string{"task_dependency_hash"},
		PackageDir:           turbopath.AnchoredUnixPath("package_dir"),
		HashOfFiles:          "hash_of_files",
		ExternalDepsHash:     "external_deps_hash",
		Task:                 "task",
		Outputs: TaskOutputs{
			Inclusions: []string{"inclusions"},
			Exclusions: []string{"exclusions"},
		},
		PassThruArgs:    []string{"pass_thru_args"},
		Env:             []string{"env"},
		ResolvedEnvVars: env.EnvironmentVariablePairs{},
		PassThroughEnv:  []string{"pass_thru_env"},
		EnvMode:         util.Infer,
		DotEnv:          []turbopath.AnchoredUnixPath{"dotenv"},
	}

	hash, err := HashTaskHashable(&taskHashable)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "5b222af1dea5828e", hash)
}

func Test_CapnpHashGlobal(t *testing.T) {
	globalHashable := GlobalHashable{
		GlobalCacheKey: "global_cache_key",
		GlobalFileHashMap: map[turbopath.AnchoredUnixPath]string{
			turbopath.AnchoredUnixPath("global_file_hash_map"): "global_file_hash_map",
		},
		RootExternalDepsHash: "root_external_deps_hash",
		Env:                  []string{"env"},
		ResolvedEnvVars:      env.EnvironmentVariablePairs{},
		PassThroughEnv:       []string{"pass_through_env"},
		EnvMode:              util.Infer,
		FrameworkInference:   true,

		DotEnv: []turbopath.AnchoredUnixPath{"dotenv"},
	}

	hash, err := HashGlobalHashable(&globalHashable)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "afa6b9c8d52c2642", hash)
}

func Test_CapnpLockfilePackages(t *testing.T) {
	packages := []lockfile.Package{}

	hash, err := HashLockfilePackages(packages)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "3cbace99d7f9f070", hash)
}

func Test_CapnpLockfilePackages2(t *testing.T) {
	packages := []lockfile.Package{
		{
			Key:     "key",
			Version: "version",
			Found:   false,
		},
	}

	hash, err := HashLockfilePackages(packages)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "ae101a620fb8d207", hash)
}

func Test_CapnpLockfilePackages_InOrder(t *testing.T) {
	packages := []lockfile.Package{
		{
			Key:     "key",
			Version: "version",
			Found:   false,
		},
		{
			Key:     "zey",
			Version: "version",
			Found:   false,
		},
	}

	hash, err := HashLockfilePackages(packages)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "e1c49e53fdbeb38a", hash)
}

func Test_CapnpLockfilePackages_OutOfOrder(t *testing.T) {
	packages := []lockfile.Package{
		{
			Key:     "zey",
			Version: "version",
			Found:   false,
		},
		{
			Key:     "key",
			Version: "version",
			Found:   false,
		},
	}

	hash, err := HashLockfilePackages(packages)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "a9da37ee949583bd", hash)
}

func Test_CapnpFileHashes_Empty(t *testing.T) {
	fileHashes := map[turbopath.AnchoredUnixPath]string{}

	hash, err := HashFileHashes(fileHashes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "a6dd8f3ed2853e94", hash)
}

func Test_CapnpFileHashes_NonEmpty(t *testing.T) {
	fileHashes := map[turbopath.AnchoredUnixPath]string{
		turbopath.AnchoredUnixPath("a"): "b",
		turbopath.AnchoredUnixPath("c"): "d",
	}

	hash, err := HashFileHashes(fileHashes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "0f75c29eda3ab994", hash)
}

func Test_CapnpFileHashes_OrderResistant(t *testing.T) {
	fileHashes := map[turbopath.AnchoredUnixPath]string{
		turbopath.AnchoredUnixPath("c"): "d",
		turbopath.AnchoredUnixPath("a"): "b",
	}

	hash, err := HashFileHashes(fileHashes)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "0f75c29eda3ab994", hash)
}
