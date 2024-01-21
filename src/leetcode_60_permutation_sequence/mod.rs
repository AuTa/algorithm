pub struct Solution;
/*
 * @lc app=leetcode id=60 lang=rust
 *
 * [60] Permutation Sequence
 */

// @lc code=start
impl Solution {
    pub fn get_permutation(n: i32, k: i32) -> String {
        let n = n as usize;
        let factorial = (1..=n - 1).fold(vec![1], |mut acc, x| {
            acc.push(acc[x - 1] * x);
            acc
        });
        let mut n_vec: Vec<_> = (1..=n).collect();
        let mut k = k as usize - 1;
        let acc = (0..n).rev().fold(0, |acc, n| {
            let kpre = k;
            k %= factorial[n];
            acc * 10 + n_vec.remove(kpre / factorial[n])
        });
        acc.to_string()
    }
}
// @lc code=end
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn main() {
        let r = Solution::get_permutation(3, 3);
        assert_eq!(r, "213");
    }
}
