package day19

import "strings"

var Input = `47: 72 10 | 5 112
117: 72 129 | 5 63
104: 5 46 | 72 70
85: 109 5 | 112 72
96: 81 5 | 120 72
49: 48 72 | 41 5
93: 5 44 | 72 118
57: 102 72 | 109 5
102: 5 72
110: 68 5 | 112 72
134: 77 72 | 9 5
9: 72 5
97: 5 51 | 72 15
114: 5 75 | 72 68
17: 5 121 | 72 107
126: 102 5 | 21 72
56: 36 5 | 113 72
109: 72 5 | 5 72
77: 72 72 | 72 5
28: 72 127 | 5 111
3: 19 72 | 105 5
45: 72 109 | 5 9
10: 72 72 | 5 72
36: 50 5 | 112 72
116: 72 126 | 5 51
63: 82 72 | 24 5
123: 72 50 | 5 70
89: 71 5 | 50 72
21: 5 5
53: 29 72 | 59 5
122: 68 5 | 9 72
20: 72 109 | 5 70
108: 86 5 | 5 72
46: 86 86
69: 77 72
48: 94 72 | 33 5
130: 72 75 | 5 77
13: 125 72 | 85 5
107: 5 89 | 72 124
39: 9 72 | 71 5
78: 75 72 | 70 5
29: 122 72 | 101 5
60: 102 5 | 75 72
27: 22 5 | 98 72
95: 68 72 | 46 5
90: 134 72 | 69 5
76: 72 102 | 5 102
38: 5 45 | 72 123
121: 72 110 | 5 40
24: 72 52 | 5 67
74: 130 5 | 25 72
128: 72 10 | 5 9
7: 5 9 | 72 50
1: 72 102 | 5 75
51: 72 109 | 5 50
43: 104 72 | 78 5
101: 50 72 | 70 5
12: 72 77 | 5 46
87: 72 131 | 5 74
50: 72 72
65: 5 108 | 72 70
131: 134 72 | 95 5
72: "b"
70: 72 86 | 5 72
6: 72 100 | 5 73
62: 5 50 | 72 102
88: 5 49 | 72 26
80: 12 5 | 1 72
31: 5 35 | 72 117
41: 90 5 | 91 72
5: "a"
33: 135 72 | 62 5
42: 96 5 | 88 72
64: 5 99 | 72 119
82: 80 72 | 103 5
26: 5 64 | 72 6
11: 42 31
120: 5 37 | 72 32
115: 34 72 | 39 5
75: 72 72 | 5 5
19: 58 72 | 115 5
44: 77 5 | 112 72
61: 112 5 | 9 72
98: 72 71 | 5 108
91: 5 47 | 72 55
23: 72 72 | 5 86
127: 72 75 | 5 68
8: 42
58: 114 5 | 47 72
92: 46 5 | 10 72
94: 65 5 | 76 72
55: 86 77
66: 79 5 | 127 72
124: 72 68 | 5 70
37: 54 5 | 56 72
35: 5 4 | 72 3
30: 97 5 | 28 72
99: 62 72 | 1 5
54: 5 135 | 72 1
129: 87 72 | 53 5
103: 72 18 | 5 78
111: 108 5 | 75 72
119: 72 92 | 5 14
16: 5 10 | 72 70
135: 46 5 | 112 72
34: 71 86
0: 8 11
52: 128 5 | 60 72
112: 5 72 | 5 5
67: 5 83 | 72 7
79: 86 46
59: 5 20 | 72 133
22: 5 9 | 72 71
32: 38 72 | 13 5
113: 102 5 | 112 72
100: 5 16 | 72 15
18: 102 72 | 102 5
81: 72 17 | 5 84
71: 72 5 | 5 5
2: 5 93 | 72 27
105: 116 5 | 43 72
83: 72 46 | 5 70
68: 72 72 | 86 5
125: 77 5 | 108 72
4: 5 2 | 72 30
15: 9 72
40: 68 5 | 50 72
106: 130 5 | 132 72
132: 5 75 | 72 70
133: 9 72 | 21 5
84: 106 72 | 66 5
86: 72 | 5
14: 108 72 | 112 5
118: 72 77 | 5 102
73: 5 61 | 72 57
25: 72 68 | 5 23

aaaaaaaababbabababaaaababaaaabba
babbbababbbbbbaaababbbbabbbaaaabbaaaaabbbbbbaabaabbbaabb
abbabbaababbabbbaabbbabbabbbabbbbabbbabb
abaabaaabaabbbaaaabbbbaabbbbbabbbbaaabab
aaaaaabbabbbabbaabaaaabbbababbbb
bbaababaabababababaababbabbaaaba
bbabbbbaabbababbaabbbaba
aabababaaabaaaababaabbabaabaaabb
bababbbabbbbbbbaabababaa
abababaaabbabbbaabaabaaaaaaababbabaabaababbaaabbbaabaabaabaaabbbbabbbbaaabbabbbb
abbbabbbaabaaaaabbbabbba
abbaabbabababbaaaaabbbab
ababaaabaababbbababaaabbbaaabbaabbaabaab
babaaaabbabaaaababbbbbba
bbaaababababbbbbbabaaaaababbbbba
abbbbbaabbbabbaaaabaaaaabbbbaaaaaaababab
bababbaaabaabbbabaabbbbbbbabbaaabaabaabb
aabaaaaaaabbbabbbaabaababaaaabbaabbbbbba
abbabababbbaabaabbbabbabbbaabbaaabbaabab
aaaaabbbbabbbbabbaaaaaba
abbbbaaaabaabaabbabaaabbbbbbbbbb
bbaabbbbbabbabaaabbabbbababaabbabbbbabaa
aabbbbabaaabaaaaaaaaaaab
aaabbbbbabbbbbaaaaaabbabaabbbabbabbababbbaaaabab
abbaabbbbbbabaababbabbaabbbaaaaaaabbaaab
aabaabbbaababbbaabaaaabaaababaabababaaba
aabaabbbabababbbaababbaa
abbbabbbbaabaababbaababbabaaabbb
ababbbabbbabbbaaabbabbab
babaaaaaabbabbaabaaabaab
baababbababaabbbbaabaabb
bbbbbbbabbababbbbaabbbaababbbbababbbabbbabbabaaabaaabbaabaaabbba
abaaabaabbaababaabbaabab
baabbaaabbbbbbaaabbabaab
aaaaaabbaaaabbbabaababbaababbaabbababbbaaaaababbaaabaaaababbbabbabbbabba
bbbbbbbaaabbaabaaabbbaababaabbbbabaaabbabaabbbabbaaaaabb
abbbbbaabbbbbaaaabbbbbba
bbaababaababaabbbabbababbbbbbaba
baababbaaaaaaabaabababaa
aabbbabbabbbbbababaabbbbbbbbbbbb
baabaaaabababbaabbbbbbbaaaaaaabaaaabbabbaabbaababaabaaab
bbbaababababbababbaaabaababbabaa
aababbbaabbbbaaaaababbbb
aabbaaaaaaaabbaabaababbabaaaaaba
aaabbbbaaabbbbabaabbaabbbbaaabbb
bbbabababbbbbbaabaaaabbabbbbaaabbabbbabbaabaaaab
aabbbabbbbbbbbbbaaaaaaababbbbbbb
aaabbbbbabbababaabbbbbabbaaababbabbbabba
baaabbbbabbaabbbbaaaaaabbaaababaabbbbabbaaabbabbabbababa
abababbbbabaababbabaabbbbabbaabb
aaaabaaababbaaaaaaabbaaabbaabbbbbabbaabb
abbaabbbaaaababbabaaaababbbababbabaabbababbaaaab
babbabaaababaababaaaaabababbbbba
baaabbbbaaaaaabaabbbbbbb
bbbbaaaaaaabbbbbaaaabaab
baaababbabbaaabbabbaaabbbabaaaabaaabbaaabbbbbbabaaaaaaaabbbbbabaaaabaabb
bbbaaabbbbabbabbbbabbbaabbaabbababbabababbbbbaabaaaaaaabaabaabababaaabaaabbbaaaa
bbbbaaabbbbaaaaabbbaaaaaaabbaabbbbbababa
ababaaaabaabababbababaaa
aabbbbabaaabbabaababbaaa
bbbaaabbbbaaaaabbaaaabbbbaaabbbbbaaaaaaaaaababbaababbbaaaabbbaabaabbabbb
abbbbaaababbbaaaabaabbbbbbaaabaabababbbaabbaabbabbabbabb
bbbaabaaababaabbabaaaaaabbaaaaaabbaabbbabaaabaaaaaaabaabbbbabbba
ababbbababaaaababaabbabb
baabbbabaaabbbabaaaabbbbbaaaababbbabaabbbbababbabbbabbab
abbaabbbbbaabaabbaaabbaabbaaaabb
aabbabbbaabbbaabbabbaaababbbbabb
aabbaabaabaaaaaaaabbbaba
ababbaaaabbabbbaabbaaaba
bababbbaababbbabbabbabbbbabababa
baabaabbaabbbbbababaabaabbababbbabbbbaabbabbbaaabababbab
bbbaaabbbababbaaabbabbbbbbbbabbb
bbabbbbabaabbababababbababbbaababbbaaababaaaaaba
abbbabbaabbbbaaaaaaaababbbbababaaaababbbbbbabababbabbbbbbaaaabbababaaaaaaabbbbba
bbabbaabbaabbabaabbaaabb
abbaababbbbbbabaaaaaaaaaabbbabbbabaaaaaa
abbaaaaabbababbbbbbbbabbabababbaabbaabab
aaabababbaabbbbbaaaaaabbbaaababbaababbbbaabbaabbbbabaabbbabbbbabaabaaaba
bbbabaaaaabbbbbabbbbbaaabaaabbaa
aabbabbaaaaaababbbabbabababbbbaa
bbbbabaaaabbabaababababbaabbbbababbabbabaabaaaab
baabbbabbbaababbbabababbaabababbbbbabbbbbabbaabaabaabababbabbaaabbbbabaabbabbbbaaabbaaaabababaab
abababababaabaaaaabbbaba
bbaaaaaaaababbbabbabaaaabaaaaabb
babbbaababbaabbbbbabaaaa
aababbbababaababababababbbbbaabaaabaabababbbabbaaabaaabbaaabababbaabaabb
ababaaaababbaabaabbbbabbabbbabbabaaaabaababaabbbbbabaabbbbbbbbabaaabbbab
abaabbbbbaaabababbabbabaaabbbbbbbababbbaabaabababaaabbabbabbabab
baaabaababaaabaaabbaabbbabbbabaabaabbbbaaabaaabbaaabbbaa
babaaaaaabaabaaabbaabaaabbaabbabaaababaa
bbbabaaaaaabababbbabbabbaaaaabaa
bbbbabaababaabbababbbbbabaababababaababaabbabbaaabbabbbbaaaababbbbbbbbbabababbbbabbbabbabababbab
baaabbbbabbabbbbbbbbaabbabababaa
aabbaaaaaabababbabbbaabb
bbbbbaaaabbbbaabaaabbaab
baabaababaabaabababababbabbbbaab
bbaabbabbbbabaababbbbabbabbabbbaaabaabaaababbaaa
baaabbbbaaaaaaaabbbabbaaaaababaa
ababaaaababbabbbabababba
aabbabaabbabaaaaaabbbaaababaaabaabbaaaabbbbaaaaaaaaaaaab
bbababbbbabbbabaabbabbab
baababbbaaaaaaaaaabaabaaabbaaababaaabbba
babaaaababbaabbbbbabbaaa
bbaaababaaaaababaababaab
aabbabbbbaabaabbaabbababaabbaabaabababab
abaaabaabbaaaaabbbabbabaabbbaababaabbabb
bbbabaaaaaaabaaabaabaaab
baabaababbaaaababaaaaaaabbbabaab
bbabaaaababaaaaaaabbaabababbbaaababaababbababbbbababbbba
aaaababbbbabbbaabaabbbbb
aabababbababaaabaabbbbbb
bbababbaabbabbaaabaaabbb
ababbbaaaabbabbabbbbaaab
bbbbaaaaaabbabbaaabaabab
babbbbbbbbbbaaaaabbbbaabaaabbbaaababbaaa
baabaaaaaabaabbbabaabaaababaabaababbaabb
babababbbbbaabbbaaabaabaaabbaababaababaa
abbabbaabbabaaaaabbbbbab
aaabbbbbabbaaaaabbbbabaa
abaaaaaaaaaaaaaababbabbababbaabbbbbbbaab
bbbbbbaaaabbaaaaabbaaaaabbabbaabbaabaabbabbaaabbbaabbabb
babbbbaaabbaaababbaababaabbabaaaabbabaab
abbabababbbbbabababbbbbaaaabbaaababaababbabbbaaabbababbabaabbaaaaaaaabab
baaabbaaaababbaabababaaa
bbaaaabaaaabbababaaabaaabbabbabb
babaabbaabbabbababbbbbaabbaaabaaaabaaabbbababaababbabbbbaabbbaaabaaababb
abbbabbbabbbabbababaabbaababbaaabaaabbabaaaaabaabaaaaabbababbbaa
abbbabbabbaabaaabbbaaabbaabaaaaaababababaabaabaa
abbababbababbababbbbbabbbaabaaaaabbaabab
abbaababbbbaaaabbabbbbbabbbaabba
ababbababbaabbbaabbbbbbaababbbababbaabba
babbaaaabaababbaaaabaabb
aaabbabaaaababbaabbbbaaaaabbaabbbbabbbbb
aaabaaaaaabbaaaaaaababbbbaaaaaabbabbabaa
aaabbbbbaababbabbbbaabbbbbaaaabb
aababbbababbabbbabaaaabb
aaabbbbababbbaabaababbabbbaabbbaabbaabab
aaabbbbaababbbabbabaaaaabbbabaab
aabbbabbaaababbabababaab
aababaaaaaaaaabaaaaabaaa
baabbaaabbbaaabbbaaababa
aababbabbabbaabaaaaaaaab
abbaabbbbabbabbbababbbaa
aabbbbaaaabaabbbaaaabbba
aababbabaabaabaababaabaaaaababab
babbaaabaaaababbbaababaa
ababababababaaababbbaaaaabaabaaaabbabbba
babbaabaabbbababaaabbbbbbbbabababbbbbbbb
bbbbbbbaabbabbaabbbbabba
aaabbabbbbababbabbaabbaa
aababbabababaaabbabaaaaaaaabbaaabaaaabaa
aaaaababbabbbaaabbaaaababbbbabba
abbaabbbaaaababaabbbaaabbbbbabab
baaabbbbbbbaabbbabaaabaabaaaabab
aabbbbaabbaaaaabaabaabbbabbbababaabbbabbbaabbbbb
abbbabaabbabaabbbbabbbaabaaabaabbabaabbabbbaaaababbaabbabaabaaababbabbab
bbaabaaaabbaaaabaaabbaba
aaaaaabaaabaaabaaababaaaaaababaa
aaaababaaaaabbabbaabbabaabbaabbababaababbababbab
aabbbbaaaabaaabaabbabbba
babbabbababaaabbbbbbaababbbaabbabaababaabbbabaababbaaaab
bababbaaaaaabbaabbbbaaab
babbaababbbababbabbabbbbabbabbabbaaaabbaaaaababbbbaabbbaaabbabba
bbaababbbbbaabaabbbabbaabbabaaaabbbaababbaabbaab
bbabaaaaabbbbabaababbaaababbabaabaaaaabb
bbabababbbaaaababaaababb
aaaaabbbbabbaabaabaabbaa
abababbbaababbbabbbababbbaabbbab
aaaaaaabbaabaabbabbabbbb
babaaaaaaababaaababbbbba
aaaaabababaabaaaabaaaabb
bbabbbbaaabaaababaaaaabb
bbbaaabbabbbbaabbbbaabaaaabababaababbbbb
babbbaababbbbaabaabbbbbb
ababaaabbbabbababaabababbbbbabbb
bbbaabaaabbaabbbaaaabababbbbbabbaababaaabbbbbaba
aabababbbaaabaababaaabaababaabbb
abbbaaababaaaaaaababbbba
aaaababbbbbabbaaabaabaaaabbaababaabbaaab
aaaabaaaaabaabaabaababba
bbbabaaabbbababbbbabababababbbbbabbaaaaabbbaabaaabaabbab
abaabaabababaaabbaaaaaabbabbbababbabbbaabaaababbbbbbaabb
babbaaababbbbaabbaabbbab
bbbaabbbabaabaaabaaabbba
babaabaaababaaaabababaaa
babababbababbbabbaaaaabb
bbabbaababbaaaabbabbabbaaababbaabaaabbaabbaabbbabbbabaababbbabaabaaababb
abaabbbbbababababaaabaab
babbbaabbaaaabababbbbbabbbbaabba
aabaaabaaababaaaababaaabbabbbbbbabbaaaababaaaaabbbabaaab
ababaaabbbaabbbababbabaa
aaaaaaaabaaabbababaaaaaaabbbbbbb
bbaaaababbabbababbbbbaba
baaaaaaaaaabaabaaababbbabbaabbab
bbbaaabbbabbbbbbbabaabaabbaaabba
abbababaabbababbabbbbbba
bbaabababbabbbbaabbaabbbbbaaaabb
bbaabababaababbbabbbabaabaababbabbbabbba
aabbabbaaaaababbbaabbbbb
baabbabaababababbaaaabab
abbaabbaaabbaabaaaabbaaa
aabaaabaaaabaabababaaabbbbaaaaaabaabaabbbaabbbab
bbbabaaababbbabaabbbbbabaababbbaaaaaabbbaaabababbababaababbaabaa
aabbaabbbaaabbbbabaaaaab
abaabaababababbbabababba
babaababaaaaaaaaaabbaaaaabbaaabbaaababab
baaabbabababaaaaabbaaabb
abbaabbababbabbabaabbaaababbaaabbabbabbbbaaabaabaaaaabbabbaabbabbbbbaabb
bbaaaabaabbbbaabbabbababbbaabababbababbabbbbaaab
babaababbabbbbbbbaaababaabababaa
aabaabbbbbbababbbaaabbabababbbabababaaababaabbabaabbabbbabbbbabbababbbba
aaabbbbaaabaabbbabbbbbaabaaabbaa
abbaaaaababbabbabbbbbbab
abbbbaaaabbabbabababbaaabbbaaaba
aabbbbbabbabaabaaaaaaabb
baaaaaaababaababbabbbaabaabbbabbbbbabbaabbaaabbb
aaaaabababbabbbaabbabbaaabaaababaaabbabbbabbbbabbabbbabbbbabaaaaaaaabaab
ababaaabbaababbbbbbbbbab
baaabbbbbbaabbbaaabbabbabbbbabba
aabaabaaaababbbabaabbbbb
bbbabbabaabaabaababaababbbbaaabbbabbbaaa
bbabaaabaabbaaabbbabbaaababaabaaabbabbbabbbababbbbabbabbaabababbbbbbbbbababbbaabaaabaaabaabbbaba
aaabbabbabbababaababaaba
aababbabaabbaababbaababbbbbabbaaabaabbab
aaababbaabaaaabaabaaabba
abaaaababaaabbbbbabbbbaa
bbaaaabaaaabaabaaabbaabbabbabbbb
aaabaaaabbbbaaaabbaaaabb
bbbababbaabaaaaabbbaaaaaabbabbbbaabbbaba
aabbbabbabaaaabaaaababaa
abbbbbbaabaabbbabbabbbabbbbaabaabbaabbbbbbbaabaabaaaababaabbabbaaaaababbbbbbaabaabbbbbaa
ababababaabbbaaaaabbaaaababbbaaabaabbaabbaabbabbaaaabbbabbbbabbb
bbaababababbaaaabaaabbbabbbbbababababaab
babbaabababaabbbbbaabbaa
aabbaabbbbababababbabbba
abaabaabaaaabababbbbaaabbbbbaabaabbbaabbbbaabaaa
ababbbabbabbabbbaaabbbab
aaaabaaaaaabbbbabbbaaaaaabaaababbaabbababbbaaaaa
bababbbaabbaabbabbbaabab
abbbbbabbabbbbbbbbabbaabbaabaaaaabbbbababaababaaaaaaaaab
bbbbbbbaaaabbbbaabbabaaabaaaaaaabbbabbbbbbbbaabbabaaabba
abbbbaababbabbbbbbabbaaaaabbaaab
bbabbbbbbbabaabaaabbaaababbabaaa
abbababbabaabaaabbaaaabbbbbbaaababaabbbaababbbba
aabbabaababbbabaaaaaaaab
baaababbabaabbbbaaabaaab
bbbabababbababbbbabbbbbbabbaabbbbaaaaaba
aabbabbabbbababbbbaababbabbabababababbba
abbabbaabaaaaaaabbbbbaab
aaaabbaabbabaaaabbbbaaaababbbaabaabaaaaaabaaabbb
baaaaabbbbabaaabbaabbbabababaaababbbbaabbaabbaaaabbabbab
bbaababbabaaaaaaaaabbaaa
abbbaabbabaaaabbbababbbbaaaaaaabbaabaababbaaababaaababab
babbbbabababaabbbabbaabb
abbabbbaabbaaabbaabbbabaabaaaaabbbaabbabaabbbbaaaabaabbb
aabbabbbabbaaaaabbaabbbabbabbabbaaabbbbabbabababbbbbbaaaabbaababaabbbaba
aabababbbabaababbaababaa
aababbabbbaaabbaaabbbbabbbaaaaabbabababa
aaabaabaaaabbaaabbaabaaa
babaabbbbbabaaaabbaabaab
bababbaaabaaabbabbabaaab
aababbabbbbaabbbababaabbaaaabbabaababbba
babababbbbaababbbabaabba
abaababbbbabaaaaabbbaabb
aaaabbabbabbabbaaaababaa
bbbabababaabababaaaabaab
aaaaaaaabbabababbbaabbbaaabaabaaababaaaa
bbabbbbabbbabaaababbbaabbbaabaaa
ababbbabbabababbbaabaaaaaabbaaab
ababaabbaaaaababbbaaabab
aabbbbbbababbabbbbbaaaaa
aaababaabbababbbbbabaaaabbabaaabbabbabbababbbbabbbbaabbb
bbbabbbbbaabababbbaabbab
abaaaababbaabbbabbaaaabb
bbaaaabbbaabbbbbabaaabba
baabaaabaaaababbaabbabbb
aabaaaabbbabbaabaaaaababaaabbbaababbbbabbaabbbbaaabbaaaaabbabbabaaabbabaaaabbbbb
bbbbbaabbaabaaaabbaaababababbbba
abbbabbbabababbbababbbba
bbabaaaabaaabbbbbababbaaaababaaaababbabbaababbaaaaaabaab
aabaaabaaaaabaaababaaaaaabbbabbbaaabbbaa
bbbbbbbaabbbaaaabaaaaabb
aaaabbababbababbabbabbaababaabbbbaababababbbbbbbbaaaabab
aabbaabababaaababaabaaab
abaaababababbbabbbabbbaaaabbababbaaabababbabaabaaabbbbbbbbbbbaaaababbbbaaaaaaababababbab
ababaaabbaaaaaabbaaaaaabbabaaaabaaabaaaabaaababb
abbaabababaabbaaaabbbabaaaabaaaaabaaaabbbaaaabaa
aaaabbbaaaabbbabbabbbbbabbaaabba
ababbbaaaabbabaabbaaabba
abbaabbbaabbaabbaaaaabaaaaabaabbbbbabababababaabababbbbababababb
abababbbbbbababaaaaabbabaababbbb
aaaabbaabbaababaabbaaaba
abbbbbaababaababbbaabbab
babaabbbabbbabbbbbaabbbababbbaaa
aabaaabbababbbbbbbbbbabaababbbbababababbaaaaababbbaaaaabaaaababb
bbbababaaaaabaabbabaabbbbbaabbaabbbbabab
bbbabbaaababbbbbbabbaabb
bbabbaabaabbaaaabbbbaabb
aababbabbbabbbaababbabbbbabbbaabbbaabbbb
ababbbaaabaababbaaaabaaaaaabbaabaababbaa
ababaaaababbbbabbabaabaabbbababbbbabaaab
abaaaaaaababbbabbbbbbbab
aaabaababbabbababbaabaab
baababbaabaababbabbabbab
ababbbbbababbababbabbaabababbbbaaaabbaaabbbabbbaaabbbbbb
babbbababbbbbbbabbabbbbb
abbbbbaababababbbabbabbaabaabaabaabaabba
ababaaababbbaaaaabbbbaaaaaababbaaaabbaaa
babbabbabbabaaabbbbabbbbbbbaabbabbaaaabb
abbaabbbbabbababbabaaabbaaaabbaaabababbbbaaabaabaaababab
babbabbabbbaababbbababababbbababbaaabababbababaaaaaaabbaabbbaaabbabbabbabbabbbba
abbaabbaaaaabbaaaaababab
abaaabaabaaaaaaaabbbabbbbabaababbabbabbbbbbbabaa
aabaaaaaaabbbaaabbaaabaababaabab
aaababbabbbaaababaababaa
babbbbaababbbabbaabbbbabbabbbabaaaaaabbbabbbbbabbbabaaababbaaaabaaaababb
bbaaaaabaaaaaabaaaabbaba
aaaababbabbababbbbbabbabbabaaaaaaabaaaabababbbba
bbbababaaaaaaaaabbaaabba
aabbabbabbaaaaabbaabbbab
bbbabbabbaaabbabaabbabaaabbbbbabbbaaabab
bababbbaaaaababababaababbaababaa
aaaabbaabaaaaaaabbbbabaa
aabbbbaaaabaabaaaabbbabbaaaaaaab
ababbaabaaababbabbaaaabaaababaab
babbbbbbaaababbbabbabaaa
abbbabbbbabbaaaabbbaabab
baaaaaabbbababaaabbaabaaaababababbbbbaab
abbbbaaaaaabbbbbababbabb
ababbbaaaaababbaaaabbabaababbaabbaabbabbbaabbbba
bbaaaaaababaaaabbbabaaba
baabbbaabbaababababababa
bbaaababaaaaaaaabbbbbabbbabbabaaabbaaababbbaaaabbbababaa
ababbaaabbbbbbbbbbbaababbbabbbbaaabaaaaaabbaaaabababbaba
abbabbaaaabbbbbaabaabbbb
abbaaaaaaaaababbbabbaabbbbbaaaaabbaabaab
aabaaabaabbbababbbaaaaaabbabbbba
aaabbbaabbaaabababaababbabababbbbaaabbab
babbaaabbaabbbaabaabababbbabbbabaaabaabb
baaaabababbbbaaaaabababaabbbbabbbbbbbbabbbaabbabbaabbaaaaabbababbbbababbabaaabaabbabbbba
aabaabaabbaaabbbaaabbbaa
bbbababbbaabbbaabbabbbab
bbbbbaaaabaababbabbbabba
bbaaaabaabaabaaaabbbaabb
ababbbbbaaabbababbaaabba
babbaaabbbaaababbbbbbbbb
bbbababbbbababbbaaabbbab
aaaaababaabbbbaaabaaabaabaabaabababbbbaa
bbbbaaabababbbbabaaababa
abbbabaabaabaabaabaaabba
babaabaaababbaabbbbbaaaaaabaabaaaabaabba
ababaaababaaabaaaabaaabb
bbabaaaaaabbaaaaaaaaaaab
bbbaaabbbabaabaaaabbbbaaaabbbaba
abbababbbabbaaaaaabbbaab
aaaababbbabbaaabaabbbaaabaababaa
bbbbbbbaabbbbbaaaabbabbabaaaabbabbbbbbab
abaaabaababaaaabbabbbbabbbbaabaababbabaabbabbbbabbabaaaabbbbabbbbbabababaaabaaab
abbabababaaabbbbbbbaaaab
aabaabbbabbbbaabbaaaaaabaabbaabbbbbabbbbabaaaabbbaaaabaaaabaaaababbaabaa
bbbaabbbaaabaababbbbaaaaaaaaaaabababbaaaaaabbaaabbbbbbabaababaab
baaabaabbabbabbaabbaabba
bbbabbaabbaaaabaaaabaabb
babaabbabbbaabababbaaabaabaaaaaababaabbbaabaaabaaaabaaba
aaaaaabaababbababbbbbabbabbbbbaabaabababaaababaaaabaabba
babbbabaaaababbbbaababaa
abbbababbabaaabbabbbaabaababababbbabbbabaaaabbbbaabbbbba
baababbababaaaabbababaabbbaaaaba
abbababaaaaabaabababbbabaaaaaaaaababaababaaaaaababbbababaabaabbbbbabbbaaabbbabba
aaaaaaaaabbbaaaaabaaabab
bbaaaaaababaaaabaabbaabb
abbbabbbbabbaaaabbabbbaabbbbbbaaaaabbaabaabaaababbaaabbaaabaaaaabbaababbabbbaabbaababaaa
bbabbbabaabbbaabbaaabaab
aabababbbabaaaababaaaaaaaaababbbaaaaabaabbaaabbabbbbaaabaabbaaaabaabbabb
abababbbbbbaabababaaaaabbbaaabbbababbbab
aabbaaaabbabbabbabbbbbabbbabaaababbbbaaabaaababaaabbbababbbbabbbbabbbbbbbbaaaaabaaabbabbaaabbbab
bbaabbbababbaaabbababbbb
abbbbaabbbaaaababbaabaaa
abaababbaaaababaaabbaaaaaaaaababaabbaaab
bbbbaaaaaabbabaabbaaaaaaabbabbab
bbaaababaaaabaaababaaabaabbabbab
bbaabbaabaabaabbbbbabababbbbabbbbbbbabaa
bbabbbbaaaaaabaabbbaababaabaabab
baaaaaaaaaabbabbbabbbaabbaabbabaabbbababbbbbabaabbabaabbbbbbaaab
bbbaabbbbababaabaababbbbaabbabbbaaabbbbabaabbbab
aaabbabbbbaaaaaabbaabbbaaaabaaaabbbbaabbbaabbabb
babbababbabababbbbbbbbaabaabbabb
babbabbbbaaaaaaabbabaabb
ababaabbababaabbbababbbaaaabaaaa
aababbabbbaaababaaabbbaa
aabaabbbaabbaaaaababbababaaaaabb
bbbaabbbabbbabaabbabbbab
babbabbbaabbaaaabbababbbbbaabbbaaaaabbbbaabaaaab
aaabaababaaabbababbabaabbbabababaababbbbabbbaabaaabbbbbb
abbabababaaaaaabbbaaaaaababaabaabaabbaaaabbbaabaaabbabbbbaaaaabbbaabaaab
aabaabaabbabaaaabbbbabba
bbbbaaaaaaaaababbbaabbbaaaaabbaaaaabbbbbaabbbabaaaabaabbaaaaaabbabbbbabb
ababbaabbbbbaaaabbbaabbaabbaaabb
aaaabaaabbbabbbbaaabaaab
bbabbabaabbbaaaabaababbabbbbaabb
baabaababbabaaaaababbabaaaabbbbaaaaaabaa
aabbbabbaaabbbbbaaaaabaa
aaaaabbbbbaabababbbabbababababbbabaaabab
ababbbaaabaaaababbbbbbaaabbbabbaababababaaabbabaabbaabbabbbababbaabaabaa
baaababaaabaababbaabaabb`

var ExampleInput1 = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

var ExampleInput2 = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

var ExampleInput2R = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31 | 42 11 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42 | 42 8
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

func ReadRule(input string) (RuleName, Rule) {
	split1 := strings.Split(input, ": ")
	name := RuleName(split1[0])
	choices := strings.Split(split1[1], " | ")
	ruleChoices := make(Rule, len(choices))
	for i, choice := range choices {
		tokens := strings.Fields(choice)
		ruleChoices[i] = make(RuleList, len(tokens))
		for j, token := range tokens {
			switch token[0] {
			case '"':
				ruleChoices[i][j] = RuleLiteral(token[1 : len(token)-1])
			default:
				ruleChoices[i][j] = RuleName(token)
			}
		}
	}
	return name, ruleChoices
}

func ReadRules(input string) Grammar {
	blocks := strings.Split(input, "\n\n")
	lines := strings.Split(blocks[0], "\n")
	grammar := make(Grammar)

	for _, line := range lines {
		name, choices := ReadRule(line)
		grammar[name] = choices
	}

	return grammar
}

func ReadWords(input string) []string {
	blocks := strings.Split(input, "\n\n")
	return strings.Split(blocks[1], "\n")
}
