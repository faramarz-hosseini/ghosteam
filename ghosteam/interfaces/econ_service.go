package interfaces

import (
	"encoding/json"
	"fmt"
	"io"
)

type EconService struct {
	*Base
}

type GetTradeHistoryResponse struct {
	Response struct {
		TotalTrades int
		More        bool
		Trades      []struct {
			TradeID        string
			SteamIDOther   string `json:"steamid_other"`
			TimeInit       int    `json:"time_init"`
			Status         int
			AssetsReceived []struct {
				AppID        int
				ContextID    string
				AssetID      string
				Amount       string
				ClassID      string
				InstanceID   string
				NewAssetID   string `json:"new_assetid"`
				NewContextID string `json:"new_contextid"`
			} `json:"assets_received"`
			AssetsGiven []struct {
				AppID        int
				ContextID    string
				AssetID      string
				Amount       string
				ClassID      string
				InstanceID   string
				NewAssetID   string `json:"new_assetid"`
				NewContextID string `json:"new_contextid"`
			} `json:"assets_given"`
		}
		Descriptions []struct {
			AppID           int
			ClassID         string
			InstanceID      string
			Currency        bool
			BackgroundColor string
			IconURL         string `json:"icon_url"`
			IconURLLarge    string `json:"icon_url_large"`
			Descriptions    []struct {
				Type  string
				Value string
				Color string
			}
			Tradable                    bool
			FraudWarnings               []string
			Name                        string
			NameColor                   string `json:"name_color"`
			Type                        string
			MarketName                  string `json:"market_name"`
			MarketHashName              string `json:"market_hash_name"`
			Commodity                   bool
			MarketTradableRestriction   int `json:"market_tradable_restriction"`
			MarketMarketableRestriction int `json:"market_marketable_restriction"`
			Marketable                  bool
		}
	}
}

func (i *EconService) GetTradeHistory(
	maxTrades uint32, startAfterTime uint32,
	startAfterTradeID uint64, navigatingBack bool,
	getDescriptions bool, includeFailed bool,
	includeTotal bool,
) (*GetTradeHistoryResponse, error) {
	resp, err := i.httpClient.Get(
		fmt.Sprintf(
			`%s/IEconService/GetTradeHistory/v1/?key=%s`+
				`&input_json={"max_trades":%d,"start_after_time":%d,`+
				`"start_after_tradeid":%d,"navigating_back":%t,"get_descriptions":%t,`+
				`"include_failed":%t,"include_total":%t}`,
			i.steamAPIEndpoint, i.apiKey,
			maxTrades, startAfterTime,
			startAfterTradeID, navigatingBack,
			getDescriptions, includeFailed,
			includeTotal,
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

	var tradeHistoryResp GetTradeHistoryResponse
	if err := json.Unmarshal(body, &tradeHistoryResp); err != nil {
		return nil, err
	}

	return &tradeHistoryResp, err
}


