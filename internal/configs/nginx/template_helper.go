package nginx

import (
        "strings"
        "text/template"
)

func headerListToCIMap(headers []Header) map[string]string {
        ret := make(map[string]string)

        for _, header := range headers {
                ret[strings.ToLower(header.Name)] = header.Value
        }

        return ret
}

func hasCIKey(key string, d map[string]string) bool {
        _, ok := d[strings.ToLower(key)]
        return ok
}


func toLower(s string) string {
        return strings.ToLower(s)
}

var helperFunctions = template.FuncMap{
        "headerListToCIMap": headerListToCIMap,
        "hasCIKey":          hasCIKey,
        "toLower":           toLower,
}
