section_titles=(
6.1.Consolidating-Data-Clumps
6.2.Making-Sense-of-Conditionals
6.3.Replacing-Conditionals-with-Polymorphism
6.4.Transitioning-Between-Types
6.5.Making-the-Easy-Change
6.6.Defending-the-Domain
6.7.Summary
)

for title in ${section_titles[@]}; do
  section_dir="$(dirname -- "${BASH_SOURCE[0]}")/$title"
  mkdir -p $section_dir
  touch "$section_dir/note.md"
done

git add .

git commit -F- <<EOF
Setup directories.

bash $(dirname -- "${BASH_SOURCE[0]}")/setup.sh
EOF
