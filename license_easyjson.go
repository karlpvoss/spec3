// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package spec3

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

func easyjson967143c7DecodeGithubComGoOpenapiSpec3(in *jlexer.Lexer, out *License) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "url":
			out.URL = string(in.String())
		case "extensions":
			(out.Extensions).UnmarshalEasyJSON(in)
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
func easyjson967143c7EncodeGithubComGoOpenapiSpec3(out *jwriter.Writer, in License) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	if true {
		const prefix string = ",\"extensions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Extensions).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v License) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson967143c7EncodeGithubComGoOpenapiSpec3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v License) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson967143c7EncodeGithubComGoOpenapiSpec3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *License) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson967143c7DecodeGithubComGoOpenapiSpec3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *License) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson967143c7DecodeGithubComGoOpenapiSpec3(l, v)
}
