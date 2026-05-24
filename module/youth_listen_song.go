package module

import "moekoe-go/util"

func init() { Register("/youth/listen/song", YouthListenSong) }
func YouthListenSong(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL:"/youth/v2/report/listen_song",Method:"POST",
		Data: map[string]interface{}{"mixsongid":toIntDefault(params,"mixsongid",666075191)},
		Params: map[string]interface{}{"clientver":10566},EncryptType:"android",Cookie:cookies,
		Headers:map[string]string{"user-agent":"Android13-1070-10566-201-0-ReportPlaySongToServerProtocol-wifi","content-type":"application/json; charset=utf-8"},
	})
}