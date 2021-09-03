package user

type UserPermission struct {
	UserUUID  string  `json:"userUUID"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Channel   Channel `json:"channel"`
	Role      Role    `json:"role"`
}

type Channel struct {
	ChannelUUID string `json:"channelUUID"`
	ChannelName string `json:"channelName"`
	Store       Store  `json:"store"`
}

type Store struct {
	StoreUUID string `json:"storeUUID"`
	StoreName string `json:"storeName"`
}

type Role struct {
	RoleUUID    string `json:"roleUUID"`
	RoleName    string `json:"roleName"`
	Permissions []Pms  `json:"permissions"`
}

type Pms struct {
	MenuName string `json:"menuName,omitempty"`
	Scope    Scope  `json:"scope,omitempty"`
}

type Scope struct {
	CanView   bool `json:"view"`
	CanCreate bool `json:"create"`
	CanUpdate bool `json:"update"`
	CanDelete bool `json:"delete"`
	CanExport bool `json:"export"`
}
