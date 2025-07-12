package vrchat

import (
	"time"
)

// DetailedInstanceResponse is a detailed version of InstanceResponse with additional fields
type DetailedInstanceResponse struct {
	Active          bool   `json:"active"`
	AgeGate         bool   `json:"ageGate"`
	Capacity        int    `json:"capacity"`
	ClientNumber    string `json:"clientNumber"`
	ContentSettings struct {
		AllowAvatarCopying bool `json:"allowAvatarCopying"`
		AllowUploads       bool `json:"allowUploads"`
		Platform           struct {
			Android           bool `json:"android"`
			Standalonewindows bool `json:"standalonewindows"`
		} `json:"platform"`
	} `json:"contentSettings"`
	DisplayName              string   `json:"displayName"`
	Id                       string   `json:"id"`
	InstanceId               string   `json:"instanceId"`
	Location                 string   `json:"location"`
	NUsers                   int      `json:"n_users"`
	Name                     string   `json:"name"`
	OwnerId                  string   `json:"ownerId"`
	PhotonRegion             string   `json:"photonRegion"`
	Platforms                []string `json:"platforms"`
	PlayerPersistenceEnabled bool     `json:"playerPersistenceEnabled"`
	Private                  bool     `json:"private"`
	QueueEnabled             bool     `json:"queueEnabled"`
	RecommendedCapacity      int      `json:"recommendedCapacity"`
	RoleRestricted           bool     `json:"roleRestricted"`
	Strict                   bool     `json:"strict"`
	UserCount                int      `json:"userCount"`
	World                    struct {
		AuthorId            string    `json:"authorId"`
		AuthorName          string    `json:"authorName"`
		Capacity            int       `json:"capacity"`
		CreatedAt           time.Time `json:"created_at"`
		Description         string    `json:"description"`
		Favorites           int       `json:"favorites"`
		Featured            bool      `json:"featured"`
		Heat                int       `json:"heat"`
		Id                  string    `json:"id"`
		ImageUrl            string    `json:"imageUrl"`
		LabsPublicationDate string    `json:"labsPublicationDate"`
		Name                string    `json:"name"`
		Namespace           string    `json:"namespace"`
		Occupants           int       `json:"occupants"`
		Organization        string    `json:"organization"`
		Popularity          int       `json:"popularity"`
		PreviewYoutubeId    string    `json:"previewYoutubeId"`
		PublicationDate     string    `json:"publicationDate"`
		ReleaseStatus       string    `json:"releaseStatus"`
		Tags                []string  `json:"tags"`
		ThumbnailImageUrl   string    `json:"thumbnailImageUrl"`
		UnityPackages       []struct {
			AssetUrl       string `json:"assetUrl"`
			AssetUrlObject struct {
				Url string `json:"url"`
			} `json:"assetUrlObject"`
			AssetVersion    int    `json:"assetVersion"`
			CreatedAt       string `json:"created_at"`
			Id              string `json:"id"`
			Platform        string `json:"platform"`
			PluginUrl       string `json:"pluginUrl"`
			PluginUrlObject struct {
				Url string `json:"url"`
			} `json:"pluginUrlObject"`
			UnitySortNumber int    `json:"unitySortNumber"`
			UnityVersion    string `json:"unityVersion"`
		} `json:"unityPackages"`
		UpdatedAt time.Time `json:"updated_at"`
		Version   int       `json:"version"`
		Visits    int       `json:"visits"`
	} `json:"world"`
	Users []struct {
		AllowAvatarCopying             bool     `json:"allowAvatarCopying"`
		Bio                            string   `json:"bio"`
		BioLinks                       []string `json:"bioLinks"`
		CurrentAvatarImageUrl          string   `json:"currentAvatarImageUrl"`
		CurrentAvatarThumbnailImageUrl string   `json:"currentAvatarThumbnailImageUrl"`
		DateJoined                     string   `json:"date_joined"`
		DeveloperType                  string   `json:"developerType"`
		DisplayName                    string   `json:"displayName"`
		FriendKey                      string   `json:"friendKey"`
		Id                             string   `json:"id"`
		InstanceId                     string   `json:"instanceId"`
		IsFriend                       bool     `json:"isFriend"`
		LastActivity                   string   `json:"last_activity"`
		LastLogin                      string   `json:"last_login"`
		LastPlatform                   string   `json:"last_platform"`
		Location                       string   `json:"location"`
		Note                           string   `json:"note"`
		ProfilePicOverride             string   `json:"profilePicOverride"`
		State                          string   `json:"state"`
		Status                         string   `json:"status"`
		StatusDescription              string   `json:"statusDescription"`
		Tags                           []string `json:"tags"`
		TravelingToInstance            string   `json:"travelingToInstance"`
		TravelingToLocation            string   `json:"travelingToLocation"`
		TravelingToWorld               string   `json:"travelingToWorld"`
		UserIcon                       string   `json:"userIcon"`
		WorldId                        string   `json:"worldId"`
	} `json:"users"`
	GroupAccessType   string    `json:"groupAccessType"`
	HasCapacityForYou bool      `json:"hasCapacityForYou"`
	Nonce             string    `json:"nonce"`
	ClosedAt          time.Time `json:"closedAt"`
	HardClose         bool      `json:"hardClose"`
}
