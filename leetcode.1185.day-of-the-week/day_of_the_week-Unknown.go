/*
 * @lc app=leetcode id=1185 lang=golang
 *
 * [1185] Day of the Week
 *
 * https://leetcode.com/problems/day-of-the-week/description/
 *
 * algorithms
 * Easy (64.33%)
 * Likes:    58
 * Dislikes: 686
 * Total Accepted:    14.4K
 * Total Submissions: 22.3K
 * Testcase Example:  '31\n8\n2019'
 *
 * Given a date, return the corresponding day of the week for that date.
 *
 * The input is given as three integers representing the day, month and year
 * respectively.
 *
 * Return the answer as one of the following values {"Sunday", "Monday",
 * "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.
 *
 *
 * Example 1:
 *
 *
 * Input: day = 31, month = 8, year = 2019
 * Output: "Saturday"
 *
 *
 * Example 2:
 *
 *
 * Input: day = 18, month = 7, year = 1999
 * Output: "Sunday"
 *
 *
 * Example 3:
 *
 *
 * Input: day = 15, month = 8, year = 1993
 * Output: "Sunday"
 *
 *
 *
 * Constraints:
 *
 *
 * The given dates are valid dates between the years 1971 and 2100.
 *
 */
package leetcode1185

// @lc code=start
import "time"

func dayOfTheWeek(day int, month int, year int) string {
	const BaseDay = 4 // 1971/01/01 算第一天，为周五，前一天是周四，基础量就是 4
	MonthDays := [...]int{-1, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	yearDays := 0
	for y := 1971; y < year; y++ { // 最后一年的天数靠月和天来计算
		yearDays += 365
		if isLeapYear(y) {
			yearDays += 1
		}
	}
	monthDays := day // 最后一个月的天数就是 day
	for m := 1; m < month; m++ {
		monthDays += MonthDays[m]
	}
	if month > 2 && isLeapYear(year) { // 闰年且月份在二月之后需要加一天
		monthDays += 1
	}
	countDays := yearDays + monthDays + BaseDay
	return time.Weekday(countDays % 7).String()
}

// isLeapYear 计算是否是闰年。
// 闰年定义：年份是 4 的倍数，不是 100 的倍数；年份是 400 的倍数。
func isLeapYear(year int) bool {
	return (year%4) == 0 && (year%100) != 0 || (year%400) == 0
}

// @lc code=end

func dayOfTheWeekKim(day int, month int, year int) string {
	var w int
	if month == 1 || month == 2 {
		month += 12
		year--
	}
	w = (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400 + 1) % 7
	return time.Weekday(w).String()
}
