section_titles=(
5.1.Selecting-the-Target-Code-Smel
5.2.Extracting-Clases
5.3.Apreciating-Imutability
5.4.Asuming-Fast-Enough
5.5.Creating-BotleNumbers
5.6.Recognizing-Liskov-Violations
5.7.Sumary
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

