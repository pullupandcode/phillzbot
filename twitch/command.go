package twitch

type TwitchCommandRole int

const (
	Broadcaster TwitchCommandRole = iota
	Moderator
	VIP
	Subscribers
	Follower
	Viewer
)

func (role TwitchCommandRole) String() string {
	return [...]string{"BROADCASTER", "MODERATOR", "VIP", "SUBSCRIBERS", "FOLLOWERS", "VIEWERS"}[role]
}

type TwitchCommand struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Role        TwitchCommandRole `json:"role"`
	Value       string            `json:"value"`
	Cooldown    int               `json:"cooldown"`
}
