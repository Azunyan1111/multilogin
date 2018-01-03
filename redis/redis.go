package redis


var oneTime map[string]string = map[string]string{}

func Get(key string) string {
	v, ok := oneTime[key]
	if ok {
		return v
	} else {
		return ""
	}
}

func Set(key string,value string){
	oneTime[key] = value
}