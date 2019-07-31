package gameobject

import (
	"fmt"
	"time"
)

type CodeListComponent struct {
	*CustomComponent
}

func NewCodeListComponent() *CodeListComponent {
	cl := &CodeListComponent{&CustomComponent{ct: CodeListComponentType}}
	cl.CustomComponent.value = make([]CodeComponent, 0)
	return cl
}

func (cl *CodeListComponent) Add(code CodeComponent) {
	sl := cl.value.([]CodeComponent)
	cl.value = append(sl, code)
}

func (cl *CodeListComponent) Remove(code CodeComponent) error {
	sl := cl.value.([]CodeComponent)
	idx := -1
	for i, cc := range sl {
		if cc == code {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("CodeListComponent doesn't have Code sent for removal")
	}
	sl[idx] = sl[len(sl)-1]
	sl = sl[:len(sl)-1]
	cl.value = sl
	return nil
}

func (cl *CodeListComponent) Update(obj *GameObject, elapsed time.Duration) {
	sl := cl.value.([]CodeComponent)
	for _, cc := range sl {
		if cc.Initialized() {
			cc.Update(obj, elapsed)
			return
		}
		cc.Init()
	}
}
