## 必要的术语解释

+ 字符集(CharacterSet), 字面上的理解就是字符的集合, 例如ASCII字符集, 定义了128个字符; GB2312定义了7445个字符. 而计算机系统中提到的字符集准确来说, 指的是已编号的字符的有序集合(不一定连续)

+ 字符集是 `字符` 和 `计算机里数字编号` 的映射关系的集合

+ 字符码(CodePoint)指的就是字符集中每个字符的数字编号

+ 例如 ASCII字符集用0-127这连续的128个数字分别表示128个字符
+ GBK字符集使用区位码的方式为每个字符编号, 首先定义一个94X94的矩阵, 行称为"区", 列称为"位", 然后将所有国标汉字放入矩阵当中, 这样每个汉字就可以用唯一的"区位"码来标识了
    + 例如"中"字被放到54区第48位, 因此字符码就是5448
+ Unicode中将字符集按照一定的类别划分到0-16这17个层面(Planes)中, 每个层面中拥有`2^16=65536`个字符码, 因此Unicode总共拥有的字符码, 也即是Unicode的字符空间总共有`17*65536=1114112`个

+ 字符(码)(CodePoint)指的是字符集中各个字符的数字编号

+ 编码: 将字符转换成字节流

+ 解码: 将字节流解析为字符

+ 字符编码(CharacterEncoding)是将字符集中的字符码映射为字节流的**一种具体实现方案**

+ 例如ASCII字符编码规定使用单字节中低位的7个比特去编码所有的字符
    + 例如'A'的编号是65, 用单字节表示就是0x41, 因此写入存储设备的时候就是`b'01000001'`

+ GBK编码则是将区位码(GBK的字符码)中的区码和位码的分别加上0xA0(160)的偏移(之所以要加上这样的偏移, 主要是为了和ASCII码兼容)
    + 例如刚刚提到的"中"字, 区位码是5448, 十六进制是0x3630, 区码和位码分别加上0xA0的偏移之后就得到0xD6D0, 这就是"中"字的GBK编码结果

+ 代码页(CodePage): 一种字符编码具体形式
    + 早期字符相对少, 因此通常会使用类似表格的形式将字符直接映射为字节流, 然后通过查表的方式来实现字符的编解码
    + 现代操作系统沿用了这种方式. 例如Windows使用936代码页、Mac系统使用EUC-CN代码页实现GBK字符集的编码, 名字虽然不一样, 但对于同一汉字的编码肯定是一样的

+ 大小端: 表示写操作中 `先写入高位字节` 或 `先写入低位字节`
    + 大小端的说法源自《格列佛游记》, 我们知道, 鸡蛋通常一端大一端小, 小人国的人们对于剥蛋壳时应从哪一端开始剥起有着不一样的看法
    + 同样, 计算机界对于传输多字节字(由多个字节来共同表示一个数据类型)时, 是先传高位字节(大端)还是先传低位字节(小端)也有着不一样的看法, 这就是计算机里头大小端模式的由来了
    + 无论是写文件还是网络传输, 实际上都是往流设备进行写操作的过程, 而且这个写操作是从流的低地址向高地址开始写(这很符合人的习惯), **对于多字节字来说, 如果先写入高位字节, 则称作大端模式, 反之则称作小端模式**
    + 也就是说, 大端模式下, 字节序和流设备的地址顺序是相反的, 而小端模式则是相同的
    + 一般网络协议都采用大端模式进行传输, Windows操作系统采用UTF-16小端模式

## 简单理解

计算机只认得0和1, 但是人类使用的语言里包含成千上万的符号, 要使用计算机语言描述世界, 就需要将这成千上万的符号映射到计算机里面

字符集就是这样的映射

最初, 字符集在英语世界里, 最初只需要`1个字节的空间`就能映射(当时)所需使用的所有字符(ASCII字符集), 甚至可以空出很多没有统一定义的"空"位

那时候的字符编解码系统非常简单, 就是简单的查表过程

例如将字符序列编码为二进制流写入存储设备, 只需要在ASCII字符集中依次找到字符对应的字节, 然后直接将该字节写入存储设备即可

