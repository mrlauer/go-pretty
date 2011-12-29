/*
Pretty-printer for go values.
Defines one function, Pretty
*/
package pretty

import(
    "fmt"
    "reflect"
    "strings"
)

func addIndent(s string, indent string) string {
    return strings.Replace(s, "\n", "\n" + indent, -1)
}

func getInterfaceDammit(v reflect.Value) interface{} {
    if v.CanInterface() {
        return v.Interface()
    }
    switch v.Kind() {
    case reflect.Bool:
        return v.Bool()
    case reflect.Int:
        return v.Int()
    case reflect.Int8:
        return v.Int()
    case reflect.Int16:
        return v.Int()
    case reflect.Int32:
        return v.Int()
    case reflect.Int64:
        return v.Int()
    case reflect.Uint:
        return v.Uint()
    case reflect.Uint8:
        return v.Uint()
    case reflect.Uint16:
        return v.Uint()
    case reflect.Uint32:
        return v.Uint()
    case reflect.Uint64:
        return v.Uint()
    case reflect.Float32:
        return v.Float()
    case reflect.Float64:
        return v.Float()
    case reflect.Complex64:
        return v.Complex()
    case reflect.Complex128:
        return v.Complex()
    case reflect.String:
        return v.String()
    }
    return "?"
}

func pretty(s interface{}, prefix, indent string) string {
    v := reflect.ValueOf(s)

    var result string

    if str, ok := s.(fmt.Stringer); ok {
        return str.String()
    }

    switch v.Kind() {
    case reflect.Slice:
        fallthrough
    case reflect.Array:
        n := v.Len()
        result = fmt.Sprintf("[%d]%s[\n", n, v.Type().Elem())
        for i := 0; i<n; i++ {
            f := pretty(v.Index(i).Interface(), prefix, indent)
            result += indent + addIndent(f, indent) + "\n"
        }
        result += "]"
        return result
    case reflect.Ptr:
        return fmt.Sprintf("&%s", pretty(v.Elem().Interface(), prefix, indent))
    case reflect.Struct:
        n := v.NumField()
        result = fmt.Sprintf("%v{\n", v.Type())
        for i := 0; i<n; i++ {
            sf := v.Type().Field(i)
            f := pretty(getInterfaceDammit(v.Field(i)), prefix, indent)
            result += fmt.Sprintf("%s%s: %s\n", indent, sf.Name, addIndent(f, indent))
        }
        result += "}"
    case reflect.Map:
        result = fmt.Sprintf("map[%s]%s[\n", v.Type().Key(), v.Type().Elem())
        keys := v.MapKeys()
        for _, k := range keys {
            e := v.MapIndex(k)
            keyStr := addIndent(pretty(k, prefix, indent), indent)
            elemStr := addIndent(pretty(e, prefix, indent), indent)
            result += fmt.Sprintf("%s%s: %s\n", indent, keyStr, elemStr)
        }
        result += "]"
    default:
        result = fmt.Sprintf("%v", s)
    }
    return result
}

func Pretty(s interface{}, indent string) string {
    return pretty(s, "", indent)
}