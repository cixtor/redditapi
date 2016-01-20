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

type KarmaList struct {
	Kind string  `json:"kind"`
	Data []Karma `json:"data"`
}

type Karma struct {
	SubReddit    string `json:"sr"`
	LinkKarma    int    `json:"link_karma"`
	CommentKarma int    `json:"comment_karma"`
}

type Preferences struct {
	DefaultThemeSubreddit  string   `json:"default_theme_sr"`
	ThreadedMessages       bool     `json:"threaded_messages"`
	HideDowns              bool     `json:"hide_downs"`
	ShowStylesheets        bool     `json:"show_stylesheets"`
	ShowLinkFlair          bool     `json:"show_link_flair"`
	CredditAutorenew       bool     `json:"creddit_autorenew"`
	ShowTrending           bool     `json:"show_trending"`
	PrivateFeeds           bool     `json:"private_feeds"`
	MonitorMentions        bool     `json:"monitor_mentions"`
	ShowSnoovatar          bool     `json:"show_snoovatar"`
	Research               bool     `json:"research"`
	IgnoreSuggestedSort    bool     `json:"ignore_suggested_sort"`
	NumComments            int      `json:"num_comments"`
	Clickgadget            bool     `json:"clickgadget"`
	UseGlobalDefaults      bool     `json:"use_global_defaults"`
	LabelNsfw              bool     `json:"label_nsfw"`
	AffiliateLinks         bool     `json:"affiliate_links"`
	Over18                 bool     `json:"over_18"`
	EmailMessages          bool     `json:"email_messages"`
	LiveOrangereds         bool     `json:"live_orangereds"`
	HighlightControversial bool     `json:"highlight_controversial"`
	NoProfanity            bool     `json:"no_profanity"`
	DomainDetails          bool     `json:"domain_details"`
	CollapseLeftBar        bool     `json:"collapse_left_bar"`
	Lang                   string   `json:"lang"`
	HideUps                bool     `json:"hide_ups"`
	PublicServerSeconds    bool     `json:"public_server_seconds"`
	AllowClicktracking     bool     `json:"allow_clicktracking"`
	HideFromRobots         bool     `json:"hide_from_robots"`
	Compress               bool     `json:"compress"`
	StoreVisits            bool     `json:"store_visits"`
	ThreadedModmail        bool     `json:"threaded_modmail"`
	MinLinkScore           int      `json:"min_link_score"`
	MediaPreview           string   `json:"media_preview"`
	EnableDefaultThemes    bool     `json:"enable_default_themes"`
	ContentLangs           []string `json:"content_langs"`
	ShowPromote            bool     `json:"show_promote"`
	MinCommentScore        int      `json:"min_comment_score"`
	PublicVotes            bool     `json:"public_votes"`
	Organic                bool     `json:"organic"`
	CollapseReadMessages   bool     `json:"collapse_read_messages"`
	ShowFlair              bool     `json:"show_flair"`
	MarkMessagesRead       bool     `json:"mark_messages_read"`
	ForceHTTPS             bool     `json:"force_https"`
	HideAds                bool     `json:"hide_ads"`
	Beta                   bool     `json:"beta"`
	Newwindow              bool     `json:"newwindow"`
	Numsites               int      `json:"numsites"`
	LegacySearch           bool     `json:"legacy_search"`
	Media                  string   `json:"media"`
	ShowGoldExpiration     bool     `json:"show_gold_expiration"`
	HighlightNewComments   bool     `json:"highlight_new_comments"`
	DefaultCommentSort     string   `json:"default_comment_sort"`
	HideLocationbar        bool     `json:"hide_locationbar"`
}

type TrophyList struct {
	Kind string     `json:"kind"`
	Data TrophyData `json:"data"`
}

type TrophyData struct {
	Trophies []TrophyInfo `json:"trophies"`
}

type TrophyInfo struct {
	Kind string `json:"kind"`
	Data Trophy `json:"data"`
}

type Trophy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon70      string `json:"icon_70"`
	AwardID     string `json:"award_id"`
	Description string `json:"description"`
	Icon40      string `json:"icon_40"`
	URL         string `json:"url"`
}
type Comment struct {
	Json      CommentJson `json:"json"`
	ErrorMsg  string      `json:"message"`
	ErrorCode int         `json:"error"`
}

type CommentJson struct {
	Data   CommentData `json:"data"`
	Errors []string    `json:"errors"`
}

type CommentData struct {
	Things []CommentThings `json:"things"`
}

