package day20

import "strings"

var Input = `Tile 2789:
.#..#...##
##...#....
#.#.##..#.
.#.##.##.#
....#....#
#.#......#
.....#....
#.........
...#.....#
###.####.#

Tile 3041:
###.#.....
#.........
.......#.#
..#.#.....
.#.#...#.#
#.#.#..#.#
.#.....#.#
...#.....#
......#.#.
####.#.##.

Tile 2663:
#..##..#..
..........
#...#.....
#.....#...
#..#.#...#
..#......#
....#..#..
.........#
#.........
.####..#..

Tile 2477:
#....#..##
.#........
#.........
.###.....#
..##....##
........#.
.....#..#.
.....#....
......#...
#######.##

Tile 3673:
........#.
###.......
#..#.#..#.
#.#..#....
#.....#..#
#........#
..#...#.#.
...#....#.
#..#...#.#
##.####..#

Tile 1031:
.####...##
..#......#
#........#
#..#......
....##.#..
##........
......##..
##.....#.#
....#.....
.##.###.#.

Tile 3119:
....##.#.#
#........#
#.....#...
#..#..##.#
#..#...###
#.#.###...
..#......#
...#.#.#.#
..#..#...#
##...#..##

Tile 3347:
#.##..#.#.
.#.#.#...#
##..#..#..
.....##..#
.....#...#
....#.....
.......#..
..........
..#...#...
##..##.#.#

Tile 1747:
#.##...#.#
##..#....#
......#.#.
#..#....#.
#..##...##
#.##...#..
..#.#..#.#
..........
....##.##.
#...#....#

Tile 1693:
..####...#
#.....#.##
.....#....
#....#....
...#.....#
..........
#.#....#..
#..#.....#
#....#.#.#
...####.#.

Tile 1093:
###...#.#.
#.....#.#.
.......##.
..#.#...#.
#..#...###
#..###.#.#
....#...#.
.......#.#
...##.....
#.####..#.

Tile 1223:
.#...#..##
........#.
#..#......
#.##.....#
###.##...#
##...#..##
#.#.#.....
#.##......
#....#....
###.######

Tile 1307:
.#..#....#
###..#...#
..#.......
........##
.#......#.
###.##....
##.......#
##....#...
..##.#.#..
##..#..#.#

Tile 1123:
###...#...
#....#...#
#......#..
#.#..#...#
.#.#..#...
#........#
.#...#.#..
........##
...#......
######.##.

Tile 1019:
..#..#.#..
.#..##....
#####.#..#
#...#....#
#....#...#
#.........
#.....#...
.....#..#.
....#....#
#...#.#.#.

Tile 2083:
..####....
....#.....
...##...#.
.#.....###
#..##.##.#
..........
#.......#.
#....#.###
.##....##.
#######..#

Tile 1361:
..##.##...
##..#....#
#.........
##....#...
...#.....#
#....#..#.
.........#
#....#....
#.....#...
..########

Tile 2927:
#..#...#.#
#..#..#...
..##.#...#
#...###...
##...#.#.#
......#..#
.........#
..##....##
..#......#
#..#######

Tile 2111:
.###.#.##.
#.........
...#...#.#
.........#
.....#..#.
##........
..##.....#
##........
##......##
##.##.#..#

Tile 1481:
....#.....
#.........
.........#
#.........
#.....##.#
...##....#
#.#...#...
#..#.#...#
#...#.#...
#....#.#.#

Tile 3931:
#..#....##
.....#....
......#...
...#.##..#
...###.#.#
#..#.#....
.##.......
#....#....
#...#.#..#
#.##...###

Tile 3697:
#.#.######
##...#....
##.....#.#
#........#
...#......
#.#.......
...##.....
.........#
.#...##...
...#..###.

Tile 3089:
#.#..#..##
.........#
#.........
.......#..
#.......#.
.......##.
#....###..
..........
....#.....
##...#.#..

Tile 1823:
...##...#.
...#......
..##....##
...#......
.#.#......
###...#.##
#.....#...
#......##.
..#.#.#.##
##.....###

Tile 2843:
#####.#...
........#.
#..#...#.#
#.......##
##..##.#..
#.........
........#.
.......#..
.#.##....#
..#.#..#.#

Tile 2377:
.#.##..###
...#..##..
.#...#..##
...#...#..
....#...##
##....#..#
..#.....##
#..#......
..#..#...#
..##.#...#

Tile 2069:
#########.
#.......#.
#......#.#
#........#
#....#.#..
#.........
#........#
.##......#
#..##..#.#
.####.....

Tile 2551:
..##..####
.##....##.
.#.......#
##.....##.
....#####.
...#.....#
...#.#.###
...##....#
....#....#
..#..#..#.

Tile 1193:
###...####
##..###.#.
.#.#.#...#
..##...##.
#...#....#
.......#.#
#.......#.
..#.#.....
#..##....#
###.#..#.#

Tile 1451:
#.#.###.#.
#...#....#
....#.....
.........#
....#..#..
..........
.#........
........##
.........#
#...##..#.

Tile 1567:
#..##....#
..........
...##..#.#
#..#####..
#.......##
##..##.#..
##.......#
#.....##..
....###..#
.#.#.####.

Tile 2707:
#####.#.#.
......#..#
##........
#.....#..#
...#.....#
...#......
.##.....#.
#......#..
#..####..#
..#..#..#.

Tile 3863:
.#.###...#
...#......
.#.#..#..#
#..#..##..
..#.#..##.
......#.##
##..#.....
#......#.#
#.........
.....##.##

Tile 3331:
#.####..#.
#...#..###
#..#...##.
...#...#.#
#...#...##
#.#......#
.........#
....#...##
#..#...#..
#.#...####

Tile 2423:
#..##.#.##
#....#....
..........
#..#.....#
..........
#...#.....
......#..#
...#..#...
..#.#....#
...###..##

Tile 2207:
...#...#..
#....#....
#..#.....#
........#.
#........#
.......#..
##....#..#
........#.
#.....#..#
....#.....

Tile 1913:
..#.#..#..
#.........
..........
.........#
#.......#.
#..#...#..
..........
...####...
...#.#..##
#######.#.

Tile 2963:
#..##...##
..........
...#.#....
#........#
#.#....#..
..#..#...#
#....#...#
......#..#
#.......#.
#.#.#.#..#

Tile 3733:
##.##.#...
...#.#....
......#...
#.........
#.........
#.........
..#...##.#
.....#....
#....#....
##.#...##.

Tile 2239:
...###.#.#
##....#...
##.#......
#........#
#..#.....#
..........
....#.#.##
##........
........#.
.#.#####.#

Tile 1063:
..#####.##
#....#...#
.........#
#..#..#...
..#..#...#
#......#..
..##.#..##
.........#
#.#.#..#.#
...#..####

Tile 2749:
..#.##.#.#
...#.....#
.......#.#
#.#.......
.........#
..........
.........#
###.#..#..
#.##...##.
###.#....#

Tile 1187:
##.#.....#
#....#...#
#.#..#....
#.###.....
.......#..
...#......
##........
..........
..........
.#########

Tile 1613:
####......
#.#.#...#.
##........
#......#..
..#.##...#
.###.#....
...#.....#
..........
#........#
.#........

Tile 2521:
##...#..#.
.#....#..#
..##...#..
..#.#.....
#....#....
....##....
#.#.#.....
.....#...#
........#.
.#..##.#..

Tile 3671:
###..#....
.#.....###
#...#..#.#
......##..
##..#.##.#
...#.##..#
.##...#.##
.#.#.#.#.#
..........
.#...##...

Tile 1459:
...##..###
...##..#.#
..##..#.#.
#.##..#..#
....#....#
#.##.##..#
##......##
....##.###
...#..#..#
###.#..#.#

Tile 2999:
###..##...
.........#
#.....#...
.#.#.....#
#.#.#..###
.#.....#..
#..#......
..#.#.....
#..#.#.#.#
.####...#.

Tile 1973:
#.####.##.
#...#....#
..##..#...
.........#
####..#...
#.##....#.
#.#.#..###
....##...#
#..#.....#
.###.#..##

Tile 1999:
###.##...#
.........#
.......#..
........#.
#.#.....##
....#.....
...#....##
#...##....
..........
..##..####

Tile 2371:
###.##...#
...#......
..........
...##....#
........#.
#.#..##...
......#...
##......##
#........#
.##.###.##

Tile 3793:
#..##..#..
#.......#.
...#.#.#..
.#......##
.......###
.#........
.....#....
....#....#
....#..###
.#.###..#.

Tile 2791:
#..#####.#
#...##....
#........#
#.....#...
........##
...#.....#
..##.#...#
........#.
#........#
.#..#.####

Tile 2903:
##....##.#
.#.......#
.#..#.....
#...#.###.
#..#.#.##.
#.........
......#...
.#.....#.#
.....#...#
.#......##

Tile 1619:
.....###..
..##...#..
#.#.#....#
##...##...
#.....##.#
.#...####.
.#.......#
#......#.#
#.#..##...
..#.##...#

Tile 2267:
##....###.
#...###.##
..##.#...#
##.....#..
#.....#.#.
.#.#...###
.##....#.#
#.#......#
#.#..#....
..###..#..

Tile 1429:
#.#..#.#..
.......#.#
#.......#.
....#..###
#.#......#
#...#.##.#
#.....###.
#.....#.##
.....#....
###....##.

Tile 3593:
#..##...#.
..#......#
#......###
....#.....
#.........
##.......#
##........
..#.#....#
...#.#....
###.######

Tile 2801:
.##.#.....
.........#
#........#
#...#.#..#
.##...#..#
..#.#.....
#.........
#.##..#...
.....#...#
#...##.#.#

Tile 1997:
.###......
#......#..
#.....##..
..#..##...
##.....#.#
.##.....##
..#.....##
.......#.#
......##..
..#..##.##

Tile 1499:
....#.##.#
#.#.....#.
...#.#...#
.........#
#.........
..........
....#.....
#........#
..#...#...
#..##.###.

Tile 3191:
.#..###.##
.....#..##
..........
.#..#..##.
#....###.#
.....#...#
..#.#.....
.....#.#..
#...#.#..#
...###....

Tile 3121:
.#.#.#..#.
...#..##..
#........#
.#...#..##
#.#.......
.....#....
..#.#....#
#.....#..#
#..###...#
.#####....

Tile 3023:
.....##..#
......#.#.
....#.####
#.##.....#
..........
#.........
......#...
#......#..
...#....#.
....#.##..

Tile 2531:
..#..##...
..#..#...#
.....#....
.#..#..#.#
#..#..#.#.
........##
.......#.#
###.....##
#..##.#...
##...#####

Tile 3001:
#####...##
#.#..#.###
...#.....#
#....#....
##.....#..
.##.....#.
#.#..#....
....#...#.
#.........
..#.#.##.#

Tile 1399:
#.##.###.#
...##.....
...#.....#
#.#.#....#
.#.#....#.
##..#....#
..........
...#.....#
.##.....#.
..###....#

Tile 2251:
#..#..#...
.....#..##
..##.....#
#..##....#
##.##.###.
...#....#.
..#..#....
...#...#..
..........
...#.....#

Tile 2659:
.###.##..#
.........#
..#.....#.
...###...#
.#...#...#
#..##....#
#.......#.
......#..#
#...#..#.#
##.##.####

Tile 3067:
#..###.###
#.##......
....#.....
.........#
...#....##
......#...
.......#..
#.#.#....#
#..#...#.#
......##.#

Tile 1163:
..#.#....#
#.........
.....#...#
.#.......#
#...#..###
#.......##
#..#..###.
#...#....#
.........#
.#.#.#....

Tile 3433:
#.##...###
#.........
#....#..##
#..#..##..
........#.
###...#.#.
.#..##.#..
..##.....#
..#..#....
####...##.

Tile 3637:
###..###.#
..........
.......#..
.....#....
#.....#...
......#...
.#........
.#..#.....
#..#.#.#.#
######..#.

Tile 3253:
#..##.#...
.......###
.#........
#........#
..........
...#....##
##.......#
##.......#
#.#......#
########.#

Tile 3271:
#.##.#...#
#..#..#...
......#.##
#.....#.#.
.....#..##
#..#.....#
..........
#...#.##..
.......#..
#..#.##..#

Tile 2657:
.#..#..###
#...#.....
.....#....
###..##...
#...#..#..
#..#.....#
.#........
....#..#..
....####.#
.#.#.#...#

Tile 1087:
.#.#..####
..#...#..#
...###.#.#
.##.#....#
...#.....#
...#.#....
##..#.....
##.###..#.
#.####.#.#
##.....#..

Tile 1721:
.#.#...###
.....#...#
..#..##..#
#..#..#.#.
#.#......#
.#....#...
..#..#.#.#
#..#..###.
#........#
.....#.###

Tile 3557:
#...##.#..
..###...#.
..#.##....
##.#..#.##
#.......#.
.###.##..#
##.#......
#..#...###
..###....#
##.##....#

Tile 1471:
...##.#..#
....##...#
.#...#...#
......###.
....#.#..#
##...#...#
#...#.#...
##....#...
.........#
#.#..##..#

Tile 2237:
..#.####.#
........##
#.........
....##.#.#
.##....#..
..#...#...
##..#.....
...#....#.
.......#.#
##...#.#..

Tile 3779:
##..#.##.#
..##.....#
#.........
...#...##.
..###...##
.##..#.##.
#......##.
.#.......#
....#.....
###.....#.

Tile 1129:
.#..#.#..#
....#.....
##..#.#...
#.....#..#
#.#..#..##
#....#...#
.#.#...#..
........##
#.##......
#.....##.#

Tile 3929:
###..#..##
#........#
..#..#....
.........#
#.#.......
..........
#..###....
#.#..#..#.
#.........
###.......

Tile 1871:
###...####
#....#...#
#.........
#...#.....
...#...#..
#........#
#......#..
....#...##
.......#.#
#...######

Tile 2137:
#..#..##..
.....#..#.
.....#....
....###...
#.#...#...
.....#.#.#
######..##
#..##.#...
.......#..
.#.###....

Tile 2591:
####.####.
.#......##
#.....#...
..........
#.#..#.###
.#...###.#
#.........
#....#...#
#.........
#.#...#...

Tile 3049:
..#....#.#
###.#.....
#.#.......
....#....#
.....#..#.
.........#
#...#..#.#
#...#.#.##
.....#....
..#...#..#

Tile 2753:
##.#.#...#
#..#.#.#..
#........#
#...##....
....#.....
#.##......
.......###
#....#....
...#...#.#
###..#..##

Tile 1951:
.#.##.##.#
.........#
......##..
.....#.#.#
###..#....
##..#.#..#
......#...
#.........
.#.#.##...
#.#.#..#.#

Tile 2339:
.......#.#
.........#
#.#.....#.
#.#.###..#
#.....#...
..........
...##.....
##......##
..#.#....#
.#...#...#

Tile 1039:
.######..#
.#.......#
#.........
#..##.#..#
#.......#.
#.........
...#......
#..#..#...
...#.....#
.#..###.#.

Tile 1249:
#.####.###
#.......#.
#...#...#.
......#.#.
#.#.#....#
..#.#....#
.....#.#.#
#....#....
#.##.#..#.
#.#.#.##.#

Tile 2579:
#.##.###.#
#.#.#....#
#.........
##........
#........#
#.........
.##.......
..#.......
#.....#..#
##.#..##..

Tile 3499:
#.##.#..##
#.....#.##
....#...##
.....#...#
......#...
.......#.#
.#....#...
......#...
#.....#...
....#.###.

Tile 2879:
#.#.#..###
#........#
#....#.###
.#.#.....#
....#....#
....#...#.
#..#..##.#
#.........
#.....#.#.
.###.#..##

Tile 2027:
#####..###
......#.#.
........##
.....#.##.
#.........
#..#...#..
#.#......#
####....##
#...#.##.#
##.#.#.#.#

Tile 1987:
##.##.#.##
#...##...#
.........#
#.#.###..#
.#.##....#
....#.....
.....#...#
..........
.........#
#..#..##.#

Tile 3659:
#.##.#.###
..#.....#.
......##..
.........#
....#.....
#..#..#.#.
#.........
##........
#.........
#.##.....#

Tile 2297:
##.##..#..
...##.....
#..#.#.#.#
#....#...#
....##....
#....#...#
##......#.
#....#..##
##...##...
###..##.#.

Tile 3739:
##...###.#
#......#..
#....#....
.###..#...
...##...##
#..#......
....#.##..
#.....####
##.#......
##.#...##.

Tile 3539:
.#....##.#
##....#...
#......#..
#.....#.#.
##.......#
##...#....
......#...
#..#......
.....#####
..###.....

Tile 3329:
#.###.####
..#.......
..........
..........
#.#.......
##......#.
#......##.
#..#....#.
.....#...#
#.###.#.#.

Tile 1201:
#..#..####
#.#...####
.#...##..#
##.#.#..#.
..........
#...#..#..
..#..##...
...##.#...
#.....#...
##.#.#..#.

Tile 3559:
##....#.##
#.....#..#
##.......#
#.#.......
#.........
........#.
##.......#
......##.#
#.##..#...
#.....#...

Tile 3259:
...#.#####
.##.......
#..#.#...#
....#....#
..........
.#.#.....#
.#....#..#
.........#
#.##.#...#
#..#####.#

Tile 3407:
...##.#..#
.#..#.....
#.....#...
#.#.......
#.......##
#....#...#
..........
#....##.#.
#.##...###
.#..##.#.#

Tile 1051:
...#...#.#
..........
.........#
..#....#.#
#.#####.#.
...#.#...#
..##...#.#
......#.#.
#........#
##..##.##.

Tile 3467:
.#.##.#..#
#...##....
#.#..##.#.
###...##..
#.#.#..#.#
.....#....
..........
...#..##..
###....##.
.####..#.#

Tile 1847:
....#...#.
.....##..#
.#........
.#.......#
........#.
#......#..
#..#.##...
.##...####
........##
.#...###..

Tile 3623:
#....####.
#...#..#..
.....##...
#.....#...
#......##.
#.##.#...#
.......###
.....#.#.#
......#..#
######....

Tile 1907:
##...#.#.#
###.#..#..
#..#....##
..#.....#.
........#.
#.#...#...
.#.......#
#......#..
.....##..#
###..#.##.

Tile 2837:
#..#....##
.....#...#
....#.#...
#....#...#
.....##.#.
#...#.#.##
.....#....
.#...#....
..#.#.....
###....##.

Tile 1571:
##..####.#
..........
##..#.....
#.........
#....###..
#.#.#.#..#
#..###...#
#.....#..#
#........#
.#####....

Tile 1097:
#...#.##.#
#.##.....#
#....#...#
...#...#..
.###.....#
...#...#..
.#.#..#...
#.#....#..
.........#
.#.##....#

Tile 2803:
###...##..
#.........
.......#..
.....#...#
##........
.......#.#
......#...
##....####
#.....#.#.
.##..####.

Tile 1319:
.#.###.#..
...#......
...##.#...
.##.##.#.#
.#.......#
.....#...#
#.##......
#......#..
##.......#
....###.##

Tile 1663:
##.#..####
##...#....
#....#..##
.........#
####......
.#...#....
#........#
#..##.#...
#...#....#
##.##.#.##

Tile 1709:
.#..##..##
..#.#.....
.....#....
#........#
.#..#..#.#
......####
#.###...##
..#..#..#.
.#.......#
###..##..#

Tile 1171:
#..###.###
#.#......#
#....#....
....##.#..
#...#....#
###......#
...#.#####
...##....#
.........#
.#..####..

Tile 3469:
##....#...
..#..##...
...#..#...
#.#....##.
#.........
#.........
.....#...#
..#.......
.....#...#
.######.##

Tile 2311:
.#..###.##
.#.......#
#.#.......
#.........
...#...#.#
.........#
#..#......
.....#....
.##....##.
####.#..#.

Tile 2539:
#.########
#......###
#......#.#
#...###..#
#....#..#.
.#.....###
.#...#....
###.......
.........#
.####.####

Tile 1439:
.#...#####
.##.#...#.
##...#.#..
#.#....#.#
..#...#.#.
.........#
..#.#....#
.##...#.##
##...#.###
#.#..#...#

Tile 2383:
#####..#.#
#.#......#
..#......#
..#...#...
#.........
#.........
##.#.....#
#.#..#....
..........
#.##...#.#

Tile 1559:
...#....#.
#.........
...###.#..
#.##....#.
..#....#.#
##.#.#..##
#........#
##.....#.#
#...#....#
.#.##.#..#

Tile 3061:
###..#####
.#....#..#
......#.##
#...#....#
#...#.....
...#...#..
.....##.#.
#......#..
......#..#
#.##....##

Tile 1511:
#.##..#.#.
#........#
##.#..#...
.#.......#
.##.#...##
#..#.#..##
..........
.##...#.##
#........#
..#.#.....

Tile 3343:
#....#.##.
.........#
#.....#..#
#.......#.
#........#
.........#
.....#.#..
.##.......
#.....#...
#....###..

Tile 1277:
#.#...#.##
..##.....#
....##....
...#....##
##.......#
#....#...#
#........#
.........#
#........#
.###.#....

Tile 1291:
#.#..##...
####..#...
#..#.....#
.........#
#..#.....#
..........
##.##....#
.#.#...###
.#.#....#.
.###....##

Tile 3109:
.#..######
........##
..#.....#.
###....#.#
#...#.....
.........#
...#......
#........#
.#...#...#
.#..##.#.#

Tile 3821:
....###...
#...##...#
#.#.##....
#........#
.......###
...##.#..#
.....#.#..
....##...#
...##.....
#.##..#..#

Tile 2677:
...#....##
#.#.#....#
..........
.#..###...
##..#..#..
...##.....
#.....#.##
..#......#
##.#.....#
####.##..#

Tile 1933:
#.#.#...##
#...#..#..
#.##...##.
.......#..
#.......##
#.......#.
.......#..
#....#..##
....#..#..
#.####.#.#

Tile 2819:
....#...##
...#......
#......#..
..........
....#.....
....#....#
.......#.#
..#......#
#.....#..#
.#...##..#

Tile 1733:
..##..##.#
......#...
...#..#..#
..#..#....
#......#.#
#........#
##.......#
....#..##.
##.....##.
#....#..#.

Tile 1783:
##.#....##
.........#
..........
......#..#
#.........
#..#.#....
#....#.#.#
##.......#
..........
.###..##.#

Tile 2473:
.##.##..##
#......#..
#........#
#.#..#.#.#
#..#..#...
#........#
..##......
..#.#....#
......#.##
#.###.#.##

Tile 2767:
#......###
..........
.#.#....##
#..#......
#......#.#
..###...#.
#.....#.#.
#...#.....
..........
###.....#.

Tile 1091:
##..###.##
#.....#..#
#.##.....#
..........
#........#
#..#......
.........#
#...#..#..
.###....#.
#...#.#..#

Tile 1609:
#.#.#....#
#.....#...
..#...#..#
##.......#
#...#.##.#
#..#......
........#.
.#.......#
#..#......
#.#...#.#.

Tile 3769:
...#.##.##
##........
#....#....
..........
.....##.##
.#........
#...##...#
#.........
......#...
#####..#..

Tile 2011:
.#.####.##
##..#.....
#..#.#..#.
#.#.....#.
#........#
#.....#..#
..#......#
#......#.#
..#.....#.
.#.#.##.##`

var ExampleInput1 = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`

func ReadInput(input string) []*ImageTile {
	blocks := strings.Split(input, "\n\n")
	result := make([]*ImageTile, len(blocks))
	for i, block := range blocks {
		result[i] = NewImageTile(block)
	}
	return result
}