解码二进制流的过程也是类似

这些空位被不同的人/组织/机构做出了不同的解释(OEM字符集是硬件制造商的, 比如IBM做PC时ASCII不够用, 扩展了一些, 固化在硬件中), 这时候, 如果两拨人进行信息传输, 那麻烦就大了, 因为在两台机器上的映射关系可能不一样(*这是一个问题*)

不同的OEM字符集导致人们无法跨机器交流各种文档

但是随着计算机/互联网发展, 字符种类变多了, 或者说全世界的字符做了并集, 1个字节的空间所能表示的字符数不够用了(*这是另一个问题*)

为了解决这些问题, 发展出了`多字节字符集(MBCS)和中文字符集`

可以通过第一个字节标识某段字符使用什么码表(CodePage)进行解析, 确定后, 再去对应的CodePage进行查表解码

之后有组织制定了ANSI标准, 国家标准, ISO标准

有了标准, 混乱会减少一些, 操作系统在发布的时候, 通常会预装这些标准的字符集和一些平台专用的字符集, 只要你写的, 或者看的文档使用这些标准字符集写成, 通用性就能得到保证

但是还有个问题: 没法在一份文档里展示所有的字符(因为一份文档只能使用上面说的一种标准)

因此, 终极标准`Unicode字符集`登场了

**Unicode字符集涵盖了目前人类使用的所有字符, 并为每个字符进行统一编号, 分配唯一的CodePoint**

+ 编码系统的变化

+ 在Unicode出现之前, 所有的字符集都是和具体编码方案绑定在一起的, 都是直接将字符和最终字节流绑定死了

+ 例如ASCII编码系统规定使用7比特来编码ASCII字符集; GB2312以及GBK字符集, 限定了使用最多2个字节来编码所有字符, 并且规定了字节序

+ 这样的编码系统通常用简单的查表, 也就是通过代码页就可以直接将字符映射为存储设备上的字节流了

+ 这种方式的缺点在于, 字符和字节流之间耦合得太紧密了, 从而限定了字符集的扩展能力. 假设以后火星人入住地球了, 要往现有字符集中加入火星文就变得很难甚至不可能了, 而且很容易破坏现有的编码规则

+ 因此Unicode在设计上考虑到了这一点, **将字符集和字符编码方案分离开**

+ **也就是说, 虽然每个字符在Unicode字符集中都能找到唯一确定的编号(字符码, 又称Unicode码), 但是决定最终字节流的却是具体的字符编码方式**

+ **例如同样是对Unicode字符"A"进行编码, `UTF-8字符编码`得到的字节流是`0x41`, 而`UTF-16`(大端模式)得到的是`0x00 0x41`**

+ 常见的Unicode编码
    + UCS-2/UTF-16
    + UTF-8
    + GB18030

## Unicode相关的常见问题

### Unicode是两个字节吗？

+ Unicode只定义了一个庞大的、全球通用的字符集, 并为每个字符规定了唯一确定的编号, 具体存储为什么样的字节流, 取决于字符编码方案
+ Unicode只定义了一个全球通用的字符集, 并为他们确定了唯一编号, 具体的编码方式和Unicode无关
+ 推荐的Unicode编码方式(方案)是UTF-16和UTF-8

## 乱码问题

+ 乱码指的是程序显示出来的字符文本无法用任何语言去解读
+ 一般情况下会包含大量`解码失败替换字符`或者`?`
+ 乱码问题是所有计算机用户或多或少会遇到的问题
+ 造成乱码的原因就是使用了错误的字符编码方式去解码字节流
+ **因此当我们在思考任何跟文本显示有关的问题时, 请时刻保持清醒: 当前使用的字符编码是什么?**
    + 读的时候, 解码方式是什么? 和 写的时候的编码方式一致吗?
    + 写的时候, 编码方式是什么?

+ 只有这样, 我们才能正确分析和处理乱码问题

+ 乱码是指字符解码失败显示的无法识别的字符; 乱码产生的原因是: 使用错误的字符编码方式去解码字节流

+ **It does not make sense to have a string without knowing what encoding it uses.**
