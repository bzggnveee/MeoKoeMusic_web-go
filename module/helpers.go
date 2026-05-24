package module

import ("fmt"; "sort"; "strconv"; "strings"; "moekoe-go/util")

func toString(v interface{}) string {
	switch val := v.(type) {
	case string: return val
	case float64:
		if val==float64(int64(val)) { return strconv.FormatInt(int64(val),10) }
		return strconv.FormatFloat(val,'f',-1,64)
	case int: return strconv.Itoa(val)
	case int64: return strconv.FormatInt(val,10)
	case bool: if val{return "true"}; return "false"
	default: return fmt.Sprintf("%v",v)
	}
}

func parseAnyInt(s string) int {
	f, err := strconv.ParseFloat(s, 64)
	if err == nil { return int(f) }
	n, _ := strconv.Atoi(s)
	return n
}

func toInt(v interface{}) int {
	switch val := v.(type) {
	case float64: return int(val)
	case int: return val
	case int64: return int(val)
	case string: return parseAnyInt(val)
	default: return 0
	}
}
func toBool(v interface{}) bool {
	switch val := v.(type) {
	case bool: return val
	case string: return val=="true"||val=="1"
	case float64: return val!=0
	default: return false
	}
}
func toLowerCase(s string) string { return strings.ToLower(s) }
func toUpper(s string) string { return strings.ToUpper(s) }
func toIntDefault(m map[string]interface{}, key string, def int) int { if v,ok:=m[key];ok{return toInt(v)}; return def }
func toStringDefault(m map[string]interface{}, key string, def string) string { if v,ok:=m[key];ok{return toString(v)}; return def }
func getVal(params map[string]interface{}, cookies map[string]string, key string) string {
	if v,ok:=params[key];ok{return toString(v)}
	if v,ok:=cookies[key];ok{return v}
	return ""
}
func getValInt(params map[string]interface{}, cookies map[string]string, key string) int {
	if v,ok:=params[key];ok{return toInt(v)}
	if v,ok:=cookies[key];ok{return parseAnyInt(v)}
	return 0
}
func dfidVal(params map[string]interface{}, cookies map[string]string) string {
	if v,ok:=params["dfid"];ok{return toString(v)}
	if v,ok:=cookies["dfid"];ok&&v!=""{return v}
	return "-"
}
func splitByPipe(s string) []string { if s==""{return nil}; return strings.Split(s,"|") }
func splitCSV(s string) []string { if s==""{return nil}; return strings.Split(s,",") }
func kugouSign(dataMap map[string]interface{}) string {
	const str="*s&iN#G70*"
	keys:=make([]string,0,len(dataMap))
	for k:=range dataMap{keys=append(keys,k)}
	sort.Strings(keys)
	var parts []string
	for _,k:=range keys{
		v:=dataMap[k]; var vs string
		switch val:=v.(type){case string:vs=val; default:vs=fmt.Sprintf("%v",val)}
		parts=append(parts,fmt.Sprintf("%s=%s",k,vs))
	}
	return util.CryptoMD5(strings.Join(parts,"&")+str)[8:24]
}
func rsaEncrypt2(data map[string]interface{}) string { return util.RsaEncrypt2(data) }
func playlistAesEncrypt(data map[string]interface{}) util.AesEncryptResult { return util.PlaylistAesEncrypt(data) }
func playlistAesDecrypt(str, key string) interface{} { return util.PlaylistAesDecrypt(str,key) }