package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type SteamUser struct {
	*Base
}

type GetPlayerSummariesResponse struct {
	Response struct {
		Players []struct {
			SteamID                  string
			CommunityVisibilityState int
			ProfileState             int
			PersonAName              string
			CommentPermission        int
			ProfileURL               string
			Avatar                   string
			AvatarMedium             string
			AvatarFull               string
			LastLogOff               int
			TimeCreated              int
			PersonAStateFlags        int
			LocCountryCode           string
			LocStateCode             string
			LocCityID                int
			RealName                 string
			PrimaryClanID            string
			GameID                   int
			GameServerID             string
			GameExtraInfo            string
		}
	}
}

func (i *SteamUser) GetPlayerSummaries(steamIDs []string) (*GetPlayerSummariesResponse, error) {
	resp, err := i.httpClient.Get(
		fmt.Sprintf(
			"%s/ISteamUser/GetPlayerSummaries/v2/?key=%s&steamids=%s",
			i.steamAPIEndpoint,
			i.apiKey,
			strings.Join(steamIDs, ","),
		),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var playSummaryResp GetPlayerSummariesResponse
	if err := json.Unmarshal(body, &playSummaryResp); err != nil {
		return nil, err
	}

	return &playSummaryResp, err
}

type GetFriendListResponse struct {
	FriendsList struct {
		Friends []struct {
			SteamID      string
			Relationship string
			FriendSince  int `json:"friend_since"`
		}
	}
}

func (i *SteamUser) GetFriendList(steamID string) (*GetFriendListResponse, error) {
	resp, err := i.httpClient.Get(
		fmt.Sprintf(
			"%s/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend",
			i.steamAPIEndpoint,
			i.apiKey,
			steamID,
		),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getFriendListResp GetFriendListResponse
	if err := json.Unmarshal(body, &getFriendListResp); err != nil {
		return nil, err
	}

	return &getFriendListResp, err
}

type GetPlayerBansResponse struct {
	Players []struct {
		SteamID          string
		CommunityBanned  bool
		VACBanned        bool
		NumberOfVACBans  int
		DaysSinceLastBan int
		NumberOfGameBans int
		EconomyBan       string
	}
}

func (i *SteamUser) GetPlayerBans(steamIDs []string) (*GetPlayerBansResponse, error) {
	resp, err := i.httpClient.Get(
		fmt.Sprintf(
			"%s/ISteamUser/GetPlayerBans/v1/?key=%s&steamids=%s",
			i.steamAPIEndpoint,
			i.apiKey,
			strings.Join(steamIDs, ","),
		),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getPlayerBansResp GetPlayerBansResponse
	if err := json.Unmarshal(body, &getPlayerBansResp); err != nil {
		return nil, err
	}

	return &getPlayerBansResp, err
}

type GetUserGroupListResponse struct {
	Response struct {
		Groups []struct {
			GID string
		}
	}
}

func (i *SteamUser) GetUserGroupList(steamID string) (*GetUserGroupListResponse, error) {
	resp, err := i.httpClient.Get(
		fmt.Sprintf(
			"%s/ISteamUser/GetUserGroupList/v1/?key=%s&steamid=%s",
			i.steamAPIEndpoint,
			i.apiKey,
			steamID,
		),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getUserGroupListResp GetUserGroupListResponse
	if err := json.Unmarshal(body, &getUserGroupListResp); err != nil {
		return nil, err
	}

	return &getUserGroupListResp, err
}

type ResolveVanityURLResponse struct {
	Response struct {
		SteamID string
	}
}

type resolveVanityURLOptions struct {
	urlType int32
}
type ResolveVanityURLOpt func(opts *resolveVanityURLOptions)

func ResolveVanityURLWithURLType(urlType int32) ResolveVanityURLOpt {
	return func(opt *resolveVanityURLOptions) {
		opt.urlType = urlType
	}
}

func (i *SteamUser) ResolveVanityURL(vanityURL string, opts ...ResolveVanityURLOpt) (*ResolveVanityURLResponse, error) {
	optionals := resolveVanityURLOptions{urlType: 1}
	for _, opt := range opts {
		opt(&optionals)
	}

	resp, err := i.httpClient.Get(
		fmt.Sprintf(
			"%s/ISteamUser/ResolveVanityURL/v1/?key=%s&vanityurl=%s&url_type=%d",
			i.steamAPIEndpoint,
			i.apiKey,
			vanityURL,
			optionals.urlType,
		),
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return nil, nil
}
