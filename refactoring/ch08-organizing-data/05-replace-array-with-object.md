You have an Array in which certain elements mean different things.

*Replace the Array with an object that has a field for each element.*

+ 注意: 不是我最初臆想的 建一个继承自 `Array` 的新类, 再对数组的 数据/行为 进行封装

+ "One by one", mind the migration
    + add the `[]` and `[]=` methods
    + migrate one by one leaning on tests
    + remove the `[]` and `[]=` methods

+ Refactor with Deprecation
