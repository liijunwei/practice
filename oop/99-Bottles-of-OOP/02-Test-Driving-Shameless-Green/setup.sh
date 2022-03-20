section_titles=(
2.01.Understanding-Testing
2.02.Writing-the-First-Test
2.03.Removing-Duplication
2.04.Understanding-Transformations
2.05.Tolerating-Duplication
2.06.Hewing-to-the-Plan
2.07.Exposing-Responsibilities
2.08.Choosing-Names
2.09.Revealing-Intentions
2.10.Writing-Cost-Effective-Tests
2.11.Avoiding-the-Echo-Chamber
2.12.Considering-Options
2.13.Summary
)

for title in ${section_titles[@]}; do
  mkdir -p "oop/99-Bottles-of-OOP/02-Test-Driving-Shameless-Green/$title"
  touch "oop/99-Bottles-of-OOP/02-Test-Driving-Shameless-Green/$title/note.md"
done

git add .

git commit -F- <<EOF
Setup directories.

bash oop/99-Bottles-of-OOP/02-Test-Driving-Shameless-Green/setup.sh
EOF

git push

