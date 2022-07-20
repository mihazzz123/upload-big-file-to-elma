// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package enum

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

func easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesEnum(in *jlexer.Lexer, out *EnumItem) {
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
		case "name":
			out.Name = string(in.String())
		case "checked":
			out.Checked = bool(in.Bool())
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
func easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesEnum(out *jwriter.Writer, in EnumItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix[1:])
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"checked\":"
		out.RawString(prefix)
		out.Bool(bool(in.Checked))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnumItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesEnum(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnumItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesEnum(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnumItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesEnum(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnumItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesEnum(l, v)
}
func easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesEnum1(in *jlexer.Lexer, out *EnumData) {
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
		case "enumItems":
			if in.IsNull() {
				in.Skip()
				out.EnumItems = nil
			} else {
				in.Delim('[')
				if out.EnumItems == nil {
					if !in.IsDelim(']') {
						out.EnumItems = make([]EnumItem, 0, 1)
					} else {
						out.EnumItems = []EnumItem{}
					}
				} else {
					out.EnumItems = (out.EnumItems)[:0]
				}
				for !in.IsDelim(']') {
					var v1 EnumItem
					(v1).UnmarshalEasyJSON(in)
					out.EnumItems = append(out.EnumItems, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesEnum1(out *jwriter.Writer, in EnumData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"enumItems\":"
		out.RawString(prefix[1:])
		if in.EnumItems == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.EnumItems {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EnumData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesEnum1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EnumData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesEnum1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EnumData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesEnum1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EnumData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesEnum1(l, v)
}