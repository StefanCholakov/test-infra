package incubator

import (
	"path"
	"testing"

	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/jobsuite"
)

const (
	jobBasePath = "./../../../../prow/jobs/incubator"
)

var components = []struct {
	name              string
	image             string
	suite             func(config *jobsuite.Config) jobsuite.Suite
	additionalOptions []jobsuite.Option
}{
	{
		name:  "connector",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "director",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "gateway",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "schema-migrator",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "connectivity-adapter",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
		},
	},
	{
		name:  "pairing-adapter",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "external-services-mock",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "system-broker",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "compass-console",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			func(suite *jobsuite.Config) {
				suite.Path = "compass-console/compass"
			},
			jobsuite.JobFileSuffix("ui"),
			jobsuite.CompassConsoleRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "ord-service",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			func(suite *jobsuite.Config) {
				suite.Path = "ord-service/components/ord-service"
			},
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassORDServiceRepo(),
			jobsuite.AllReleases(),
		},
	},
	{
		name:  "operations-controller",
		image: tester.ImageGolangBuildpack1_16,
		suite: tester.NewGenericComponentSuite,
		additionalOptions: []jobsuite.Option{
			jobsuite.JobFileSuffix("generic"),
			jobsuite.CompassRepo(),
			jobsuite.AllReleases(),
		},
	},
}

func TestComponentJobs(t *testing.T) {
	testedConfigurations := make(map[string]struct{})
	repos := map[string]struct{}{}
	for _, component := range components {
		t.Run(component.name, func(t *testing.T) {
			opts := []jobsuite.Option{
				jobsuite.CompassComponent(component.name, component.image),
				jobsuite.KymaRepo(),
				jobsuite.AllReleases(),
			}
			opts = append(opts, component.additionalOptions...)
			cfg := jobsuite.NewConfig(opts...)
			suite := component.suite
			if suite == nil {
				suite = tester.NewComponentSuite
			}
			ts := suite(cfg)
			if pathProvider, ok := ts.(jobsuite.JobConfigPathProvider); ok {
				testedConfigurations[path.Clean(pathProvider.JobConfigPath())] = struct{}{}
			}
			repos[cfg.Repository] = struct{}{}
			ts.Run(t)
		})
	}
	t.Run("All Files covered by test", jobsuite.CheckFilesAreTested(repos, testedConfigurations, jobBasePath, []string{"components", "compass"}))
}
