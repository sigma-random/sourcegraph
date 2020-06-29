package graphqlbackend

import (
	"context"

	"github.com/sourcegraph/sourcegraph/internal/search/query"
)

func (*schemaResolver) ParseSearchQuery(ctx context.Context, args *struct {
	Query       string
	PatternType string
}) (*JSONValue, error) {
	var searchType query.SearchType
	switch args.PatternType {
	case "literal":
		searchType = query.SearchTypeLiteral
	case "structural":
		searchType = query.SearchTypeStructural
	case "regexp", "regex":
		searchType = query.SearchTypeRegex
	default:
		searchType = query.SearchTypeLiteral
	}
	queryInfo, err := query.ProcessAndOr(args.Query, searchType)
	if err != nil {
		return nil, err
	}
	json, err := queryInfo.JSON()
	if err != nil {
		return nil, err
	}
	return &JSONValue{Value: string(json)}, nil
}
