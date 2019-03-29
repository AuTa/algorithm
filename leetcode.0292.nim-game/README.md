# [292. Nim Game](https://leetcode.com/problems/nim-game/)

You are playing the following Nim Game with your friend: There is a heap of stones on the table, each time one of you take turns to remove 1 to 3 stones. The one who removes the last stone will be the winner. You will take the first turn to remove the stones.

Both of you are very clever and have optimal strategies for the game. Write a function to determine whether you can win the game given the number of stones in the heap.

### Example:

```md
**Input:** 4
**Output:** false 
**Explanation:** If there are 4 stones in the heap, then you will never win the game;
             No matter 1, 2, or 3 stones you remove, the last stone will always be 
             removed by your friend.
```

## 思路

对于 x 个石头 A 拿之后的场景有三种，然后 B 拿之后有九种：
```
B: (x-1) -> A: (x-2), (x-3), (x-4)
B: (x-2) -> A: (x-3), (x-4), (x-5)
B: (x-3) -> A: (x-4), (x-5), (x-6)
```

B 选择之后 A 只要有一种选择结果是失败，那么 B 就会选择导致 A 失败的方式，所以 A 胜利取决于：  
`f(x) = (f(x-2) && f(x-3) && f(x-4)) || (f(x-3) && f(x-4) && f(x-5)) ||(f(x-4) && f(x-5) && f(x-6))`  
如果 `f(x-4) == false` 那么 `f(x)` 必然等于 `false`。

`f(1)`、`f(2)`、`f(3)` 都等于 `true`，`f(4)` 等于 `false`。`f(5)` 的时候 `(f(3) && f(2) && f(1)) == true`，所以 `f(5)` 等于 `ture`，同理 `f(6)`、`f(7)` 等于 `true`。`f(8)` 的时候，三个分支都存在 `f(4)`，所以 `f(8)` 等于 `false`。

所以 `n % 4 == 0` 时，结果为 `false`。
