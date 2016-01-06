package redditapi

import "encoding/json"

type Account struct {
	ID                      string `json:"id"`
	Name                    string `json:"name"`
	Over18                  bool   `json:"over_18"`
	LinkKarma               int    `json:"link_karma"`
	CommentKarma            int    `json:"comment_karma"`
	GoldCreddits            int    `json:"gold_creddits"`
	GoldExpiration          string `json:"gold_expiration"`
	HasMail                 bool   `json:"has_mail"`
	HasModMail              bool   `json:"has_mod_mail"`
	HasVerifiedEmail        bool   `json:"has_verified_email"`
	HideFromRobots          bool   `json:"hide_from_robots"`
	InBeta                  bool   `json:"in_beta"`
	InboxCount              int    `json:"inbox_count"`
	IsEmployee              bool   `json:"is_employee"`
	IsGold                  bool   `json:"is_gold"`
	IsMod                   bool   `json:"is_mod"`
	IsSuspended             bool   `json:"is_suspended"`
	SuspensionExpirationUtc string `json:"suspension_expiration_utc"`

	Created    json.Number     `json:"created"`
	CreatedUtc json.Number     `json:"created_utc"`
	Features   AccountFeatures `json:"features"`
}

type AccountFeatures struct {
	YoutubeScraper         bool `json:"youtube_scraper"`
	ShowNewIcons           bool `json:"show_new_icons"`
	LiveHappeningNow       bool `json:"live_happening_now"`
	AdserverReporting      bool `json:"adserver_reporting"`
	LegacySearchPref       bool `json:"legacy_search_pref"`
	MobileWebTargeting     bool `json:"mobile_web_targeting"`
	AdsAuction             bool `json:"ads_auction"`
	AdzerkDoNotTrack       bool `json:"adzerk_do_not_track"`
	StickyComments         bool `json:"sticky_comments"`
	UpgradeCookies         bool `json:"upgrade_cookies"`
	AdsPrefs               bool `json:"ads_prefs"`
	AdsAutoRefund          bool `json:"ads_auto_refund"`
	OrangeredsAsEmails     bool `json:"orangereds_as_emails"`
	ImgurGifConversion     bool `json:"imgur_gif_conversion"`
	ExpandoEvents          bool `json:"expando_events"`
	EuCookiePolicy         bool `json:"eu_cookie_policy"`
	ForceHTTPS             bool `json:"force_https"`
	MobileNativeBanner     bool `json:"mobile_native_banner"`
	DoNotTrack             bool `json:"do_not_track"`
	OutboundClicktracking  bool `json:"outbound_clicktracking"`
	ImageUploads           bool `json:"image_uploads"`
	NewLoggedinCachePolicy bool `json:"new_loggedin_cache_policy"`
	HTTPSRedirect          bool `json:"https_redirect"`
	ScreenviewEvents       bool `json:"screenview_events"`
	PauseAds               bool `json:"pause_ads"`
	GiveHstsGrants         bool `json:"give_hsts_grants"`
	NewReportDialog        bool `json:"new_report_dialog"`
	MoatTracking           bool `json:"moat_tracking"`
	SubredditRules         bool `json:"subreddit_rules"`
	Timeouts               bool `json:"timeouts"`
	MobileSettings         bool `json:"mobile_settings"`
	AdzerkReporting2       bool `json:"adzerk_reporting_2"`
	ActivityServiceWrite   bool `json:"activity_service_write"`
	AdsAutoExtend          bool `json:"ads_auto_extend"`
	InterestTargeting      bool `json:"interest_targeting"`
	PostEmbed              bool `json:"post_embed"`
	ScrollEvents           bool `json:"scroll_events"`
	AdblockTest            bool `json:"adblock_test"`
	ActivityServiceRead    bool `json:"activity_service_read"`

	SeoCommentsPageHoldout AccountSeoComments `json:"seo_comments_page_holdout"`
}

type AccountSeoComments struct {
	Variant      string `json:"variant"`
	ExperimentID int    `json:"experiment_id"`
}

type UserList struct {
	Kind string       `json:"kind"`
	Data UserListData `json:"data"`
}

type UserListData struct {
	Children []Friend `json:"children"`
}

type Friend struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Date json.Number `json:"date"`
}
