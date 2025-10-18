# statloc core

<div>
    <img src="https://img.shields.io/github/v/release/statloc/core?sort=semver&display_name=release&style=flat-square&label=latest%20release&color=purple">
    <a href="https://github.com/statloc/core/blob/master/go.mod"><img src="https://img.shields.io/badge/go_version-1.24%20|%201.25-purple?style=flat-square&label=go%20version&color=purple"></a>
    <a href="https://github.com/statloc/core/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-coffeeware-purple?style=flat-square&label=license&color=purple"></a>
    <a href="https://github.com/statloc/core/stargazers/"><img src="https://img.shields.io/github/stars/statloc/core?style=flat-square&label=stars&color=purple"></a>
    <a href="https://github.com/statloc/core/forks/"><img src="https://img.shields.io/github/forks/statloc/core?style=flat-square&color=purple"></a>
    <a href="https://github.com/statloc/core/actions/workflows/check.yml/"><img src="https://img.shields.io/github/actions/workflow/status/statloc/core/check.yml?branch=master&style=flat-square&label=checks&color=purple"></a>
</div>

### ✏️ About
- Go package for collecting project statistic (lines of code, methods, files, components, etc.)
- Core package for @statloc
- Used both by server and CLI

### ⚡ Installation
```shell
go get github.com/statloc/core
```

### 📝 Sample usage
```go
package main

import (
    "github.com/statloc/core"
    "fmt"
)

func main() {
    stats := core.GetStatistics("/path/to/project")

    result := = fmt.Sprintf(
        `Statistics(total):
lines of code: %d
files: %d`,
        stats.Total.LOC,
        stats.Total.Files,
    )

    fmt.Println(results)
}

```

### 🤝 Contributing
Follow our [guide](https://github.com/statloc/core/blob/master/CONTRIBUTING.md) to contribute to this project
