section_titles=(
3.1.Listening-to-Change
3.2.Starting-With-the-Open-Closed-Principle
3.3.Recognizing-Code-Smells
3.4.Identifying-the-Best-Point-of-Attack
3.5.Refactoring-Systematically
3.6.Following-the-Flocking-Rules
3.7.Converging-on-Abstractions
3.8.Summary
)

for title in ${section_titles[@]}; do
  section_dir="oop/99-Bottles-of-OOP/03-Unearthing-Concepts/$title"
  mkdir -p $section_dir
  touch "$section_dir/note.md"
done

git add .

git commit -F- <<EOF
Setup directories.

bash oop/99-Bottles-of-OOP/03-Unearthing-Concepts/setup.sh
EOF

git push

