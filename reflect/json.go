package reflect

import (
	"bytes"
	"fmt"
	"reflect"
)

type JsonPerson struct {
	Name    string
	age     int32
	friends []*JsonPerson
}

func JsonEntry() {
	zoro := JsonPerson{
		Name: "zoro",
		age:  10,
	}

	luffy := JsonPerson{
		Name: "luffy",
		age:  20,
		friends: []*JsonPerson{
			&zoro,
		},
	}

	allen := JsonPerson{
		Name: "allen",
		age:  19,
		friends: []*JsonPerson{
			&zoro,
			&luffy,
		},
	}

	persons := map[string]JsonPerson{
		"zoro":  zoro,
		"luffy": luffy,
		"allen": allen,
	}

	buf, err := Parse(persons)
	if err != nil {
		fmt.Println("%v\n", err)
		return
	}
	fmt.Printf("%v\n", string(buf))
}

func Parse(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := parse(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func parse(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%s", v.String())
	case reflect.Ptr:
		return parse(buf, v.Elem())
	case reflect.Slice, reflect.Array:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := parse(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(']')
	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(buf, `"%s:"`, v.Type().Field(i).Name)
			if err := parse(buf, v.Field(i)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	case reflect.Map:
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(buf, "%s", key)
			if err := parse(buf, v.MapIndex(key)); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}

	return nil
}
