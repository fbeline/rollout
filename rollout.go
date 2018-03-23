package rollout

import "hash/crc32"
import "math"

// Feature struct keeps the main information of a feature as:
// identification, the percentage of users that will be affected and status
type Feature struct {
	name       string
	percentage float64
	active     bool
}

// Rollout component struct
type Rollout struct {
	features map[string]Feature
}

// IsActive will check if a given user is active for a feature
func (r Rollout) IsActive(feature string, id string) bool {
	f, ok := r.features[feature]
	crc32q := crc32.MakeTable(0xEDB88320)
	crc32 := crc32.Checksum([]byte(id), crc32q)
	return ok && f.percentage > math.Mod(float64(crc32), 100)
}

// IsFeatureActive checks if a feature is active
func (r Rollout) IsFeatureActive(feature string) bool {
	f, ok := r.features[feature]
	return ok && f.active == true
}

// Activate active a feature
// if the feature does not exists it will ignore the action
func (r *Rollout) Activate(feature string) {
	f, ok := r.features[feature]
	if ok {
		f.active = true
		r.Set(f)
	}
}

// Deactivate deactivate a feature
// if the feature does not exists it will ignore the action
func (r *Rollout) Deactivate(feature string) {
	f, ok := r.features[feature]
	if ok {
		f.active = false
		r.Set(f)
	}
}

// Set upsert a feature inside rollout component
func (r *Rollout) Set(feature Feature) {
	r.features[feature.name] = feature
}

// Create is factory used to create a new Rollout
func Create(features []Feature) *Rollout {
	r := Rollout{}
	r.features = make(map[string]Feature)
	for _, v := range features {
		r.features[v.name] = v
	}
	return &r
}
