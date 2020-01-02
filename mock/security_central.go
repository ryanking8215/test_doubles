package mock

type Closer interface {
	Close()
}

type SecurityCentral struct {
	window Closer
	door   Closer
}

func NewSecurityCentral(w, d Closer) *SecurityCentral {
	return &SecurityCentral{
		window: w,
		door:   d,
	}
}

func (sc *SecurityCentral) SecurityOn() {
	sc.window.Close()
	sc.door.Close()
}
