package common

// 超过200票才能算是热门文章
const REQUIRE_VOTE_DAILY = 200

// 一天多少秒
const ONE_DAY_IN_SECONDS = 86400

// 通过每天秒数除所需票数计算每票分数
const VOTE_SCORE = ONE_DAY_IN_SECONDS / REQUIRE_VOTE_DAILY

// 一周多少秒
const ONE_WEEK_IN_SECONDS = 7 * ONE_DAY_IN_SECONDS
