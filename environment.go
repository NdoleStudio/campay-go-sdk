package campay

// Environment is the URL of the campay environment
type Environment string

var (
	// DevEnvironment is the development Environment
	DevEnvironment = Environment("https://demo.campay.net/api")

	// ProdEnvironment is the production Environment
	ProdEnvironment = Environment("https://www.campay.net/api")
)

func (e Environment) String() string {
	return string(e)
}
