// Code generated by reactGen. DO NOT EDIT.

package react

// TableProps are the props for a <table> component
type TableProps struct {
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

func (t *TableProps) assign(v *_TableProps) {

	v.ClassName = t.ClassName

	v.DangerouslySetInnerHTML = t.DangerouslySetInnerHTML

	if t.DataSet != nil {
		for dk, dv := range t.DataSet {
			v.o.Set("data-"+dk, dv)
		}
	}

	if t.ID != "" {
		v.ID = t.ID
	}

	if t.Key != "" {
		v.Key = t.Key
	}

	if t.OnChange != nil {
		v.o.Set("onChange", t.OnChange.OnChange)
	}

	if t.OnClick != nil {
		v.o.Set("onClick", t.OnClick.OnClick)
	}

	if t.Ref != nil {
		v.o.Set("ref", t.Ref.Ref)
	}

	v.Role = t.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = t.Style.hack()

}