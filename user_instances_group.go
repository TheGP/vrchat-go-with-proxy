package vrchat

// UserInstancesGroupsResponseFull mirrors the full JSON response for /users/{userId}/instances/groups
// Only the most relevant fields are included for now; expand as needed.
type UserInstancesGroupsResponseFull struct {
	FetchedAt string                  `json:"fetchedAt"`
	Instances []UserGroupInstanceFull `json:"instances"`
}

type UserGroupInstanceFull struct {
	Active                   bool               `json:"active"`
	AgeGate                  bool               `json:"ageGate"`
	CanRequestInvite         bool               `json:"canRequestInvite"`
	Capacity                 int                `json:"capacity"`
	ClientNumber             string             `json:"clientNumber"`
	DisplayName              string             `json:"displayName"`
	Full                     bool               `json:"full"`
	GameServerVersion        int                `json:"gameServerVersion"`
	ID                       string             `json:"id"`
	InstanceId               InstanceId         `json:"instanceId"`
	Location                 string             `json:"location"`
	NUsers                   int                `json:"n_users"`
	Name                     string             `json:"name"`
	OwnerId                  UserId             `json:"ownerId"`
	Permanent                bool               `json:"permanent"`
	PhotonRegion             string             `json:"photonRegion"`
	PlayerPersistenceEnabled bool               `json:"playerPersistenceEnabled"`
	Region                   InstanceRegion     `json:"region"`
	SecureName               string             `json:"secureName"`
	ShortName                string             `json:"shortName"`
	Tags                     []string           `json:"tags"`
	Type                     string             `json:"type"`
	WorldId                  WorldId            `json:"worldId"`
	Hidden                   string             `json:"hidden"`
	Friends                  string             `json:"friends"`
	Private                  string             `json:"private"`
	QueueEnabled             bool               `json:"queueEnabled"`
	QueueSize                int                `json:"queueSize"`
	RecommendedCapacity      int                `json:"recommendedCapacity"`
	RoleRestricted           bool               `json:"roleRestricted"`
	Strict                   bool               `json:"strict"`
	UserCount                int                `json:"userCount"`
	GroupAccessType          GroupAccessType    `json:"groupAccessType"`
	HasCapacityForYou        bool               `json:"hasCapacityForYou"`
	Nonce                    string             `json:"nonce"`
	ClosedAt                 string             `json:"closedAt"`
	HardClose                bool               `json:"hardClose"`
	World                    UserInstanceWorld  `json:"world"`
	Users                    []UserInstanceUser `json:"users"`
}

type UserInstanceWorld struct {
	AuthorId               UserId                     `json:"authorId"`
	AuthorName             string                     `json:"authorName"`
	Capacity               int                        `json:"capacity"`
	RecommendedCapacity    int                        `json:"recommendedCapacity"`
	CreatedAt              string                     `json:"created_at"`
	DefaultContentSettings map[string]bool            `json:"defaultContentSettings"`
	Description            string                     `json:"description"`
	Favorites              int                        `json:"favorites"`
	Featured               bool                       `json:"featured"`
	Heat                   int                        `json:"heat"`
	ID                     string                     `json:"id"`
	ImageUrl               string                     `json:"imageUrl"`
	Instances              [][]interface{}            `json:"instances"`
	LabsPublicationDate    string                     `json:"labsPublicationDate"`
	Name                   string                     `json:"name"`
	Namespace              string                     `json:"namespace"`
	Occupants              int                        `json:"occupants"`
	Organization           string                     `json:"organization"`
	Popularity             int                        `json:"popularity"`
	PreviewYoutubeId       string                     `json:"previewYoutubeId"`
	PrivateOccupants       int                        `json:"privateOccupants"`
	PublicOccupants        int                        `json:"publicOccupants"`
	PublicationDate        string                     `json:"publicationDate"`
	ReleaseStatus          string                     `json:"releaseStatus"`
	StoreId                string                     `json:"storeId"`
	Tags                   []string                   `json:"tags"`
	ThumbnailImageUrl      string                     `json:"thumbnailImageUrl"`
	UnityPackages          []UserInstanceUnityPackage `json:"unityPackages"`
	UpdatedAt              string                     `json:"updated_at"`
	UrlList                []string                   `json:"urlList"`
	Version                int                        `json:"version"`
	Visits                 int                        `json:"visits"`
	UdonProducts           []string                   `json:"udonProducts"`
}

type UserInstanceUnityPackage struct {
	ID                  string                 `json:"id"`
	AssetUrl            string                 `json:"assetUrl"`
	AssetUrlObject      map[string]interface{} `json:"assetUrlObject"`
	AssetVersion        int                    `json:"assetVersion"`
	CreatedAt           string                 `json:"created_at"`
	ImpostorizerVersion string                 `json:"impostorizerVersion"`
	PerformanceRating   string                 `json:"performanceRating"`
	Platform            string                 `json:"platform"`
	PluginUrl           string                 `json:"pluginUrl"`
	PluginUrlObject     map[string]interface{} `json:"pluginUrlObject"`
	UnitySortNumber     int                    `json:"unitySortNumber"`
	UnityVersion        string                 `json:"unityVersion"`
	WorldSignature      string                 `json:"worldSignature"`
	ImpostorUrl         string                 `json:"impostorUrl"`
	ScanStatus          string                 `json:"scanStatus"`
	Variant             string                 `json:"variant"`
}

type UserInstanceUser struct {
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	CurrentAvatarImageUrl          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
	DeveloperType                  string   `json:"developerType"`
	DisplayName                    string   `json:"displayName"`
	FallbackAvatar                 string   `json:"fallbackAvatar"`
	ID                             string   `json:"id"`
	IsFriend                       bool     `json:"isFriend"`
	LastPlatform                   string   `json:"last_platform"`
	LastLogin                      string   `json:"last_login"`
	ProfilePicOverride             string   `json:"profilePicOverride"`
	Pronouns                       string   `json:"pronouns"`
	Status                         string   `json:"status"`
	StatusDescription              string   `json:"statusDescription"`
	Tags                           []string `json:"tags"`
	UserIcon                       string   `json:"userIcon"`
	Username                       string   `json:"username"`
	Location                       string   `json:"location"`
	FriendKey                      string   `json:"friendKey"`
}
