def get_suffix_array_of(string)
  upper_limit = string.size - 1
  string.chars.map.with_index {|_char, index| {k: index, v: string[index..upper_limit]}}.sort_by {|e| e[:v]}
end

get_suffix_array_of('HORSE') # => ["E", "HORSE", "ORSE", "RSE", "SE"]
get_suffix_array_of('camel') # => ["amel", "camel", "el", "l", "mel"]
get_suffix_array_of('to be or not to be')
get_suffix_array_of('aacaagtttacaagc')
