#https://stackoverflow.com/questions/9280336/mysql-query-to-extract-domains-from-urls
SELECT
CONCAT(
".", # having the . in the end result
SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(SUBSTRING_INDEX(url,
'?', 1), # split on url params to remove weirdest stuff first
'://', -1), # remove protocal http:// https:// ftp:// ...
'/', 1), # split on path
':', 2), # split on user:pass
'@', 1), # split on user:port@
':', 1), # split on port
'www.', -1), # remove www.
'.', 10), # keep TLD + domain name
'.', -1), # keep tld
'/', 1)
) as tld,
COUNT(*) as counts,
SUM(score) as scores

FROM hackernewsstats
WHERE url != ""
AND YEAR(createdAt)= ?
GROUP BY tld
HAVING tld != ""
order by counts DESC
