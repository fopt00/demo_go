package reflect

import (
	"bytes"
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	age     int32
	friends []*Person
}

func JsonEntry() {
	zoro := Person{
		Name: "zoro",
		age:  10,
	}
	luffy := Person{
		Name: "luffy",
		age:  18,
		friends: []*Person{
			&zoro,
		},
	}
	allen := Person{
		Name: "allen",
		age:  19,
		friends: []*Person{
			&luffy,
			&zoro,
		},
	}

	// 我们的目标是将这个 map 序列化成 json 串
	persons := map[string]Person{
		"allen": allen,
		"luffy": luffy,
		"zoro":  zoro,
	}

	// Parse 就是我们要编写的函数
	buf, err := Parse(persons)
	if err != nil {
		fmt.Printf("%v\n", err)
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

// parse 是一个内置的递归函数
func parse(buf *bytes.Buffer, v reflect.Value) error {
	// 取到变量的类型的类别
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, `"%s"`, v.String())
	case reflect.Ptr:
		// 如果是指针，则解引用后继续解析
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
			// key 为结构体字段名称
			fmt.Fprintf(buf, `"%s":`, v.Type().Field(i).Name)
			// value 需要递归解析，因为 value 也需要序列化成 json 串
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
			fmt.Fprintf(buf, `"%s":`, key)
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