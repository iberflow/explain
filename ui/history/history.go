package history

const OptionSelect = "nav"
const CommandChange = "cmd"
const CursorMove = "cur"

type Item struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func NewItem(t string, v interface{}) *Item {
	return &Item{
		Type:  t,
		Value: v,
	}
}

type listenFunc func(i *Item)

type History struct {
	items            []*Item
	index            int
	listenFunc       listenFunc
	disableEventOnce bool
}

func NewHistory() *History {
	return &History{
		items: make([]*Item, 0),
		index: 0,
	}
}

func (h *History) DisableNextEvent() {
	h.disableEventOnce = true
}

func (h *History) Add(t string, value interface{}) {
	// some events
	if h.disableEventOnce {
		h.disableEventOnce = false
		return
	}

	item := NewItem(t, value)

	// if current index is lower than the number of items
	// we traveled back in history
	// all future history items need to be replaced with this item
	if h.index+1 < len(h.items) {
		h.items = append(h.items[0:h.index+1], item)
	} else {
		h.items = append(h.items, item)
	}

	h.index++

	h.dispatch(item)
}

func (h *History) Index() int {
	return h.index
}

func (h *History) Count() int {
	return len(h.items)
}

func (h *History) GoForward() {
	if h.index < len(h.items)-1 {
		h.index++
		h.dispatch(NewItem(CursorMove, h.index))
	}
}

func (h *History) GoBack() {
	if h.index > 0 {
		// before navigation, after a new item is added to history
		// we get index = len(items)
		// so going back just selects the last item
		// to fix it go back 2 steps
		if h.index == len(h.items) {
			h.index = h.index - 2
		} else {
			h.index--
		}
		h.dispatch(NewItem(CursorMove, h.index))
	}
}

func (h *History) GetItem() *Item {
	if len(h.items) > 0 && h.index < len(h.items) {
		return h.items[h.index]
	}

	return nil
}

func (h *History) SetListenFunc(fn listenFunc) {
	h.listenFunc = fn
}

func (h *History) dispatch(item *Item) {
	if h.listenFunc != nil {
		h.listenFunc(item)
	}
}
