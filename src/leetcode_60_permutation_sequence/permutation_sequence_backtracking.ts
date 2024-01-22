/*
 * @lc app=leetcode id=60 lang=typescript
 *
 * [60] Permutation Sequence
 */

// @lc code=start
function getPermutation(n: number, k: number): string {
    let factorial = [1]
    let nVec: number[] = []
    for (let x = 1; x <= n - 1; x++) {
        factorial.push(factorial[x - 1] * x)
        nVec.push(x)
    }
    nVec.push(n)
    k = k - 1
    let acc = 0
    for (let i = n - 1; i >= 0; i--) {
        acc = acc * 10 + nVec.splice(k / factorial[i], 1)[0]
        k %= factorial[i]
    }
    return acc.toString()
}
// @lc code=end
