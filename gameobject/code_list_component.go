package gameobject

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CodeListComponent struct {
	*CustomComponent
}

func NewCodeListComponent() *CodeListComponent {
	cl := &CodeListComponent{&CustomComponent{ct: CodeListComponentType}}
	cl.CustomComponent.value = make([]*CodeListComponentFunc, 0)
	return cl
}

func (cl *CodeListComponent) Add(f *CodeListComponentFunc) {
	sl := cl.value.([]*CodeListComponentFunc)
	cl.value = append(sl, f)
}

func (cl *CodeListComponent) Remove(f *CodeListComponentFunc) error {
	sl := cl.value.([]*CodeListComponentFunc)
	idx := -1
	for i, ff := range sl {
		if ff == f {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("CodeListComponent doesn't have CodeListComponentFunc sent for removal")
	}
	sl[idx] = sl[len(sl)-1]
	sl = sl[:len(sl)-1]
	cl.value = sl
	return nil
}

func (cl *CodeListComponent) Update(elapsed time.Duration) {
	sl := cl.value.([]*CodeListComponentFunc)
	for _, f := range sl {
		f.f(cl.obj, elapsed)
	}
}

type CodeListComponentFunc struct {
	uuid uuid.UUID
	name string
	f    func(obj *GameObject, elapsed time.Duration) error
}

func NewCodeListComponentFunc(
	name string, f func(*GameObject, time.Duration) error,
) *CodeListComponentFunc {
	uuid, _ := uuid.NewRandom()
	return &CodeListComponentFunc{
		uuid: uuid,
		name: name,
		f:    f,
	}
}