type CommentThings struct {
	Kind string           `json:"kind"`
	Data CommentThingData `json:"data"`
}

type CommentThingData struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	LinkID              string   `json:"link_id"`
	ParentID            string   `json:"parent_id"`
	SubredditID         string   `json:"subreddit_id"`
	Edited              bool     `json:"edited"`
	BannedBy            string   `json:"banned_by"`
	RemovalReason       string   `json:"removal_reason"`
	Likes               bool     `json:"likes"`
	Replies             string   `json:"replies"`
	UserReports         []string `json:"user_reports"`
	Saved               bool     `json:"saved"`
	Gilded              int      `json:"gilded"`
	Archived            bool     `json:"archived"`
	ReportReasons       []string `json:"report_reasons"`
	Author              string   `json:"author"`
	Score               int      `json:"score"`
	ApprovedBy          string   `json:"approved_by"`
	Controversiality    int      `json:"controversiality"`
	Body                string   `json:"body"`
	AuthorFlairCSSClass string   `json:"author_flair_css_class"`
	BodyHTML            string   `json:"body_html"`
	Subreddit           string   `json:"subreddit"`
	ScoreHidden         bool     `json:"score_hidden"`
	Stickied            bool     `json:"stickied"`
	Created             float64  `json:"created"`
	AuthorFlairText     string   `json:"author_flair_text"`
	CreatedUtc          float64  `json:"created_utc"`
	Distinguished       string   `json:"distinguished"`
	ModReports          []string `json:"mod_reports"`
	NumReports          int      `json:"num_reports"`
	Downs               int      `json:"downs"`
	Ups                 int      `json:"ups"`
}

type Info struct {
	Kind string   `json:"kind"`
	Data InfoData `json:"data"`
}

type InfoData struct {
	Modhash  string         `json:"modhash"`
	After    string         `json:"after"`
	Before   string         `json:"before"`
	Children []InfoChildren `json:"children"`
}

type InfoChildren struct {
	Kind string           `json:"kind"`
	Data InfoChildrenData `json:"data"`
}

type InfoChildrenData struct {
	ContestMode         bool     `json:"contest_mode"`
	BannedBy            string   `json:"banned_by"`
	Domain              string   `json:"domain"`
	Subreddit           string   `json:"subreddit"`
	SelftextHTML        string   `json:"selftext_html"`
	Selftext            string   `json:"selftext"`
	Likes               int      `json:"likes"`
	SuggestedSort       string   `json:"suggested_sort"`
	UserReports         []string `json:"user_reports"`
	Saved               bool     `json:"saved"`
	ID                  string   `json:"id"`
	Gilded              int      `json:"gilded"`
	Clicked             bool     `json:"clicked"`
	ReportReasons       string   `json:"report_reasons"`
	Author              string   `json:"author"`
	Name                string   `json:"name"`
	Score               int      `json:"score"`
	ApprovedBy          string   `json:"approved_by"`
	Over18              bool     `json:"over_18"`
	RemovalReason       string   `json:"removal_reason"`
	Hidden              bool     `json:"hidden"`
	Thumbnail           string   `json:"thumbnail"`
	SubredditID         string   `json:"subreddit_id"`
	Edited              bool     `json:"edited"`
	LinkFlairCSSClass   string   `json:"link_flair_css_class"`
	AuthorFlairCSSClass string   `json:"author_flair_css_class"`
	Downs               int      `json:"downs"`
	ModReports          []string `json:"mod_reports"`
	Archived            bool     `json:"archived"`
	IsSelf              bool     `json:"is_self"`
	HideScore           bool     `json:"hide_score"`
	Spoiler             bool     `json:"spoiler"`
	Permalink           string   `json:"permalink"`
	Locked              bool     `json:"locked"`
	Stickied            bool     `json:"stickied"`
	Created             float64  `json:"created"`
	URL                 string   `json:"url"`
	AuthorFlairText     string   `json:"author_flair_text"`
	Quarantine          bool     `json:"quarantine"`
	Title               string   `json:"title"`
	CreatedUtc          float64  `json:"created_utc"`
	LinkFlairText       string   `json:"link_flair_text"`
	Distinguished       string   `json:"distinguished"`
	NumComments         int      `json:"num_comments"`
	Visited             bool     `json:"visited"`
	NumReports          int      `json:"num_reports"`
	Ups                 int      `json:"ups"`

	// SecureMedia         interface{} `json:"secure_media"`
	// Media               interface{} `json:"media"`
}
