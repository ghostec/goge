package gameobject

import (
	"github.com/ghostec/goge/store"
	"github.com/google/uuid"
)

const CodeListComponentType = ComponentType("gameobject.component.CodeList")

type CodeListComponent struct {
	*CustomComponent
	codeStore *store.Store
}

func NewCodeListComponent() *CodeListComponent {
	cl := &CodeListComponent{
		CustomComponent: &CustomComponent{ct: CodeListComponentType},
		codeStore:       store.New(),
	}
	return cl
}

func (cl *CodeListComponent) Add(code CodeComponent) CodeComponentKey {
	uuid, _ := uuid.NewRandom()
	key := CodeComponentKey(uuid.String())
	cl.codeStore.Set(store.Key(key), code)
	return key
}

func (cl *CodeListComponent) Remove(key CodeComponentKey) {
	cl.codeStore.Unset(store.Key(key))
}

func (cl *CodeListComponent) Update(ctx *Context) {
	for _, c := range cl.codeStore.All() {
		cc := c.(CodeComponent)
		if cc.Initialized() {
			cc.Update(ctx)
			return
		}
		cc.Init(ctx)
	}
}
