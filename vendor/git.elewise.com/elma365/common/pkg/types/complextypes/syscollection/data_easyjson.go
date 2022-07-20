// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package syscollection

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

func easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesSyscollection(in *jlexer.Lexer, out *Data) {
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
		case "namespace":
			out.Namespace = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "linkedFieldCode":
			out.LinkedFieldCode = string(in.String())
		case "isDependent":
			out.IsDependent = bool(in.Bool())
		case "bindings":
			if in.IsNull() {
				in.Skip()
				out.Bindings = nil
			} else {
				in.Delim('[')
				if out.Bindings == nil {
					if !in.IsDelim(']') {
						out.Bindings = make([]json.RawMessage, 0, 2)
					} else {
						out.Bindings = []json.RawMessage{}
					}
				} else {
					out.Bindings = (out.Bindings)[:0]
				}
				for !in.IsDelim(']') {
					var v1 json.RawMessage
					if data := in.Raw(); in.Ok() {
						in.AddError((v1).UnmarshalJSON(data))
					}
					out.Bindings = append(out.Bindings, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "filter":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Filter).UnmarshalJSON(data))
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
func easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesSyscollection(out *jwriter.Writer, in Data) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"namespace\":"
		out.RawString(prefix[1:])
		out.String(string(in.Namespace))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"linkedFieldCode\":"
		out.RawString(prefix)
		out.String(string(in.LinkedFieldCode))
	}
	{
		const prefix string = ",\"isDependent\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsDependent))
	}
	{
		const prefix string = ",\"bindings\":"
		out.RawString(prefix)
		if in.Bindings == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Bindings {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Raw((v3).MarshalJSON())
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"filter\":"
		out.RawString(prefix)
		out.Raw((in.Filter).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Data) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesSyscollection(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Data) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGitElewiseComElma365CommonPkgTypesComplextypesSyscollection(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Data) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesSyscollection(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Data) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGitElewiseComElma365CommonPkgTypesComplextypesSyscollection(l, v)
}
