package rollout

import (
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestRollout_IsActive(t *testing.T) {
	features := []Feature{
		Feature{"foo", 100, false},
		Feature{"bar", 50, true},
		Feature{"baz", 0, true},
	}
	r := Create(features)

	id, _ := uuid.NewV4()

	type fields struct {
		features map[string]Feature
	}
	type args struct {
		feature string
		id      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"rollout 100%", fields{r.features}, args{"foo", id.String()}, true},
		{"rollout 0%", fields{r.features}, args{"baz", id.String()}, false},
		{"rollout 50% out", fields{r.features}, args{"bar", "d0b7b9df-9fa8-4f72-885b-3f1cd0d705d5"}, false},
		{"rollout 50% int", fields{r.features}, args{"bar", "8975b460-4446-4af2-965a-ba020092e1ca"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rollout{
				features: tt.fields.features,
			}
			if got := r.IsActive(tt.args.feature, tt.args.id); got != tt.want {
				t.Errorf("Rollout.IsActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRollout_IsFeatureActive(t *testing.T) {
	features := []Feature{
		Feature{"foo", 0.5, false},
		Feature{"bar", 0.5, true},
	}
	r := Create(features)
	type fields struct {
		features map[string]Feature
	}
	type args struct {
		feature string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"active", fields{r.features}, args{"foo"}, false},
		{"deactivated", fields{r.features}, args{"bar"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rollout{
				features: tt.fields.features,
			}
			if got := r.IsFeatureActive(tt.args.feature); got != tt.want {
				t.Errorf("Rollout.IsFeatureActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRollout_Activate(t *testing.T) {
	features := []Feature{Feature{"foo", 0.5, false}}
	r := Create(features)
	r.Activate("foo")
	assert.Equal(t, true, r.features["foo"].Active)
}

func TestRollout_Deactivate(t *testing.T) {
	features := []Feature{Feature{"foo", 0.5, true}}
	r := Create(features)
	r.Deactivate("foo")
	assert.Equal(t, false, r.features["foo"].Active)
}

func TestRollout_Set(t *testing.T) {
	features := []Feature{Feature{"foo", 0.5, true}}
	expected := Feature{"foo", 0.7, true}
	r := Create(features)
	r.Set(expected)
	assert.Equal(t, expected, r.features["foo"])
}

func TestCreate(t *testing.T) {
	f := Feature{"foo", 0.5, true}
	features := []Feature{f}
	r := Create(features)
	assert.Equal(t, f, r.features["foo"])
}
