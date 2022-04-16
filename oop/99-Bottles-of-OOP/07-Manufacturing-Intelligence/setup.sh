section_titles=(
7.1.Contrasting-the-Concrete-Factory-with-Shameless-Green
7.2.Fathoming-Factories
7.3.Opening-the-Factory
7.4.Supporting-Arbitrary-Class-Names
7.5.Dispersing-The-Choosing-Logic
7.6.Self-registering-Candidates
7.7.Auto-registering-Candidates
7.8.Summary
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
