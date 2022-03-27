toc = <<-EOF
5.1.Selecting the Target  Code  Smell
5.2.Extracting  Classes
5.3.Appreciating  Immutability
5.4.Assuming  Fast  Enough
5.5.Creating  BottleNumbers
5.6.Recognizing Liskov  Violations
5.7.Summary
EOF

puts toc.split("\n").map(&:squeeze).map {|line| line.gsub(/\s/, '-')}

