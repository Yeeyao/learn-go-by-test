package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode"
)

/*

func Join(a []string, sep string) string 在字符串数组中两个字符串之间插入指定的字符串内容
func LastIndexAny(s, chars string) int 在chars中找存在于s中的最后出现的字符串的位置
func LastIndexByte(s string, c byte) int 同上，不同的数据类型
func LastIndexFunc(s string, f func(rune) bool) int 同上，增加函数判断
func Map(mapping func(rune) rune, s string) string
func SplitAfterN(s, sep string, n int) []string 指定split的数量
func SplitN(s, sep string, n int) []string 指定split数量
func Title(s string) string 直接将字符串内容的每个单词首字母变成Unicode Title
func ToLower(s string) string 将字符串的Unicode字母都变成小写
func ToLowerSpecial(c unicode.SpecialCase, s string) string 根据第一个条件来将特定的字母转换为小写
func ToTitleSpecial(c unicode.SpecialCase, s string) string 同上，只是变成了title
func ToUpper(s string) string 变成大写
func ToUpperSpecial(c unicode.SpecialCase, s string) string 有条件的变成大写
func TrimLeftFunc(s string, f func(rune) bool) string 同trimleft只是使用func判断cutset中的字符
func TrimRightFunc(s string, f func(rune) bool) string 同trimright只是使用func判断cutset中的字符

*/

/*

type Builder 使用写方法来高效构建一个字符串，最小化内存复制，零值可被使用，不要复制一个非零 Builder

type Reader 通过读取一个字符串实现 io.Reader 等接口，Reader 的零值行为时 Reader 读取一个空字符串

type Replacer Replacer 用 replacements 来替换掉一个字符串列表，被多个 go routine 并发使用时是安全的

*/

func StrFun(a, b string) bool {
	return strings.Contains(a, b)
}

func FieldsFun(a string) []string {
	return strings.Fields(a)
}

// 用chars中的每个去遍历s中的每一个字符然后判断
func IndexAnyFun(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// substr出现在s中则返回最后出现在s字符串的字符位置
// substr未出现则返回-1
func LastIndexFun(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// 对s中每个满足mapping(rune)的字符替换为mapping(rune)的返回值
// 如果mapping(rune)返回值为负数，则从返回的字符串中删除对应的字符
func MapFun(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

// 重复字符 会判断count非负以及产生的字符串最大长度问题
func RepeatFun(s string, count int) string {
	return strings.Repeat(s, count)
}

// 替换字符 n 为-1则默认是全部替换，同时，从s中找到匹配的才进行替换
// old为空，则在返回的字符串中的所有字符后添加一个new
func ReplaceFun(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// s全部替换
func ReplaceAllFun(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// 返回的字符串数组中的字符串将删除sep
// 进行字符串的拆分，如果s不含有sep且sep非空则直接返回包含s的数组
// 如果sep为空，则将s每个都拆分 如果s和sep都为空则返回空的slice
func SplitFun(s, sep string) []string {
	return strings.Split(s, sep)
}

// 在sep之后切分
// 返回的字符串数组中的字符串保留sep
// 其余条件同上
func SplitAfterFun(s, sep string) []string {
	return strings.SplitAfter(s, sep)
}

// 下面的删除cutset中的字符包括空格等
// 删除cutset间的字符 如果只有开头则只删除开头，之后结尾则只删除结尾
// 输入的cutset不含结尾和开头则不处理 同时，中间不能有其他字符
func TrimFun(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

// 从cutset中选择满足条件的rune
func TrimFuncFun(s string, f func(rune) bool) string {
	return strings.TrimFunc(s, f)
}

// 从左边开始删除存在于cutset的内容
func TrimLeftFun(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// 去除开头，没有匹配则不处理
func TrimPrefixFun(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// 从右边开始删除存在于cutset的内容
func TrimRightFun(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// 删除所有的空格
func TrimSpaceFun(s string) string {
	return strings.TrimSpace(s)
}

// 删除后缀
func TrimSuffixFun(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

func TestStr(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want bool) {
		// 这里告诉测试套件这个方法是辅助函数，测试失败时报告的行号是函数调用内而不是辅助函数中
		t.Helper()
		if got != want {
			t.Errorf("got '%t' want '%t'", got, want)
		}
	}

	t.Run("contains test1", func(t *testing.T) {
		got := StrFun("b", "b")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("contains test2", func(t *testing.T) {
		got := StrFun("c", "b")
		want := false
		assertCorrectMessage(t, got, want)
	})

	t.Run("contains test3", func(t *testing.T) {
		got := StrFun("cb", "")
		want := true
		assertCorrectMessage(t, got, want)
	})

	t.Run("fields test", func(t *testing.T) {
		got := FieldsFun("foo bar baz")
		want := []string{"foo", "bar", "baz"}
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("index any test", func(t *testing.T) {
		got := IndexAnyFun("chicken", "nai")
		want := 2
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("last index test", func(t *testing.T) {
		got := LastIndexFun("go gopher", "hereeea")
		want := -1
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("map test", func(t *testing.T) {
		rot13 := func(r rune) rune {
			switch {
			case r >= 'A' && r <= 'Z':
				return 'A' + (r-'A'+13)%26
			case r >= 'a' && r <= 'z':
				return 'a' + (r-'a'+13)%26
			}
			return r
		}
		const testStr = "'Twas brillig and the slithy gopher..."
		got := MapFun(rot13, testStr)
		want := "'Gjnf oevyyvt naq gur fyvgul tbcure..."
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

	})

	t.Run("repeat test", func(t *testing.T) {
		got := RepeatFun("go gopher ", 3)
		want := "go gopher go gopher go gopher "
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("replace test", func(t *testing.T) {
		got := ReplaceFun("eesdfgo gopher", "go", "ha", -1)
		want := "eesdfha hapher"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("replace all test", func(t *testing.T) {
		got := ReplaceAllFun("eesdfgo gopher", "go", "ha")
		want := "eesdfha hapher"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("split test", func(t *testing.T) {
		got := SplitFun("eesdfgo gopher", "go")
		want := []string{"eesdf", " ", "pher"}
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("split after test", func(t *testing.T) {
		got := SplitAfterFun("eesdfgo gopher", "go")
		want := []string{"eesdfgo", " go", "pher"}
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("trim test", func(t *testing.T) {
		got := TrimFun("###eesdfgo gopher!!", "#!")
		want := "eesdfgo gopher"
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("trim func test", func(t *testing.T) {
		cutsetFun := func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		}
		got := TrimFuncFun("###eesdfgo gopher!!", cutsetFun)
		want := "eesdfgo gopher"
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("trim left test", func(t *testing.T) {
		got := TrimLeftFun("###eesdfgo gopher!!", "!#")
		want := "eesdfgo gopher!!"
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("trim prefix test", func(t *testing.T) {
		got := TrimPrefixFun("###eesdfgo gopher!!", "###eesd")
		fmt.Println(got)
		want := "fgo gopher!!"
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("trim right test", func(t *testing.T) {
		got := TrimRightFun("###eesdfgo gopher\r\t\n!!", "!")
		want := "###eesdfgo gopher\r\t\n"
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("trim space test", func(t *testing.T) {
		got := TrimSpaceFun("###eesdfgo gopher!!\n\t\r\n")
		want := "###eesdfgo gopher!!"
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("trim suffix test", func(t *testing.T) {
		got := TrimSuffixFun("###eesdfgo gopher!!\n\t\r\n", "\n\t\r\n")
		want := "###eesdfgo gopher!!"
		// deep equal or compare one by one
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

}
