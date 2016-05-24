package event

type Flag int

const (
	Modified Flag = 1 << iota
	Deleted
	Truncated
	Reopened
	Dying

	All = Modified | Deleted | Truncated | Reopened | Dying
)

func (e Flag) String() string {
	switch e {
	case Modified:
		return "Modified"
	case Deleted:
		return "Moved/Deleted"
	case Truncated:
		return "Truncated"
	case Reopened:
		return "Re-opened"
	case Dying:
		return "Dying"
	}
	return "Unknown"
}

// Has return that a given flag 'f' is set on 'e'.
func (e Flag) Has(f Flag) bool {
	return e&f > 0
}

type Event struct {
	Filename string
	Flag     Flag
}

type Notification struct {
	sendto chan<- Event
	flag   Flag
}

func NewNotification(to chan<- Event, e Flag) *Notification {
	return &Notification{sendto: to, flag: e}
}

func Notify(oc Flag, fp string, dests []*Notification) {
	for _, no := range dests {
		if no.sendto == nil {
			continue
		}
		if no.flag.Has(oc) {
			continue
		}
		select {
		case no.sendto <- Event{Filename: fp, Flag: ev}:
		default:
		}
	}
}
