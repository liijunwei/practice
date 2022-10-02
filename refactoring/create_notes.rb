# https://www.pearson.com/en-us/subject-catalog/p/refactoring-ruby-edition-ruby-edition/P200000009348/9780321984135?tab=table-of-contents

require 'fileutils'

hash = {}

hash["ch06-composing-methods"] = <<EOF
Extract Method . . . 102
Inline Method . . . 108
Inline Temp . . . 110
Replace Temp with Query . . . 111
Replace Temp with Chain . . . 114
Introduce Explaining Variable . . . 117
Split Temporary Variable . . . 121
Remove Assignments to Parameters . . . 124
Replace Method with Method Object . . . 127
Substitute Algorithm . . . 131
Replace Loop with Collection Closure Method . . . 133
Extract Surrounding Method . . . 135
Introduce Class Annotation . . . 139
Introduce Named Parameter . . . 142
Remove Named Parameter . . . 147
Remove Unused Default Parameter . . . 150
Dynamic Method Definition . . . 152
Replace Dynamic Receptor with Dynamic Method Definition . . . 158
Isolate Dynamic Receptor . . . 160
Move Eval from Runtime to Parse Time . . . 165
EOF

hash["ch07-moving-features-between-objects"] = <<EOF
Move Method . . . 167
Move Field . . . 172
Extract Class . . . 175
Inline Class . . . 179
Hide Delegate . . . 181
Remove Middle Man . . . 185
EOF

hash["ch08-organizing-data"] = <<EOF
Self Encapsulate Field . . . 188
Replace Data Value with Object . . . 191
Change Value to Reference . . . 194
Change Reference to Value . . . 198
Replace Array with Object . . . 201
Replace Hash with Object . . . 206
Change Unidirectional Association to Bidirectional . . . 210
Change Bidirectional Association to Unidirectional . . . 213
Replace Magic Number with Symbolic Constant . . . 217
Encapsulate Collection . . . 219
Replace Record with Data Class . . . 224
Replace Type Code with Polymorphism . . . 225
Replace Type Code with Module Extension . . . 232
Replace Type Code with State or Strategy . . . 239
Replace Subclass with Fields . . . 251
Lazily Initialized Attribute . . . 255
Eagerly Initialized Attribute . . . 257
EOF

hash["ch09-simplifying-conditional-expressions"] = <<EOF
Decompose Conditional . . . 261
Recompose Conditional . . . 264
Consolidate Conditional Expression . . . 265
Consolidate Duplicate Conditional Fragments . . . 268
Remove Control Flag . . . 269
Replace Nested Conditional with Guard Clauses . . . 274
Replace Conditional with Polymorphism . . . 279
Introduce Null Object . . . 284
Introduce Assertion . . . 292
EOF

hash["ch10-making-method-calls-simpler"] = <<EOF
Rename Method . . . 298
Add Parameter . . . 300
Remove Parameter . . . 302
Separate Query from Modifier . . . 303
Parameterize Method . . . 307
Replace Parameter with Explicit Methods . . . 310
Preserve Whole Object . . . 313
Replace Parameter with Method . . . 317
Introduce Parameter Object . . . 320
Remove Setting Method . . . 324
Hide Method . . . 327
Replace Constructor with Factory Method . . . 328
Replace Error Code with Exception . . . 332
Replace Exception with Test . . . 337
Introduce Gateway . . . 341
Introduce Expression Builder . . . 346
EOF

hash["ch11-dealing-with-generalization"] = <<EOF
Pull Up Method . . . 353
Push Down Method . . . 356
Extract Module . . . 357
Inline Module . . . 362
Extract Subclass . . . 363
Introduce Inheritance . . . 368
Collapse Heirarchy . . . 371
Form Template Method . . . 372
Replace Inheritance with Delegation . . . 386
Replace Delegation with Hierarchy . . . 389
Replace Abstract Superclass with Module . . . 392
EOF

hash["ch12-big-refactorings"] = <<EOF
The Nature of the Game . . . 397
Why Big Refactorings Are Important . . . 398
Four Big Refactorings . . . 398
Tease Apart Inheritance . . . 399
Convert Procedural Design to Objects . . . 405
Separate Domain from Presentation . . . 406
Extract Hierarchy . . . 412
EOF

def base_folder
  "#{ENV['HOME']}/practice/refactoring"
end

def print_or_create_sub_chs(ch, sub_chs)
  sub_chs.each do |sub_ch|
    filepath = "#{base_folder}/#{ch}/#{sub_ch}"

    if ENV['mode']
      FileUtils.touch filepath
    else
      puts filepath
    end
  end
end

hash.each do |ch, sub_chs|
  formated_sub_chapters = sub_chs.split("\n").map(&:strip).map {|e| e.split(' . . . ')[0]}.map(&:downcase).map {|e| e.gsub(' ', '-')}.map.with_index {|e, i| "#{(i+1).to_s.rjust(2, "0")}-#{e}.md"}
  print_or_create_sub_chs(ch, formated_sub_chapters)
end

