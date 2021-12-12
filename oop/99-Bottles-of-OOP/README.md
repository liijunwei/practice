Recommended by YimingChen https://yiming.dev/blog/2017/12/08/book-review-99-bottles-of-oop/

# What This Book Is About

This book is about writing cost-effective, maintainable, and pleasing code.

Chapter 1 explores how to decide if code is "good enough." This chapter uses metrics to compare several possible solutions to the 99 Bottles problem. It introduces a type of solution known as Shameless Green, and argues that although Shameless Green is neither clever nor changeable, it is the best initial solution to many problems.

Chapter 2 is a primer for Test-Driven Development (TDD), which is used to find Shameless Green. This chapter is concerned with deciding what to test, and with creating tests that happily tolerate changes to the underlying code.

# Who Should Read This Book 9

Chapter 3 introduces a new requirement (six-pack), which leads to a discussion of how to decide where to start when changing code. This chapter examines the Open/Closed Principle, and then explores code smells. The chapter then defines a simple set of Flocking Rules which guide a step-by- step refactoring of code.

Chapter 4 continues the step-by-step refactoring begun in Chapter 3. It iteratively applies the Flocking Rules, eventually stumbles across the need for the Liskov Substitution Principle, and ultimately unearths a deeply hidden abstraction.

Chapter 5 inventories the existing code for smells, chooses the most prominent one, and uses it to trigger the creation of a new class. Along the way it takes a hard look at immutability, performance, and caching.

Chapter 6 is not yet available. This chapter performs a miracle which not only removes all conditionals, but also allows you to finally implement the new six-pack requirement without altering any existing code.

