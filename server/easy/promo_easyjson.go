// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
	sql "database/sql"
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

func easyjson47107b8bDecodeServerEasy(in *jlexer.Lexer, out *RespPromo) {
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
		case "Type":
			out.Type = int(in.Int())
		case "Discount":
			out.Discount = uint(in.Uint())
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
func easyjson47107b8bEncodeServerEasy(out *jwriter.Writer, in RespPromo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"Discount\":"
		out.RawString(prefix)
		out.Uint(uint(in.Discount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RespPromo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson47107b8bEncodeServerEasy(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RespPromo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson47107b8bEncodeServerEasy(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RespPromo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson47107b8bDecodeServerEasy(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RespPromo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson47107b8bDecodeServerEasy(l, v)
}
func easyjson47107b8bDecodeServerEasy1(in *jlexer.Lexer, out *DBGetPromo) {
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
		case "ID":
			out.ID = uint(in.Uint())
		case "Code":
			out.Code = string(in.String())
		case "PromoType":
			out.PromoType = int(in.Int())
		case "Sale":
			easyjson47107b8bDecodeDatabaseSql(in, &out.Sale)
		case "RestaurantId":
			easyjson47107b8bDecodeDatabaseSql(in, &out.RestaurantId)
		case "ActiveFrom":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ActiveFrom).UnmarshalJSON(data))
			}
		case "ActiveTo":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.ActiveTo).UnmarshalJSON(data))
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
func easyjson47107b8bEncodeServerEasy1(out *jwriter.Writer, in DBGetPromo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"PromoType\":"
		out.RawString(prefix)
		out.Int(int(in.PromoType))
	}
	{
		const prefix string = ",\"Sale\":"
		out.RawString(prefix)
		easyjson47107b8bEncodeDatabaseSql(out, in.Sale)
	}
	{
		const prefix string = ",\"RestaurantId\":"
		out.RawString(prefix)
		easyjson47107b8bEncodeDatabaseSql(out, in.RestaurantId)
	}
	{
		const prefix string = ",\"ActiveFrom\":"
		out.RawString(prefix)
		out.Raw((in.ActiveFrom).MarshalJSON())
	}
	{
		const prefix string = ",\"ActiveTo\":"
		out.RawString(prefix)
		out.Raw((in.ActiveTo).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DBGetPromo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson47107b8bEncodeServerEasy1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DBGetPromo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson47107b8bEncodeServerEasy1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DBGetPromo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson47107b8bDecodeServerEasy1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DBGetPromo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson47107b8bDecodeServerEasy1(l, v)
}
func easyjson47107b8bDecodeDatabaseSql(in *jlexer.Lexer, out *sql.NullString) {
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
		case "String":
			out.String = string(in.String())
		case "Valid":
			out.Valid = bool(in.Bool())
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
func easyjson47107b8bEncodeDatabaseSql(out *jwriter.Writer, in sql.NullString) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"String\":"
		out.RawString(prefix[1:])
		out.String(string(in.String))
	}
	{
		const prefix string = ",\"Valid\":"
		out.RawString(prefix)
		out.Bool(bool(in.Valid))
	}
	out.RawByte('}')
}
