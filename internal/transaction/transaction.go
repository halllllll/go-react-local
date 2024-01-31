package transaction

import (
	"context"
)

type Transaction interface {
	DoTx(context.Context, func(context.Context) error) error
}
