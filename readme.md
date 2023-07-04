# Go Linear Programming

![Coverage](coverage.svg)

## Introduction

- [Linear Programming](https://en.wikipedia.org/wiki/Linear_programming)
- [The Art of Linear Programming](https://youtu.be/E72DWgKP_1Y)
- Find a vector `[b1, b2, ..., bn]`
- That dot product with a vector `[a1, a2, ..., an]` subject to the constraints
- And dot product with a vector `[c1, c2, ..., cn]` maximize a function

## Knapsack Problem

- [Knapsack Problem](https://en.wikipedia.org/wiki/Knapsack_problem)
- A list of objects where each object has a weight and a value: `Object{Weight, Value}`
- Find a binary vector `[b1, b2, ..., bn]`
- Then do dot product with the vector `objects.Weight` with the constraint: `b1 + b2 + ... + bn <= capacity`
- And maximize dot product with the vector `objects.Value`

## Run

```bash
make test
```
