section_titles=(
9.1.Choosing-Which-Units-to-Test
9.2.Reorganizing-Tests
9.3.Seeking-Context-Independence
9.4.Communicating-With-the-Future
9.5.Summary
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

git push

