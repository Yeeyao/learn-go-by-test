package main

/*
	%v 值的默认格式
	%T 值的类型
	%t 布尔类型

	%b 二进制
	%c 字符类型
	%o 八进制
	%d 十进制

	%e %E 科学计数法的浮点数
	%f 十进制的浮点数，不含指数域

	%s 未解释的 string 或者 slice 的字节
	%q 一个双引号字符串，使用 Go 语法安全转义的

	%p Slice 第 0 个元素的地址

	%p 0x 开头的地址

*/

/*
	打印输出

	精度使用 Unicode code pointes 即 runes 来计算长度（C 语言使用 bytes 来计算长度）
	对绝大多数的值而言，宽度是输出的 runes 的最小数量，如果不足则将会使用空格来补齐

	对于 strings 以及 byte slices 和 byte arrays，精度限制了输入的格式，在必要时将会被截断
	一般使用 runes 来计算长度，但是对于格式是 %x %X 的，使用 bytes 来计算长度

	对于浮点数值，长度是所有位（包括十进制以及小数数位）的最大的长度，精度是十进制后的数量，除了 %g 以及 %G，尽量使用有效位的最大数量（结尾的0会被移除）
	%e %f %#g 的默认精度是 6，%g 则是针对每个值的所需的最小数量

	在格式字母前的宽度是可选的，如果不指定，则默认是足够表示打印的值的宽度

	%9.f 表示精度是0
	%9.2f 表示精度是2
	%9f 表示使用默认精度

	对于复数，长度和精度都是单独作用到两个部分的

	其他 flag: -: 在右边加空格 # , ,: 为符号位预留一个空格（% d） 0: 对数字填充 0 而不是空格

	对每个 Printf-like 函数，存在一个 Print 函数，不会使用任何格式，等同于对每个操作数都使用 %v 格式，另一个是 Println 函数

	对于 strings 以及 slice，有分别递归格式化元素以及格式化整个结构的区别

	显式参数索引 在 Printf, Sprintf 以及 Fprintf 中，默认对调用中的连续参数逐个使用格式化，但是，使用记号 [n] 可以指定针对第 n 个元素格式化
	同时，对于浮点数的打印，可以传递长度以及精度参数

*/

/*

	读取输入

	Scan,Scanf,Scanln 从 os.Stdin 中读取，Fscan, Fscanf 以及 Fscanln 从一个特殊的 io.Reader中读取，同时，前面部分从一个参数字符串中读取

	Scan,Fscan,Sscan 把输入中的 newlines 当做 spaces 对待
	Scanln,Fscanln 以及 Sscanln 在遇到一个 newline 时停下，同时要求读取的对象需要以 newline 或者 EOF 结尾
	Scanf,Fscanf 以及 Sscanf 从一个格式化字符串中解析参数，类似 Printf。
	中间一段没有仔细读或者说暂时没看懂。。。

	通过 verbs 处理输入是一种 隐式 space-delimited：除了 %c，所有 verb 都从输入中的第一个非空格开始读取，同时，%s 在遇到第一个空格或者 newline 时停止

	0b,0o,0x 前缀格式也可以使用

	可以指定长度，但是无法指定精度

	scanning functions 中 \r\n 和 \n一样对待

	所有的参数必须是指向基本类型的指针或者是 the Scanner interface 的实现

	Scanf,Fscanf,Sscanf 不需要消耗掉其整个输入

	最后一段没有看懂

*/

// func fprintfFun(w io.Writer, a ...interface{}) (int, error) {
// 	return fmt.Fprint(w, a)
// }
