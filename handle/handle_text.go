package handle

import (
	"html"
	"regexp"
	"strings"
)

/****************评论过滤*******************/
func replaceEmoji(com string) string {
	com = html.EscapeString(com)
	com = strings.Replace(com, "[大笑]", `<img title="大笑" src="https://i.loli.net/2019/08/05/wbmvoD5ycFJ79KA.gif">`, -1)
	com = strings.Replace(com, "[喷血]", `<img title="喷血" src="https://i.loli.net/2019/08/05/EbY8FCLmKcQIWHz.gif">`, -1)
	com = strings.Replace(com, "[抠鼻]", `<img title="抠鼻" src="https://i.loli.net/2019/08/05/xfhvsjE284nPVip.gif">`, -1)
	com = strings.Replace(com, "[吐]", `<img title="吐" src="https://i.loli.net/2019/08/05/98nzCEk4pHriBV1.gif">`, -1)
	com = strings.Replace(com, "[微笑]", `<img  title="微笑" src="https://i.loli.net/2019/08/05/YoNItSivPL2ZJAn.gif">`, -1)
	com = strings.Replace(com, "[笑哭]", `<img  title="笑哭" src="https://i.loli.net/2019/08/05/PX6M1eFxrYVZNdC.gif">`, -1)
	com = strings.Replace(com, "[斜眼笑]", `<img  title="斜眼笑" src="https://i.loli.net/2019/08/05/vcDieIRrf25KlAz.gif">`, -1)
	com = strings.Replace(com, "[阴险]", `<img  title="阴险" src="https://i.loli.net/2019/08/05/tNqeoXlZcywU41R.gif">`, -1)
	com = strings.Replace(com, "[doge]", `<img  title="doge" src="https://i.loli.net/2019/08/05/sRAJywb56qphYK1.gif">`, -1)
	com = strings.Replace(com, "[狗头]", `<img  title="狗头" src="https://i.loli.net/2019/08/05/Ww6GxbhD1Vamq8A.gif">`, -1)
	com = strings.Replace(com, "[猪头]", `<img  title="猪头" src="https://i.loli.net/2019/08/05/jGbMDPfFphlrXuz.gif">`, -1)
	com = strings.Replace(com, "[马]", `<img  title="马" src="https://i.loli.net/2019/08/05/MnKwdZi4eBVj3LD.gif">`, -1)
	com = strings.Replace(com, "[牛]", `<img  title="牛" src="https://i.loli.net/2019/08/05/AguMSPy1DTN7BVt.gif">`, -1)
	com = strings.Replace(com, "[啤酒]", `<img  title="啤酒" src="https://i.loli.net/2019/08/05/q8i1MQUCE5YrOxF.gif">`, -1)
	return com
}

/******************正则替换*****************/
func regexpReplaceAllString(content, rule, key string) string {
	r1 := regexp.MustCompile(rule)
	r2 := r1.ReplaceAllString(content, key)
	return r2
}
