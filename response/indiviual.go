package response

type UpdateUserInforesp struct {
	Msg     string `json:"msg"`
	YNLogin bool   `json:"YNLogin"`
}
type UpdateUserAvatarresp struct {
	Msg     string `json:"msg"`
	YNLogin bool   `json:"YNLogin"`
	Success bool   `json:"success"`
}
