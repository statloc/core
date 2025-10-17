package matching

import "regexp"

func FindMatch[valueType any](s string, mapping map[string]valueType) (value valueType, exists bool) {
    for key, value := range mapping {
        matches, _ := regexp.MatchString(key, s)
        if matches {
            return value, true
        }
    }
    return
}
