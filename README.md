# rollout [![Build Status](https://scrutinizer-ci.com/g/fbeline/rollout/badges/build.png?b=master)](https://scrutinizer-ci.com/g/fbeline/rollout/build-status/master) [![Code Coverage](https://scrutinizer-ci.com/g/fbeline/rollout/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/fbeline/rollout/?branch=master) [![](https://godoc.org/github.com/nathany/looper?status.svg)](https://godoc.org/github.com/fbeline/rollout)

Feature based rollout for Golang.

## Installation

`go get github.com/fbeline/rollout`

## Introduction

rollout is a library used to create feature rollouts based on percentage.

usage examples:

- You have a new feature and you want to expose that only for 5% of your userbase.

- A/B tests.

The library is free of any persistence system, but I strongly recommend to you to persist the rollout state at disk instead make it hardcoded. In that way, you will be able to fast manipulate the rollout percentages and status.

## How to use

import the rollout package

```go
import "github.com/fbeline/rollout"
```

### Creating a rollout

```go
var foo = rollout.Feature{Name: "foo", Percentage: 0.5, Active: true, Whitelist: []string{}}
var bar = rollout.Feature{Name: "bar", Percentage: 0.7, Active: true, Whitelist: []string{}}
var features = []rollout.Feature{foo, bar}
var r = rollout.Create(features)
```

### Check if a feature is active for a given user

```go
<rollout instance>.IsActive("featureName", "UserId")
```

### Upsert a feature

The feature name is used as a unique key. If the feature exists it will be updated otherwise created.

```go
var newFoo = rollout.Feature{Name: "foo", Percentage: 0.8, Active: true}
<rollout instance>.Set(newFoo)
```

### Checking if a feature is active

```go
<rollout instance>.IsFeatureActive("featureName")
```

### Activate a feature

```go
<rollout instance>.Activate("featureName")
```

### Deactivate a feature

```go
<rollout instance>.Deactivate("featureName")
```

### Get a feature

The first value (f) is assigned the value stored under the feature name.The second value (ok) is a bool that is true if the feature exists, and false if not.

```go
f, ok := <rollout instance>.Get("featureName")
```

### Get all features

```go
<rollout instance>.GetAll()
```

## License

MIT
