# page 132

# Around Alias 环绕别名

# 思路:
# 1. 给方法定义一个别名
# 2. 重新定义这个方法
# 3. 在新方法中调用老的方法

# 警告: 环绕别名存在两个潜在的陷阱
# 1. 环绕别名是猴子补丁, 可能会破坏已有的代码(使用过的技巧越强, 就越需要写更多的测试代码)
# 2. 与文件加载有关, 你永远不该把一个环绕别名加载两次

class String
  alias :real_length :length # 注意这个得放在方法定义的前面

  def length
    real_length > 5 ? "long" : "short"
  end
end

require 'rspec'

RSpec.describe String do
  describe "enhanced string" do
    let(:str) { "War and Peace" }

    it 'uses original length()' do
      expect(str.real_length).to eq(13)
    end

    it 'uses enhanced length()' do
      expect(str.length).to eq("long")
    end
  end
end
