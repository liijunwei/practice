# page 132

# Around Alias 环绕别名

# 思路:
# 先给方法命名一个别名, 然后重新定义它
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
