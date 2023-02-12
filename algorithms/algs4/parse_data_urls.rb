require 'net/http'
require 'nokogiri' # https://rubygems.org/gems/nokogiri

class HTMLParser
  attr_reader :url

  def initialize(url)
    @url = url
  end

  def parse
    content = Net::HTTP.get(URI(url))
    document = Nokogiri.HTML(content)
    rows = document.xpath('/html/body/div/center/table').search('tr')

    filenames = rows.select { |e| !e.search("td").empty? }.map { |row| parse_filename(row) }
    puts filenames.map {|e| "#{url}#{e}"}
  end

  private

  def parse_filename(row)
    file, _size, _description, _format, _source = row.search('td').map { |e| e.content }

    file
  end
end

HTMLParser.new('https://introcs.cs.princeton.edu/java/data/').parse
