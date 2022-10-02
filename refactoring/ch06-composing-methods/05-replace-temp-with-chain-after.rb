class Select
  def self.with_option(option)
    s = self.new
    s.options << option
    s
  end

  def options
    @options ||= []
  end

  def and(option)
    options << option
    self
  end
end

select = Select.with_option(1999).and(2000).and(2001).and(2002)
p select
