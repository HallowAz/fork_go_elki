// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package dto

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	entity "server/internal/domain/entity"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson16134a91DecodeServerEasy(in *jlexer.Lexer, out *StringSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(StringSlice, 0, 8)
			} else {
				*out = StringSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *string
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(string)
				}
				*v1 = string(in.String())
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson16134a91EncodeServerEasy(out *jwriter.Writer, in StringSlice) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				out.String(string(*v3))
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v StringSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeServerEasy(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StringSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeServerEasy(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StringSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeServerEasy(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StringSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeServerEasy(l, v)
}
func easyjson16134a91DecodeServerEasy1(in *jlexer.Lexer, out *RestaurantWithCategoriesSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(RestaurantWithCategoriesSlice, 0, 8)
			} else {
				*out = RestaurantWithCategoriesSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *RestaurantWithCategories
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(RestaurantWithCategories)
				}
				(*v4).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson16134a91EncodeServerEasy1(out *jwriter.Writer, in RestaurantWithCategoriesSlice) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v RestaurantWithCategoriesSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeServerEasy1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RestaurantWithCategoriesSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeServerEasy1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RestaurantWithCategoriesSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeServerEasy1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RestaurantWithCategoriesSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeServerEasy1(l, v)
}
func easyjson16134a91DecodeServerEasy2(in *jlexer.Lexer, out *RestaurantWithCategoriesAndProductsSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(RestaurantWithCategoriesAndProductsSlice, 0, 8)
			} else {
				*out = RestaurantWithCategoriesAndProductsSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 *RestaurantWithCategoriesAndProducts
			if in.IsNull() {
				in.Skip()
				v7 = nil
			} else {
				if v7 == nil {
					v7 = new(RestaurantWithCategoriesAndProducts)
				}
				(*v7).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson16134a91EncodeServerEasy2(out *jwriter.Writer, in RestaurantWithCategoriesAndProductsSlice) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			if v9 == nil {
				out.RawString("null")
			} else {
				(*v9).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v RestaurantWithCategoriesAndProductsSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeServerEasy2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RestaurantWithCategoriesAndProductsSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeServerEasy2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RestaurantWithCategoriesAndProductsSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeServerEasy2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RestaurantWithCategoriesAndProductsSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeServerEasy2(l, v)
}
func easyjson16134a91DecodeServerEasy3(in *jlexer.Lexer, out *RestaurantWithCategoriesAndProducts) {
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
		case "Name":
			out.Name = string(in.String())
		case "Rating":
			out.Rating = float32(in.Float32())
		case "CommentsCount":
			out.CommentsCount = int(in.Int())
		case "Icon":
			out.Icon = string(in.String())
		case "Categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Categories = append(out.Categories, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "MinDeliveryTime":
			out.MinDeliveryTime = int(in.Int())
		case "MaxDeliveryTime":
			out.MaxDeliveryTime = int(in.Int())
		case "DeliveryPrice":
			out.DeliveryPrice = int(in.Int())
		case "Products":
			if in.IsNull() {
				in.Skip()
				out.Products = nil
			} else {
				in.Delim('[')
				if out.Products == nil {
					if !in.IsDelim(']') {
						out.Products = make([]*entity.Product, 0, 8)
					} else {
						out.Products = []*entity.Product{}
					}
				} else {
					out.Products = (out.Products)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *entity.Product
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(entity.Product)
						}
						(*v11).UnmarshalEasyJSON(in)
					}
					out.Products = append(out.Products, v11)
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
func easyjson16134a91EncodeServerEasy3(out *jwriter.Writer, in RestaurantWithCategoriesAndProducts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"CommentsCount\":"
		out.RawString(prefix)
		out.Int(int(in.CommentsCount))
	}
	{
		const prefix string = ",\"Icon\":"
		out.RawString(prefix)
		out.String(string(in.Icon))
	}
	{
		const prefix string = ",\"Categories\":"
		out.RawString(prefix)
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.Categories {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"MinDeliveryTime\":"
		out.RawString(prefix)
		out.Int(int(in.MinDeliveryTime))
	}
	{
		const prefix string = ",\"MaxDeliveryTime\":"
		out.RawString(prefix)
		out.Int(int(in.MaxDeliveryTime))
	}
	{
		const prefix string = ",\"DeliveryPrice\":"
		out.RawString(prefix)
		out.Int(int(in.DeliveryPrice))
	}
	{
		const prefix string = ",\"Products\":"
		out.RawString(prefix)
		if in.Products == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Products {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RestaurantWithCategoriesAndProducts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeServerEasy3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RestaurantWithCategoriesAndProducts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeServerEasy3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RestaurantWithCategoriesAndProducts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeServerEasy3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RestaurantWithCategoriesAndProducts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeServerEasy3(l, v)
}
func easyjson16134a91DecodeServerEasy4(in *jlexer.Lexer, out *RestaurantWithCategories) {
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
		case "Name":
			out.Name = string(in.String())
		case "Rating":
			out.Rating = float32(in.Float32())
		case "CommentsCount":
			out.CommentsCount = int(in.Int())
		case "Icon":
			out.Icon = string(in.String())
		case "Categories":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.Categories = append(out.Categories, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "MinDeliveryTime":
			out.MinDeliveryTime = int(in.Int())
		case "MaxDeliveryTime":
			out.MaxDeliveryTime = int(in.Int())
		case "DeliveryPrice":
			out.DeliveryPrice = int(in.Int())
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
func easyjson16134a91EncodeServerEasy4(out *jwriter.Writer, in RestaurantWithCategories) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"CommentsCount\":"
		out.RawString(prefix)
		out.Int(int(in.CommentsCount))
	}
	{
		const prefix string = ",\"Icon\":"
		out.RawString(prefix)
		out.String(string(in.Icon))
	}
	{
		const prefix string = ",\"Categories\":"
		out.RawString(prefix)
		if in.Categories == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Categories {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"MinDeliveryTime\":"
		out.RawString(prefix)
		out.Int(int(in.MinDeliveryTime))
	}
	{
		const prefix string = ",\"MaxDeliveryTime\":"
		out.RawString(prefix)
		out.Int(int(in.MaxDeliveryTime))
	}
	{
		const prefix string = ",\"DeliveryPrice\":"
		out.RawString(prefix)
		out.Int(int(in.DeliveryPrice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RestaurantWithCategories) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeServerEasy4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RestaurantWithCategories) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeServerEasy4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RestaurantWithCategories) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeServerEasy4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RestaurantWithCategories) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeServerEasy4(l, v)
}
func easyjson16134a91DecodeServerEasy5(in *jlexer.Lexer, out *MenuTypeWithProductsSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(MenuTypeWithProductsSlice, 0, 8)
			} else {
				*out = MenuTypeWithProductsSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v19 *MenuTypeWithProducts
			if in.IsNull() {
				in.Skip()
				v19 = nil
			} else {
				if v19 == nil {
					v19 = new(MenuTypeWithProducts)
				}
				(*v19).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v19)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson16134a91EncodeServerEasy5(out *jwriter.Writer, in MenuTypeWithProductsSlice) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v20, v21 := range in {
			if v20 > 0 {
				out.RawByte(',')
			}
			if v21 == nil {
				out.RawString("null")
			} else {
				(*v21).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v MenuTypeWithProductsSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeServerEasy5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MenuTypeWithProductsSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeServerEasy5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MenuTypeWithProductsSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeServerEasy5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MenuTypeWithProductsSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeServerEasy5(l, v)
}
func easyjson16134a91DecodeServerEasy6(in *jlexer.Lexer, out *MenuTypeWithProducts) {
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
		case "MenuType":
			if in.IsNull() {
				in.Skip()
				out.MenuType = nil
			} else {
				if out.MenuType == nil {
					out.MenuType = new(entity.MenuType)
				}
				easyjson16134a91DecodeServerInternalDomainEntity(in, out.MenuType)
			}
		case "Products":
			if in.IsNull() {
				in.Skip()
				out.Products = nil
			} else {
				in.Delim('[')
				if out.Products == nil {
					if !in.IsDelim(']') {
						out.Products = make([]*entity.Product, 0, 8)
					} else {
						out.Products = []*entity.Product{}
					}
				} else {
					out.Products = (out.Products)[:0]
				}
				for !in.IsDelim(']') {
					var v22 *entity.Product
					if in.IsNull() {
						in.Skip()
						v22 = nil
					} else {
						if v22 == nil {
							v22 = new(entity.Product)
						}
						(*v22).UnmarshalEasyJSON(in)
					}
					out.Products = append(out.Products, v22)
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
func easyjson16134a91EncodeServerEasy6(out *jwriter.Writer, in MenuTypeWithProducts) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"MenuType\":"
		out.RawString(prefix[1:])
		if in.MenuType == nil {
			out.RawString("null")
		} else {
			easyjson16134a91EncodeServerInternalDomainEntity(out, *in.MenuType)
		}
	}
	{
		const prefix string = ",\"Products\":"
		out.RawString(prefix)
		if in.Products == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v23, v24 := range in.Products {
				if v23 > 0 {
					out.RawByte(',')
				}
				if v24 == nil {
					out.RawString("null")
				} else {
					(*v24).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MenuTypeWithProducts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson16134a91EncodeServerEasy6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MenuTypeWithProducts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson16134a91EncodeServerEasy6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MenuTypeWithProducts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson16134a91DecodeServerEasy6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MenuTypeWithProducts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson16134a91DecodeServerEasy6(l, v)
}
func easyjson16134a91DecodeServerInternalDomainEntity(in *jlexer.Lexer, out *entity.MenuType) {
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
		case "Name":
			out.Name = string(in.String())
		case "RestaurantID":
			out.RestaurantID = uint(in.Uint())
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
func easyjson16134a91EncodeServerInternalDomainEntity(out *jwriter.Writer, in entity.MenuType) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.ID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"RestaurantID\":"
		out.RawString(prefix)
		out.Uint(uint(in.RestaurantID))
	}
	out.RawByte('}')
}
