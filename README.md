# go-ffrobe
simple wrapper for ffprobe command 

#CODE
file := &FFProbeMeta{}
meta := file.SetFile("path to file").GetMeta()
fmt.Println((meta[0].(map[string]interface{}))["duration"])
#CODE
