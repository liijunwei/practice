+ 读者应该注意到, 本届中在讨论外部变量时 谨慎地使用了定义(define) 与 声明(declaration) 这两个词
    + "定义" 表示创建变量或分配存储单元
    + "声明" 指的是说明变量的性质, 但并不分配存储单元

+ 过分依赖外部变量会导致一定的风险, 因为他会使程序中的数据关系模糊不清
    + 外部变量的值可能会被意外地或者不经意地修改, 而程序的修改又变得十分困难

