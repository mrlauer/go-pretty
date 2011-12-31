/*
Pretty-printer for go values.
Defines one function, 
    func Pretty(interface{}, indent) string
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
    }
    return v.String()
}

func pretty(v reflect.Value, indent string) string {

    var result string

    // If it's a Stringer, and we can get to it, use that
    if(v.CanInterface()) {
        if s, ok := v.Interface().(fmt.Stringer); ok {
            return s.String()
        }
    }

    switch v.Kind() {
    case reflect.Interface:
        if v.IsNil() {
            return "nil"
        }
        return pretty(v.Elem(), indent)
    case reflect.Slice:
        fallthrough
    case reflect.Array:
        n := v.Len()
        result = fmt.Sprintf("[%d]%s[\n", n, v.Type().Elem())
        for i := 0; i<n; i++ {
            f := pretty(v.Index(i), indent)
            result += indent + addIndent(f, indent) + "\n"
        }
        result += "]"
        return result
    case reflect.Ptr:
        if v.IsNil() {
            return "nil"
        }
        return fmt.Sprintf("&%s", pretty(v.Elem(), indent))
    case reflect.Struct:
        n := v.NumField()
        result = fmt.Sprintf("%v{\n", v.Type())
        for i := 0; i<n; i++ {
            sf := v.Type().Field(i)
            f := pretty(v.Field(i), indent)
            result += fmt.Sprintf("%s%s: %s\n", indent, sf.Name, addIndent(f, indent))
        }
        result += "}"
    case reflect.Map:
        result = fmt.Sprintf("map[%s]%s[\n", v.Type().Key(), v.Type().Elem())
        keys := v.MapKeys()
        for _, k := range keys {
            e := v.MapIndex(k)
            keyStr := addIndent(pretty(k, indent), indent)
            elemStr := addIndent(pretty(e, indent), indent)
            result += fmt.Sprintf("%s%s: %s\n", indent, keyStr, elemStr)
        }
        result += "]"
    default:
        result = fmt.Sprintf("%v", getInterfaceDammit(v))
    }
    return result
}

// Pretty-print a value
func Pretty(s interface{}, indent string) string {
    return pretty(reflect.ValueOf(s), indent)
}
