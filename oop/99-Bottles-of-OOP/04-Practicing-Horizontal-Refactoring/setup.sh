section_titles=(
4.01.Replacing-Difference-With-Sameness
4.02.Equivocating-About-Names
4.03.Deriving-Names-From-Responsibilities
4.04.Choosing-Meaningful-Defaults
4.05.Seeking-Stable-Landing-Points
4.06.Obeying-the-Liskov-Substitution-Principle
4.07.Taking-Bigger-Steps
4.08.Discovering-Deeper-Abstractions
4.09.Depending-on-Abstractions
4.10.Summary
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

