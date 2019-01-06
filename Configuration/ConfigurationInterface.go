package Configuration

import (
	"go/types"
)

type ConfigurationInterface interface {
	getVersion() string
	getSteps() types.Array
	getCaches() types.Array
}
