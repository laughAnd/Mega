package vm

import (
	"../model"
)

type ProfileViewModel struct {
	BaseViewModel
	Posts          []model.Post
	Editable       bool
	IsFollow       bool
	FollowersCount int
	FollowingCount int
	ProfileUser    model.User
	BasePageViewModel
}

type ProfileViewModelOp struct{}

func (ProfileViewModelOp) GetVM(sUser, pUser string, page, limit int) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	posts, total, _ := model.GetPostsByUserIDPageAndLimit(u.ID, page, limit)
	v.ProfileUser = *u
	v.Editable = sUser == pUser
	v.SetBasePageViewModel(total, page, limit)
	if !v.Editable {
		v.IsFollow = u.IsFollowedByUser(sUser)
	}
	v.FollowersCount = u.FollowersCount()
	v.FollowingCount = u.FollowingCount()

	v.Posts = *posts
	v.SetCurrentUser(sUser)
	return v, err
}

// Follow func : A follow B
func Follow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.Follow(b)
}
func UnFollow(a, b string) error {
	u, err := model.GetUserByUsername(a)
	if err != nil {
		return err
	}
	return u.UnFollow(b)
}

func (ProfileViewModelOp) GetPopupVM(sUser, pUser string) (ProfileViewModel, error) {
	v := ProfileViewModel{}
	v.SetTitle("Profile")
	u, err := model.GetUserByUsername(pUser)
	if err != nil {
		return v, err
	}
	v.ProfileUser = *u
	v.Editable = sUser == pUser
	if !v.Editable {
		v.IsFollow = u.IsFollowedByUser(sUser)
	}
	v.FollowingCount = u.FollowingCount()
	v.FollowersCount = u.FollowersCount()
	v.SetCurrentUser(sUser)

	return v, nil
}