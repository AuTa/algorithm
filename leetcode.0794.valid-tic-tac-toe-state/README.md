# [794. Valid Tic-Tac-Toe State](https://leetcode.com/problems/valid-tic-tac-toe-state/)

A Tic-Tac-Toe board is given as a string array `board`. Return True if and only if it is possible to reach this board position during the course of a valid tic-tac-toe game.

The board is a 3 x 3 array, and consists of characters `" "`, `"X"`, and `"O"`.  The `" "` character represents an empty square.

Here are the rules of Tic-Tac-Toe:

- Players take turns placing characters into empty squares (" ").
- The first player always places "X" characters, while the second player always places "O" characters.
- "X" and "O" characters are always placed into empty squares, never filled ones.
- The game ends when there are 3 of the same (non-empty) character filling any row, column, or diagonal.
- The game also ends if all squares are non-empty.
- No more moves can be played if the game is over.
```md
Example 1:
Input: board = ["O  ", "   ", "   "]
Output: false
Explanation: The first player always plays "X".

Example 2:
Input: board = ["XOX", " X ", "   "]
Output: false
Explanation: Players take turns making moves.

Example 3:
Input: board = ["XXX", "   ", "OOO"]
Output: false

Example 4:
Input: board = ["XOX", "O O", "XOX"]
Output: true
```
**Note:**

- `board` is a length-3 array of strings, where each string `board[i]` has length 3.
- Each `board[i][j]` is a character in the set `{" ", "X", "O"}`.

## 思路

排除法。  

排除不满足要求的即可（累计玩家的胜利次数）。
- X 要么和 O 的次数一样，要么比 O 多一次；
- 由于一步棋最多让两条线连起来，一个玩家的胜利次数不能大于 2；
- 两个玩家不能同时胜利；
- X 胜利时，X 的次数应该比 O 的次数多 1；
- O 胜利时，两个人的次数应该相等。
