#!/usr/bin/env ruby
# frozen_string_literal: true

require "benchmark"
require_relative "persistent_vector"
require_relative "persistent_vector_basic"

def measure(label)
  GC.start
  alloc0 = GC.stat(:total_allocated_objects)
  t = Benchmark.realtime { yield }
  allocs = GC.stat(:total_allocated_objects) - alloc0
  puts format("%-22s %10.1f ms   allocs=%8d", label, t * 1000, allocs)
end

def bench(klass, name)
  puts "== #{name} =="

  measure("push 到 32") do
    v = klass.empty
    31.times { |i| v = v.push(i) }
  end

  measure("push 到 100_000") do
    v = klass.empty
    100_000.times { |i| v = v.push(i) }
  end

  big = klass.from_a((0...100_000).to_a)

  measure("随机 nth × 200k") do
    rng = Random.new(1)
    200_000.times { big[rng.rand(100_000)] }
  end

  measure("尾部 nth × 200k") do
    rng = Random.new(1)
    200_000.times { big[99_968 + rng.rand(32)] }
  end

  measure("pop 到空") do
    v = klass.from_a((0...100_000).to_a)
    100_000.times { v = v.pop }
  end
end

bench(PersistentVectorNoTail, "无 tail（基线）")
bench(PersistentVector,       "有 tail（优化后）")
