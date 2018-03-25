# rollout

Feature based rollout for Golang.

## Installation

`go get github.com/fbeline/rollout`

## Introduction

This project is a small library to create feature rollout based on percentage.
Example: You have a new feature that you wish to only 5% of your users to have access to it.

The library is free of any persistence system. But I strongly recommend to you to persist the component configuration at disk instead make it hardcoded.*

*in case of a bug or the act of increasing and decreasing the percentage you could just change a file in the disk or a value in the database.

## How to use

import the rollout package

```go
import "github.com/fbeline/rollout"
```

### Creating a rollout

```go
var foo = rollout.Feature{Name: "foo", Percentage: 0.5, Active: true}
var bar = rollout.Feature{Name: "bar", Percentage: 0.7, Active: true}
var features = []rollout.Feature{foo, bar}
var r = rollout.Create(features)
```

### Checking if a user is in rollout

`IsActive` will returns `true` if the given customer identification is in the rollout and false if it's not.

```go
IsActive("featureName", "UserId")
```

### Upsert a feature

The feature name is used as a unique key. If the feature exists it will be updated otherwise created.

```go
var newFoo = rollout.Feature{Name: "foo", Percentage: 0.8, Active: true}
r.Set(newFoo)
```

### Checking if a feature is active

```go
IsFeatureActive("featureName")
```

### Activating a feature

```go
Activate("featureName")
```

### Deactivating a feature

```go
Deactivate("featureName")
```

## License

MIT
