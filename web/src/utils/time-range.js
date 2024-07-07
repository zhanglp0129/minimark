const now = new Date()

// 今天
export const todayRange = []
const todayStart = new Date(`${now.getFullYear()}-${now.getMonth()+1}-${now.getDate()} 00:00:00`)
const todayEnd = new Date(`${now.getFullYear()}-${now.getMonth()+1}-${now.getDate()} 23:59:59`)
todayRange.push(todayStart, todayEnd)

// 昨天
export const yesterdayRange = []
const yesterday = new Date()
yesterday.setDate(yesterday.getDate()-1)
const yesterdayStart = new Date(`${yesterday.getFullYear()}-${yesterday.getMonth()+1}-${yesterday.getDate()} 00:00:00`)
const yesterdayEnd = new Date(`${yesterday.getFullYear()}-${yesterday.getMonth()+1}-${yesterday.getDate()} 23:59:59`)
yesterdayRange.push(yesterdayStart, yesterdayEnd)

// 近一周，7天，直到昨天23:59:59
export const weekRange = []
const week = new Date()
week.setDate(week.getDate()-7)
const weekStart = new Date(`${week.getFullYear()}-${week.getMonth()+1}-${week.getDate()} 00:00:00`)
weekRange.push(weekStart, yesterdayEnd)

// 近一月，30天，直到昨天23:59:59
export const monthRange = []
const month = new Date()
month.setDate(month.getDate()-30)
const monthStart = new Date(`${month.getFullYear()}-${month.getMonth()+1}-${month.getDate()} 00:00:00`)
monthRange.push(monthStart, yesterdayEnd)

// 近一年，365天，知道昨天23:59:59
export const yearRange = []
const year = new Date()
year.setDate(year.getDate()-365)
const yearStart = new Date(`${year.getFullYear()}-${year.getMonth()+1}-${year.getDate()} 00:00:00`)
yearRange.push(yearStart, yesterdayEnd)