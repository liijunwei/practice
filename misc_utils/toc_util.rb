def sanitize_filename(filename)
  filename = filename.split /(?<=.)\.(?=[^.])(?!.*\.[^.])/m
  filename = filename.map { |s| s.gsub /[^a-z0-9\-]+/i, '_' }.map(&:downcase)

  return filename.join '.'
end

def get_toc_directories(dir_base, toc_string)
  list = toc_string.split("\n")
  list = list.each_with_index.map {|e, i| "#{dir_base}/#{i.to_s.rjust(2, '0')}.#{sanitize_filename(e)}"}

  return list
end

def make_toc_directories(list)
  list.each {|dir_name| FileUtils.mkdir_p dir_name}
end

def touch_note_file(list)
  list.each {|dir_name| FileUtils.touch "#{dir_name}/note01.md"}
end

toc = <<-EOF
UNIX PROCESSES
Introduction
Primer
Processes Have IDs
Processes Have Parents
Processes Have File Descriptors
Processes Have Resource Limits
Processes Have an Environment
Processes Have Arguments
Processes Have Names
Processes Have Exit Codes
Processes Can Fork
Orphaned Processes
Processes Are (CoW) Friendly
Processes Can Wait
Zombie Processes
Processes Can Get Signals
Processes Can Communicate
Daemon Processes
Spawning Terminal Processes
Ending
Appendix: How Resque Manages Processes
Appendix: How Unicorn Reaps Worker Processes
Appendix: Preforking Servers
Appendix: Spyglass
EOF

dir_base = "/Users/lijunwei/practice/ruby/working-with-unix-processes"
toc_list = get_toc_directories(dir_base, toc)
make_toc_directories(toc_list)
touch_note_file(toc_list)
