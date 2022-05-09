// Code generated by "callbackgen -type Pivot"; DO NOT EDIT.

package pivot

import ()

func (inc *Pivot) OnUpdate(cb func(valueLow float64, valueHigh float64)) {
	inc.UpdateCallbacks = append(inc.UpdateCallbacks, cb)
}

func (inc *Pivot) EmitUpdate(valueLow float64, valueHigh float64) {
	for _, cb := range inc.UpdateCallbacks {
		cb(valueLow, valueHigh)
	}
}
