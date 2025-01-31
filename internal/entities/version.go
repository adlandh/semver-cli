package entities

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/adlandh/semver-cli/internal/enum/phases"
	"github.com/adlandh/semver-cli/internal/utils/str"
)

type Version struct {
	Prefix      string
	Major       uint
	Minor       uint
	Patch       uint
	Phase       phases.Phase
	PatchNumber uint
}

func NewVersion(v string) (*Version, error) {
	version := &Version{}
	err := version.Set(v)
	if err != nil {
		return nil, err
	}
	return version, nil
}

func (v *Version) String() string {
	if v.Phase.IsRelease() {
		return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
	}
	return fmt.Sprintf("v%d.%d.%d-%s.%d", v.Major, v.Minor, v.Patch, v.Phase, v.PatchNumber)
}

func (v *Version) Set(value string) error {
	tagPattern, _ := regexp.Compile(`^(v?)(\d+)\.(\d+)\.(\d+)((-(alpha|beta|rc))\.(\d+))?$`)
	if !tagPattern.MatchString(value) {
		return errors.New("invalid version format")
	}

	parts := tagPattern.FindStringSubmatch(value)
	v.Prefix = parts[1]
	v.Major = str.ParseUIntOrDefault(parts[2])
	v.Minor = str.ParseUIntOrDefault(parts[3])
	v.Patch = str.ParseUIntOrDefault(parts[4])
	v.Phase = phases.ValueOf(parts[7])
	v.PatchNumber = str.ParseUIntOrDefault(parts[8])

	return nil
}

func (v *Version) Type() string {
	return "version"
}
