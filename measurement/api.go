package measurement

import (
	"time"
)

const (
	ENDPOINT string = "www.google-analytics.com/collect"
)

// Tracking ID
type Tracker struct {
	Name    string // Name of the tracker e.g. rollup, main
	TrackID string `json:"tid"` // Tracking ID e.g. UA-XXXXXXX
}

// Base structure holding basic information used by subsequent calls
type UserData struct {
	Protov    int    `json:"v"`   // Protocol version
	ClientID  string `json:"cid"` // Client ID
	UserID    string `json""`     // User ID e.g. SSO ID
	UserIP    string `json:"uip"` // User IP override
	UserAgent string `json:"ua"`  // End-user's user-agent override
	DataID    int    `json:"z"`   // Cachebuster uid to get around caching
	Delay     int    `json:"qt"`  // Delay in transmission (if sitting in queue)
}

type PageviewMessage struct {
	// 'pageview', 'screenview', 'event', 'transaction', 'item', 'social', 'exception', 'timing'.
	Type     string          `json:"t"`  // t=pageview
	Hostname string          `json:"dh"` // Hostname.
	Path     string          `json:"dp"` // Page path
	Title    string          `json:"dt"` // Page title
	LinkID   `json:"linkid"` // Link ID used to diff between multiple links
}

func (m *PageviewMessage) Flatten() (output string, ok bool) {

	return
}

type EventMessage struct {
	Type     string `json:"t"`  // t=event
	Category string `json:"ec"` // Event Category. Required.
	Action   string `json:"ea"` // Event Action. Required.
	Label    string `json:"el"` // Event label.
	Value    string `json:"ev"` // Event value.
}

type UserSession struct {
	Base         *UserData
	Started      time.Time
	MessageCount int
	Source       string
	Medium       string
	Campaign     string
	Trackers     []*Tracker // list of trackers to transmit to
}

func (s *UserSession) Transmit(message string) (ok bool) {
	return
}

// Send pageview
func (s *UserSession) Pageview(m *PageviewMessage) (ok bool) {
	message, ok := m.Flatten()
	if !ok {
		return
	}
	if sent := s.Transmit(message); sent {
		if s.MessageCount == 0 {
			s.Started = time.Now()
		}
		s.MessageCount += 1
		ok = true
	}
	return
}
