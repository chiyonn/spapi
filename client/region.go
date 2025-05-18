package client

type RegionConfig struct {
	MarketplaceID string
	BaseURL       string
	AWSRegion     string
}

var countryRegionMap = map[string]RegionConfig{
	// North America
	"US": {MarketplaceID: "ATVPDKIKX0DER", BaseURL: "https://sellingpartnerapi-na.amazon.com", AWSRegion: "us-east-1"},
	"CA": {MarketplaceID: "A2EUQ1WTGCTBG2", BaseURL: "https://sellingpartnerapi-na.amazon.com", AWSRegion: "us-east-1"},
	"MX": {MarketplaceID: "A1AM78C64UM0Y8", BaseURL: "https://sellingpartnerapi-na.amazon.com", AWSRegion: "us-east-1"},
	"BR": {MarketplaceID: "A2Q3Y263D00KWC", BaseURL: "https://sellingpartnerapi-na.amazon.com", AWSRegion: "us-east-1"},

	// Europe
	"IE": {MarketplaceID: "A28R8C7NBKEWEA", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"ES": {MarketplaceID: "A1RKKUPIHCS9HS", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"UK": {MarketplaceID: "A1F83G8C2ARO7P", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"FR": {MarketplaceID: "A13V1IB3VIYZZH", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"BE": {MarketplaceID: "AMEN7PMS3EDWL", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"NL": {MarketplaceID: "A1805IZSGTT6HS", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"DE": {MarketplaceID: "A1PA6795UKMFR9", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"IT": {MarketplaceID: "APJ6JRA9NG5V4", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"SE": {MarketplaceID: "A2NODRKZP88ZB9", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"ZA": {MarketplaceID: "AE08WJ6YKNBMC", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"PL": {MarketplaceID: "A1C3SOZRARQ6R3", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"EG": {MarketplaceID: "ARBP9OOSHTCHU", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"TR": {MarketplaceID: "A33AVAJ2PDY3EV", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"SA": {MarketplaceID: "A17E79C6D8DWNP", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"AE": {MarketplaceID: "A2VIGQ35RCS4UG", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},
	"IN": {MarketplaceID: "A21TJRUUN4KGV", BaseURL: "https://sellingpartnerapi-eu.amazon.com", AWSRegion: "eu-west-1"},

	// Far East
	"SG": {MarketplaceID: "A19VAU5U5O7RUS", BaseURL: "https://sellingpartnerapi-fe.amazon.com", AWSRegion: "us-west-2"},
	"AU": {MarketplaceID: "A39IBJ37TRP1C6", BaseURL: "https://sellingpartnerapi-fe.amazon.com", AWSRegion: "us-west-2"},
	"JP": {MarketplaceID: "A1VC38T7YXB528", BaseURL: "https://sellingpartnerapi-fe.amazon.com", AWSRegion: "us-west-2"},
}

