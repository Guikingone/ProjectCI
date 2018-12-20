package Build

import (
	"go/types"
	"time"
)

type BuildInterface interface {
	getConfiguration()
	getStartDate() time.Time
	getEndDate() time.Time
	getTags() types.Array
	isValid() bool
}
