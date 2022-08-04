package campay

// Environment is the URL of the campay environment
type Environment string

var (
	// DevEnvironment is the development Environment
	DevEnvironment = Environment("https://demo.campay.net")

	// ProdEnvironment is the production Environment
	ProdEnvironment = Environment("https://www.campay.net")
)

func (e Environment) String() string {
	return string(e)
}
