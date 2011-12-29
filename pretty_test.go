package pretty

import(
    "testing"
    "reflect"
    "fmt"
)

func testBody(s interface{}, expected string, t *testing.T) {
    result := Pretty(s, " ")
    if result != expected {
        t.Errorf("Bad string:\n%s %v\nproduced\n%s\nrather than\n%s\n", 
            reflect.TypeOf(s), s, result, expected)
    }
}

func TestString(t *testing.T) {
    testBody("This is a test", "This is a test", t)
}

func TestString2(t *testing.T) {
    testBody("This is\na test", "This is\na test", t)
}

func TestInt(t *testing.T) {
    testBody(42, "42", t)
}

func TestArray(t *testing.T) {
    a := [3]int{4, 5, 6}
    expected := `[3]int[
 4
 5
 6
]`
    testBody(a, expected, t)
}

func TestSlice(t *testing.T) {
    a := [3]int{4, 5, 6}
    expected := `[3]int[
 4
 5
 6
]`
testBody(a[:], expected, t)
}

type StringerStruct struct {
    AField  int
}

func (s StringerStruct)String() string {
    return fmt.Sprintf("StringerStruct: %s", s.AField)
}

func TestStringer(t *testing.T) {
    s := StringerStruct{42}
    testBody(s, s.String(), t)
}

func TestPtr(t *testing.T) {
    s := 42
    testBody(&s, "&42", t)
}

type Nested struct {
    Something int
}

type TStruct struct {
    AField int
    aPrivateField int
    NestedStr Nested
    NestedPtr *Nested
    privateNested Nested
}

func TestStruct(t *testing.T) {
    n := Nested{4}
    s := TStruct{42, 43, Nested{3}, &n, Nested{5}}
    expected := `pretty.TStruct{
 AField: 42
 aPrivateField: 43
 NestedStr: pretty.Nested{
  Something: 3
 }
 NestedPtr: &pretty.Nested{
  Something: 4
 }
 privateNested: pretty.Nested{
  Something: 5
 }
}`

    testBody(s, expected, t)
}

func TestMap(t *testing.T) {
    m := map[string]string{ "foo" : "bar", "baz": "wox" }
    expected :=
`map[string]string[
 foo: bar
 baz: wox
]`
    testBody(m, expected, t)
}
