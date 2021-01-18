package apperr

// APPErr アプリ内部で取り扱うエラー
type APPErr interface {
	CodeValue() ErrCode
	Error() string
	UnWrap() error
}

// ErrCode APPErr 識別型
type ErrCode int

// ErrCode 種類
const (
	UnknownErrCode ErrCode = iota // 0 無効値
	ConfigErrCode
	InternalErrCode
)

var errCodeValueMap = map[ErrCode]string{
	ConfigErrCode:   "configErrCode",
	InternalErrCode: "internalErrCode",
}

// String ErrCode 文字列を返却
func (c ErrCode) String() string {
	return errCodeValueMap[c]
}
