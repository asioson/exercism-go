// react implements routines of a basic reactive system
package react

type reactor struct {
	rList []*cell
}

type cell struct {
	isInput   bool
	value     int
	prevValue int
	compute   func() int
	cbMap     map[int]func(int)
	react     *reactor
}

type canceler struct {
	id            int
	whereToCancel *cell
}

// Cancel removes a callback function.
func (c *canceler) Cancel() {
	delete(c.whereToCancel.cbMap, c.id)
}

// Value returns the current value of the cell.
func (c *cell) Value() int {
	if c.isInput {
		return c.value
	} else {
		return c.compute()
	}
}

// SetValue sets the value of the cell.
func (c *cell) SetValue(value int) {
	c.value = value
	for _, ccell := range c.react.rList {
		if len(ccell.cbMap) > 0 && ccell.Value() != ccell.prevValue {
			for _, cbfunc := range ccell.cbMap {
				cbfunc(ccell.Value())
			}
			ccell.prevValue = ccell.Value()
		}
	}
}

// AddCallback adds a callback which will be called when the value changes and
// returns a Canceler which can be used to remove the callback.
func (c *cell) AddCallback(callback func(int)) Canceler {
	c.cbMap[len(c.cbMap)] = callback
	return &canceler{id: len(c.cbMap) - 1, whereToCancel: c}
}

// New creates a new Reactor
func New() Reactor {
	return &reactor{rList: make([]*cell, 0)}
}

// CreateInput creates an input cell linked into the reactor with the given initial value.
func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{isInput: true, value: initial, react: r}
}

// CreateCompute1 creates a compute cell which computes its value based on
// the dep cell. The compute function will only be called if the value of
// the passed cell changes.
func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	var c cell
	c.compute = func() int { return compute(dep.Value()) }
	c.prevValue = c.compute()
	c.cbMap = make(map[int]func(int))
	r.rList = append(r.rList, &c)
	return &c
}

// CreateCompute2 creates a compute cell which computes its value based on
// the dep1 and dep2 cells. The compute function will only be called if the value of
// the passed cell changes.
func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	var c cell
	c.compute = func() int { return compute(dep1.Value(), dep2.Value()) }
	c.prevValue = c.compute()
	c.cbMap = make(map[int]func(int))
	r.rList = append(r.rList, &c)
	return &c
}
