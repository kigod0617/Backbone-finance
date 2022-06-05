// Code generated by "callbackgen -type ActiveOrderBook"; DO NOT EDIT.

package bbgo

import (
	"github.com/c9s/bbgo/pkg/types"
)

func (b *ActiveOrderBook) OnFilled(cb func(o types.Order)) {
	b.filledCallbacks = append(b.filledCallbacks, cb)
}

func (b *ActiveOrderBook) EmitFilled(o types.Order) {
	for _, cb := range b.filledCallbacks {
		cb(o)
	}
}
