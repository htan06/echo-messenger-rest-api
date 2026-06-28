package user

type UpdateInfoReq struct {
	FirstName     string  `json:"first_name"`
	LastName      *string `json:"last_name,omitempty"`
	AvatarURL     string  `json:"avatar_url"`
	CoverPhotoURL string  `json:"cover_photo_url"`
}

type ChangeReadStatusReq struct {
	Status bool `json:"status"`
}

type UpdateUsernameReq struct {
	Username string `json:"username"`
}

type UserInfoResp struct {
	ID            int64   `json:"id"`
	Username      string  `json:"username"`
	FirstName     string  `json:"first_name"`
	LastName      *string `json:"last_name,omitempty"`
	AvatarURL     *string `json:"avatar_url,omitempty"`
	CoverPhotoURL *string `json:"cover_photo_url,omitempty"`
}
