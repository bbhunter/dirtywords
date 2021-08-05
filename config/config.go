package config

type CollInfo []struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TimeGate string `json:"timegate"`
	CDXAPI   string `json:"cdx-api"`
}

type OTXResult struct {
	HasNext    bool `json:"has_next"`
	ActualSize int  `json:"actual_size"`
	URLList    []struct {
		Domain   string `json:"domain"`
		URL      string `json:"url"`
		Hostname string `json:"hostname"`
		HTTPCode int    `json:"httpcode"`
		PageNum  int    `json:"page_num"`
		FullSize int    `json:"full_size"`
		Paged    bool   `json:"paged"`
	} `json:"url_list"`
}

const OTXResultsLimit = 200

type PageInfo struct {
	Pages    int `json:"pages"`
	PageSize int `json:"pageSize"`
	Blocks   int `json:"blocks"`
}

type UrlInfo struct {
	URL   string `json:"url"`
	Error string `json:"error"`
}

type Wayback [][]string
