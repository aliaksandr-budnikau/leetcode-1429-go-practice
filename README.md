# 1429. First Unique Number

Go practice project. The solution is intentionally **not implemented**.

Problem: https://leetcode.doocs.org/en/lc/1429/#description

## Task

Implement a data structure that supports:

- `Constructor(nums []int) FirstUnique` — initialize it with the given sequence.
- `ShowFirstUnique() int` — return the earliest number in the queue whose frequency is exactly one; return `-1` if no such number exists.
- `Add(value int)` — append a number to the queue.

The order is the order in which numbers were initially provided and subsequently added.

## Constraints

- `1 <= nums.length <= 100000`
- `1 <= nums[i], value <= 100000000`
- At most `50000` calls to `ShowFirstUnique` and `Add`.

## Run tests

```bash
go test ./...
```

The starter implementation is only a compileable stub, so the tests are expected to fail until you implement the data structure.

## Suggested workflow

1. Implement the struct fields in `first_unique.go`.
2. Implement `Constructor`, `ShowFirstUnique`, and `Add`.
3. Run `go test ./...` after each change.
4. Do not change the tests to make them pass.
