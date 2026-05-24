package module

import "moekoe-go/util"

func init() { Register("/playlist/tags", PlaylistTags) }
func PlaylistTags(params map[string]interface{}, cookies map[string]string, requestFn func(util.RequestConfig) (*util.Response, error)) (*util.Response, error) {
	return requestFn(util.RequestConfig{
		URL: "/pubsongs/v1/get_tags_by_type", Method: "POST",
		Data: map[string]interface{}{"tag_type": "collection", "tag_id": 0, "source": 3},
		EncryptType: "android", Cookie: cookies,
	})
}