package measurement

import (
	"time"
)

// main lib structure
type Measurement struct {
	Sender      *dispatcher
	UserSession []*UserSession
}

func (m *Measurement) Init() {
	m.Sender = Dispatcher(5)
}

// Tracking ID
type Tracker struct {
	Name    string // Name of the tracker e.g. rollup, main
	TrackID string `param:"tid"` // Tracking ID e.g. UA-XXXXXXX
}

// Base structure holding basic information used by subsequent calls
type UserData struct {
	Protov    int    `param:"v"`            // Protocol version
	ClientID  string `param:"cid"`          // Client ID
	UserID    string `param"uid"`           // User ID e.g. SSO ID
	UserIP    string `param:"uip"`          // User IP override
	UserAgent string `param:"ua"`           // End-user's user-agent override
	DataID    int    `param:"z"`            // Cachebuster uid to get around caching
	Delay     int    `param:"qt,omitempty"` // Delay in transmission (if sitting in queue)
}

// Representation of a user session
type UserSession struct {
	Base         *UserData
	Started      time.Time
	MessageCount int
	Source       string
	Medium       string
	Campaign     string
	Trackers     []*Tracker // list of trackers to transmit to
}

// Pageview struct
type PageviewMessage struct {
	// 'pageview', 'screenview', 'event', 'transaction', 'item', 'social', 'exception', 'timing'.
	Type     string `param:"t"`                // t=pageview
	Hostname string `param:"dh"`               // Hostname.
	Path     string `param:"dp"`               // Page path
	Title    string `param:"dt"`               // Page title
	LinkID   int    `param:"linkid,omitempty"` // Link ID used to diff between multiple links
}

// Event struct
type EventMessage struct {
	Type     string `param:"t"`            // t=event
	Category string `param:"ec"`           // Event Category. Required.
	Action   string `param:"ea"`           // Event Action. Required.
	Label    string `param:"el"`           // Event label.
	Value    string `param:"ev,omitempty"` // Event value.
}

// Custom dimensions / metrics
type CustomDM struct {
	Dimensions map[string]string
	Metrics    map[string]interface{}
}

func (s *UserSession) Transmit(message string) bool {
	// return Dispatcher(message)
	return false
}

// Send pageview
func (s *UserSession) Pageview(m *PageviewMessage) (ok bool) {
	message := Flatten(m).Encode()

	if sent := s.Transmit(message); sent {
		if s.MessageCount == 0 {
			s.Started = time.Now()
		}
		s.MessageCount += 1
		ok = true
	}
	return
}
