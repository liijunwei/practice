page 140

```bash
tig oop/99-Bottles-of-OOP/04-Practicing-Horizontal-Refactoring/lib/bottles.rb
```

```ruby
def container(number)
  if number == 1
    "bottle"
  else
    "bottles"
  end
end

def pronoun(number)
  if number == 1
    "it"
  else
    "one"
  end
end

def quantity(number)
  if number == 0
    "no more"
  else
    number
  end
end
```

+ Notice the similarities in the above methods. Each has a single responsibility.

+ They are identical in shape.

+ All take the same argument.

+ Each contains a conditional and that conditional tests the argument against a specific value; it checks to see if the argument is equal to something, as opposed to greater or less than something.

+ These methods are incredibly consistent, and this did not happen by accident --- they’re a direct result of the refactoring rules.


