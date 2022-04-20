+ disperses [dɪˈspɜːrsɪz]  分散

+ If you would like the factory to simply continue working when new candidates appear, you have two basic options.
    1. The factory could dynamically figure out which classes belong on its list, or
    2. classes who want to be on the list could explicitly ask the factory to put them there.

如果你想要让工厂能够在新加类后能正常工作, 你有两个简单的选择

1. 工厂能够动态的识别出哪些类能正常工作
2. 工厂提供注册的方法, 需要新加的方法自动注册

+ Choice #1 above is possible only if there’s something about the candidate classes that allows the factory to identify them, and this may not always be true.

第一种选择只有在 新增的类里有某些特征, 让工厂能识别出来的 条件下才是可能的, 但这种条件不可能总是满足

+ Choice #2, however, is always an option. If candidates are willing to depend on knowing the name of the factory, they can assume responsibility for putting themselves on the list. Lists like these are often referred to as registries.

第二种选择却总是可行的

如果新增的类知道自己将被工厂制造, 那么它可以假设自己有职责将自己像被制造这个信息告诉工厂

这种方法常常叫做 "注册"

