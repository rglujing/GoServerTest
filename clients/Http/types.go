package HttpClient;

type CreateAuCode_ParaReq struct {
        Appid         string       `json:"appid"`
        Callbackurl   string       `json:"callbackurl"`
        SetExpire     uint32       `json:"setExpire,omitempty"`
        Uname         string       `json:"uname"`
        Usertid       string       `json:"usertid"`
        ReqData       interface{}  `json:"reqdata,omitempty"`
}

type CreateAuCode_ParaResp_Suc struct {
	Result   int                 `json:"result"`
	AuCode   string              `json:"Aucode"`
	ReqData  interface{}         `json:"req,omitempty"`
}

type RequestToken_ParaReq struct {
	Appid         string       `json:"appid"`
	Secret        string       `json:"secret"`
	AuCode        string       `json:"AuCode"`
	CallBackUrl   string       `json:"callbackurl"`
	ReqData       interface{}  `json:"reqdata,omitempty"`
}

type RequestToken_ParaResp_Suc struct {
	Result        int           `json:"result"`
	AsToken       string        `json:"AsToken"`
	Expireln      uint32        `json:"expireln"`
	Usertid       string        `json:"usertid"`
	ReqData       interface{}   `json:"req,omitempty"`
}

type  AccessToken_ParaReq        struct {
	AsToken    string          `json:"AsToken"`
	ReqData    interface{}     `json:"req,omitempty"`
}

type  AccessToken_ParaResp_Suc  struct {
	Result     int             `json:"result"`
	Uname      string          `json:"uname"`
	Usertid    string          `json:"usertid"`
	Expireln   uint32          `json:"expireln"`
	CacheData  interface{}     `json:"cachedata,omitempty"`
	ReqData    interface{}     `json:"req,omitempty"`
}

type  RefreshToken_ParaReq        struct {
	Appid     string          `json:"appid"`
	Secret    string          `json:"secret"`
	AsToken   string          `json:"AsToken"`
	AddExpire uint32          `json:"addExpire,omitempty"`
	ReqData   interface{}     `json:"reqdata,omitempty"`
}

type   RefreshToken_ParaResp_Suc   struct {
	Result    int             `json:"result"`
	AsToken   string          `json:"AsToken"`
	Expireln  uint32          `json:"expireln"`
	Usertid   string          `json:"usertid"`
	ReqData   interface{}     `json:"req,omitempty"`
}

type  InspectToken_ParaReq        struct {
	AsToken    string          `json:"AsToken"`
	ReqData    interface{}     `json:"req,omitempty"`
}

type   InspectToken_ParaResp_Suc   struct {
	Result    int             `json:"result"`
	Appid     string          `json:"appid"`
	Usertid   string          `json:"usertid"`
	CreateTM  uint32          `json:"createTM"`
	Expireln  uint32          `json:"expireln"`
	ReqData   interface{}     `json:"req,omitempty"`
}

type   RequestSession_ParaReq  struct {
        Appid         string       `json:"appid"`
	Secret        string       `json:"secret"`
	AuCode        string       `json:"AuCode"`
	CallBackUrl   string       `json:"callbackurl"`
	CacheData     interface{}  `json:"cachedata"`
	ReqData       interface{}  `json:"reqdata,omitempty"`
}

type RequestSession_ParaResp_Suc struct {
	Result        int           `json:"result"`
	AsToken       string        `json:"AsToken"`
	Expireln      uint32        `json:"expireln"`
	Usertid       string        `json:"usertid"`
	ReqData       interface{}   `json:"req,omitempty"`
}

type   SetSessionData_ParaReq  struct {
        Appid         string       `json:"appid"`
	Secret        string       `json:"secret"`
	AsToken       string       `json:"AsToken"`
	CacheData     interface{}  `json:"cachedata"`
	ReqData       interface{}  `json:"reqdata,omitempty"`
}

type   SetSessionData_ParaResp_Suc struct {
	Result        int           `json:"result"`
	AsToken       string        `json:"AsToken"`
	Expireln      uint32        `json:"expireln"`
	Usertid       string        `json:"usertid"`
	ReqData       interface{}   `json:"req,omitempty"`
}

type   Common_ParaResp_Fail   struct {
	Result    int            `json:"result"`
	Msg       string         `json:"msg"`
	ReqData   interface{}    `json:"req,omitempty"`
}
