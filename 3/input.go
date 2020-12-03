package main

var ForrestSlice = []string{
	"....#............#.###...#.#.#.",
	".#.#....##.........#.....##.#..",
	".#..#.#...####.##..#......#..##",
	"......#...#...#.......#........",
	"........#...###..#.#....#....#.",
	"..##.....#.....#.#.........#.#.",
	".##.......#.#.#...#..#...##...#",
	"...##.....#....##....#...###.#.",
	"..#...#......##.#.##.....#.#..#",
	".#....#.###.........#..........",
	".#.#..##.....###.....###....#.#",
	"....###....#......#...#......##",
	"......##...#.##.........#.#..#.",
	"##.#....##...#..##....#.#..#.##",
	".#...#..#.....#.#.......#...#..",
	"..........#..###.###......##..#",
	"..#.##.#..#......#.......###.#.",
	"...#...#.#.#..#...#.#..........",
	"#........#..#..#.#....#.##..###",
	"#...#.....#..####.........####.",
	".....###..........#.#...##...#.",
	".....#...#..#.......#....##.#..",
	"...........#..##.....#...#..#..",
	"......##..#........#...........",
	"#.#..#.#.#..#.....#....#.....#.",
	"..#....##....##...#.....#......",
	".#.#....#..#.#......#..###...#.",
	".......#...#.#....##..#..#..#..",
	".#.#.#.......#....#.#.#.......#",
	".#..........#.##.#...#..#.#.##.",
	"..#.#..........#.#....##.#.##..",
	"###..#..#.#...##.#.#..#........",
	"##....#...#....#....#...#.#....",
	"#...#.#....#.##..##...#.#......",
	"......#...#.###......##....#...",
	".................#.###......#..",
	"##..#....#....##...###.#.#..###",
	"..#..........#..####..##..#...#",
	".#......#..#.#...........##.#..",
	".#..#......#...#.#.#..#.#.#.#.#",
	".#......###.....#.#.#......##..",
	"#..........#.##...#...........#",
	"..#....#.##....#.........#.....",
	".#..##....#...##.........#..#..",
	"....##..#.###..#.#...#..###..#.",
	"..#......#........#...#.#......",
	"........#..#..#..#...#.##......",
	".##.#.#......#...#.........#...",
	"#..###.#...#....###.##..###....",
	"........##.............#....#..",
	"...#...............#....#.#....",
	"#..........#..#..#.#.....#...#.",
	".#.............#...#.......#..#",
	".#..#..#...#........##.........",
	".....#.#..#.#..#..##.........#.",
	"..#..##...#....#.#...#.###..#..",
	"#...........##.....#...#.##....",
	"#.#.#.#........##......#...#.#.",
	"......#..#.###.#...#.##.##....#",
	".#....#...#....#........#....#.",
	"..#.#..........#..##.......#..#",
	".....#...##..#................#",
	".#...............##...#.##...##",
	"#.####....##.....#.......#.##..",
	"......#.##.#...##..###..#.#....",
	".#.##.#...##..#.......#.#..#...",
	"#...#.##..........##..........#",
	"#.###...#...#..#.....#.#.##..##",
	".##.....#....#...##.....##.....",
	"...#........#..###.###...#.....",
	"##..#....#.....#...#.#....#.#..",
	"#....#....#.#..........#...#..#",
	"...##..#......#..#..#..#..#....",
	".....##...#..####..##.........#",
	".....#..#.#...#..#....##..##...",
	"..#.......##.#..#.##...#.#....#",
	".#..#.#...##..##....#..#......#",
	"..##.##..##...###..#....#...#..",
	"........##.......##...##.....##",
	".#....###...#..#..#..#.......#.",
	"#.###............#....##.....#.",
	"..........#...#...##..#...#....",
	"..#......#.##.......#....##..##",
	"..#..###.....#...#.......#.....",
	"#.#...##.....#...#....#.......#",
	"....##.##.#....#.....#.#....#..",
	"...#....#.###............#..###",
	"#..##..#.........##.....#.#...#",
	"....#.......##......#....#...#.",
	"....#..##.#..........#.........",
	"....#...#.###.......#...#.#....",
	"#..#..#...#.......##...#..#.##.",
	"#.......#...##.##......#.......",
	"##..##...##...#......#...#...##",
	"..#...#.#.####.#...##.....##...",
	"#...#..#..#...##......#.#..#..#",
	"..##..##.#.#..#...####.....###.",
	".#........#..##.###...#.##.#...",
	"........#..#...##......#.#....#",
	"..#...###.......##..##..#....#.",
	".##...#.#..#.##.......##.###...",
	"#....#.#.#........#....#..#.##.",
	"....#.##.#.##..#.#####.....###.",
	"#.#..#..#...#.#..#.......#.#...",
	"....#...#....###...............",
	".###.#.....#.#.......###......#",
	"##...#.#.###....##..#...##.....",
	"...#.#..#.###.#.......#...#.#..",
	".#...#....#...#..####....###...",
	"..#....##.....##.#.#.##....#...",
	"#....#..##.......#...##.##....#",
	".##..#.......#..#....###.......",
	"#.##.....##.#.........#......##",
	".####.#...#.....#..#...#.##..#.",
	"....#...........#.....#........",
	".#............##...#.......#.#.",
	"#....#.##........#....#.#..#..#",
	"#....#.##....#...##...#..#..#..",
	"...#..#.####.#....#............",
	"....#......#.........#.........",
	"#....##....###.....#......#.#..",
	"...#..#....#........###..#...#.",
	"..#.#........#.#.#.###..#.#.#..",
	".....###.....##.#....###.#.....",
	"##.#....#....##...##.###.#.##..",
	".###.#..#.......##...#...##....",
	".#...###........#.......##.##..",
	"#......####...#...##.#.######..",
	"....##.............#..##.##...#",
	"...........#..##.#...#.#.#...#.",
	"###.......#.##..#....#...#....#",
	".........#.....#.#.#..##.#.....",
	"#...##..#....#..#.............#",
	"...#.......#.##.............#.#",
	".....#..#...##......####..#....",
	".#.#.#.....#...####..#...##...#",
	"#...#.#..#..#.#..#.##..........",
	".....#.##..#.#.##..#.#.#....#.#",
	"...##..#...#...#..#....#.......",
	"........#.#..#...#...#.#...#...",
	"##..#........#..#.....#......##",
	".........#..#...#......#......#",
	"..#.#.#........##...#.##.....##",
	".###....##....#...#....#..#....",
	".#.............###...#..##..###",
	".##.##.##.......###.........#.#",
	"..#..###...#...#....#..#.#..#.#",
	"......#..#.#..#.....#.#........",
	"......#...####...#.#.....#.....",
	".#...##.......#..#......#...#..",
	"#..#...#.......###..#..#.#.#.#.",
	".....#.....###.##....#.#.##.#.#",
	"#........#....##...#..#.##..#..",
	"...#.#........##....#.#..###.#.",
	"#...#...##.........#........###",
	"##...#.##..##...#.....#.###.#..",
	"#.###.#.#..#...........##..#...",
	"........#.......#..#..#.###....",
	"#........#....#......###.......",
	"..#.###.######...#.###..#......",
	"...#...######..#.....#....#.#..",
	"..#.......#..#..#.........#...#",
	".#...#..##.##.........##.......",
	".........#.#.##.#..#....#.#...#",
	"#.......#....#......#.....###.#",
	"##..............#.###........#.",
	"..#.##..#.##.....#...#.#.#..###",
	"..#.#......#..#..##.#........#.",
	"..#.....#...#.#...#...###..#.#.",
	".......#...........#..#..#.#.##",
	".......#...##..#.###...........",
	".#........#.###.#..#..#..#..#..",
	"##.#.##....#..###..#.##.##...#.",
	".....#....##.#........#.#.#....",
	"....##....#..#..#....##....#.#.",
	"#.....##....#.....#.###.#....#.",
	".#.##.##..#..#...#...........##",
	"...#..###..#.....##....#.......",
	"...#..##..###.#..#..#.#........",
	"......##..#.......#..##.....###",
	".#...##.#.#.#......#...#.#.#.##",
	"....#.#....#...#........#...#..",
	"....#.#......#.#.###.#.#.##.#..",
	"#..#........###..#..#..#.....#.",
	"...#....#...##...#........##.##",
	".....#..#..#.....#....#.#...#..",
	"..#.###....#.#..##......#.##.#.",
	"..####......#..#.#.#..#.#####..",
	".......##..#..###.#............",
	"..###.#........#..........##.##",
	"#.#.........#.##.#......#..#...",
	"...#.....#.....##..#..##.##..#.",
	"#.#.##....#.......###....##....",
	"...##.#..#...##.#..#......#..#.",
	"..##.........#.##.#####...#.#..",
	".#....#...#....#.#.....##...###",
	"##.....#..####............###.#",
	"......#...........#....#.......",
	".#......#.....##...........###.",
	"#......##.......#.#.#..##.....#",
	"...###.#.....##.#...#.#....#.#.",
	"...###.......#...#.............",
	"..#..#.#....#.#.###.#.#.##..##.",
	"..##...#..#.#..##.#.##....##...",
	"..#...........#..#....#....#...",
	"#.##...........#..#.#..##.#.#..",
	"...##...##................#..#.",
	".#...#.##......#.#......#.####.",
	"#.##....#....#.........#....###",
	".....###........#.#.#.##...#.##",
	".....#....#.#....#.........#..#",
	"..#...#.#.#.#...#...#...##.#..#",
	"###.......#.....#.............#",
	"#.####.#.......#.#.#.#..#.#....",
	"#..#..#####......#....#..##....",
	"...............#.....#.#....###",
	".###.....#...#.##..#.#..#.#####",
	"#.##.....#......##.......##....",
	"..........###.......#...#.#....",
	"..#.#..#...##.....#........#.#.",
	"........##.##....#####.#.#..##.",
	"..##.#.#...#####..........#.#.#",
	"#.........#......##...#.....#..",
	".##.#........#...#..##...#...#.",
	".......##..#...#.....#.##......",
	"....#.#...##..##..#....##......",
	"#........#..........##..####.#.",
	"...###...#.#.###.#...#....#.#.#",
	".....##.#.....#........#.#....#",
	"#.......#....#...##..#......#..",
	"...#..........#.#.#...#.#.###.#",
	"....##.....#.##..#.#.#.........",
	"#.##..##..#....#.........#...#.",
	".###..........#...##.#..#......",
	".....####.............##...###.",
	".#..#....#..#...#..#...........",
	"#..#..##..#...#.##..#.###.#...#",
	"......#.#..###...#..#.....#..#.",
	"##.##......#...#.#...#.........",
	"....##.#.......#.#..##....#.#.#",
	"#..##..#...###.#....##.........",
	".............#.#....#...##..#..",
	"..#....#...#.....#.##.#..##..##",
	"##.#..##.#..##.#.#.##.#...#.#..",
	".##.#..#.#........##.#...##....",
	"#.........##....##..#......#...",
	".#.#.......##...#..#......###.#",
	"........#.#.#.#......#....#..#.",
	"...##..#...#...#.##..#....#.#..",
	"...#.#.#.#.......#.......###..#",
	"...#..##..#####.#.....##.#..#..",
	".......#.#.....#.....#...#...##",
	"#...#...#......##.#....##......",
	"#.....#.#.#.....#....#......#..",
	"..#..#.##.#......##..#.#..#..##",
	"####...#.....#....#.#..........",
	"....#.....###...#...##......#..",
	".....#....#...#............#...",
	"...#...#..##.........#...#...##",
	"#.#..#.#...##.#.......#..#.#...",
	".#.....#...##.............#...#",
	".....#..##..#....#......#.##..#",
	"....#...###.................#..",
	"...###...#....#...#...#........",
	"....#.##.#.......#..#..........",
	"...#..#......#.#...###...#.#...",
	"..#.#..#...#.......#.......#.#.",
	".#.#...#.#.##........#.........",
	"...#..#...#....#.#.#.#.#..###..",
	".#..##......#.#.##..#.##....#..",
	"#....#.......##.....#.#........",
	"..###..#.#.#.......##....#.....",
	"........#.#.#....##...##..#....",
	"#....##.#....#...##..##...#....",
	"...#..##.#.....#...#.....##....",
	".#.#..#..#...#....#..##.#....#.",
	"##.#.##....#.....#....#....#.#.",
	".##......#............##...#...",
	"#..##.#.####.#.#....#..#..#.#.#",
	"#...##...#......##....###.....#",
	"..#.##.....#....#....#......#..",
	".##.#...#.....#.#.#.#........##",
	".#..#....#.#...........#...#...",
	"#.....#..#.....#.#.##.#.....#..",
	"....#.....#..#.#....###........",
	".....###...........#.#..##.#.#.",
	"....###....#.......###..#...#.#",
	".###.....#...##.#...##........#",
	"..#..#.#..#...#.#...#.#..#...#.",
	"#.####.......#....##.#..#.#..#.",
	"....#.#.##.#..###.........##.#.",
	"..#..#.#....#....#.##..........",
	"..##.###..#.#..#.#......#....#.",
	".#..#.....##...#.#......##.#..#",
	"#.#....#..#.#.#........#.###...",
	"...#....##....##..###.###.#.#..",
	"..#....#.....#....##.#.........",
	"#.......#....#.........##..#...",
	".#..#...#.#..#.#....#.#........",
	"...#..###...#.....#......##....",
	"..#...........#.....#..........",
	"....###.#...#......#...#..#....",
	".....#.##..#..#....#.......#..#",
	"....##..#.#.#..............#.#.",
	".#.#..#..#.#......#...#.#......",
	"....#.......#.##....##.#.#.#..#",
	"............#.#.#.....##.......",
	"........#...##.#..#......#...##",
	".........#...#...#....#...#.##.",
	"..#.....#......#......#.....#..",
	"#....#...##..#.#....#.#...#.###",
	".......#..#..#..#.#...#.....#.#",
	"...#.#...#......##.....#..#....",
	"...#.#.####..##.#..#...........",
	"..##..##....#.....####...#....#",
	"###.......#...##.#...#...#...#.",
	".##..#.....#..####......#....#.",
	"#.....#..#..##..##...#..#..#...",
	".#....#.....#...####..####.....",
	"..#....#...#......#........#.#.",
	"##.#.......#..#.....#..##..##..",
	".#..#..#.#.#...#....##...#.##.#",
	"##...#..#....#.........##......",
}