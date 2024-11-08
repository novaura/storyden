// Code generated by enumerator. DO NOT EDIT.

package instance_info

import (
	"database/sql/driver"
	"fmt"
)

type Capability struct {
	v capabilityEnum
}

var (
	CapabilityNone   = Capability{capabilityNone}
	CapabilitySemdex = Capability{capabilitySemdex}
)

func (r Capability) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, r.v)
	case 'q':
		fmt.Fprintf(f, "%q", r.String())
	default:
		fmt.Fprint(f, r.v)
	}
}
func (r Capability) String() string {
	return string(r.v)
}
func (r Capability) MarshalText() ([]byte, error) {
	return []byte(r.v), nil
}
func (r *Capability) UnmarshalText(__iNpUt__ []byte) error {
	s, err := NewCapability(string(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func (r Capability) Value() (driver.Value, error) {
	return r.v, nil
}
func (r *Capability) Scan(__iNpUt__ any) error {
	s, err := NewCapability(fmt.Sprint(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func NewCapability(__iNpUt__ string) (Capability, error) {
	switch __iNpUt__ {
	case string(capabilityNone):
		return CapabilityNone, nil
	case string(capabilitySemdex):
		return CapabilitySemdex, nil
	default:
		return Capability{}, fmt.Errorf("invalid value for type 'Capability': '%s'", __iNpUt__)
	}
}