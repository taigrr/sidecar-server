package types

type TabSet struct {
	Tabs []Tab `json:"tabs"`
}
type MutedInfo struct {
	IsMuted bool `json:"muted"`
}

type Tab struct {
	IsActive          bool      `json:"active"`
	IsAudible         bool      `json:"audible"`
	IsAutoDiscardable bool      `json:"autoDiscardable"`
	IsDiscarded       bool      `json:"discarded"`
	FaviconURL        string    `json:"favIconUrl"`
	GroupID           int       `json:"groupId"`
	Height            int       `json:"height"`
	IsHighlighted     bool      `json:"highlighted"`
	ID                int       `json:"id"`
	IsIncognito       bool      `json:"incognito"`
	Index             int       `json:"index"`
	MutedInfo         MutedInfo `json:"mutedInfo"`
	IsPinned          bool      `json:"pinned"`
	IsSelected        bool      `json:"selected"`
	LoadedStatus      string    `json:"status"`
	Title             string    `json:"title"`
	URL               string    `json:"url"`
	Width             int       `json:"width"`
	WindowID          int       `json:"windowId"`
}

type Actions struct {
	Close []int `json:"close"`
}
