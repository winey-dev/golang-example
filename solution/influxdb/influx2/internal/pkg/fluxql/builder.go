package fluxql

import (
	"strings"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	indent     = " "
)

type Builder interface {
	From(bucket string) Builder
	Range(start, end time.Time) Builder
	RangeString(startTime string) Builder
	Filter(k, v string) Builder
	Filters(kv []FilterKeyValue) Builder
	KeepColumns(cols []string) Builder
	Measurement(measurement string) Builder
	Field(fieldName string) Builder
	Builder() Object
}

type builder struct {
	start       time.Time
	startString string
	end         time.Time
	from        string
	measurement string
	fieldName   string
	filter      map[string]string
	filters     [][]FilterKeyValue
	keepColumns []string
}

func NewBuilder() Builder {
	b := new(builder)
	b.filter = make(map[string]string)
	return b
}

func (b *builder) From(bucket string) Builder {
	b.from = bucket
	return b
}

func (b *builder) Range(start, end time.Time) Builder {
	b.start = start
	b.end = end
	return b
}

func (b *builder) RangeString(startTime string) Builder {
	b.startString = startTime
	return b
}

func (b *builder) Filter(k, v string) Builder {
	if k == "__measurement" {
		b.measurement = v
	} else if k == "__field" {
		b.fieldName = v
	} else {
		b.filter[k] = v
	}
	return b
}

func (b *builder) Filters(kv []FilterKeyValue) Builder {
	b.filters = append(b.filters, kv)
	return b
}

func (b *builder) KeepColumns(cols []string) Builder {
	b.keepColumns = append(b.keepColumns, cols...)
	return b
}

func (b *builder) Measurement(measurement string) Builder {
	b.measurement = measurement
	return b
}

func (b *builder) Field(fieldName string) Builder {
	b.fieldName = fieldName
	return b
}

func (b *builder) Builder() Object {
	o := new(object)
	// from start: from (bucket: "bucket")
	o.b.WriteString("from(bucket: \"")
	o.b.WriteString(b.from)
	o.b.WriteString("\")\n")
	// from end

	// range start
	o.b.WriteString(indent)
	o.b.WriteString("|> range(")
	o.b.WriteString("start: ")

	if b.startString != "" {
		o.b.WriteString(b.startString)
	} else if !b.start.IsZero() {
		o.b.WriteString(b.start.Format(time.RFC3339))

		if b.end.IsZero() {
			o.b.WriteString(", ")
			o.b.WriteString("stop: now()")

		} else {
			o.b.WriteString(", ")
			o.b.WriteString("stop: ")
			o.b.WriteString(b.end.Format(time.RFC3339))
		}
	}
	o.b.WriteString(")\n")

	// range end

	if b.measurement != "" {
		o.b.WriteString(indent)
		o.b.WriteString("|> filter(fn: (r) => r._measurement == \"")
		o.b.WriteString(b.measurement)
		o.b.WriteString("\")\n")
	}

	if b.fieldName != "" {
		o.b.WriteString(indent)
		o.b.WriteString("|> filter(fn: (r) => r._field == \"")
		o.b.WriteString(b.fieldName)
		o.b.WriteString("\")\n")
	}

	for k, v := range b.filter {
		o.b.WriteString(indent)
		o.b.WriteString("|> filter(fn: (r) => r.")
		o.b.WriteString(k)
		o.b.WriteString(" == \"")
		o.b.WriteString(v)
		o.b.WriteString("\")\n")
	}

	// |> filter(fn: (r) => r.tag == ems and r.pod == pod-1)
	for _, filters := range b.filters {
		o.b.WriteString(indent)
		o.b.WriteString("|> filter(fn: (r) => ")
		for _, filter := range filters {
			o.b.WriteString("r.")
			o.b.WriteString(filter.Key)
			o.b.WriteString(" == \"")
			o.b.WriteString(filter.Value)
			o.b.WriteString("\" ")
			o.b.WriteString(filter.Op.String())
		}
		o.b.WriteString(")\n")
	}

	if len(b.keepColumns) != 0 {
		o.b.WriteString(indent)
		o.b.WriteString("|> keep(columns: [")
		for i := 0; i < len(b.keepColumns); i++ {
			o.b.WriteString("\"")
			o.b.WriteString(b.keepColumns[i])
			o.b.WriteString("\"")
			if i < len(b.keepColumns)-1 {
				o.b.WriteString(",")
			}
		}
		o.b.WriteString("]")
	}

	/*
		o.b.WriteString(indent)
		o.b.WriteString("|> yield(name: \"mean\")")
	*/

	return o
}

type Object interface {
	String() string
}

type object struct {
	b strings.Builder
}

func (o *object) String() string {
	return o.b.String()
}

type FilterKeyValue struct {
	Key   string
	Value string
	Op    Operator
}

type Operator int32

const (
	END Operator = 0
	OR  Operator = 1
	AND Operator = 2
)

func (o Operator) String() string {
	if o == OR {
		return "or "
	} else if o == AND {
		return "and "
	}
	return ""
}
