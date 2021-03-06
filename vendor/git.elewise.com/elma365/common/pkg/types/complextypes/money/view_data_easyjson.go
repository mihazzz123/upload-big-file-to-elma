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

func easyjson8a9057a0DecodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(in *jlexer.Lexer, out *ViewData) {
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
			out.Currency = string(in.String())
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
func easyjson8a9057a0EncodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(out *jwriter.Writer, in ViewData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"currency\":"
		out.RawString(prefix[1:])
		out.String(string(in.Currency))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ViewData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a9057a0EncodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ViewData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a9057a0EncodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ViewData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a9057a0DecodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ViewData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a9057a0DecodeGitElewiseComElma365CommonPkgTypesComplextypesMoney(l, v)
}
