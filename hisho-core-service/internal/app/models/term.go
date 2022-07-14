package models

import "github.com/alexadastra/hisho/hisho-core-service/pkg/api"

// Term is a wrapper around Term Enum
type Term struct {
	Value    string
	ValueInt int64
}

// TermFromProto translates proto Enum to the internal struct
func TermFromProto(apiTerm *api.Term) *Term {
	return &Term{
		Value:    apiTerm.String(),
		ValueInt: int64(api.Term_value[apiTerm.String()]),
	}
}

// TermToProto translates internal struct into proto Enum
func TermToProto(term *Term) string {
	return api.Term_name[int32(term.ValueInt)]
}
