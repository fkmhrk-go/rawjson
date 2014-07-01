rawjson
=======

Just a support library of "encoding/json". This library may be useful when you want to handle Raw Json Object / Array. 

You know we can handle Raw Json as follows:

```
s := "{\"name\":\"fkm\",\"age\":29,\"lang\":[\"Java\",\"Go\"]}"
r := strings.NewReader(s)
doc := json.NewDecoder(r)
// decode as raw json
var obj map[string]interface{}
doc.Decode(&obj)
```


Example
=======

```
package main

import (
        "fmt"
        "github.com/fkmhrk-go/rawjson"
)

func main() {
        s := "{\"name\":\"fkm\"}"
	json, _ := rawjson.ObjectFromString(s)
	val, _ := json.String("name")
	fmt.Printf("name is %s", val)
}
```
