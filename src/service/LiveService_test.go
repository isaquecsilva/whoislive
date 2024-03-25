package service

import (
	"strings"
	"testing"

	"github.com/isaquecsilva/whoislive/src/domain"
	"github.com/stretchr/testify/assert"
)

const rawData string = `[
    {
        "data":
        {
            "personalSections":
            [
                {
                    "type": "POPULAR_SECTION",
                    "title":
                    {
                        "localizedFallback": "Popular Channels",
                        "localizedTokens":
                        [],
                        "__typename": "PersonalSectionTitle"
                    },
                    "items":
                    [
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "36340781",
                                "login": "tarik",
                                "displayName": "tarik",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/f04d2a14-8d63-4cd5-a469-7ec2cd6e5ce3-profile_image-70x70.png",
                                "primaryColorHex": null,
                                "broadcastSettings":
                                {
                                    "id": "36340781",
                                    "title": "SEN vs GEN - Masters Madrid Grand Final - Twitter @tarik #MastersCostream",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "50712785261",
                                "broadcaster":
                                {
                                    "id": "36340781",
                                    "broadcastSettings":
                                    {
                                        "id": "36340781",
                                        "title": "SEN vs GEN - Masters Madrid Grand Final - Twitter @tarik #MastersCostream",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 218123,
                                "game":
                                {
                                    "id": "516575",
                                    "slug": "valorant",
                                    "displayName": "VALORANT",
                                    "name": "VALORANT",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "490592527",
                                "login": "valorant",
                                "displayName": "VALORANT",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/e0751cdf-bd99-490d-af4f-4fcb6dc2491f-profile_image-70x70.png",
                                "primaryColorHex": "FA2929",
                                "broadcastSettings":
                                {
                                    "id": "490592527",
                                    "title": "GEN vs. SEN — VCT Masters Madrid — Grand Final",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "42127320952",
                                "broadcaster":
                                {
                                    "id": "490592527",
                                    "broadcastSettings":
                                    {
                                        "id": "490592527",
                                        "title": "GEN vs. SEN — VCT Masters Madrid — Grand Final",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 140951,
                                "game":
                                {
                                    "id": "516575",
                                    "slug": "valorant",
                                    "displayName": "VALORANT",
                                    "name": "VALORANT",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "181077473",
                                "login": "gaules",
                                "displayName": "Gaules",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/1f4ff3f0-3ca3-4344-96f5-40993325b886-profile_image-70x70.png",
                                "primaryColorHex": "00A7E1",
                                "broadcastSettings":
                                {
                                    "id": "181077473",
                                    "title": "Pain vs NAVI PGL CS2 Major Copenhagen 2024 - !Sorteio - Siga Gaules nas redes sociais",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "42442487049",
                                "broadcaster":
                                {
                                    "id": "181077473",
                                    "broadcastSettings":
                                    {
                                        "id": "181077473",
                                        "title": "Pain vs NAVI PGL CS2 Major Copenhagen 2024 - !Sorteio - Siga Gaules nas redes sociais",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 137249,
                                "game":
                                {
                                    "id": "32399",
                                    "slug": "counter-strike",
                                    "displayName": "Counter-Strike",
                                    "name": "Counter-Strike",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "21681549",
                                "login": "pgl",
                                "displayName": "PGL",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/c7496e0d-a668-421a-9567-f7ec89e68d64-profile_image-70x70.png",
                                "primaryColorHex": "000000",
                                "broadcastSettings":
                                {
                                    "id": "21681549",
                                    "title": "NAVI vs. paiN Gaming - PGL CS2 Major Copenhagen 2024 - Elimination Stage - Day 4",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "42126155080",
                                "broadcaster":
                                {
                                    "id": "21681549",
                                    "broadcastSettings":
                                    {
                                        "id": "21681549",
                                        "title": "NAVI vs. paiN Gaming - PGL CS2 Major Copenhagen 2024 - Elimination Stage - Day 4",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 102844,
                                "game":
                                {
                                    "id": "32399",
                                    "slug": "counter-strike",
                                    "displayName": "Counter-Strike",
                                    "name": "Counter-Strike",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "738000896",
                                "login": "evelone2004",
                                "displayName": "evelone2004",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/4af340ba-8d13-44e8-9e72-14c4c86d1f20-profile_image-70x70.png",
                                "primaryColorHex": null,
                                "broadcastSettings":
                                {
                                    "id": "738000896",
                                    "title": "NAVI [0:0] PAIN | PGL CS2 MAJOR COPENHAGEN 2024",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "43887677579",
                                "broadcaster":
                                {
                                    "id": "738000896",
                                    "broadcastSettings":
                                    {
                                        "id": "738000896",
                                        "title": "NAVI [0:0] PAIN | PGL CS2 MAJOR COPENHAGEN 2024",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 74017,
                                "game":
                                {
                                    "id": "32399",
                                    "slug": "counter-strike",
                                    "displayName": "Counter-Strike",
                                    "name": "Counter-Strike",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "445529741",
                                "login": "kyedae",
                                "displayName": "Kyedae",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/d7bab577-5fec-4ce7-9410-6bd46e0bce87-profile_image-70x70.png",
                                "primaryColorHex": "FFF147",
                                "broadcastSettings":
                                {
                                    "id": "445529741",
                                    "title": "SEN vs GenG GRAND FINALS! - Valorant Masters #MastersCostream",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "50712739037",
                                "broadcaster":
                                {
                                    "id": "445529741",
                                    "broadcastSettings":
                                    {
                                        "id": "445529741",
                                        "title": "SEN vs GenG GRAND FINALS! - Valorant Masters #MastersCostream",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 64050,
                                "game":
                                {
                                    "id": "516575",
                                    "slug": "valorant",
                                    "displayName": "VALORANT",
                                    "name": "VALORANT",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "43683025",
                                "login": "ohnepixel",
                                "displayName": "ohnePixel",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/5742b015-e6ed-4f7c-a1dd-87cd88fe1eb9-profile_image-70x70.png",
                                "primaryColorHex": "7C6EFF",
                                "broadcastSettings":
                                {
                                    "id": "43683025",
                                    "title": "🔴NAVI vs PAIN | CS2 MAJOR PGL COPENHAGEN 2024 [ELIMINATION STAGE]🔴",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "42126593592",
                                "broadcaster":
                                {
                                    "id": "43683025",
                                    "broadcastSettings":
                                    {
                                        "id": "43683025",
                                        "title": "🔴NAVI vs PAIN | CS2 MAJOR PGL COPENHAGEN 2024 [ELIMINATION STAGE]🔴",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 53328,
                                "game":
                                {
                                    "id": "32399",
                                    "slug": "counter-strike",
                                    "displayName": "Counter-Strike",
                                    "name": "Counter-Strike",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "96116107",
                                "login": "mixwell",
                                "displayName": "Mixwell",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/6d3b4513-6824-4912-848c-dc046ee262ad-profile_image-70x70.png",
                                "primaryColorHex": null,
                                "broadcastSettings":
                                {
                                    "id": "96116107",
                                    "title": "[ES] GRAN FINAL SEN vs GenG - #VALORANTMasters #Watchparty",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "50712781389",
                                "broadcaster":
                                {
                                    "id": "96116107",
                                    "broadcastSettings":
                                    {
                                        "id": "96116107",
                                        "title": "[ES] GRAN FINAL SEN vs GenG - #VALORANTMasters #Watchparty",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 49998,
                                "game":
                                {
                                    "id": "516575",
                                    "slug": "valorant",
                                    "displayName": "VALORANT",
                                    "name": "VALORANT",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "121606712",
                                "login": "kingsleague",
                                "displayName": "kingsleague",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/881bf9b5-a6bb-419a-950a-b15a36e0dd0e-profile_image-70x70.png",
                                "primaryColorHex": "010A61",
                                "broadcastSettings":
                                {
                                    "id": "121606712",
                                    "title": "👑 Kings League InfoJobs 👑 Kunisports vs Ultimate Móstoles - JORNADA 9 ⚽ #KINGSLEAGUEJ9",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "50712249341",
                                "broadcaster":
                                {
                                    "id": "121606712",
                                    "broadcastSettings":
                                    {
                                        "id": "121606712",
                                        "title": "👑 Kings League InfoJobs 👑 Kunisports vs Ultimate Móstoles - JORNADA 9 ⚽ #KINGSLEAGUEJ9",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 40613,
                                "game":
                                {
                                    "id": "518203",
                                    "slug": "sports-1",
                                    "displayName": "Sports",
                                    "name": "Sports",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        },
                        {
                            "trackingID": "",
                            "promotionsCampaignID": "",
                            "user":
                            {
                                "id": "50985620",
                                "login": "papaplatte",
                                "displayName": "Papaplatte",
                                "profileImageURL": "https://static-cdn.jtvnw.net/jtv_user_pictures/04abc1b4-7bad-4b55-8da8-c0f1cf031bda-profile_image-70x70.png",
                                "primaryColorHex": "7298E6",
                                "broadcastSettings":
                                {
                                    "id": "50985620",
                                    "title": "2V2 OLYMPIADE MIT DEM STEUERFLÜCHTLING , NOTY & DEM KATZENMANN (@nooreax @letshugotv @dimeax)",
                                    "__typename": "BroadcastSettings"
                                },
                                "__typename": "User"
                            },
                            "label": "NONE",
                            "content":
                            {
                                "id": "50713474605",
                                "broadcaster":
                                {
                                    "id": "50985620",
                                    "broadcastSettings":
                                    {
                                        "id": "50985620",
                                        "title": "2V2 OLYMPIADE MIT DEM STEUERFLÜCHTLING , NOTY & DEM KATZENMANN (@nooreax @letshugotv @dimeax)",
                                        "__typename": "BroadcastSettings"
                                    },
                                    "__typename": "User"
                                },
                                "viewersCount": 37280,
                                "game":
                                {
                                    "id": "509658",
                                    "slug": "just-chatting",
                                    "displayName": "Just Chatting",
                                    "name": "Just Chatting",
                                    "__typename": "Game"
                                },
                                "type": "live",
                                "__typename": "Stream"
                            },
                            "__typename": "PersonalSectionChannel"
                        }
                    ],
                    "__typename": "PersonalSection"
                }
            ]
        },
        "extensions":
        {
            "durationMilliseconds": 54,
            "operationName": "PersonalSections",
            "requestID": "01HSRY673A9C0E5JQR4Q5MKEQD"
        }
    }
]`

func TestGetLiveChannels(test *testing.T) {
	test.Run("should successfuly parse a twitch json data", func(t *testing.T) {
		ls := NewLiveService()
		channels, err := ls.parseData(strings.NewReader(rawData))

		assert.Nil(t, err)
		assert.IsType(t, domain.Channels{}, channels)
		assert.NotEmpty(t, channels)
	})
}