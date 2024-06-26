package friend

import "errors"

func (x *GetPaginationFriendsReq) Check() error {
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}

func (x *ApplyToAddFriendReq) Check() error {
	if x.ToUserID == "" {
		return errors.New("toUserID is empty")
	}
	if x.FromUserID == "" {
		return errors.New("fromUserID is empty")
	}
	return nil
}

func (x *ImportFriendReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserIDs == nil {
		return errors.New("friendUserIDS is empty")
	}
	return nil
}

func (x *GetPaginationFriendsApplyToReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	return nil
}

func (x *GetDesignatedFriendsReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserIDs == nil {
		return errors.New("friendUserIDS is empty")
	}
	return nil
}

func (x *AddBlackReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.BlackUserID == "" {
		return errors.New("blackUserID is empty")
	}
	return nil
}

func (x *RemoveBlackReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.BlackUserID == "" {
		return errors.New("blackUserID is empty")
	}
	return nil
}

func (x *GetPaginationBlacksReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	return nil
}

func (x *IsFriendReq) Check() error {
	if x.UserID1 == "" {
		return errors.New("userID1 is empty")
	}
	if x.UserID2 == "" {
		return errors.New("userID2 is empty")
	}
	return nil
}

func (x *IsBlackReq) Check() error {
	if x.UserID1 == "" {
		return errors.New("userID1 is empty")
	}
	if x.UserID2 == "" {
		return errors.New("userID2 is empty")
	}
	return nil
}

func (x *DeleteFriendReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("userID1 is empty")
	}
	if x.FriendUserID == "" {
		return errors.New("userID2 is empty")
	}
	return nil
}

func (x *RespondFriendApplyReq) Check() error {
	if x.ToUserID == "" {
		return errors.New("toUserID is empty")
	}
	if x.FromUserID == "" {
		return errors.New("fromUserID is empty")
	}
	return nil
}

func (x *SetFriendRemarkReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserID == "" {
		return errors.New("fromUserID is empty")
	}
	return nil
}

func (x *GetPaginationFriendsApplyFromReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	if x.Pagination == nil {
		return errors.New("pagination is empty")
	}
	if x.Pagination.PageNumber < 1 {
		return errors.New("pageNumber is invalid")
	}
	return nil
}

func (x *GetFriendIDsReq) Check() error {
	if x.UserID == "" {
		return errors.New("userID is empty")
	}
	return nil
}
func (x *GetDesignatedFriendsApplyReq) Check() error {
	if x.FromUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.ToUserID == "" {
		return errors.New("toUserID is empty")
	}
	return nil
}
func (x *UpdateFriendsReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.FriendUserIDs == nil {
		return errors.New("friendUserIDs is empty")
	}
	return nil
}
func (x *GetSpecifiedFriendsInfoReq) Check() error {
	if x.OwnerUserID == "" {
		return errors.New("ownerUserID is empty")
	}
	if x.UserIDList == nil {
		return errors.New("userIDList is empty")
	}
	return nil
}