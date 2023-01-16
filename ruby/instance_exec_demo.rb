class Demo
  def foo(&block)
    block.call
  end

  def bar(&block)
    instance_exec(&block)
  end
end

# rspec instance_exec_demo.rb
RSpec.describe Demo do
  context = self

  let(:r1) { Demo.new.foo { self } }
  let(:r2) { Demo.new.bar { self } }

  it 'has different context' do
    expect(r1).to be_a(context)
    expect(r2).to be_a(Demo)
  end
end
