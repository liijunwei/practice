require 'fileutils'

def setup_tutorial_chapter(chapter, chapter_num, padding = 2)
  folder_name = "scheme/Yet-Another-Scheme-Tutorial/#{chapter_num.to_s.rjust(padding, "0")}.#{chapter}"
  FileUtils.mkdir_p(folder_name)

  filename = "#{folder_name}/note.md"
  FileUtils.touch(filename)
end

def print_toc(chapter, chapter_num, padding = 2)
  puts "- [ ] #{chapter_num.to_s.rjust(padding, "0")}.#{chapter}"
end

chapters = %w[
  installing_mit_scheme
  using_scheme_as_a_calculator
  making_lists
  defining_functions
  branching
  local_variables
  looping
  higher_order_functions
  input_output
  assignment
  character_string
  symbol
  association_list_and_hash_table
  vector_and_structure
  defining_syntaxes
  continuation
  lazy_evaluation
  nondeterminism
]

chapters.each_with_index do |chapter, index|
  setup_tutorial_chapter(chapter, index+1)
end

chapters.each_with_index do |chapter, index|
  print_toc(chapter, index+1)
end


