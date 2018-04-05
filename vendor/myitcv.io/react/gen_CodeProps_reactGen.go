// Code generated by reactGen. DO NOT EDIT.

package react

// CodeProps defines the properties for the <code> element
type CodeProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
}

func (c *CodeProps) assign(v *_CodeProps) {

	v.ClassName = c.ClassName

	v.DangerouslySetInnerHTML = c.DangerouslySetInnerHTML

	if c.DataSet != nil {
		for dk, dv := range c.DataSet {
			v.o.Set("data-"+dk, dv)
		}
	}

	if c.ID != "" {
		v.ID = c.ID
	}

	if c.Key != "" {
		v.Key = c.Key
	}

	if c.OnChange != nil {
		v.o.Set("onChange", c.OnChange.OnChange)
	}

	if c.OnClick != nil {
		v.o.Set("onClick", c.OnClick.OnClick)
	}

	if c.Ref != nil {
		v.o.Set("ref", c.Ref.Ref)
	}

	v.Role = c.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = c.Style.hack()

}