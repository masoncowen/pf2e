package data

import (
    "os"
    "encoding/json"
    "internal/constants"
)

func LoadCampaign(campaignPath string) (constants.Campaign, error) {
    campaign := constants.Campaign{}
    jsonData, err := os.ReadFile(campaignPath)
    if err != nil {
        return campaign, err
    }

    err = json.Unmarshal([]byte(jsonData), &campaign)
    if err != nil {
        return campaign, err
    }
    campaign.Path = campaignPath
    return campaign, nil
}
