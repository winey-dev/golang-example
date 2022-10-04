package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

func (Person) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "person"},
	}
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32) primary key",
		}).StorageKey("user_id"),

		field.String("first_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32) not null",
		}),

		field.String("second_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32) not null",
		}),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return nil
}
