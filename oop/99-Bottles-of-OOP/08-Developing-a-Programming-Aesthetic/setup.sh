section_titles=(
8.1.Appreciating-the-Mechanical-Process
8.2.Clarifying-Responsibilities-with-Pseudocode
8.3.Extracting-the-Verse
8.4.Coding-by-Wishful-Thinking
8.5.Inverting-Dependencies
8.6.Obeying-the-Law-of-Demeter
8.7.Identifying-What-The-Verse-Method-Wants
8.8.Pushing-Object-Creation-to-the-Edge
8.9.Summary
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
