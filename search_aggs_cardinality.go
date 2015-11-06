// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://quorumsco.mit-license.org/license.txt for details.

package elastic

// CardinalityAggregation is a single-value metrics aggregation that
// calculates an approximate count of distinct values.
// Values can be extracted either from specific fields in the document
// or generated by a script.
// See: http://www.elasticsearch.org/guide/en/elasticsearch/reference/current/search-aggregations-metrics-cardinality-aggregation.html
type CardinalityAggregation struct {
	field              string
	script             string
	scriptFile         string
	lang               string
	format             string
	params             map[string]interface{}
	subAggregations    map[string]Aggregation
	precisionThreshold *int64
	rehash             *bool
}

func NewCardinalityAggregation() CardinalityAggregation {
	a := CardinalityAggregation{
		params:          make(map[string]interface{}),
		subAggregations: make(map[string]Aggregation),
	}
	return a
}

func (a CardinalityAggregation) Field(field string) CardinalityAggregation {
	a.field = field
	return a
}

func (a CardinalityAggregation) Script(script string) CardinalityAggregation {
	a.script = script
	return a
}

func (a CardinalityAggregation) ScriptFile(scriptFile string) CardinalityAggregation {
	a.scriptFile = scriptFile
	return a
}

func (a CardinalityAggregation) Lang(lang string) CardinalityAggregation {
	a.lang = lang
	return a
}

func (a CardinalityAggregation) Format(format string) CardinalityAggregation {
	a.format = format
	return a
}

func (a CardinalityAggregation) Param(name string, value interface{}) CardinalityAggregation {
	a.params[name] = value
	return a
}

func (a CardinalityAggregation) SubAggregation(name string, subAggregation Aggregation) CardinalityAggregation {
	a.subAggregations[name] = subAggregation
	return a
}

func (a CardinalityAggregation) PrecisionThreshold(threshold int64) CardinalityAggregation {
	a.precisionThreshold = &threshold
	return a
}

func (a CardinalityAggregation) Rehash(rehash bool) CardinalityAggregation {
	a.rehash = &rehash
	return a
}

func (a CardinalityAggregation) Source() interface{} {
	// Example:
	//	{
	//    "aggs" : {
	//      "author_count" : {
	//        "cardinality" : { "field" : "author" }
	//      }
	//    }
	//	}
	// This method returns only the "cardinality" : { "field" : "author" } part.

	source := make(map[string]interface{})
	opts := make(map[string]interface{})
	source["cardinality"] = opts

	// ValuesSourceAggregationBuilder
	if a.field != "" {
		opts["field"] = a.field
	}
	if a.script != "" {
		opts["script"] = a.script
	}
	if a.scriptFile != "" {
		opts["script_file"] = a.scriptFile
	}
	if a.lang != "" {
		opts["lang"] = a.lang
	}
	if a.format != "" {
		opts["format"] = a.format
	}
	if len(a.params) > 0 {
		opts["params"] = a.params
	}
	if a.precisionThreshold != nil {
		opts["precision_threshold"] = *a.precisionThreshold
	}
	if a.rehash != nil {
		opts["rehash"] = *a.rehash
	}

	// AggregationBuilder (SubAggregations)
	if len(a.subAggregations) > 0 {
		aggsMap := make(map[string]interface{})
		source["aggregations"] = aggsMap
		for name, aggregate := range a.subAggregations {
			aggsMap[name] = aggregate.Source()
		}
	}

	return source
}
