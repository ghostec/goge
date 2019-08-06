package event

type Subscriber interface{}

type Dispatcher struct {
	toDispatch    []*Event
	subscriptions map[Key]map[Subscriber]func(*Event)
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		toDispatch:    make([]*Event, 0),
		subscriptions: map[Key]map[Subscriber]func(*Event){},
	}
}

func (d *Dispatcher) Subscribe(s Subscriber, k Key, onEvent func(*Event)) {
	if d.subscriptions[k] == nil {
		d.subscriptions[k] = map[Subscriber]func(*Event){}
	}
	d.subscriptions[k][s] = onEvent
}

func (d *Dispatcher) Unsubscribe(s Subscriber, k Key) {
	if d.subscriptions[k] == nil || d.subscriptions[k][s] == nil {
		return
	}
	delete(d.subscriptions[k], s)
}

func (d *Dispatcher) Dispatch(e *Event) {
	d.toDispatch = append(d.toDispatch, e)
}

func (d *Dispatcher) Process() {
	for _, e := range d.toDispatch {
		for _, onEvent := range d.subscriptions[e.key] {
			onEvent(e)
		}
	}
	d.toDispatch = make([]*Event, 0)
}
