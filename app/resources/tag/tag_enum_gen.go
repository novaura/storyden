// Code generated by enumerator. DO NOT EDIT.

package tag

import (
	"database/sql/driver"
	"fmt"
)

type TagFillRule struct {
	v tagFillRuleEnum
}

var (
	TagFillRuleEnumNone = TagFillRule{tagFillRuleEnumNone}
	TagFillRuleQuery    = TagFillRule{tagFillRuleQuery}
	TagFillRuleReplace  = TagFillRule{tagFillRuleReplace}
)

func (r TagFillRule) Format(f fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprint(f, r.v)
	case 'q':
		fmt.Fprintf(f, "%q", r.String())
	default:
		fmt.Fprint(f, r.v)
	}
}
func (r TagFillRule) String() string {
	return string(r.v)
}
func (r TagFillRule) MarshalText() ([]byte, error) {
	return []byte(r.v), nil
}
func (r *TagFillRule) UnmarshalText(__iNpUt__ []byte) error {
	s, err := NewTagFillRule(string(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func (r TagFillRule) Value() (driver.Value, error) {
	return r.v, nil
}
func (r *TagFillRule) Scan(__iNpUt__ any) error {
	s, err := NewTagFillRule(fmt.Sprint(__iNpUt__))
	if err != nil {
		return err
	}
	*r = s
	return nil
}
func NewTagFillRule(__iNpUt__ string) (TagFillRule, error) {
	switch __iNpUt__ {
	case string(tagFillRuleEnumNone):
		return TagFillRuleEnumNone, nil
	case string(tagFillRuleQuery):
		return TagFillRuleQuery, nil
	case string(tagFillRuleReplace):
		return TagFillRuleReplace, nil
	default:
		return TagFillRule{}, fmt.Errorf("invalid value for type 'TagFillRule': '%s'", __iNpUt__)
	}
}