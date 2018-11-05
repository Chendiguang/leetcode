def median(A, B):
    m, n = len(A), len(B)
    if m > n:
        A, B, m, n = B, A, n, m
    if m == 0:
        raise ValueError
    
    imin, imax, half_len = 0, m, (m + n + 1) / 2

def longestPalindrome(string):
    if len(string) <= 1:
        return string
    # 格式化字符串
    s = '^'
    for char in string:
        s += '#' + char
    s += '#$'
    print(s)
    p = [0 for _ in range(len(s))]
    # 当前的中心点center，拥有最长回文字符串的中心索引maxCenterIdx，右边界right
    center, maxCenterIdx, right = 0, 0, 0
    for i in range(1, len(s)-1):
        # p[i] = (right > i)? min(right-i, p[j]) : 0
        j = 2 * center - i
        print(j)
        if right > i:
            p[i] = min(right-i, p[j])

        # p[i] = min(right-i, p[j]) if right > i else 0

        # 尝试扩展以i为中心的回文字符串
        # print(i, s[i+1+p[i]], s[i-1-p[i]])
        if s[i+1+p[i]] == s[i-1-p[i]]:
            p[i] += 1
            
        # 更新右边界和当前的的回文字符串的中心
        if i+p[i] > right:
            center = i
            right = i + p[i]
        # 更新最长回文的中心索引
        if p[maxCenterIdx] < p[i]:
            maxCenterIdx = i

    print(p)
    start = (maxCenterIdx - 1 - p[maxCenterIdx]) // 2
    end = start + p[maxCenterIdx]
    return string[start:end]

