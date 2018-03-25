package rollout

import "hash/crc32"
import "math"

// Feature struct keeps the main information of a feature as:
// identification, the percentage of users that will be affected and status
type Feature struct {
	Name       string
	Percentage float64
	Active     bool
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
	return ok && f.Percentage > math.Mod(float64(crc32), 100)
}

// IsFeatureActive checks if a feature is active
func (r Rollout) IsFeatureActive(feature string) bool {
	f, ok := r.features[feature]
	return ok && f.Active
}

// Activate active a feature
// if the feature does not exists the action is ignored
func (r *Rollout) Activate(feature string) {
	f, ok := r.features[feature]
	if ok {
		f.Active = true
		r.Set(f)
	}
}

// Deactivate deactivate a feature
// if the feature does not exists the action is ignored
func (r *Rollout) Deactivate(feature string) {
	f, ok := r.features[feature]
	if ok {
		f.Active = false
		r.Set(f)
	}
}

// Set upsert a feature inside rollout component
func (r *Rollout) Set(feature Feature) {
	r.features[feature.Name] = feature
}

// Create is function used to create a new Rollout
func Create(features []Feature) *Rollout {
	r := Rollout{}
	r.features = make(map[string]Feature)
	for _, v := range features {
		r.features[v.Name] = v
	}
	return &r
}

// Get a feature by name
func (r Rollout) Get(feature string) (Feature, bool) {
	f, ok := r.features[feature]
	return f, ok
}

// GetAll get all features
func (r Rollout) GetAll() []Feature {
	var features []Feature
	for _, el := range r.features {
		features = append(features, el)
	}
	return features
}
