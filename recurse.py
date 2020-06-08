grid = [
    [5,3,4,6,7,8,9,1,2],
    [6,7,2,1,9,5,3,4,8],
    # etc
]

def solve():
    global grid
    for y in range(9):
        for x in range(9):
            if grid[y][x] == 0:
                for n in range(1,10):
                    if possible(y,x,n):
                        grid[y][x] = n
                        solve()
                        grid[y][x] = 0
                return
    print(np.matrix(grid))
    input("More?")                    