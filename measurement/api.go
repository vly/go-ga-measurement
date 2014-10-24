package measurement

/*
Protocol Version	v	v=1	The protocol version. The value should be 1.
Tracking ID	tid	tid=UA-123456-1	The ID that distinguishes to which Google Analytics property to send data.
Client ID	cid	cid=xxxxx	An ID unique to a particular user.
Hit Type	t	t=pageview	The type of interaction collected for a particular user.
*/

const (
	ENDPOINT string = "www.google-analytics.com/collect"
)

// Base structure holding basic information used by subsequent calls
type UserData struct {
	Protov    int    `json:"v"`   // Protocol version
	TrackID   string `json:"tid"` // Tracking ID e.g. UA-XXXXXXX
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

type EventMessage struct {
	Type     string `json:"t"`  // t=event
	Category string `json:"ec"` // Event Category. Required.
	Action   string `json:"ea"` // Event Action. Required.
	Label    string `json:"el"` // Event label.
	Value    string `json:"ev"` // Event value.
}

type UserSession struct {
	Base *UserData
}
