https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/

> Unicode was a brave effort to create a single character set that included every reasonable writing system on the planet and some make-believe ones like Klingon, too. Some people are under the misconception that Unicode is simply a 16-bit code where each character takes 16 bits and therefore there are 65,536 possible characters. This is not, actually, correct. It is the single most common myth about Unicode, so if you thought that, don’t feel bad.

> In fact, Unicode has a different way of thinking about characters, and you have to understand the Unicode way of thinking of things or nothing will make sense.

Until now, we’ve assumed that a letter maps to some bits which you can store on disk or in memory:

`A -> 0100 0001`

In Unicode, a letter maps to something called a code point which is still just a theoretical concept. How that code point is represented in memory or on disk is a whole nuther story.

Every platonic letter in every alphabet is assigned a magic number by the Unicode consortium which is written like this: `U+0639`.

This magic number is called a code point.
The `U+` means “Unicode” and the numbers are hexadecimal.
`U+0639` is the Arabic letter Ain.
The English letter A would be U+0041.
You can find them all using the charmap utility on Windows 2000/XP or visiting [the Unicode web site](https://home.unicode.org/).


+ So for example in Israel DOS used a code page called 862, while Greek users used 737. They were the same below 128 but different from 128 up, where all the funny letters resided.

+  UTF-8 was another system for storing your string of Unicode code points, those magic U+ numbers, in memory using 8 bit bytes.

+  In UTF-8, every code point from 0-127 is stored in a single byte. Only code points 128 and above are stored using 2, 3, in fact, up to 6 bytes.

*Here stand the takeaway*
+ **It does not make sense to have a string without knowing what encoding it uses.**
+ You can no longer stick your head in the sand and pretend that “plain” text is ASCII.
+ **There Ain’t No Such Thing As Plain Text.**

+ **If you have a string, in memory, in a file, or in an email message, you have to know what encoding it is in or you cannot interpret it or display it to users correctly.**

不应该假设收到的字符串是"普通的文本"
一定得知道文本的编码方式, 才能对文本进行正确的处理

如果获取到了一段字符串, 但是不知道它的编码方式, 那么很可能因解析错误出现乱码
