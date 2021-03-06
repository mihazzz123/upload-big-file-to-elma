// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package money

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

func easyjson8d6bd286DecodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(in *jlexer.Lexer, out *Money) {
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
		case "currency":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.C).UnmarshalJSON(data))
			}
		case "cents":
			out.V = int64(in.Int64())
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
func easyjson8d6bd286EncodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(out *jwriter.Writer, in Money) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"currency\":"
		out.RawString(prefix[1:])
		out.Raw((in.C).MarshalJSON())
	}
	{
		const prefix string = ",\"cents\":"
		out.RawString(prefix)
		out.Int64(int64(in.V))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Money) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8d6bd286EncodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Money) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8d6bd286EncodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Money) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8d6bd286DecodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Money) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8d6bd286DecodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(l, v)
}
