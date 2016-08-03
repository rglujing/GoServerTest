package token;

const (
	RedisPrefix = "AuCode:"
	BdbPrefix = "token:"

	CREATE_AUCODE = 1
	REQUEST_TOKEN = 2
	ACCESS_TOKEN  = 3
	REFRESH_TOKEN = 4
	INSPECT_TOKEN = 5
	DESTROY_TOKEN = 6
	REQUEST_SESSION = 7	
	SET_SESSION_DATA = 8

	METHOD_CREATE_AUCODE = "CreateAucode"
	METHOD_REQUEST_TOKEN = "RequestToken"
	METHOD_REFRESH_TOKEN = "RefreshToken"
	METHOD_REQUEST_SESSION = "RequestSession"
	METHOD_SET_SESSION_DATA = "SetSessionData"
)


type  Common_ParaResp_Fail   struct {
	Result    int             `json:"result"`
	Msg       string          `json:"msg"`
	ReqData   interface{}     `json:"req,omitempty"`
}

type RedisStore struct {
	Appid        string    `json:"appid"`
	SetExpire     uint32    `json:"expire"`
	Uname        string    `json:"uname"`
	Usertid      string    `json:"usertid"`
	CreateTM     int64     `json:"createtm"`
}

type BdbStore struct {
	Appid        string         `json:"appid"`
	SetExpire     uint32        `json:"expire"`
	Uname        string         `json:"uname"`
	Usertid      string         `json:"usertid"`
	CreateTM     int64          `json:"createtm"`
	CacheData    interface{}    `json:"cachedata,omitempty"`
}


type TokenErr struct {
	RetCode int
	Msg     string
}

const (
	RET_SUC = 0
	/*Third Party Error, add RetCode*/
	RET_ERR_BASE0           = 1100
	RET_INVALID_JSON        = RET_ERR_BASE0 + 1
	RET_INVALID_REDIS_IO    = RET_ERR_BASE0 + 2
	RET_INVALID_BDB_IO      = RET_ERR_BASE0 + 3
	RET_INVALID_ETCD_IO     = RET_ERR_BASE0 + 4
	RET_INVALID_IO          = RET_ERR_BASE0 + 5

	/*Self Define Error*/
	RET_INTERNAL_ERR     = 1000
	RET_INVALID_ACCESS   = RET_INTERNAL_ERR + 1
	RET_INVALID_SECRET   = RET_INTERNAL_ERR + 2
	RET_INVALID_AUCODE   = RET_INTERNAL_ERR + 3
	RET_INVALID_ASTOKEN  = RET_INTERNAL_ERR + 4
	RET_EXIPIER_OVERFLOW = RET_INTERNAL_ERR + 5
	RET_ASTOKEN_TIMEOUT  = RET_INTERNAL_ERR + 6
	RET_CACHE_OVERLEN    = RET_INTERNAL_ERR + 7
	RET_INVALID_METHOD   = RET_INTERNAL_ERR + 8
	RET_PACKAGE_OVERFLOW = RET_INTERNAL_ERR + 9
	RET_VISIT_TOO_MANY   = RET_INTERNAL_ERR + 10
	RET_INVALID_PARAM    = RET_INTERNAL_ERR + 11	

	Internal_Error  =   "err:Internal Error"
	Invalid_Mix     =   "err:appid not exist and gateway not allowed"
	Invalid_Appid   =   "err:invalid appid"
	Invalid_GW      =   "err:gateway not allowed"
	Invalid_Method  =   "err:you have no right to access this method"
	Invalid_Secret  =   "err:invalid secret"
	Invalid_AuCode  =   "err:invalid Aucode"
	Invalid_AsToken =   "err:invalid AsToken"
	Expire_Overflow =   "err:expire too max, over flow"
	AsToken_Timeout =   "err:astoken timeout"
	Cache_Overlen   =   "err:cache data is too large"
	Visit_Too_Many  =   "err:Visit too many times in these time, pls wait for a while"
	Invalid_Param   =   "err:invalid parameters"

	/*Const Error, used when receive unsupported package*/
	Conten_tlong_json   =   "{\"result\": 1009 , \"msg\":\"err:Package Size is too long\"}"
	Invalid_method_json =   "{\"result\": 1008 , \"msg\":\"err:invalid method or parameter\"}"
	/*Const Error, used when json marshal failed, this may never be used*/
	Internal_error_json =   "{\"result\": 1000    , \"msg\":\"err:Internal Error,please connect to the interface provider!\"}"
)

var (
	Err_Internal       =  &TokenErr{ RET_INTERNAL_ERR,     Internal_Error } 
	Err_Invalid_Mix    =  &TokenErr{ RET_INVALID_ACCESS,   Invalid_Mix    }
	Err_Invalid_Appid  =  &TokenErr{ RET_INVALID_ACCESS,   Invalid_Appid  }
	Err_Invalid_GW     =  &TokenErr{ RET_INVALID_ACCESS,   Invalid_GW     }
	Err_Invalid_Method =  &TokenErr{ RET_INVALID_ACCESS,   Invalid_Method }
	Err_Invalid_Secret =  &TokenErr{ RET_INVALID_SECRET,   Invalid_Secret }
	Err_Invalid_AuCode =  &TokenErr{ RET_INVALID_AUCODE,   Invalid_AuCode }	
	Err_Invalid_AsToken = &TokenErr{ RET_INVALID_ASTOKEN,  Invalid_AsToken }
	//Err_Expire_Overflow = &TokenErr{ RET_EXIPIER_OVERFLOW, Expire_Overflow }	
	Err_AsToken_Timeout = &TokenErr{ RET_ASTOKEN_TIMEOUT,  AsToken_Timeout }
	Err_Cache_Overlen   = &TokenErr{ RET_CACHE_OVERLEN,    Cache_Overlen   }
	Err_Visit_TooMany   = &TokenErr{ RET_VISIT_TOO_MANY,   Visit_Too_Many  }
	Err_Invalid_Param   = &TokenErr{ RET_INVALID_PARAM,    Invalid_Param   }
	MAX_JSON_LEN = 0
)

