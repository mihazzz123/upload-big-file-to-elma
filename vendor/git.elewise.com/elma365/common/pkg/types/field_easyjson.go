// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package types

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson30283592DecodeGitElewiseComElma365CommonPkgTypes(in *jlexer.Lexer, out *FieldView) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "sort":
			out.Sort = int(in.Int())
		case "tooltip":
			out.Tooltip = string(in.String())
		case "tooltipAsHtml":
			out.TooltipAsHTML = bool(in.Bool())
		case "system":
			out.System = bool(in.Bool())
		case "hidden":
			out.Hidden = bool(in.Bool())
		case "data":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Data).UnmarshalJSON(data))
			}
		case "disabled":
			out.Disabled = bool(in.Bool())
		case "input":
			if in.IsNull() {
				in.Skip()
				out.Input = nil
			} else {
				if out.Input == nil {
					out.Input = new(bool)
				}
				*out.Input = bool(in.Bool())
			}
		case "output":
			if in.IsNull() {
				in.Skip()
				out.Output = nil
			} else {
				if out.Output == nil {
					out.Output = new(bool)
				}
				*out.Output = bool(in.Bool())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson30283592EncodeGitElewiseComElma365CommonPkgTypes(out *jwriter.Writer, in FieldView) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Sort != 0 {
		const prefix string = ",\"sort\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Sort))
	}
	if in.Tooltip != "" {
		const prefix string = ",\"tooltip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Tooltip))
	}
	if in.TooltipAsHTML {
		const prefix string = ",\"tooltipAsHtml\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.TooltipAsHTML))
	}
	if in.System {
		const prefix string = ",\"system\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.System))
	}
	if in.Hidden {
		const prefix string = ",\"hidden\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Hidden))
	}
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Data).MarshalJSON())
	}
	if in.Disabled {
		const prefix string = ",\"disabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Disabled))
	}
	if in.Input != nil {
		const prefix string = ",\"input\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.Input))
	}
	if in.Output != nil {
		const prefix string = ",\"output\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.Output))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FieldView) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30283592EncodeGitElewiseComElma365CommonPkgTypes(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FieldView) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30283592EncodeGitElewiseComElma365CommonPkgTypes(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FieldView) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30283592DecodeGitElewiseComElma365CommonPkgTypes(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FieldView) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30283592DecodeGitElewiseComElma365CommonPkgTypes(l, v)
}
func easyjson30283592DecodeGitElewiseComElma365CommonPkgTypes1(in *jlexer.Lexer, out *Field) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "code":
			out.Code = string(in.String())
		case "type":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Type).UnmarshalJSON(data))
			}
		case "searchable":
			out.Searchable = bool(in.Bool())
		case "indexed":
			out.Indexed = bool(in.Bool())
		case "deleted":
			out.Deleted = bool(in.Bool())
		case "array":
			out.Array = bool(in.Bool())
		case "required":
			out.Required = bool(in.Bool())
		case "single":
			out.Single = bool(in.Bool())
		case "defaultValue":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Default).UnmarshalJSON(data))
			}
		case "calcByFormula":
			out.CalcByFormula = bool(in.Bool())
		case "formula":
			out.Formula = string(in.String())
		case "data":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Data).UnmarshalJSON(data))
			}
		case "view":
			(out.View).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson30283592EncodeGitElewiseComElma365CommonPkgTypes1(out *jwriter.Writer, in Field) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Raw((in.Type).MarshalJSON())
	}
	{
		const prefix string = ",\"searchable\":"
		out.RawString(prefix)
		out.Bool(bool(in.Searchable))
	}
	{
		const prefix string = ",\"indexed\":"
		out.RawString(prefix)
		out.Bool(bool(in.Indexed))
	}
	{
		const prefix string = ",\"deleted\":"
		out.RawString(prefix)
		out.Bool(bool(in.Deleted))
	}
	{
		const prefix string = ",\"array\":"
		out.RawString(prefix)
		out.Bool(bool(in.Array))
	}
	{
		const prefix string = ",\"required\":"
		out.RawString(prefix)
		out.Bool(bool(in.Required))
	}
	{
		const prefix string = ",\"single\":"
		out.RawString(prefix)
		out.Bool(bool(in.Single))
	}
	{
		const prefix string = ",\"defaultValue\":"
		out.RawString(prefix)
		out.Raw((in.Default).MarshalJSON())
	}
	{
		const prefix string = ",\"calcByFormula\":"
		out.RawString(prefix)
		out.Bool(bool(in.CalcByFormula))
	}
	{
		const prefix string = ",\"formula\":"
		out.RawString(prefix)
		out.String(string(in.Formula))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		out.Raw((in.Data).MarshalJSON())
	}
	{
		const prefix string = ",\"view\":"
		out.RawString(prefix)
		(in.View).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Field) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson30283592EncodeGitElewiseComElma365CommonPkgTypes1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Field) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson30283592EncodeGitElewiseComElma365CommonPkgTypes1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Field) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson30283592DecodeGitElewiseComElma365CommonPkgTypes1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Field) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson30283592DecodeGitElewiseComElma365CommonPkgTypes1(l, v)
}
