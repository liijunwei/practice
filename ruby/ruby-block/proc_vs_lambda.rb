require 'rspec'
require 'pry'

class ProcVsLambda
  def double(calable_object)
    calable_object.call * 2
  end

  def lambda_adder
    lambda { |a, b| a + b }
  end

  def proc_adder
    Proc.new { |a, b| a + b }
  end

  def another_double
    p = Proc.new { return 10 }
    result = p.call

    return result * 2 # unreachable code!
  end
end

RSpec.describe ProcVsLambda do
  context "comparing the return keyword" do
    describe "lambda" do
      it 'returns only from lambda --- behaves like method' do
        l = lambda { return 10 }
        expect(subject.double(l)).to eq(20)
      end
    end

    describe "proc" do
      it 'returns from where proc was defined --- behaves like block' do
        expect(subject.another_double).to eq(10)
      end
    end
  end

  context "comparing the arguments count" do
    describe "lambda" do
      it 'works well when arguments count match' do
        expect(subject.lambda_adder.call(1, 2)).to eq(3)
      end

      it 'will fail when arguments count mismatch' do
        expect { subject.lambda_adder.call(1, 2, 3) }.to raise_error ArgumentError
      end
    end

    describe "proc" do
      it 'works well when arguments count match' do
        expect(subject.proc_adder.call(1, 2)).to eq(3)
        expect(subject.proc_adder.call(3, 1)).to eq(4)
        expect(subject.proc_adder.call(2, 3)).to eq(5)
      end

      it 'works well when arguments count mismatch' do
        expect(subject.proc_adder.call(1, 2, 3)).to eq(3)
        expect(subject.proc_adder.call(3, 1, 2)).to eq(4)
        expect(subject.proc_adder.call(2, 3, 1)).to eq(5)
      end
    end
  end
end
