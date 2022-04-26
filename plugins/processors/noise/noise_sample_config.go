//go:generate go run ../../../tools/generate_plugindata/main.go
//go:generate go run ../../../tools/generate_plugindata/main.go --clean
// DON'T EDIT; This file is used as a template by tools/generate_plugindata
package noise

func (p *Noise) SampleConfig() string {
	return `# Adds noise to numerical fields
[[processors.noise]]
  ## Specified the type of the random distribution.
  ## Can be "laplacian", "gaussian" or "uniform".
  # type = "laplacian

  ## Center of the distribution.
  ## Only used for Laplacian and Gaussian distributions.
  # mu = 0.0

  ## Scale parameter for the Laplacian or Gaussian distribution
  # scale = 1.0

  ## Upper and lower bound of the Uniform distribution
  # min = -1.0
  # max = 1.0

  ## Apply the noise only to numeric fields matching the filter criteria below.
  ## Excludes takes precedence over includes.
  # include_fields = []
  # exclude_fields = []
`
}
